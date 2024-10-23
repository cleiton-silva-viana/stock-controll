package uowmock

import "github.com/stretchr/testify/mock"

type UnitOfWorkMock struct {
	mock.Mock
}

func NewUniOfWorkMock() *UnitOfWorkMock {
	return &UnitOfWorkMock{}
}

func (m *UnitOfWorkMock) Begin() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UnitOfWorkMock) Commit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UnitOfWorkMock) Rollback() error {
	args := m.Called()
	return args.Error(0)
}
