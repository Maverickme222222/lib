// Package grpc has some grpc helpers
package grpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	errorKit "github.com/kappapay/backend/lib/golang/src/kappa/errors"

	"github.com/rs/zerolog"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// CodeRPCError is a custom code to support the ErrTypeRPCError Type value
	CodeRPCError codes.Code = 1000
	// CodeBadGateway is a custom code to indicate a failure further downstream to the dependency.
	CodeBadGateway codes.Code = 1001
	// CodeUnprocessableEntity is a custom code to indicates that the server understands the content type of the request entity,
	// and the syntax of the request entity is correct, but it was unable to process the contained instructions.
	CodeUnprocessableEntity codes.Code = 1002

	// DomainRPC is a special reserved Domain for handling conversions across GRPC boundaries.
	DomainRPC errorKit.Domain = "Kappa.RPC"

	// ErrTypeRPCError is a custom code specifically to report internal RPC translate
	ErrTypeRPCError errorKit.Type = 1000

	// ReasonUnknownErrorType indicates the server side error type was not supported
	ReasonUnknownErrorType grpcInterceptorReason = "Unknown_Error_Type"

	genericMethodInfoMessage = "error while calling GRPC method: %s"
)

type (
	grpcInterceptorReason errorKit.Reason
)

func (g grpcInterceptorReason) Reason() errorKit.Reason {
	return errorKit.Reason(g)
}

func (g grpcInterceptorReason) Type() errorKit.Type {
	return errorKit.Unknown
}

func (g grpcInterceptorReason) Domain() errorKit.Domain {
	return DomainRPC
}

func (g grpcInterceptorReason) Message() string {
	return ""
}

var (
	typeToCode = map[errorKit.Type]codes.Code{
		errorKit.ErrSystem:           codes.Internal,
		errorKit.ErrUser:             codes.InvalidArgument,
		errorKit.NotAuthorized:       codes.PermissionDenied,
		errorKit.NotFound:            codes.NotFound,
		errorKit.RequestTimeout:      codes.DeadlineExceeded,
		errorKit.Unauthenticated:     codes.Unauthenticated,
		errorKit.Unimplemented:       codes.Unimplemented,
		errorKit.ServiceUnavailable:  codes.Unavailable,
		errorKit.BadGateway:          CodeBadGateway,
		errorKit.Conflict:            codes.AlreadyExists,
		errorKit.Unknown:             codes.Unknown,
		errorKit.UnprocessableEntity: CodeUnprocessableEntity,
		ErrTypeRPCError:              CodeRPCError,
	}

	codeToType = map[codes.Code]errorKit.Type{
		codes.Internal:          errorKit.ErrSystem,
		codes.InvalidArgument:   errorKit.ErrUser,
		codes.NotFound:          errorKit.NotFound,
		codes.DeadlineExceeded:  errorKit.RequestTimeout,
		codes.PermissionDenied:  errorKit.NotAuthorized,
		codes.Unauthenticated:   errorKit.Unauthenticated,
		codes.Unimplemented:     errorKit.Unimplemented,
		codes.Unavailable:       errorKit.ServiceUnavailable,
		CodeBadGateway:          errorKit.BadGateway,
		codes.AlreadyExists:     errorKit.Conflict,
		codes.Unknown:           errorKit.Unknown,
		CodeUnprocessableEntity: errorKit.UnprocessableEntity,
		CodeRPCError:            ErrTypeRPCError,
	}
)

// MapToCode maps to gokit error to grpc code if it exists or else maps tp `codes.Unknown`
func MapToCode(logger zerolog.Logger, kitError errorKit.Error) codes.Code {
	code, typeExists := typeToCode[kitError.Type]
	if !typeExists {
		logger.Error().Msgf("received unsupported Error.Type %v", kitError.Type)
		code = codes.Unknown
	}

	return code
}

// MapCodeToType takes a grpc error code and returns the appropriate gokit error type for it.
func MapCodeToType(logger zerolog.Logger, code codes.Code) errorKit.Type {
	errorType, exists := codeToType[code]
	if !exists {
		logger.Error().Msgf("received unsupported grpc/codes.Code %v", code)
		return errorKit.ErrSystem
	}

	return errorType
}

// MapToStatus maps the kit error to grpc status
func MapToStatus(logger zerolog.Logger, kitError errorKit.Error) (*status.Status, error) {
	code := MapToCode(logger, kitError)
	kitStatus := status.New(code, kitError.Message)

	errorInfoDetail := &errdetails.ErrorInfo{
		Reason:   string(kitError.Reason),
		Domain:   string(kitError.Domain),
		Metadata: kitError.Metadata,
	}

	return kitStatus.WithDetails(errorInfoDetail)
}

// ErrorMarshallingClientInterceptor is a client-side interceptor that ALWAYS returns a `kappa/errors.Error` type for
// easier consumption by service wrappers.
func ErrorMarshallingClientInterceptor(logger zerolog.Logger) grpc.UnaryClientInterceptor {
	logger.Info().Msg("Adding ErrorMarshallingClientInterceptor")
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// rpcError is the native error returned over the wire
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err == nil {
			return nil
		}

		// check if the error can be converted to a compatible Status type
		serverStatus, isStatus := status.FromError(err)
		if !isStatus {
			logger.Err(err).Msg("could not convert server error to GRPC status type")
			return errorKit.NewLocal(fmt.Errorf(genericMethodInfoMessage, method), ReasonUnknownErrorType).
				WithMessage("Received error from server that was not status.Status.").
				WithMetaPairs()
		}

		if len(serverStatus.Details()) == 0 {
			logger.Err(err).Msg("error received was non-gokit GRPC status type")
			return errorKit.NewLocal(errors.New(serverStatus.Message()), ReasonUnknownErrorType).
				WithMessage("Unknown error returned from server").
				WithMetaPairs()
		}

		errType, codeExists := codeToType[serverStatus.Code()]
		if !codeExists {
			logger.Err(err).Msgf("received unsupported error code %d", serverStatus.Code())
			errType = errorKit.ErrSystem
		}

		clientErr := errorKit.Error{
			Err:       fmt.Errorf(genericMethodInfoMessage, method),
			Type:      errType,
			Message:   serverStatus.Message(),
			Timestamp: time.Now().UTC(),
		}

		switch detail := serverStatus.Details()[0].(type) {
		case *errdetails.ErrorInfo:
			clientErr.Reason = errorKit.Reason(detail.GetReason())
			clientErr.Domain = errorKit.Domain(detail.GetDomain())
			if detail.GetMetadata() != nil {
				clientErr.Metadata = detail.GetMetadata()
			} else {
				clientErr.Metadata = map[string]string{}
			}
		default:
			logger.Err(err).Msgf("could not handle GRPC errdetail of type %T", detail)
		}

		return clientErr
	}
}

// ErrorLoggingServerInterceptor logs all errors  to the provided logger.
func ErrorLoggingServerInterceptor(logger zerolog.Logger) grpc.UnaryServerInterceptor {
	logger.Info().Msg("Adding ErrorLoggingServerInterceptor")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, handlerErr := handler(ctx, req)
		if handlerErr != nil {
			var kitError errorKit.Error
			if errors.As(handlerErr, &kitError) {
				logger.Err(handlerErr).
					Str("method", info.FullMethod).
					Dict("error.fields", zerolog.Dict().
						Str("reason", string(kitError.Reason)).
						Str("domain", string(kitError.Domain)).
						Str("type", kitError.Type.String()).
						Str("message", kitError.Message).
						Str("metadata", fmt.Sprintf("%+v", kitError.Metadata)).
						Str("timestamp", kitError.Timestamp.Format(time.RFC3339)),
					).Msgf("returning error while responding to GRPC request")
			} else {
				logger.Err(handlerErr).Str("method", info.FullMethod).Msg("returning error while responding to GRPC request")
			}
		}

		return resp, handlerErr
	}
}

// ErrorUnmarshallingServerInterceptor is a server-side interceptor that attempts to always returns a Status with LocalizedMessage
// details for easier consumption by service clients.
func ErrorUnmarshallingServerInterceptor(logger zerolog.Logger) grpc.UnaryServerInterceptor {
	logger.Info().Msg("Adding ErrorUnmarshallingServerInterceptor")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, handlerErr := handler(ctx, req)
		if handlerErr == nil {
			return resp, nil
		}

		statusMsg := fmt.Sprintf(genericMethodInfoMessage, info.FullMethod)
		var kitError errorKit.Error
		if errors.As(handlerErr, &kitError) {
			kitStatus, err := MapToStatus(logger, kitError)
			if err != nil {
				logger.Err(err).Msgf("while adding the error details from: %+v", kitError)
				return resp, status.New(CodeRPCError, "could not add details from kit "+statusMsg).Err()
			}

			return resp, kitStatus.Err()
		}

		logger.Err(handlerErr).Msg("server is attempting to unmarshal a non-kit error")
		libStatus := status.New(codes.Unknown, "Unknown error type received in server")
		var libStatusErr error
		libStatus, libStatusErr = libStatus.WithDetails(&errdetails.ErrorInfo{
			Reason:   string(ReasonUnknownErrorType),
			Domain:   string(DomainRPC),
			Metadata: map[string]string{},
		})
		if libStatusErr != nil {
			logger.Err(libStatusErr).Msgf("while adding default error details for a stdlib error.")
			return resp, status.Newf(CodeRPCError, statusMsg).Err()
		}

		return resp, libStatus.Err()
	}
}
