// Package changeeventproducer ...
package changeeventproducer

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/kappapay/backend/lib/golang/src/kappa/kafka"
	"github.com/kappapay/backend/lib/golang/src/kappa/kafka/producer"
)

//go:generate mockgen -source changeeventproducer.go -destination mock/changeeventproducer_mock.go -package mock

// Producer - A Producer will create change event messages and send these messages to a kafka topic of the same proto type
// Use this Producer when sending change events to Kafka. This package will take care of choosing the correct topic
// and serializing the data in a standard format.
// Note that the Producer will choose the topic name by taking the `topicPrefix` and concatenating it with the name of the proto.
// For example, if the prefix were "kappa.staging." and the proto was of type BusinessEntity. then the topic would
// be "kappa.staging.BusinessEntity"
type Producer interface {
	CreatedChangeEvent(ctx context.Context, key []byte, after protoreflect.ProtoMessage, metadata map[string]interface{}) error
	UpdatedChangeEvent(ctx context.Context, key []byte, before, after protoreflect.ProtoMessage, metadata map[string]interface{}) error
	DeletedChangeEvent(ctx context.Context, key []byte, before protoreflect.ProtoMessage, metadata map[string]interface{}) error
	Close() error
}

// ProducerImpl concrete implementation of Producer
type ProducerImpl struct {
	topicPrefix string
	prod        producer.KappaProducerI
}

var _ Producer = &ProducerImpl{}

// NewProducer ...
func NewProducer(topicPrefix string, prod producer.KappaProducerI) (*ProducerImpl, error) {
	return &ProducerImpl{
		topicPrefix: topicPrefix,
		prod:        prod,
	}, nil
}

// CreatedChangeEvent implements Producer
func (p *ProducerImpl) CreatedChangeEvent(ctx context.Context, key []byte, after protoreflect.ProtoMessage, metadata map[string]interface{}) error {
	event, err := kafka.NewCreatedEvent(after, metadata)
	if err != nil {
		return err
	}
	topic := p.getTopicName(after)
	return p.prod.SendToTopic(ctx, topic, key, event)
}

// DeletedChangeEvent implements Producer
func (p *ProducerImpl) DeletedChangeEvent(ctx context.Context, key []byte, before protoreflect.ProtoMessage, metadata map[string]interface{}) error {
	event, err := kafka.NewDeletedEvent(before, metadata)
	if err != nil {
		return err
	}
	topic := p.getTopicName(before)
	return p.prod.SendToTopic(ctx, topic, key, event)
}

// UpdatedChangeEvent implements Producer
func (p *ProducerImpl) UpdatedChangeEvent(ctx context.Context, key []byte, before protoreflect.ProtoMessage, after protoreflect.ProtoMessage, metadata map[string]interface{}) error {
	event, err := kafka.NewUpdatedEvent(before, after, metadata)
	if err != nil {
		return err
	}
	topic := p.getTopicName(before)
	afterTopic := p.getTopicName(after)
	if topic != afterTopic {
		return fmt.Errorf("before and after must be the same type: before type=%s, after type=%s", topic, afterTopic)
	}
	return p.prod.SendToTopic(ctx, topic, key, event)
}

func (p ProducerImpl) getTopicName(msg protoreflect.ProtoMessage) string {
	return p.topicPrefix + string(msg.ProtoReflect().Descriptor().Name())
}

// Close ...
func (p *ProducerImpl) Close() error {
	return p.prod.Close()
}
