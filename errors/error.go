// Package errors ...
package errors

// Package errors provides convenience functions to handle errors.
// Errors can be created using the Error struct and can be printed using `%s` when formatting the error to be logged.
// Errors should be embedded in the other errors using `%w` to preserve the underlying error and can be compared
// to determine if it's of a specific error Type using `Is` function. The `As` function can be used to determine if the
// error is in fact a gokit error and the `Unwrap function can be used to get to the original error.
//
// Examples: (in these examples kiterror is the backend/errors package. The standard library errors package is used as errors)
//
//    var customErr = fmt.Errorf("an error occured")
//
//    func TestErrors(t *testing.T) {
//    	err := ErrorTestHelper1()
//    	if !errors.Is(err, kiterror.ErrSystem) {
//    		t.Errorf("expected system error")
//    	}
//
//    	var ee kiterror.Error
//    	if !errors.As(err, &ee) {
//    		t.Errorf("expected system error type")
//    	}
//
//    	e := errors.Unwrap(err)
//    	if e, ok := e.(kiterror.Error); !ok || e.Err != customErr {
//    		t.Errorf("expected error to be system error")
//    	}
//    }
//
//    func ErrorTestHelper1() error {
//    	err := ErrorTestHelper2()
//    	return fmt.Errorf("wrapping the first error here: %w", err)
//    }
//
//    func ErrorTestHelper2() error {
//    	return kiterror.Error{
//    		Err:       customErr,
//    		Type:      kiterror.ErrSystem,
//    		Message:   "system error",
//    		Timestamp: time.Now(),
//    	}
//    }
//
// In addition when a service makes a grpc call to another service returning the error to the caller as kiterror.Error
// is automatically transformed using gokit's grpc server and client interceptors. For example this makes it so that an error
// of type errors.NotFound that is returned from one service can be examined by the caller as follows:
//
//    if ok := kiterror.Is(err, kiterror.NotFound); ok {
//         // the error is of type NotFound
//    }
//

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	// Unknown is the error Type when it is not clear if it's a user or system error. This should generally not be
	// used explicitly, and should be reserved as a fall-back case used by frameworks.
	Unknown Type = iota
	// ErrSystem is the error Type for system errors.
	ErrSystem
	// ErrUser is the error Type for user (caller of API) errors.
	ErrUser
	// ExpectationFailed indicates there was an expectation by the server not met by the client.
	ExpectationFailed
	// NotAuthorized is the error Type indicating that a requester does not have the appropriate role.
	NotAuthorized
	// NotFound is the error Type indicating an underlying resource could not be located.
	NotFound
	// RequestTimeout is the error Type indicating that the request timed out on the server side
	RequestTimeout
	// ServiceUnavailable is the error Type indicating that a service should be reachable but is not.
	ServiceUnavailable
	// Unauthenticated indicates an error related to identity not being verifiable.
	Unauthenticated
	// Unimplemented indicates a method is not implemented yet.
	Unimplemented
	// BadGateway indicates a dependency further downstream failed.
	BadGateway
	// Conflict indicates a request conflict with the current state of the server.
	Conflict
	// UnprocessableEntity is analogous to 422 (Unprocessable Entity) status code and indicates that the server understands the content type of the request entity,
	// and the syntax of the request entity is correct, but it was unable to process the contained instructions.
	UnprocessableEntity
	// UnsupportedMediaType indicates a request has unsupported media type
	UnsupportedMediaType
)

// String implements the fmt.Stringer interface and is useful in lieu of Error.Type not being Marshalable to JSON.
func (t Type) String() string {
	switch t {
	case Unknown:
		return "Unknown"
	case ErrSystem:
		return "ErrSystem"
	case ErrUser:
		return "ErrUser"
	case ExpectationFailed:
		return "ExpectationFailed"
	case NotAuthorized:
		return "NotAuthorized"
	case NotFound:
		return "NotFound"
	case RequestTimeout:
		return "RequestTimeout"
	case ServiceUnavailable:
		return "ServiceUnavailable"
	case Unauthenticated:
		return "Unauthenticated"
	case Unimplemented:
		return "Unimplemented"
	case BadGateway:
		return "BadGateway"
	case Conflict:
		return "Conflict"
	case UnprocessableEntity:
		return "UnprocessableEntity"
	case UnsupportedMediaType:
		return "UnsupportedMediaType"
	default:
		return "Undefined"
	}
}

// IsClientError returns true if the type of error represents client error
func (t Type) IsClientError() bool {
	switch t {
	case ErrUser, ExpectationFailed, Unauthenticated, NotAuthorized, NotFound:
		return true
	default:
		return false
	}
}

type (
	// Domain is the name of the system, or part of the system, where the error has originated.
	Domain string

	// Reason is a code that can be specified by a system generating the error
	Reason string

	// Type is a constant that indicates a Protocol level error, generally either `google.golang.org/grpc/codes` or
	// standard HTTP Status codes. See the relevant packages for additional details.
	Type int

	// Error is the object to use for internal Bread errors
	// It implements the standard library error interface
	Error struct {
		// Err is the actual error at the point of creation, the contents of which should be safe for logging only, and
		// should never be intended for consumption external to Bread.
		// Does not export as JSON to prevent the accidental sharing of internal information that might be in the stack.
		Err error `json:"-"`

		// Type specifies the the Type value agnostic to the eventual output protocol (ie. GRPC, HTTP, etc)
		// Does not export as JSON, as the value is specific to the context in which it is used (ie. GRPC, HTTP, etc)
		Type Type `json:"-"`

		// Message is a user friendly message to provide additional detail to the client if necessary. This should not be
		// substituted with or by Err, or Reason.
		Message string `json:"message"`

		// Timestamp is the time at the time where the original error was created.
		// Does not export as JSON since the default time format is not RFC3339.
		Timestamp time.Time `json:"-"`

		// Reason specifies the Reason value.
		Reason Reason `json:"reason"`

		// Domain specifies the Domain value.
		Domain Domain `json:"domain"`

		// Metadata is a map that contains additional information for context on the error.
		Metadata map[string]string `json:"metadata"`
	}
)

// Error is the implementation of the standard library error interface
func (e Error) Error() string {
	return e.Err.Error()
}

// Is corresponds to standard library errors.Is function
// and is used in a similar way
func (e Error) Is(err error) bool {
	var target Error
	if errors.As(err, &target) {
		return target.Type == e.Type
	}
	return false
}

// WithMessage is a fluent-style setter for setting the message of an Error. It is convenient to use it as follows: errors.NewUser(someErr).WithMessage("some message")
func (e Error) WithMessage(msg string) Error {
	e.Message = msg
	return e
}

// Is can be used to determine if anywhere in the chain of errors we have a certain type of error
func Is(err error, t Type) bool {
	var target Error
	if errors.As(err, &target) {
		return target.Type == t
	}

	return false
}

// Message gets the message of the error if error is of type Error otherwise return the default message
func Message(err error, def string) string {
	var target Error
	if errors.As(err, &target) {
		return target.Message
	}
	return def
}

// NewUser creates a User error from the given error
//
// Deprecated: replaced by NewUserError with better return type
func NewUser(err error) error {
	return New(ErrUser, err)
}

// NewSystem creates a System error from the given error
//
// Deprecated: replaced by NewSystemError with better return type
func NewSystem(err error) error {
	return New(ErrSystem, err)
}

// New creates an Error of the given type and underlying error,
// and sets Message to the Error() value of the error
//
// Deprecated: replaced by NewError with better return type
func New(t Type, err error) error {
	return Error{
		Err:       err,
		Type:      t,
		Message:   "",
		Timestamp: time.Now(),
	}
}

// NewUserError creates a User error from the given error
func NewUserError(err error) Error {
	return NewError(ErrUser, err)
}

// NewSystemError creates a System error from the given error
func NewSystemError(err error) Error {
	return NewError(ErrSystem, err)
}

// NewError is the entry to a builder type pattern, that creates an Error with the provided Type and sets Err to the provided error. The initial value of the message is an empty string (the zero value). To set the Message, call `WithMessage(string)`. Example: gokitErr := NewError(MyType, fmt.Errorf(...)).WithMessage("this is a custom message for my service error")
func NewError(t Type, err error) Error {
	return Error{
		Err:       err,
		Type:      t,
		Message:   "",
		Timestamp: time.Now(),
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (d Domain) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// MarshalJSON implements the json.Marshaler interface.
func (r Reason) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}
