// Package errors demonstrates using the LocalError interface to implement a fluent style of Error creation.
// nolint
package errors

import errorKit "github.com/kappapay/backend/lib/golang/src/kappa/errors"

const (
	// Usually just the Service your coding for, but there could be multiple domains that would map to sub-systems, or
	// phases of an process / workflow.
	MyServiceDomain = "MyService"

	EntityNotFound               MyServiceReason = "Service_Entity_NotFound"
	RequestValidationFailed      MyServiceReason = "Request_Validation_Failed"
	RequestExpired               MyServiceReason = "Request_Expired"
	ServiceDependencyUnavailable MyServiceReason = "Service_Dependency_Unavailable"
)

type (
	// MyServiceReason is a custom type to allow for method receivers and is the initial hook into the LocalError interface.
	MyServiceReason errorKit.Reason
)

func (m MyServiceReason) Reason() errorKit.Reason {
	return errorKit.Reason(m)
}

func (m MyServiceReason) Type() errorKit.Type {
	switch m {
	case EntityNotFound:
		return errorKit.NotFound
	case RequestValidationFailed:
		return errorKit.ErrUser
	case RequestExpired:
		return errorKit.ErrUser
	case ServiceDependencyUnavailable:
		return errorKit.ServiceUnavailable
	default:
		return errorKit.Unimplemented
	}
}

func (m MyServiceReason) Message() string {
	switch m {
	case EntityNotFound:
		return "Entity not found with provided ID."
	case RequestExpired:
		return "The request previously started, has now expired. Please resubmit."
	case RequestValidationFailed:
		return "The data provided in the request was not properly formatted. See Metadata for details."
	default:
		return "Internal Service Error"
	}
}

// Domain would typically return just a single domain, but this could also be backed by a switch statement.
func (m MyServiceReason) Domain() errorKit.Domain {
	return MyServiceDomain
}

// New is a local convenience method attached to the Reason type, so the fluent chain can begin here.
// Also consider a `Newf(string, ...interface{})` method to mirror `fmt.Errorf` functionality.
func (m MyServiceReason) New(err error) errorKit.Error {
	return errorKit.NewLocal(err, m)
}
