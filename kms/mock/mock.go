// Code generated by MockGen. DO NOT EDIT.
// Source: kms.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CompareHashedString mocks base method.
func (m *MockClient) CompareHashedString(original, compare string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareHashedString", original, compare)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CompareHashedString indicates an expected call of CompareHashedString.
func (mr *MockClientMockRecorder) CompareHashedString(original, compare interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHashedString", reflect.TypeOf((*MockClient)(nil).CompareHashedString), original, compare)
}

// Decrypt mocks base method.
func (m *MockClient) Decrypt(ciphertext string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", ciphertext)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockClientMockRecorder) Decrypt(ciphertext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockClient)(nil).Decrypt), ciphertext)
}

// Encrypt mocks base method.
func (m *MockClient) Encrypt(plaintext string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", plaintext)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockClientMockRecorder) Encrypt(plaintext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockClient)(nil).Encrypt), plaintext)
}

// GetEncryptedHashedData mocks base method.
func (m *MockClient) GetEncryptedHashedData(data string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEncryptedHashedData", data)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEncryptedHashedData indicates an expected call of GetEncryptedHashedData.
func (mr *MockClientMockRecorder) GetEncryptedHashedData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEncryptedHashedData", reflect.TypeOf((*MockClient)(nil).GetEncryptedHashedData), data)
}

// GetHashedData mocks base method.
func (m *MockClient) GetHashedData(data string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHashedData", data)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHashedData indicates an expected call of GetHashedData.
func (mr *MockClientMockRecorder) GetHashedData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashedData", reflect.TypeOf((*MockClient)(nil).GetHashedData), data)
}
