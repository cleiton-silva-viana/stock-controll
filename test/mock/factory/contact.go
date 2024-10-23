package factorymock

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

type ContactFactoryMock struct {
	mock.Mock
}

func NewContactFactoryMock() *ContactFactoryMock {
	return &ContactFactoryMock{}
}

func (c *ContactFactoryMock) Create(dto dto.CreateContactDTO) (*entity.Contact, error) {
	args := c.Called(dto)
	return args.Get(0).(*entity.Contact), args.Error(1)
}
