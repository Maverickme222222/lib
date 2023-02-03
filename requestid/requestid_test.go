package requestid

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestFromContext(t *testing.T) {
	rawUUIDString := "12344321-4444-4444-4444-654321654321"
	tests := []struct {
		name         string
		getContext   func() context.Context
		expectedUUID uuid.UUID
	}{
		{
			name: "Not Found",
			getContext: func() context.Context {
				return context.Background()
			},
			expectedUUID: uuid.Nil,
		},
		{
			name: "Not UUID Type",
			getContext: func() context.Context {
				return context.WithValue(context.Background(), contextKey, rawUUIDString)
			},
			expectedUUID: uuid.Nil,
		},
		{
			name: "Success",
			getContext: func() context.Context {
				return context.WithValue(context.Background(), contextKey, uuid.MustParse(rawUUIDString))
			},
			expectedUUID: uuid.MustParse(rawUUIDString),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualUUID := FromContext(test.getContext())
			assert.Equal(t, test.expectedUUID, actualUUID)
		})
	}
}

func TestFromContextMetadata(t *testing.T) {
	rawUUIDString := "12344321-4444-4444-4444-654321654321"
	tests := []struct {
		name          string
		getContext    func() context.Context
		expectedError error
		expectedUUID  uuid.UUID
	}{
		{
			name: "Not Found",
			getContext: func() context.Context {
				return context.Background()
			},
			expectedError: ErrNotFound,
			expectedUUID:  uuid.UUID{},
		},
		{
			name: "Success",
			getContext: func() context.Context {
				md := metadata.New(map[string]string{XHeaderKey: rawUUIDString})
				return metadata.NewIncomingContext(context.Background(), md)
			},
			expectedUUID: uuid.MustParse(rawUUIDString),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualUUID, actualError := FromContextMetadata(test.getContext())

			assert.Equal(t, test.expectedError, actualError)
			assert.Equal(t, test.expectedUUID, actualUUID)
		})
	}
}

func TestFromHTTPHeader(t *testing.T) {
	tests := []struct {
		name          string
		header        func() http.Header
		expectedError error
		expectedUUID  uuid.UUID
	}{
		{
			name: "Empty Header",
			header: func() http.Header {
				return make(http.Header)
			},
			expectedError: ErrNotFound,
			expectedUUID:  uuid.Nil,
		},
		{
			name: "Unparsable",
			header: func() http.Header {
				val := make(http.Header)
				val.Set(XHeaderKey, "jjjj1111-4444-4444-4444-123456654321")
				return val
			},
			expectedError: errors.New("invalid UUID format"),
			expectedUUID:  uuid.UUID{},
		},
		{
			name: "Success",
			header: func() http.Header {
				val := make(http.Header)
				val.Set(XHeaderKey, "11223344-4444-4444-4444-aaaabbbbcccc")
				return val
			},
			expectedError: nil,
			expectedUUID:  uuid.MustParse("11223344-4444-4444-4444-aaaabbbbcccc"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualUUID, actualError := FromHTTPHeader(test.header())
			assert.Equal(t, test.expectedError, actualError)
			assert.Equal(t, test.expectedUUID, actualUUID)
		})
	}
}

func TestSetOnContext(t *testing.T) {
	tests := []struct {
		name            string
		inputUUID       uuid.UUID
		expectedContext func() context.Context
		expectedError   error
	}{
		{
			name:      "Zero Input",
			inputUUID: uuid.UUID{},
			expectedContext: func() context.Context {
				return context.Background()
			},
			expectedError: ErrEmptyValue,
		},
		{
			name:      "Success",
			inputUUID: uuid.MustParse("00001111-4444-4444-4444-123456789abc"),
			expectedContext: func() context.Context {
				return context.WithValue(context.Background(), contextKey, uuid.MustParse("00001111-4444-4444-4444-123456789abc"))
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualCtx, actualError := SetOnContext(context.Background(), test.inputUUID)
			assert.Equal(t, test.expectedContext(), actualCtx)
			assert.Equal(t, test.expectedError, actualError)
		})
	}
}

func TestSetOnContextMetadata(t *testing.T) {
	tests := []struct {
		name            string
		inputUUID       uuid.UUID
		expectedContext func() context.Context
		expectedError   error
	}{
		{
			name:      "Zero Input",
			inputUUID: uuid.UUID{},
			expectedContext: func() context.Context {
				return context.Background()
			},
			expectedError: ErrEmptyValue,
		},
		{
			name:      "Success",
			inputUUID: uuid.MustParse("00001111-4444-4444-4444-123456789abc"),
			expectedContext: func() context.Context {
				return metadata.AppendToOutgoingContext(context.Background(), XHeaderKey, "00001111-4444-4444-4444-123456789abc")
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualCtx, actualError := SetOnContextMetadata(context.Background(), test.inputUUID)

			assert.Equal(t, test.expectedError, actualError)
			assert.Equal(t, test.expectedContext(), actualCtx)
		})
	}
}

func TestSetOnHeader(t *testing.T) {
	tests := []struct {
		name          string
		inputUUID     uuid.UUID
		expectedError error
	}{
		{
			name:          "Zero Input",
			inputUUID:     uuid.UUID{},
			expectedError: ErrEmptyValue,
		},
		{
			name:          "Success",
			inputUUID:     uuid.MustParse("aaaabbbb-4444-4444-4444-111122223333"),
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualHeader := make(http.Header)
			actualError := SetOnHeader(actualHeader, test.inputUUID)
			assert.Equal(t, test.expectedError, actualError)

			actualUUID := actualHeader.Get(XHeaderKey)

			var expectedUUID string
			if test.expectedError != nil {
				expectedUUID = ""
			} else {
				expectedUUID = test.inputUUID.String()
			}
			assert.Equal(t, expectedUUID, actualUUID)
		})
	}
}

func TestHTTPHandler_ValueInHeader(t *testing.T) {
	expectedUUID := uuid.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(XHeaderKey, expectedUUID.String())
	resp := httptest.NewRecorder()

	logWriter := new(bytes.Buffer)
	logger := zerolog.New(logWriter)

	var doAssert http.HandlerFunc = func(resp http.ResponseWriter, req *http.Request) {
		require.NotEmpty(t, resp.Header().Get(XHeaderKey))

		assert.Equal(t, expectedUUID, FromContext(req.Context()))
		assert.Equal(t, expectedUUID.String(), resp.Header().Get(XHeaderKey))

		assert.Empty(t, logWriter.String())
	}

	HTTPHandler(doAssert, logger).ServeHTTP(resp, req)
}

func TestHTTPHandler_EmptyHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	resp := httptest.NewRecorder()

	logWriter := new(bytes.Buffer)
	logger := zerolog.New(logWriter)

	var doAssert http.HandlerFunc = func(resp http.ResponseWriter, req *http.Request) {
		contextID := FromContext(req.Context())
		require.NotEqual(t, uuid.Nil, contextID)

		headerID := resp.Header().Get(XHeaderKey)
		require.NotEmpty(t, headerID)

		assert.Equal(t, headerID, contextID.String())
		assert.Contains(t, logWriter.String(), "X-Request-ID was not found in header for request to /")
	}

	HTTPHandler(doAssert, logger).ServeHTTP(resp, req)
}

func TestGRPCContext(t *testing.T) {
	rawUUIDString := "12344321-4444-4444-4444-654321654321"

	tests := []struct {
		name            string
		getContext      func() context.Context
		expectedContext func() context.Context
	}{
		{
			name: "return as is",
			getContext: func() context.Context {
				ctx, _ := SetOnContext(context.Background(), uuid.MustParse(rawUUIDString))

				md := metadata.New(map[string]string{XHeaderKey: rawUUIDString})
				ctx = metadata.NewIncomingContext(ctx, md)
				return ctx
			},
			expectedContext: func() context.Context {
				ctx, _ := SetOnContext(context.Background(), uuid.MustParse(rawUUIDString))

				md := metadata.New(map[string]string{XHeaderKey: rawUUIDString})
				ctx = metadata.NewIncomingContext(ctx, md)
				return ctx
			},
		},
		{
			name: "add to value",
			getContext: func() context.Context {
				md := metadata.New(map[string]string{XHeaderKey: rawUUIDString})
				return metadata.NewIncomingContext(context.Background(), md)
			},
			expectedContext: func() context.Context {

				md := metadata.New(map[string]string{XHeaderKey: rawUUIDString})
				ctx := metadata.NewIncomingContext(context.Background(), md)
				ctx, _ = SetOnContext(ctx, uuid.MustParse(rawUUIDString))
				return ctx
			},
		},
		{
			name: "add to metadata",
			getContext: func() context.Context {
				ctx, _ := SetOnContext(context.Background(), uuid.MustParse(rawUUIDString))
				return ctx
			},
			expectedContext: func() context.Context {
				ctx, _ := SetOnContext(context.Background(), uuid.MustParse(rawUUIDString))
				ctx = metadata.AppendToOutgoingContext(ctx, XHeaderKey, rawUUIDString)
				return ctx
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualCtx := GRPCContext(test.getContext())
			assert.Equal(t, test.expectedContext(), actualCtx)
		})
	}
}
