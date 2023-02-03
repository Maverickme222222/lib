// Package producer provides a high-level interface for producing messages to a Kafka cluster.
// go:generate mockgen -source producer.go -destination mock/producer_mock.go -package mock
package producer

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"os"

	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"go.elastic.co/apm/v2"

	kappaKafka "github.com/kappapay/backend/lib/golang/src/kappa/kafka"
)

// KappaProducer is a high-level interface for producing messages to a Kafka cluster.
type KappaProducer struct {
	Producer kappaKafka.Writer
	Topic    string
	logger   *zerolog.Logger
}

// KappaProducerI ...
type KappaProducerI interface {
	Send(ctx context.Context, key []byte, payload interface{}) error
	SendToTopic(ctx context.Context, topic string, key []byte, payload interface{}) error
	Close() error
}

// NewKappaProducer creates a new KappaProducer.
func NewKappaProducer(config kappaKafka.Config) *KappaProducer {
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	if config.DisableTLS {
		tlsConfig = nil
	}

	var saslConfig sasl.Mechanism = &plain.Mechanism{
		Username: config.Username,
		Password: config.Password,
	}
	if config.DisableSASL {
		saslConfig = nil
	}

	logger := config.Logger
	if logger == nil {
		l := zerolog.New(os.Stderr).With().Logger()
		logger = &l
	}

	return &KappaProducer{
		Topic: config.Topic,
		Producer: &kafka.Writer{
			Addr:     kafka.TCP(config.BrokerURL),
			Balancer: &kafka.LeastBytes{},
			Transport: &kafka.Transport{
				TLS:  tlsConfig,
				SASL: saslConfig,
			},
			Async: false,
		},
		logger: logger,
	}
}

// Send sends a message to Kafka.
func (k *KappaProducer) Send(ctx context.Context, key []byte, payload interface{}) error {
	return k.SendToTopic(ctx, k.Topic, key, payload)
}

// SendToTopic sends a message to Kafka.
func (k *KappaProducer) SendToTopic(ctx context.Context, topic string, key []byte, payload interface{}) error {
	span, ctx := apm.StartSpan(ctx, "kafka.producer.Send", "kafka")
	defer span.End()

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Topic: topic,
		Key:   key,
		Value: data,
	}

	if err := k.Producer.WriteMessages(ctx, msg); err != nil {
		return err
	}

	return nil
}

// Close ...
func (k *KappaProducer) Close() error {
	return k.Producer.Close()
}
