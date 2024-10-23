package factorymock

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

type CredentialFactoryMock struct {
	mock.Mock
}

func NewCredentialFactoryMock() *CredentialFactoryMock {
	return &CredentialFactoryMock{}
}

func (c *CredentialFactoryMock) Create(dto dto.CreateCredentialDTO) (*entity.Credential, error) {
	args := c.Called(dto)
	return args.Get(0).(*entity.Credential), args.Error(1)
}
