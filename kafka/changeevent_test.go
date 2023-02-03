package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
)

func TestNewCreatedEvent(t *testing.T) {
	kyc := &models.KycResultsReadyEvent{VerificationId: "a"}
	metadata := map[string]interface{}{"key": "value"}
	event, err := NewCreatedEvent(kyc, metadata)
	assert.Nilf(t, err, "error creating change event")
	require.NotNilf(t, event, "failed to create change event")

	got := &models.KycResultsReadyEvent{}
	eventType, parsedMetadata, err := ParseChangeEvent(event, nil, got)
	assert.Nilf(t, err, "got error parsing event")
	assert.Equalf(t, models.ChangeEventType_CHANGE_EVENT_TYPE_CREATED, eventType, "parsed event type is not correct")
	assert.Equal(t, metadata["key"], parsedMetadata["key"])
	assert.Nilf(t, err, "error while creating change event")
	require.NotNil(t, got, "no change event returned")
	assert.Equalf(t, kyc.VerificationId, got.VerificationId, "verificationID not converted correctly")

	assert.Equalf(t, kyc, got, "not equal after parsing change event")

}

func TestParseChangeEvent(t *testing.T) {
	tx := &models.Transaction{
		Id: "some-id",
		Source: &models.Source{
			Id:     "some-id",
			Source: &models.Source_LinkedAccount{LinkedAccount: &models.LinkedAccount{Id: "some-id"}},
		},
	}
	chEvent, err := NewCreatedEvent(tx, nil)
	assert.Nilf(t, err, "error while creating change event")

	var before, after models.Transaction
	_, _, err = ParseChangeEvent(chEvent, &before, &after)
	assert.Nilf(t, err, "error while parsing change event")
}
