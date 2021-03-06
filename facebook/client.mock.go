// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/nluangaroonc/go/src/github.com/nomkhonwaan/example-facebook-client/facebook/client.go

package facebook

import (
	http "net/http"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockClient) EXPECT() *MockClientMockRecorder {
	return _m.recorder
}

// Authenticate mocks base method
func (_m *MockClient) Authenticate(appID string, appSecret string, redirectURI string, onLoggedInRedirectURI string) http.HandlerFunc {
	ret := _m.ctrl.Call(_m, "Authenticate", appID, appSecret, redirectURI, onLoggedInRedirectURI)
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// Authenticate indicates an expected call of Authenticate
func (_mr *MockClientMockRecorder) Authenticate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Authenticate", arg0, arg1, arg2, arg3)
}

// Me mocks base method
func (_m *MockClient) Me() (map[string]interface{}, error) {
	ret := _m.ctrl.Call(_m, "Me")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Me indicates an expected call of Me
func (_mr *MockClientMockRecorder) Me() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Me")
}
