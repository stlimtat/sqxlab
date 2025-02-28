// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -destination=mock.go -package=urls -source=interface.go
//

// Package urls is a generated GoMock package.
package urls

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIUrlDiscovery is a mock of IUrlDiscovery interface.
type MockIUrlDiscovery struct {
	ctrl     *gomock.Controller
	recorder *MockIUrlDiscoveryMockRecorder
	isgomock struct{}
}

// MockIUrlDiscoveryMockRecorder is the mock recorder for MockIUrlDiscovery.
type MockIUrlDiscoveryMockRecorder struct {
	mock *MockIUrlDiscovery
}

// NewMockIUrlDiscovery creates a new mock instance.
func NewMockIUrlDiscovery(ctrl *gomock.Controller) *MockIUrlDiscovery {
	mock := &MockIUrlDiscovery{ctrl: ctrl}
	mock.recorder = &MockIUrlDiscoveryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUrlDiscovery) EXPECT() *MockIUrlDiscoveryMockRecorder {
	return m.recorder
}

// Discover mocks base method.
func (m *MockIUrlDiscovery) Discover(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discover", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Discover indicates an expected call of Discover.
func (mr *MockIUrlDiscoveryMockRecorder) Discover(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discover", reflect.TypeOf((*MockIUrlDiscovery)(nil).Discover), ctx)
}
