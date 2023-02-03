// Package kafka ...
package kafka

//go:generate mockgen -source models.go -destination mock/models_mock.go -package mock

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
)

// Config holds all configuration options for the Kafka producer.
type Config struct {
	BrokerURL string `json:"broker_url"`
	Topic     string `json:"topic"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	GroupID   string `json:"groupID"`

	// MaxRetries is the number of times to retry consuming a message before failing. If the dead letter topic is
	// enabled it will send the message to the dead letter topic after failing.
	MaxRetries *int `json:"max_retires"`

	// EnableDeadLetterTopic messages that fail to be consumed will be sent to the dead letter topic.
	// The dead letter topic is the topic name with appended with "-dl"
	EnableDeadLetterTopic bool `json:"enable_dead_letter_topic"`

	Logger *zerolog.Logger

	// Should be used for testing and development only
	// DisableTLS Disable TLS connection when connecting to Kafka. Useful for local testing
	DisableTLS bool `json:"disable_tls"`

	// DisableSASL Set to true if no username/password is need when connecting. Useful testing/development.
	DisableSASL bool `json:"disable_sasl"`
}

// Writer ...
type Writer interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

// Reader ...
type Reader interface {
	FetchMessage(ctx context.Context) (kafka.Message, error)
	CommitMessages(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}
