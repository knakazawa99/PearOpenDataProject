// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/auth.go

// Package usecase is a generated GoMock package.
package usecase

import (
	entity "api/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// RequestEmail mocks base method.
func (m *MockAuth) RequestEmail(auth entity.Auth) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestEmail", auth)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestEmail indicates an expected call of RequestEmail.
func (mr *MockAuthMockRecorder) RequestEmail(auth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestEmail", reflect.TypeOf((*MockAuth)(nil).RequestEmail), auth)
}
