package grpc

import (
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/rs/zerolog"
	"go.elastic.co/apm/module/apmgrpc/v2"
	"google.golang.org/grpc"
)

// ServerInterceptorsWithTracing packages the standard Server interceptors into a chained middleware for ease-of-use.
func ServerInterceptorsWithTracing(logger zerolog.Logger, additional ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	logger.Info().Msg("Initializing default ServerInterceptors")
	interceptors := make([]grpc.UnaryServerInterceptor, 0)
	interceptors = append(interceptors, logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger)))
	interceptors = append(interceptors, apmgrpc.NewUnaryServerInterceptor(), recovery.UnaryServerInterceptor())

	// While these are called in their index order (ErrUnmarshal, ErrLog, others...), it's also
	// important to note where in the interceptor the handler itself is invoked.
	// For ErrorUnmarshal and ErrorLogging, the logic is called after the handler is invoked, because the logic in those
	// interceptors depends on the response.
	// So for the purposes of serialized action, the ordering is actually:  ErrorLogging, ErrorUnmarshalling.

	interceptors = append(interceptors, ErrorUnmarshallingServerInterceptor(logger), ErrorLoggingServerInterceptor(logger))

	interceptors = append(interceptors, additional...)
	return grpc.ChainUnaryInterceptor(interceptors...)
}

// ClientInterceptorsWithTracing packages the standard Client interceptors into a chained middleware for ease-of-use.
func ClientInterceptorsWithTracing(logger zerolog.Logger, additional ...grpc.UnaryClientInterceptor) grpc.DialOption {
	logger.Info().Msg("Initializing default ClientInterceptors")
	interceptors := make([]grpc.UnaryClientInterceptor, 0)

	interceptors = append(interceptors, logging.UnaryClientInterceptor(grpczerolog.InterceptorLogger(logger)))

	interceptors = append(interceptors, apmgrpc.NewUnaryClientInterceptor())

	interceptors = append(interceptors, ErrorMarshallingClientInterceptor(logger))
	interceptors = append(interceptors, additional...)

	return grpc.WithChainUnaryInterceptor(interceptors...)
}
