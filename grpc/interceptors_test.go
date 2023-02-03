package grpc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	grpcTest "github.com/grpc-ecosystem/go-grpc-middleware/testing"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	grpcLib "google.golang.org/grpc"

	testproto "github.com/grpc-ecosystem/go-grpc-middleware/testing/testproto"
	errorKit "github.com/kappapay/backend/lib/golang/src/kappa/errors"
)

const (
	errorCodeReturnedKit uint32 = iota
	errorCodeReturnedLib
	errorCodeWrappedLib
	errorCodeServer

	domainGoKitTest errorKit.Domain = "GoKitTestDomain"
	reasonGoKitTest errorKit.Reason = "GoKit_Test_Reason"
)

var (
	testKitUserError = errorKit.Error{
		Err:     errors.New("this is an internal error"),
		Type:    errorKit.ErrUser,
		Message: "front-end facing message for the user",
		Domain:  domainGoKitTest,
		Reason:  reasonGoKitTest,
		Metadata: map[string]string{
			"boo":  "far",
			"this": "that",
		},
	}

	testKitServerError = errorKit.Error{
		Err:     errors.New("this is an internal error"),
		Type:    errorKit.ErrSystem,
		Message: "something failed",
		Domain:  domainGoKitTest,
		Reason:  reasonGoKitTest,
	}

	testLibError = errorKit.Error{
		Err:     errors.New("this is a standard lib error"),
		Type:    errorKit.Unknown,
		Message: "Error occurred while calling remote system",
	}
)

type assertion func()

type expectation func(interceptorsTestSuite) assertion

type testCase struct {
	name          string
	expectRequest expectation
}

func (s interceptorsTestSuite) sendErrorRequest(errorReturnCode uint32) (time.Time, error) {
	ctx := context.Background()
	timestamp := time.Now().UTC()

	resp, err := s.Client.PingError(ctx, &testproto.PingRequest{
		Value:             timestamp.Format(time.RFC3339),
		ErrorCodeReturned: errorReturnCode,
	})
	s.Require().Nil(resp)
	_, isKitError := err.(errorKit.Error)
	s.Require().True(isKitError)

	return timestamp, err
}

func (s interceptorsTestSuite) assertKitErrorEqual(expected, actual errorKit.Error) bool {
	return s.Equal(expected.Err, actual.Err) &&
		s.Equal(expected.Type, actual.Type) &&
		s.Equal(expected.Message, actual.Message) &&
		s.Equal(expected.Timestamp.Format(time.RFC3339), actual.Timestamp.Format(time.RFC3339)) &&
		s.Equal(expected.Domain, actual.Domain) &&
		s.Equal(expected.Reason, actual.Reason) &&
		s.Equal(expected.Metadata, actual.Metadata)
}

func expectErrorRequest(errorReturnCode uint32) expectation {
	return func(s interceptorsTestSuite) assertion {
		actualTimestamp, actualError := s.sendErrorRequest(errorReturnCode)
		s.Require().Error(actualError)

		var actual errorKit.Error
		s.Require().True(errors.As(actualError, &actual))

		expectedTimestamp, parseErr := time.Parse(time.RFC3339, actualTimestamp.Format(time.RFC3339))
		if parseErr != nil {
			panic(parseErr)
		}

		switch errorReturnCode {
		case errorCodeReturnedKit:
			expected := testKitUserError
			expected.Err = fmt.Errorf(genericMethodInfoMessage, "/mwitkow.testproto.TestService/PingError")
			expected.Timestamp = expectedTimestamp
			return func() {
				s.assertKitErrorEqual(expected, actual)
				s.Contains(s.logWriter.String(), `"message":"returning error while responding to GRPC request`)
				s.Contains(s.logWriter.String(), `"error.fields"`)
			}
		case errorCodeReturnedLib:
			expected := errorKit.Error{
				Err:       fmt.Errorf(genericMethodInfoMessage, "/mwitkow.testproto.TestService/PingError"),
				Type:      errorKit.Unknown,
				Message:   "Unknown error type received in server",
				Timestamp: expectedTimestamp,
				Reason:    errorKit.Reason(ReasonUnknownErrorType),
				Domain:    DomainRPC,
				Metadata:  map[string]string{},
			}
			return func() {
				s.assertKitErrorEqual(expected, actual)
				s.Contains(s.logWriter.String(), `"message":"server is attempting to unmarshal a non-kit error`)
				s.Contains(s.logWriter.String(), `"message":"returning error while responding to GRPC request`)
			}
		case errorCodeWrappedLib:
			expected := testKitUserError
			expected.Err = fmt.Errorf(genericMethodInfoMessage, "/mwitkow.testproto.TestService/PingError")
			expected.Timestamp = expectedTimestamp
			return func() {
				s.assertKitErrorEqual(expected, actual)
				s.Contains(s.logWriter.String(), `"message":"returning error while responding to GRPC request`)
				s.Contains(s.logWriter.String(), `"error.fields"`)
			}
		default:
			panic(fmt.Sprintf("unexpected errorCode in assertion %d", errorReturnCode))
		}
	}
}

type interceptorsTestSuite struct {
	grpcTest.InterceptorTestSuite

	logWriter      *bytes.Buffer
	logger         zerolog.Logger
	serverDisabled bool

	tests []testCase
}

func newTestSuite() *interceptorsTestSuite {
	s := new(interceptorsTestSuite)

	s.logWriter = new(bytes.Buffer)
	s.logger = zerolog.New(s.logWriter).Level(zerolog.ErrorLevel)

	return s
}

func TestDefault(t *testing.T) {
	s := newTestSuite()

	s.ClientOpts = append([]grpcLib.DialOption{}, ClientInterceptorsWithTracing(s.logger))
	s.ServerOpts = append([]grpcLib.ServerOption{}, ServerInterceptorsWithTracing(s.logger))
	// required to override the provided Service
	s.TestService = testService{}

	s.tests = []testCase{
		{
			name:          "ServerReturnsKitError",
			expectRequest: expectErrorRequest(errorCodeReturnedKit),
		},
		{
			name:          "ServerReturnsLibError",
			expectRequest: expectErrorRequest(errorCodeReturnedLib),
		},
		{
			name:          "ServerReturnsWrappedError",
			expectRequest: expectErrorRequest(errorCodeWrappedLib),
		},
	}

	suite.Run(t, s)
}

// TestClientOnlyInterceptors tests the scenario when the Remote server does not utilize the default Interceptors.
func TestClientOnly(t *testing.T) {
	s := newTestSuite()
	s.serverDisabled = true

	s.ClientOpts = append([]grpcLib.DialOption{}, ClientInterceptorsWithTracing(s.logger))
	s.ServerOpts = []grpcLib.ServerOption{}
	s.TestService = testService{} // required to override the provided Service

	errAssertion := func(expected errorKit.Error, actualError error) func() {
		return func() {
			expectedKit := errorKit.Error{
				Err:       expected.Err,
				Type:      errorKit.Unknown,
				Message:   "Unknown error returned from server",
				Timestamp: expected.Timestamp,
				Domain:    DomainRPC,
				Reason:    errorKit.Reason(ReasonUnknownErrorType),
				Metadata:  map[string]string{},
			}

			var actual errorKit.Error
			s.Require().True(errors.As(actualError, &actual))
			s.assertKitErrorEqual(expectedKit, actual)

			s.NotEmpty(s.logWriter.String())
		}
	}

	s.tests = []testCase{
		{
			name: "ServerReturnsKitError",
			expectRequest: func(s interceptorsTestSuite) assertion {
				timestamp, actual := s.sendErrorRequest(errorCodeReturnedKit)
				expected := testKitUserError
				expected.Timestamp = timestamp
				return errAssertion(expected, actual)
			},
		},
		{
			name: "ServerReturnsLibError",
			expectRequest: func(s interceptorsTestSuite) assertion {
				timestamp, actual := s.sendErrorRequest(errorCodeReturnedLib)
				expected := testLibError
				expected.Timestamp = timestamp
				return errAssertion(expected, actual)
			},
		},
	}

	suite.Run(t, s)
}

func (s interceptorsTestSuite) TestInterceptors() {
	for _, test := range s.tests {
		s.Run(test.name, func() {
			assertRequest := test.expectRequest(s)
			assertRequest()
			s.logWriter.Reset()
		})
	}
}

type testService struct{}

func (t testService) PingEmpty(ctx context.Context, _ *testproto.Empty) (*testproto.PingResponse, error) {
	return &testproto.PingResponse{
		Value: "testing",
	}, nil
}

func (t testService) Ping(ctx context.Context, _ *testproto.PingRequest) (*testproto.PingResponse, error) {
	return &testproto.PingResponse{
		Value:   "testing",
		Counter: 0,
	}, nil
}

func (t testService) PingError(_ context.Context, req *testproto.PingRequest) (*testproto.Empty, error) {
	switch req.ErrorCodeReturned {
	case errorCodeReturnedKit:
		err := testKitUserError
		timestamp, parseErr := time.Parse(time.RFC3339, req.Value)
		if parseErr != nil {
			panic(parseErr)
		}
		err.Timestamp = timestamp
		return nil, err
	case errorCodeReturnedLib:
		return nil, testLibError.Err
	case errorCodeWrappedLib:
		err := testKitUserError
		timestamp, parseErr := time.Parse(time.RFC3339, req.Value)
		if parseErr != nil {
			panic(parseErr)
		}
		err.Timestamp = timestamp
		return nil, fmt.Errorf("this is the outer error: %w", err)
	case errorCodeServer:
		err := testKitServerError
		err.Timestamp = time.Now()
		return nil, err
	default:
		panic("Unsupported ErrorCodeReturned")
	}
}

func (t testService) PingList(_ *testproto.PingRequest, _ testproto.TestService_PingListServer) error {
	panic("implement me")
}

func (t testService) PingStream(_ testproto.TestService_PingStreamServer) error {
	panic("implement me")
}
