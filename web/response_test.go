package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRespond(t *testing.T) {

	tests := map[string]struct {
		name               string
		responseStatusCode int
		responseData       interface{}
		expectedHeaders    func() http.Header
		expectedOutput     []byte
		expectError        bool
		success            bool
		returnedError      error
	}{
		"No Content returns without error": {
			responseStatusCode: http.StatusNoContent,
			expectedHeaders: func() http.Header {
				return make(http.Header)
			},
			expectError: false,
		},
		"Responding with invalid data returns an error": {
			responseStatusCode: http.StatusOK,
			responseData:       make(chan string),
			expectError:        true,
			success:            false,
		},
		"Well-formed response completes successfully": {
			responseStatusCode: http.StatusCreated,
			responseData: struct {
				Foo string `json:"bar"`
			}{
				Foo: "BAZ",
			},
			expectedHeaders: func() http.Header {
				h := make(http.Header)
				h.Set(ContentType, ContentTypeApplicationJSON)
				return h
			},
			expectError:    false,
			expectedOutput: []byte(`{"success":true,"errors":[],"data":{"bar":"BAZ"}}`),
			success:        true,
			returnedError:  nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(tt *testing.T) {
			writer := httptest.NewRecorder()
			actualError := Respond(writer, test.responseData, test.responseStatusCode, test.success, test.returnedError)

			if test.expectError {
				require.Error(tt, actualError)
				return
			}

			actualResponse := writer.Result()
			require.NoError(tt, actualError)
			assert.Equal(tt, test.expectedHeaders(), actualResponse.Header)
			assert.Equal(tt, test.responseStatusCode, actualResponse.StatusCode)

			actualBytes := writer.Body.Bytes()
			assert.Equal(tt, test.expectedOutput, actualBytes)
		})
	}
}
