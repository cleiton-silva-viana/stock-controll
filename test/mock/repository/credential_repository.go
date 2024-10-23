package repositoryMock

import (
	"stock-controll/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

type CredentialRepositoryMock struct {
	mock.Mock
}

func NewCredentialRepositoryMock() *CredentialRepositoryMock {
	return &CredentialRepositoryMock{}
}

// Save implementa o método da interface IRepository[Credential]
func (m CredentialRepositoryMock) Save(credential entity.Credential) error {
	args := m.Called(credential)
	return args.Error(0)
}

// Update implementa o método da interface IRepository[Credential]
func (m CredentialRepositoryMock) Update(credential entity.Credential) error {
	args := m.Called(credential)
	return args.Error(0)
}

// Delete implementa o método da interface IRepository[Credential]
func (m CredentialRepositoryMock) Delete(UID string) error {
	args := m.Called(UID)
	return args.Error(0)
}

// GetByID implementa o método da interface IRepository[Credential]
func (m CredentialRepositoryMock) GetByUID(UID string) (entity.Credential, error) {
	args := m.Called(UID)
	return args.Get(0).(entity.Credential), args.Error(1)
}