// Package changeeventproducer
package changeeventproducer

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
	"github.com/kappapay/backend/lib/golang/src/kappa/kafka/producer"
	"github.com/kappapay/backend/lib/golang/src/kappa/kafka/producer/mock"
)

func TestProducerImpl_CreatedChangeEvent(t *testing.T) {
	type fields struct {
		topicPrefix string
		prod        producer.KappaProducerI
	}
	type args struct {
		ctx   context.Context
		key   []byte
		after protoreflect.ProtoMessage
	}
	tests := map[string]struct {
		fields  fields
		args    args
		wantErr bool
	}{
		"send_to_correct_topic": {
			fields:  fields{topicPrefix: "kappa.dev.", prod: stubKafkaProducer(t, "kappa.dev.KycResultsReadyEvent", gomock.Any())},
			args:    args{after: &models.KycResultsReadyEvent{}},
			wantErr: false,
		},
		"valid_payload": {
			fields: fields{
				topicPrefix: "kappa.dev.",
				prod: stubKafkaProducer(t, "kappa.dev.KycResultsReadyEvent", &models.ChangeEvent{
					Type:  models.ChangeEventType_CHANGE_EVENT_TYPE_CREATED,
					After: toMap(&models.KycResultsReadyEvent{})})},
			args:    args{after: &models.KycResultsReadyEvent{}},
			wantErr: false,
		},
		"error_case": {
			fields:  fields{topicPrefix: "kappa.dev.", prod: failureStubKafkaProducer(t)},
			args:    args{after: &models.KycResultsReadyEvent{}},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := &ProducerImpl{
				topicPrefix: tt.fields.topicPrefix,
				prod:        tt.fields.prod,
			}
			if err := p.CreatedChangeEvent(tt.args.ctx, tt.args.key, tt.args.after, nil); (err != nil) != tt.wantErr {
				t.Errorf("ProducerImpl.CreatedChangeEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func stubKafkaProducer(t *testing.T, topicName string, payload interface{}) producer.KappaProducerI {
	ctrl := gomock.NewController(t)
	mockKappaProducer := mock.NewMockKappaProducerI(ctrl)
	mockKappaProducer.EXPECT().SendToTopic(gomock.Any(), topicName, gomock.Any(), payload)
	return mockKappaProducer
}

// This stub always returns an error
func failureStubKafkaProducer(t *testing.T) producer.KappaProducerI {
	ctrl := gomock.NewController(t)
	mockKappaProducer := mock.NewMockKappaProducerI(ctrl)
	mockKappaProducer.EXPECT().SendToTopic(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("network error"))
	return mockKappaProducer
}

func toMap(v interface{}) *structpb.Struct {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	mapResult := map[string]interface{}{}
	err = json.Unmarshal(jsonBytes, &mapResult)
	if err != nil {
		panic(err)
	}

	result, err := structpb.NewStruct(mapResult)
	if err != nil {
		panic(err)
	}
	return result

}
