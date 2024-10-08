// Code generated by MockGen. DO NOT EDIT.
// Source: internal/infrastructure/unit_work/uow.go
//
// Generated by this command:
//
//	mockgen -source=internal/infrastructure/unit_work/uow.go -destination=test/mock/uow/mock_uow.go -package=uowmock
//

// Package uowmock is a generated GoMock package.
package uowmock

import (
	reflect "reflect"
	persistence "stock-controll/internal/infrastructure/persistence"

	gomock "go.uber.org/mock/gomock"
)

// MockIUnitOfWork is a mock of IUnitOfWork interface.
type MockIUnitOfWork struct {
	ctrl     *gomock.Controller
	recorder *MockIUnitOfWorkMockRecorder
}

// MockIUnitOfWorkMockRecorder is the mock recorder for MockIUnitOfWork.
type MockIUnitOfWorkMockRecorder struct {
	mock *MockIUnitOfWork
}

// NewMockIUnitOfWork creates a new mock instance.
func NewMockIUnitOfWork(ctrl *gomock.Controller) *MockIUnitOfWork {
	mock := &MockIUnitOfWork{ctrl: ctrl}
	mock.recorder = &MockIUnitOfWorkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUnitOfWork) EXPECT() *MockIUnitOfWorkMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockIUnitOfWork) Begin() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(error)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockIUnitOfWorkMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockIUnitOfWork)(nil).Begin))
}

// Commit mocks base method.
func (m *MockIUnitOfWork) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockIUnitOfWorkMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockIUnitOfWork)(nil).Commit))
}

// ContactRepository mocks base method.
func (m *MockIUnitOfWork) ContactRepository() persistence.ISQLContact {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContactRepository")
	ret0, _ := ret[0].(persistence.ISQLContact)
	return ret0
}

// ContactRepository indicates an expected call of ContactRepository.
func (mr *MockIUnitOfWorkMockRecorder) ContactRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContactRepository", reflect.TypeOf((*MockIUnitOfWork)(nil).ContactRepository))
}

// CredentialRepository mocks base method.
func (m *MockIUnitOfWork) CredentialRepository() persistence.ISQLCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialRepository")
	ret0, _ := ret[0].(persistence.ISQLCredential)
	return ret0
}

// CredentialRepository indicates an expected call of CredentialRepository.
func (mr *MockIUnitOfWorkMockRecorder) CredentialRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialRepository", reflect.TypeOf((*MockIUnitOfWork)(nil).CredentialRepository))
}

// Rollback mocks base method.
func (m *MockIUnitOfWork) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockIUnitOfWorkMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockIUnitOfWork)(nil).Rollback))
}

// UserRepository mocks base method.
func (m *MockIUnitOfWork) UserRepository() persistence.ISQLUser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRepository")
	ret0, _ := ret[0].(persistence.ISQLUser)
	return ret0
}

// UserRepository indicates an expected call of UserRepository.
func (mr *MockIUnitOfWorkMockRecorder) UserRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRepository", reflect.TypeOf((*MockIUnitOfWork)(nil).UserRepository))
}
