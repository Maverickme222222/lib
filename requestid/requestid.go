// Package requestid handles the marshalling of an existing RequestID to and from a context.Context.
package requestid

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/metadata"
)

type contextKeyType struct{}

var (
	// ErrNotFound indicates no request id was found in the Context or Header
	ErrNotFound = errors.New("request id not found")
	// ErrEmptyValue = errors.New("request id is empty") indicates the uuid provided is a Zero value type.
	ErrEmptyValue = errors.New("request id is empty")

	// contextKey is used by the context value
	contextKey = contextKeyType{}
)

const (
	// XHeaderKey is used to correlate all the logs of a request that may  end up at various services/modules.
	// end up at various services/modules. All logs must print the request ID
	//This key is used for BOTH GRPC and HTTP, as Istio relies on this key as well in the metadata to propagate across
	// service calls regardless of protocol.
	XHeaderKey = "X-Request-ID"
)

// FromContext attempts to retrieve the existing request id as a valid and parsed UUID from the context values.
// If the value does not exist in the context, or for some reason is not of type uuid.UUID, the uuid.Nil value will be
// returned.
func FromContext(ctx context.Context) uuid.UUID {
	val := ctx.Value(contextKey)

	if val == nil {
		return uuid.Nil
	}

	if id, ok := val.(uuid.UUID); ok {
		return id
	}

	return uuid.Nil
}

// FromContextMetadata attempts to retrieve the existing request id as a valid and parsed UUID from the context metadata.
// An error may be returned if the id is not in the context metadata, or if it cannot be parsed.
func FromContextMetadata(ctx context.Context) (uuid.UUID, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		vals := md.Get(XHeaderKey)
		if len(vals) != 0 {
			// take first value
			val := vals[0]

			id, err := uuid.Parse(val)
			if err != nil {
				return uuid.Nil, err
			}

			return id, nil
		}
	}

	return uuid.UUID{}, ErrNotFound
}

// FromHTTPHeader attempts to retrieve the request id from the Header and parse it into a valid UUID.
// An error may be returned if the id is not in the Header, or if it cannot be parsed.
func FromHTTPHeader(header http.Header) (uuid.UUID, error) {
	val := header.Get(XHeaderKey)
	if val == "" {
		return uuid.Nil, ErrNotFound
	}

	id, err := uuid.Parse(val)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// SetOnContext attempts to set the provided uuid on the provided context, returning a new populated Context.
// An error may be returned if the uuid is "zero".
func SetOnContext(ctx context.Context, id uuid.UUID) (context.Context, error) {
	if id == uuid.Nil {
		return ctx, ErrEmptyValue
	}
	return context.WithValue(ctx, contextKey, id), nil
}

// SetOnContextMetadata attempts to set the provided uuid on the provided context metadata, returning a new populated Context.
// An error may be returned if the uuid is "zero".
func SetOnContextMetadata(ctx context.Context, id uuid.UUID) (context.Context, error) {
	if id == uuid.Nil {
		return ctx, ErrEmptyValue
	}

	return metadata.AppendToOutgoingContext(ctx, XHeaderKey, id.String()), nil
}

// SetOnHeader attempts to set the provided uuid on the provided Header. This is generally used to ensure
// the request id is set on the response Header.
// An error may be returned if the uuid is "zero".
func SetOnHeader(header http.Header, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrEmptyValue
	}
	header.Set(XHeaderKey, id.String())
	return nil
}

// HTTPHandler checks the http.Request header for the request id. If it exists, sets it in the context, and sets it back
// in the Response Header.
func HTTPHandler(next http.Handler, logger zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := FromHTTPHeader(r.Header)
		if err != nil {
			// Log this as an error, but provide a UUID so at least we have something till this is fixed.
			id = uuid.New()
			if err == ErrNotFound {
				logger.Err(err).Msgf("%s was not found in header for request to %s", XHeaderKey, r.RequestURI)
			} else {
				logger.Err(err).Msgf("%s could not be retrieved from the header for request to %s", XHeaderKey, r.RequestURI)
			}
		}

		if err := SetOnHeader(w.Header(), id); err != nil {
			logger.Err(err).Msgf("could not set %s header on response to request %s", XHeaderKey, r.RequestURI)
		}

		nextCtx, err := SetOnContext(r.Context(), id)
		if err != nil {
			logger.Err(err).Msgf("could not set request id on request context for request to %s", r.RequestURI)
		}

		r = r.WithContext(nextCtx)

		next.ServeHTTP(w, r)
	})
}

// GRPCContext checks the context values and metadata for the request id.
// If there is an existing request id, the context is returned as is.
// Otherwise, a new request id is set as a string in both the context values and metadata.
//
// Deprecated: This should be covered by Client and Server GRPC interceptors and HTTP middleware. More importantly, this
// can cause a massive bug when coupled with some versions of requestid.FromContextMetadata due to how values are added
// and retrieved from the Metadata.
func GRPCContext(ctx context.Context) context.Context {
	contextRequestID := FromContext(ctx)
	metadataRequestID, _ := FromContextMetadata(ctx)

	if contextRequestID == uuid.Nil && metadataRequestID == uuid.Nil {
		requestID := uuid.New()

		ctx = context.WithValue(ctx, contextKey, requestID)
		ctx = metadata.AppendToOutgoingContext(ctx, XHeaderKey, requestID.String())

		return ctx
	}

	if contextRequestID == uuid.Nil {
		return context.WithValue(ctx, contextKey, metadataRequestID)
	}

	if metadataRequestID == uuid.Nil {
		return metadata.AppendToOutgoingContext(ctx, XHeaderKey, contextRequestID.String())
	}

	return ctx
}
