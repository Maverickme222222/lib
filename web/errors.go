package web

import (
	"fmt"

	kiterror "github.com/kappapay/backend/lib/golang/src/kappa/errors"
)

const (
	// DomainKappaWeb is the default value of the errors.Domain type for the Relay service.
	DomainKappaWeb kiterror.Domain = "KappaWeb"
	// RequestValidation indicates that some part of the request contents could not be validated.
	RequestValidation KappaWebReason = "Request_Validation"
	// RequestMediaTypeValidation indicates the media type is not supported
	RequestMediaTypeValidation KappaWebReason = "Request_Media_Type_Validation"
)

type (
	// KappaWebReason is the local type to override the errors.Reason type and implement the errors.LocalError interface.
	KappaWebReason kiterror.Reason
)

// Reason implements the errors.LocalError interface.
func (r KappaWebReason) Reason() kiterror.Reason {
	return kiterror.Reason(r)
}

// Domain implements the errors.LocalError interface.
func (r KappaWebReason) Domain() kiterror.Domain {
	return DomainKappaWeb
}

// Message implements the errors.LocalError interface and just defaults to the reason code.
func (r KappaWebReason) Message() string {
	switch r {
	case RequestValidation:
		return "Validation of request failed."
	case RequestMediaTypeValidation:
		return "Validation of request media type failed."
	default:
		return string(r)
	}
}

// Type implements the errors.LocalError interface.
func (r KappaWebReason) Type() kiterror.Type {
	switch r {
	case RequestValidation:
		return kiterror.ErrUser
	case RequestMediaTypeValidation:
		return kiterror.UnsupportedMediaType
	default:
		return kiterror.ErrSystem
	}
}

// New creates a new errors.Error from the provided ApplicationReason.
func (r KappaWebReason) New(err error) kiterror.Error {
	return kiterror.NewLocal(err, r)
}

// Newf is a convenience method to implement the same signature as fmt.Errorf
func (r KappaWebReason) Newf(format string, args ...interface{}) kiterror.Error {
	return kiterror.NewLocal(fmt.Errorf(format, args...), r)
}
