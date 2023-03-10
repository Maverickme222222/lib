package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/golang/gddo/httputil/header"
)

// DecodeJSONBody decodes the body from the request
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			return RequestMediaTypeValidation.Newf("Content-Type header is not application/json")
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return RequestValidation.Newf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return RequestValidation.Newf("Request body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			return RequestValidation.Newf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return RequestValidation.Newf("Request body contains unknown field %s", fieldName)

		case errors.Is(err, io.EOF):
			return RequestValidation.Newf("Request body must not be empty")

		case err.Error() == "http: request body too large":
			return RequestValidation.Newf("Request body must not be larger than 1MB")

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return RequestValidation.Newf("Request body must only contain a single JSON object")
	}

	return nil
}
