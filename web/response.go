package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	kitErrors "github.com/kappapay/backend/lib/golang/src/kappa/errors"
	"github.com/kappapay/backend/lib/golang/src/kappa/shared"
)

const (
	// ContentType is an HTTP Header key to indicate the type of content that will be in a request or response.
	ContentType = "Content-Type"

	// ContentTypeApplicationJSON is an RFC standard type for UTF-8 JSON content.
	ContentTypeApplicationJSON = "application/json; charset=utf-8"
)

type (
	//FallbackErrorHandler lets callers define how to handle errors that result from failures in handling Error in RespondError.
	FallbackErrorHandler interface {
		HandleError(w http.ResponseWriter, err error)
	}

	// FallbackErrorHandlerFn is a function definition that allows callers to specify how a Fatal error should be handled by the
	// RespondError function.
	FallbackErrorHandlerFn func(w http.ResponseWriter, err error)

	// code is a local custom type to handle proper marhsaling of the kiterror.Type
	code kitErrors.Type

	// KappaHTTPResponse ...
	KappaHTTPResponse struct {
		Success bool        `json:"success"`
		Errors  []error     `json:"errors"`
		Data    interface{} `json:"data"`
	}
)

var (
	typeToStatus = map[kitErrors.Type]int{
		kitErrors.ErrUser:              http.StatusBadRequest,
		kitErrors.ErrSystem:            http.StatusInternalServerError,
		kitErrors.ExpectationFailed:    http.StatusExpectationFailed,
		kitErrors.NotAuthorized:        http.StatusForbidden,
		kitErrors.NotFound:             http.StatusNotFound,
		kitErrors.RequestTimeout:       http.StatusRequestTimeout,
		kitErrors.ServiceUnavailable:   http.StatusServiceUnavailable,
		kitErrors.Unauthenticated:      http.StatusUnauthorized,
		kitErrors.Unimplemented:        http.StatusNotImplemented,
		kitErrors.Unknown:              http.StatusInternalServerError,
		kitErrors.BadGateway:           http.StatusBadGateway,
		kitErrors.Conflict:             http.StatusConflict,
		kitErrors.UnprocessableEntity:  http.StatusUnprocessableEntity,
		kitErrors.UnsupportedMediaType: http.StatusUnsupportedMediaType,
	}
)

// HandleError delegates the parameters to the method receiver function type.
func (fn FallbackErrorHandlerFn) HandleError(w http.ResponseWriter, err error) {
	fn(w, err)
}

// HTTPStatus uses an internal mapping to convert the errors.Type field to the complimentary and standard HTTP Status.
func (c code) HTTPStatus() int {
	if httpStatus, hasHTTPStatus := typeToStatus[kitErrors.Type(c)]; hasHTTPStatus {
		return httpStatus
	}

	return http.StatusInternalServerError
}

// MarshalJSON returns the result of json.Marshal() on code.HTTPStatus() and is used internally to conform to the
// json.Marshaler contract
func (c code) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.HTTPStatus())
}

// Respond handles all writes to the ResponseWriter, expecting a JSON-compliant struct to Marshal, setting the
// Content-Type on the header to "application/json; charset=utf-8", and writing the provided statusCode.
// An error is returned if the struct fails to be marshalled, or the resulting bytes cannot be written to the ResponseWriter.
func Respond(w http.ResponseWriter, data interface{}, statusCode int, success bool, err error) error {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	if data == nil {
		data = map[string]interface{}{}
	}

	errors := []error{}

	if err != nil {
		errors = append(errors, err)
	}

	resp := KappaHTTPResponse{
		Success: success,
		Data:    data,
		Errors:  errors,
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set(ContentType, ContentTypeApplicationJSON)

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	if _, writeErr := w.Write(bytes); writeErr != nil {
		return writeErr
	}

	return nil
}

// RespondSuccess responds with a valid response with the provided data.
func RespondSuccess(w http.ResponseWriter, data interface{}, statusCode int) {
	if returnedErr := Respond(w, data, statusCode, true, nil); returnedErr != nil {
		return
	}
}

/*
RespondError accepts an error type and calls Respond() formatting the error to an externally-safe and HTTP-compliant
JSON representation of error.
If the err parameter is of type kappa/errors.Error, the Message field is passed to the Response body. If the Type field
has the value ErrUser, the status code will be set to 400 - BadRequest. If the Type field has the value ErrSystem, the
status code will be set to 500 - Internal Server Error.
If the err parameter is any other error type, the value "Internal Server Error" will be passed to the Response Body, and
the status code will be set to 500 - Internal Server Error.
If an error occurs while handling the input, the provided FatalHandler is invoked to handle that error, and a
http.StatusInternalServerError is returned.
*/
func RespondError(w http.ResponseWriter, err error) {
	var kitError kitErrors.Error
	if errors.As(err, &kitError) {
		if returnedErr := Respond(w, nil, code(kitError.Type).HTTPStatus(), false, kitError); returnedErr != nil {
			return
		}
		return
	}

	errResp := kitErrors.Error{
		Err:       fmt.Errorf("responded error was not of type kappa/errors.Error: %w", err),
		Type:      http.StatusInternalServerError,
		Message:   http.StatusText(http.StatusInternalServerError),
		Timestamp: time.Now().UTC(),
		Domain:    kitErrors.Domain(shared.SystemRelay),
	}

	if returnedErr := Respond(w, nil, http.StatusInternalServerError, false, errResp); returnedErr != nil {
		return
	}
}
