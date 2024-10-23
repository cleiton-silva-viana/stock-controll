package userfeature

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
	"stock-controll/internal/domain/factory"
	uowmock "stock-controll/test/mock/uow"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {
	// Arrange
	uowMock := uowmock.NewUniOfWorkMock()
	userRepositoryMock, credentialRepositoryMock, contactRepositoryMock := setupRepositoriesMocks()
	userFactoryMock, credentialFactoryMock, contactFactoryMock := setupFactoriesMocks()

	uowMock.On("Begin").Return(nil)
	uowMock.On("Commit").Return(nil)

	feature := NewUserFeatureBuilder().
		SetUnitOfWork(uowMock).
		SetFactories(userFactoryMock, credentialFactoryMock, contactFactoryMock).
		SetRepositories(userRepositoryMock, credentialRepositoryMock, contactRepositoryMock).
		Build()

	// Act
	response, err := feature.CreateUser(*userRequestDTO)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func Test_CreateUser_InvalidUserDatas(t *testing.T) {
	var tests = []struct {
		description       string
		modifyDTO         func(dto *dto.CreateUserRequestDTO)
		expectedErrorType entity.EntityError
	}{
		{
			description: "DTO with invalid datas for create a user entity",
			modifyDTO: func(dto *dto.CreateUserRequestDTO) {
				dto.Name = "invalid.name_for_user@"
			},
		},
		{
			description: "DTO with invalid datas for create a credential entity",
			modifyDTO: func(dto *dto.CreateUserRequestDTO) {
				dto.Password = "this_password_not_contais_numbers_and_is_longer"
			},
		},
		{
			description: "DTO with invalid datas for create a contact entity",
			modifyDTO: func(dto *dto.CreateUserRequestDTO) {
				dto.Email = "your.email@"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// Arrange
			uowMock := uowmock.NewUniOfWorkMock()
			uowMock.On("Begin").Return(nil)
			feature := NewUserFeatureBuilder().
				SetFactories(&factory.UserFactory{}, &factory.CredentialFactory{}, &factory.ContactFactory{}).
				Build()
			tt.modifyDTO(userRequestDTO)

			// Act
			response, err := feature.CreateUser(*userRequestDTO)

			// Assert
			assert.Nil(t, response)
			assert.IsType(t, &entity.EntityError{}, err)
		})
	}
}
