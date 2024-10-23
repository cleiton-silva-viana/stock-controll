package factorymock

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

type UserFactoryMock struct {
	mock.Mock
}

func NewUserFactoryMock() *UserFactoryMock {
	return &UserFactoryMock{}
}

func (c *UserFactoryMock) Create(dto dto.UserDTO) (*entity.User, error) {
	args := c.Called(dto)
	return args.Get(0).(*entity.User), args.Error(1)
}
