// Package consumer provides a high-level interface for reading from Kafka.
package consumer

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"

	kafkaMocks "github.com/kappapay/backend/lib/golang/src/kappa/kafka/mock"
)

func TestKappaConsumer_Read_RetryOnce(t *testing.T) {

	// Setup Mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReader := kafkaMocks.NewMockReader(ctrl)
	anyContext := gomock.AssignableToTypeOf(context.Background())
	jsonPayload := `{"key":"value"}`
	mockReader.EXPECT().FetchMessage(anyContext).Return(kafka.Message{Value: []byte(jsonPayload)}, nil).MaxTimes(1)

	logger := &zerolog.Logger{}

	k := &KappaConsumer{
		Consumer:   mockReader,
		MaxRetries: 1,
		RetryDelay: 0,
		Logger:     logger,
	}
	ctx := context.Background()
	dest := map[string]interface{}{}

	err := k.Read(ctx, &dest)
	assert.Nilf(t, err, "read should not return an error")

	// Reading a second time should not read from Kafka. This should be a retry
	err = k.Read(ctx, &dest)
	assert.Nilf(t, err, "read should not return an error")

}

func TestKappaConsumer_Read_ReadTwoMessages(t *testing.T) {

	// Setup Mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReader := kafkaMocks.NewMockReader(ctrl)
	anyContext := gomock.AssignableToTypeOf(context.Background())
	jsonPayload := `{"key":"value"}`
	mockReader.EXPECT().FetchMessage(anyContext).Return(kafka.Message{Value: []byte(jsonPayload)}, nil).Times(2)
	mockReader.EXPECT().CommitMessages(anyContext, gomock.AssignableToTypeOf(kafka.Message{})).Return(nil).Times(1)

	logger := &zerolog.Logger{}

	k := &KappaConsumer{
		Consumer:         mockReader,
		MaxRetries:       1,
		RetryDelay:       0,
		DeadLetterWriter: nil,
		Logger:           logger,
	}
	ctx := context.Background()
	dest := map[string]interface{}{}

	err := k.Read(ctx, &dest)
	assert.Nilf(t, err, "read should not return an error")

	// Some transient failure. But this is ok because we can retry up to one time.
	err = k.Read(ctx, &dest)
	assert.Nilf(t, err, "read should not return an error")

	// Mark message as consumed. Next Call to Read should fetch a new message
	k.Commit(ctx)

	err = k.Read(ctx, &dest)
	assert.Nilf(t, err, "read should not return an error")
}

func TestKappaConsumer_Read_DeadLetter(t *testing.T) {

	// Setup Mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReader := kafkaMocks.NewMockReader(ctrl)
	anyContext := gomock.AssignableToTypeOf(context.Background())
	jsonPayload := `{"key":"value"`
	badMsg := kafka.Message{Value: []byte(jsonPayload)}
	mockReader.EXPECT().FetchMessage(anyContext).Return(badMsg, nil).Times(2)
	mockReader.EXPECT().CommitMessages(anyContext, gomock.AssignableToTypeOf(kafka.Message{})).Return(nil).Times(1)

	mockDeadLetterWriter := kafkaMocks.NewMockWriter(ctrl)
	mockDeadLetterWriter.EXPECT().WriteMessages(anyContext, badMsg).Return(nil).Times(1)
	logger := &zerolog.Logger{}

	k := &KappaConsumer{
		Consumer:         mockReader,
		MaxRetries:       0,
		RetryDelay:       0,
		DeadLetterWriter: mockDeadLetterWriter,
		Logger:           logger,
	}
	ctx := context.Background()
	dest := map[string]interface{}{}

	err := k.Read(ctx, &dest)
	assert.Errorf(t, err, "expecting json unmarshaling error")

	// Since we have zero retries this Read should send the message to the dead letter topic
	err = k.Read(ctx, &dest)
	// This mock returns the same message everytime so this is give the same json unmarshaling error
	assert.Errorf(t, err, "expecting json unmarshaling error")
}
