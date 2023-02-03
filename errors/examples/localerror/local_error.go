// nolint
package main

import (
	"fmt"
	"time"

	errorKit "github.com/kappapay/backend/lib/golang/src/kappa/errors"

	"github.com/kappapay/backend/lib/golang/src/kappa/errors/examples/localerror/errors"
)

func main() {
	// Basic error without many overrides or customizations.
	fmt.Println(RawErrorString(BasicNew()))

	// Slightly more advanced with an override Message, and using Map instead of pairs.
	fmt.Println(RawErrorString(OverrideNew()))
}

func BasicNew() errorKit.Error {
	entityID := "some_db_record_id_from_request"
	if err := DatabaseLookup(entityID); err != nil {
		return errors.EntityNotFound.New(err).
			WithMetaPairs("entityID", entityID)
	}

	return errorKit.Error{}
}

func OverrideNew() errorKit.Error {
	entityID := "some_other_db_id_downstream"
	if err := DatabaseLookup(entityID); err != nil {
		return errors.RequestValidationFailed.New(err).
			WithMessage("Could not locate OtherEntity with ID").
			WithMetaMap(map[string]string{
				"OtherEntityID": entityID,
				"EntityID":      "some_db_record_id_from_request",
			})
	}

	return errorKit.Error{}
}

func DatabaseLookup(id string) error {
	return fmt.Errorf("no record found with id: %s", id)
}

func RawErrorString(err errorKit.Error) string {
	return fmt.Sprintf(`
		"Err":       "%s",
		"Type":      "%s",
		"Message":   "%s",
		"Timestamp": "%s",
		"Reason":    "%s",
		"Domain":    "%s",
		"Metadata":  "%v"
	`, err.Err.Error(), err.Type, err.Message, err.Timestamp.Format(time.RFC3339), err.Reason, err.Domain, err.Metadata)
}
