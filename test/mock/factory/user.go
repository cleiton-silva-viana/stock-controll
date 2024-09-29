// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/factory/user.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/factory/user.go -destination=test/mock/factory/user.go -package=factorymock
//

// Package factorymock is a generated GoMock package.
package factorymock

import (
	reflect "reflect"
	dto "stock-controll/internal/application/dto"
	entity "stock-controll/internal/domain/entity"

	gomock "go.uber.org/mock/gomock"
)

// MockIUserFactory is a mock of IUserFactory interface.
type MockIUserFactory struct {
	ctrl     *gomock.Controller
	recorder *MockIUserFactoryMockRecorder
}

// MockIUserFactoryMockRecorder is the mock recorder for MockIUserFactory.
type MockIUserFactoryMockRecorder struct {
	mock *MockIUserFactory
}

// NewMockIUserFactory creates a new mock instance.
func NewMockIUserFactory(ctrl *gomock.Controller) *MockIUserFactory {
	mock := &MockIUserFactory{ctrl: ctrl}
	mock.recorder = &MockIUserFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserFactory) EXPECT() *MockIUserFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserFactory) Create(userData dto.UserDTO) (*entity.User, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userData)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserFactoryMockRecorder) Create(userData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserFactory)(nil).Create), userData)
}
