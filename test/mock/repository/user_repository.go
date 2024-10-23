package repositoryMock

import (
	"stock-controll/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

// Save implementa o método da interface IRepository[User]
func (m UserRepositoryMock) Save(user entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Update implementa o método da interface IRepository[User]
func (m UserRepositoryMock) Update(user entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Delete implementa o método da interface IRepository[User]
func (m UserRepositoryMock) Delete(UID string) error {
	args := m.Called(UID)
	return args.Error(0)
}

// GetByID implementa o método da interface IRepository[User]
func (m UserRepositoryMock) GetByUID(UID string) (entity.User, error) {
	args := m.Called(UID)
	return args.Get(0).(entity.User), args.Error(1)
}
