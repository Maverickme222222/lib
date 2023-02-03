// Package consumer provides a high-level interface for reading from Kafka.
package consumer

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"os"
	"sync"
	"time"

	"go.elastic.co/apm/v2"

	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"

	kappaKafka "github.com/kappapay/backend/lib/golang/src/kappa/kafka"
)

const (
	// DeadLetterSuffix ...
	DeadLetterSuffix = "-dl"
	// DefaultRetryDelay ...
	DefaultRetryDelay = time.Second * 3
	// DefaultMaxRetries ...
	DefaultMaxRetries = 2
)

// KappaConsumer is a high-level interface for reading from Kafka.
type KappaConsumer struct {
	Consumer         kappaKafka.Reader
	MaxRetries       int
	RetryDelay       time.Duration
	DeadLetterWriter kappaKafka.Writer
	Logger           *zerolog.Logger
	Topic            string // will serve as name for apm transaction logs

	mutex sync.Mutex

	// currentMessage holds the current message being consumed while it is being consumed and retried.
	currentMessage *kafka.Message
	retried        int
}

// NewKappaConsumer creates a new Kafka consumer.
func NewKappaConsumer(config kappaKafka.Config) *KappaConsumer {

	// Use a default TLS configuration (unless TLS is disabled then set it to nil)
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	if config.DisableTLS {
		tlsConfig = nil
	}

	// Use plain SASL authentication (unless it's disabled, then set it to nil)
	var saslConfig sasl.Mechanism = &plain.Mechanism{
		Username: config.Username,
		Password: config.Password,
	}
	if config.DisableSASL {
		saslConfig = nil
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{config.BrokerURL},
		Topic:    config.Topic,
		GroupID:  config.GroupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Dialer: &kafka.Dialer{
			SASLMechanism: saslConfig,
			TLS:           tlsConfig,
		},
		CommitInterval: 0,
	})

	// Created a writer for the dead letter topic is it's enabled.
	var deadLetterWritter *kafka.Writer
	if config.EnableDeadLetterTopic {
		deadLetterWritter = &kafka.Writer{
			Addr:     kafka.TCP(config.BrokerURL),
			Topic:    config.Topic + DeadLetterSuffix,
			Balancer: &kafka.LeastBytes{},
			Transport: &kafka.Transport{
				TLS:  tlsConfig,
				SASL: saslConfig,
			},
		}
	}

	// If not logger is supplied created a default logger
	logger := config.Logger
	if logger == nil {
		l := zerolog.New(os.Stderr).With().Logger()
		logger = &l
	}

	// Use the default number of retries unless one is set on the config
	maxRetries := DefaultMaxRetries
	if config.MaxRetries != nil {
		maxRetries = *config.MaxRetries
	}

	return &KappaConsumer{
		Consumer:         reader,
		MaxRetries:       maxRetries,
		DeadLetterWriter: deadLetterWritter,
		RetryDelay:       DefaultRetryDelay,
		Logger:           logger,
		Topic:            config.Topic,
	}
}

// Read reads from Kafka. The message will be deserialized into the `dest` supplied.
//
// This method call will NOT commit the Kafka offset. The caller should call
// `Commit(ctx)` after the message has been successfully consumed. If `Read` is
// called again before `Commit` is called, then it will be assumed that the consumer
// failed to consumer the message successfully.
func (k *KappaConsumer) Read(ctx context.Context, dest interface{}) error {
	span, ctx := apm.StartSpan(ctx, "kafka.consumer.Read", "kafka")
	defer span.End()
	k.mutex.Lock()
	defer k.mutex.Unlock()

	msg, err := k.getNextMessage(ctx)
	if err != nil {
		return err
	}

	err = json.Unmarshal(msg.Value, dest)

	return err
}

// Commit will comment the offset for the current message being read.const
//
// It's important to call this method after successfully consuming a message,
// otherwise the message will be retry (ie `Read` will re-send the same message,
// until the retries have been exceeded.)
func (k *KappaConsumer) Commit(ctx context.Context) error {
	span, ctx := apm.StartSpan(ctx, "kafka.consumer.Commit", "kafka")
	defer span.End()

	k.mutex.Lock()
	defer k.mutex.Unlock()
	if k.currentMessage == nil {
		return nil
	}
	err := k.Consumer.CommitMessages(ctx, *k.currentMessage)
	k.currentMessage = nil
	k.retried = 0
	return err
}

// Close release the connection to Kafka. Ideally, the caller should `defer c.Close()` after
// creating a new Consumer
func (k *KappaConsumer) Close() error {
	err := k.Consumer.Close()
	return err
}

func (k *KappaConsumer) getNextMessage(ctx context.Context) (*kafka.Message, error) {
	// Last message was never committed.
	if k.currentMessage != nil && k.retried >= k.MaxRetries {
		err := k.handleFailedMessage(ctx, *k.currentMessage)
		if err != nil {
			return nil, err
		}
		k.currentMessage = nil
		k.retried = 0
	}
	if k.currentMessage != nil && k.retried < k.MaxRetries {
		time.Sleep(k.RetryDelay)
		k.retried++
		return k.currentMessage, nil
	}

	msg, err := k.Consumer.FetchMessage(ctx)
	if err != nil {
		return nil, err
	}

	k.currentMessage = &msg
	k.retried = 0
	return &msg, nil
}

func (k *KappaConsumer) handleFailedMessage(ctx context.Context, msg kafka.Message) error {
	k.Logger.Error().Str("topic", msg.Topic).Str("key", string(msg.Key)).Msg("failed to consume message")
	if k.DeadLetterWriter != nil {
		err := k.DeadLetterWriter.WriteMessages(ctx, kafka.Message{Key: msg.Key, Value: msg.Value})
		if err != nil {
			k.Logger.Err(err).Msg("Failed to write message to dead letter queue")
		}
	}
	err := k.Consumer.CommitMessages(ctx, msg)
	return err
}
