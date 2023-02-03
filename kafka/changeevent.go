package kafka

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
)

var (
	// ErrInvalidEvent is returned when the data provided does not match the ChangeEventType
	ErrInvalidEvent = fmt.Errorf("event invalid")
	// DefaultUnmarshaler custom unmarshaling options, specifically for ignoring new fields in the protos
	DefaultUnmarshaler = protojson.UnmarshalOptions{DiscardUnknown: true}
)

// NewChangeEvent is a helper method for creating ChangeEvents protos.
// This is useful because the Before and After fields need to be wrapped in an Any proto
func NewChangeEvent(changeType models.ChangeEventType, before, after protoreflect.ProtoMessage, metadata map[string]interface{}) (*models.ChangeEvent, error) {
	event := models.ChangeEvent{
		Type: changeType,
	}
	if before != nil {
		beforeMap, err := structToMap(before)
		if err != nil {
			return nil, err
		}
		anyWrapper, err := structpb.NewStruct(beforeMap)
		if err != nil {
			return nil, err
		}
		event.Before = anyWrapper
	}

	if after != nil {
		afterMap, err := structToMap(after)
		if err != nil {
			return nil, err
		}
		anyWrapper, err := structpb.NewStruct(afterMap)
		if err != nil {
			return nil, err
		}
		event.After = anyWrapper
	}

	if metadata != nil {
		metadataStruct, err := structpb.NewStruct(metadata)
		if err != nil {
			return nil, err
		}
		event.Metadata = metadataStruct
	}

	return &event, nil
}

// NewCreatedEvent ...
func NewCreatedEvent(after protoreflect.ProtoMessage, metadata map[string]interface{}) (*models.ChangeEvent, error) {
	return NewChangeEvent(models.ChangeEventType_CHANGE_EVENT_TYPE_CREATED, nil, after, metadata)
}

// NewUpdatedEvent ...
func NewUpdatedEvent(before, after protoreflect.ProtoMessage, metadata map[string]interface{}) (*models.ChangeEvent, error) {
	return NewChangeEvent(models.ChangeEventType_CHANGE_EVENT_TYPE_UPDATED, before, after, metadata)
}

// NewDeletedEvent ...
func NewDeletedEvent(before protoreflect.ProtoMessage, metadata map[string]interface{}) (*models.ChangeEvent, error) {
	return NewChangeEvent(models.ChangeEventType_CHANGE_EVENT_TYPE_DELETED, before, nil, metadata)
}

// ParseChangeEvent ...
func ParseChangeEvent(changeEvent *models.ChangeEvent, beforeDest, afterDest protoreflect.ProtoMessage) (models.ChangeEventType, map[string]interface{}, error) {
	var metadata map[string]interface{}
	if changeEvent.Metadata != nil {
		metadata = changeEvent.Metadata.AsMap()
	}

	switch changeEvent.Type {
	case models.ChangeEventType_CHANGE_EVENT_TYPE_CREATED:
		return models.ChangeEventType_CHANGE_EVENT_TYPE_CREATED, metadata, parseCreatedEvent(changeEvent, afterDest)

	case models.ChangeEventType_CHANGE_EVENT_TYPE_UPDATED:
		return models.ChangeEventType_CHANGE_EVENT_TYPE_UPDATED, metadata, parseUpdatedEvent(changeEvent, beforeDest, afterDest)

	case models.ChangeEventType_CHANGE_EVENT_TYPE_DELETED:
		return models.ChangeEventType_CHANGE_EVENT_TYPE_DELETED, metadata, parseDeletedEvent(changeEvent, beforeDest)
	default:
		return models.ChangeEventType_CHANGE_EVENT_TYPE_UNSPECIFIED, metadata, fmt.Errorf("unknown change type: %w", ErrInvalidEvent)
	}
}

func parseCreatedEvent(changeEvent *models.ChangeEvent, afterDest protoreflect.ProtoMessage) error {
	if changeEvent.Before != nil || changeEvent.After == nil {
		return fmt.Errorf("before should be nil and after non-nil: %w", ErrInvalidEvent)
	}
	return unmarshalStruct(changeEvent.After, afterDest)
}

func parseUpdatedEvent(changeEvent *models.ChangeEvent, beforeDest, afterDest protoreflect.ProtoMessage) error {
	if changeEvent.Before == nil || changeEvent.After == nil {
		return fmt.Errorf("before and after should be non-nil: %w", ErrInvalidEvent)
	}
	err := unmarshalStruct(changeEvent.Before, beforeDest)
	if err != nil {
		return err
	}

	err = unmarshalStruct(changeEvent.After, afterDest)
	if err != nil {
		return err
	}
	return nil
}

func parseDeletedEvent(changeEvent *models.ChangeEvent, beforeDest protoreflect.ProtoMessage) error {
	if changeEvent.Before == nil || changeEvent.After != nil {
		return fmt.Errorf("before should be non-nil and after nil: %w", ErrInvalidEvent)
	}
	return unmarshalStruct(changeEvent.Before, beforeDest)
}

func structToMap(v protoreflect.ProtoMessage) (map[string]interface{}, error) {
	jsonBytes, err := protojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func unmarshalStruct(v *structpb.Struct, dest protoreflect.ProtoMessage) error {
	jsonBytes, err := v.MarshalJSON()
	if err != nil {
		return err
	}
	return DefaultUnmarshaler.Unmarshal(jsonBytes, dest)
}
