// Package customapmgrpc adds makes sure the context is prop the downstream services
package customapmgrpc

import (
	"context"
	"strings"

	"go.elastic.co/apm/module/apmhttp/v2"
	"go.elastic.co/apm/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	elasticTraceparentHeader = strings.ToLower(apmhttp.ElasticTraceparentHeader)
	w3cTraceparentHeader     = strings.ToLower(apmhttp.W3CTraceparentHeader)
	tracestateHeader         = strings.ToLower(apmhttp.TracestateHeader)
)

// NewUnaryClientInterceptor returns a grpc.UnaryClientInterceptor that
// traces gRPC requests with the given options.
//
// The interceptor will trace spans with the "grpc" type for each request
// made, for any client method presented with a context containing a sampled
// apm.Transaction.
func NewUnaryClientInterceptor(spanContext apm.DestinationServiceSpanContext, o ...ClientOption) grpc.UnaryClientInterceptor {
	opts := clientOptions{}
	for _, o := range o {
		o(&opts)
	}
	return func(
		ctx context.Context,
		method string,
		req, resp interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		span, ctx := startSpan(ctx, method, spanContext)
		if span != nil {
			defer span.End()
		}
		return invoker(ctx, method, req, resp, cc, opts...)
	}
}

func startSpan(ctx context.Context, name string, spanContext apm.DestinationServiceSpanContext) (*apm.Span, context.Context) {
	tx := apm.TransactionFromContext(ctx)
	if tx == nil {
		return nil, ctx
	}

	traceContext := tx.TraceContext()
	propagateLegacyHeader := tx.ShouldPropagateLegacyHeader()
	if !traceContext.Options.Recorded() {
		return nil, outgoingContextWithTraceContext(ctx, traceContext, propagateLegacyHeader)
	}
	span := tx.StartSpan(name, "external.grpc", apm.SpanFromContext(ctx))
	span.Subtype = "grpc"
	span.Context.SetDestinationService(spanContext)

	if !span.Dropped() {
		traceContext = span.TraceContext()
		ctx = apm.ContextWithSpan(ctx, span)
	}
	return span, outgoingContextWithTraceContext(ctx, traceContext, propagateLegacyHeader)
}

func outgoingContextWithTraceContext(
	ctx context.Context,
	traceContext apm.TraceContext,
	propagateLegacyHeader bool,
) context.Context {
	traceparentValue := apmhttp.FormatTraceparentHeader(traceContext)
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.Pairs(w3cTraceparentHeader, traceparentValue)
	} else {
		md = md.Copy()
		md.Set(w3cTraceparentHeader, traceparentValue)
	}
	if propagateLegacyHeader {
		md.Set(elasticTraceparentHeader, traceparentValue)
	}
	if tracestate := traceContext.State.String(); tracestate != "" {
		md.Set(tracestateHeader, tracestate)
	}
	return metadata.NewOutgoingContext(ctx, md)
}

type clientOptions struct {
	tracer *apm.Tracer //nolint
}

// ClientOption sets options for client-side tracing.
type ClientOption func(*clientOptions)
