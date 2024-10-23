package repositoryMock

import (
	"stock-controll/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

type ContactRepositoryMock struct {
	mock.Mock
}

func NewContactRepositoryMock() *ContactRepositoryMock {
	return &ContactRepositoryMock{}
}

// Save implementa o método da interface IRepository[Contact]
func (m ContactRepositoryMock) Save(contact entity.Contact) error {
	args := m.Called(contact)
	return args.Error(0)
}

// Update implementa o método da interface IRepository[Contact]
func (m ContactRepositoryMock) Update(contact entity.Contact) error {
	args := m.Called(contact)
	return args.Error(0)
}

// Delete implementa o método da interface IRepository[Contact]
func (m ContactRepositoryMock) Delete(UID string) error {
	args := m.Called(UID)
	return args.Error(0)
}

// GetByID implementa o método da interface IRepository[Contact]
func (m ContactRepositoryMock) GetByUID(UID string) (entity.Contact, error) {
	args := m.Called(UID)
	return args.Get(0).(entity.Contact), args.Error(1)
}