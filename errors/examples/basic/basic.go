// This is a demonstration of just a basic example of the expanded error type.
// nolint
package main

import (
	liberr "errors"
	"fmt"

	kitErrors "github.com/kappapay/backend/lib/golang/src/kappa/errors"
)

const (
	FirstReason ExampleReason = "FirstLayerReason"

	ExamplesDomain kitErrors.Domain = "BasicExample"
)

type (
	ExampleReason kitErrors.Reason
)

func (e ExampleReason) Reason() kitErrors.Reason {
	return kitErrors.Reason(e)
}

func (e ExampleReason) Type() kitErrors.Type {
	return kitErrors.ErrSystem // this wouldn't always return the same type, but we just have one reason
}

func (e ExampleReason) Domain() kitErrors.Domain {
	return ExamplesDomain
}

func (e ExampleReason) Message() string {
	return ""
}

func main() {
	{
		err := FmtWrappedKitError()

		// Output:
		// FmtWrappedKitError().Error(): [this is the outermost error wrapping the kit Error: this is the kitErrors.Err field wrapping cause: this is the root cause] and has Type: [*fmt.wrapError]
		//
		// Here `*fmt.wrapError` is the internal Go lib type for wrapping  an error using the `fmt.Errorf(%w)` pattern.
		fmt.Printf("FmtWrappedKitError().Error(): [%s] and has Type: [%T]\n", err, err)

		// Output:
		// errors.Unwrap(FmtWrappedKitError()).Error(): [this is the kitErrors.Err field wrapping cause: this is the root cause] and has Type: [errors.Error]
		//
		// This shows how the lib `Unwrap` function works, and the type it returns.
		fmt.Printf("errors.Unwrap(FmtWrappedKitError()).Error(): [%s] and has Type: [%T]\n", liberr.Unwrap(err), liberr.Unwrap(err))

		// Output:
		// Fields for errors.Unwrap(FmtWrappedKitError()): [this is the kitErrors.Err field wrapping cause: this is the root cause]
		//
		// This shows that the unwrapped type can be manipulated as a Gokit Error if desired.
		unwrappedKitError := liberr.Unwrap(err).(kitErrors.Error)
		fmt.Printf("Fields for errors.Unwrap(FmtWrappedKitError()): [%+v]", unwrappedKitError)
	}
}

// FmtWrappedKitError shows the various wrapping of errors and how it effects the final `errors.Error()` call.
func FmtWrappedKitError() error {
	cause := liberr.New("this is the root cause")

	kitError := kitErrors.NewLocal(fmt.Errorf("this is the kitErrors.Err field wrapping cause: %w", cause), FirstReason).
		WithMessage("A user-friendly-ish type message")

	wrappedError := fmt.Errorf("this is the outermost error wrapping the kit Error: %w", kitError)

	return wrappedError
}
