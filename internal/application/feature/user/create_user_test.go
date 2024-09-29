package userfeature

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/factory"
	repositorymock "stock-controll/test/mock/repository"
	uowmock "stock-controll/test/mock/uow"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupRepositoriesMock(ctrl *gomock.Controller) (
	*uowmock.MockIUnitOfWork,
	*repositorymock.MockIUserRepository,
	*repositorymock.MockICredentialRepository,
	*repositorymock.MockIContactRepository,
) {
	unitOfWorkMock := uowmock.NewMockIUnitOfWork(ctrl)

	userRepositoryMock := repositorymock.NewMockIUserRepository(ctrl)
	credentialRepositoryMock := repositorymock.NewMockICredentialRepository(ctrl)
	contactRepositoryMock := repositorymock.NewMockIContactRepository(ctrl)

	unitOfWorkMock.EXPECT().UserRepository().Return(userRepositoryMock).AnyTimes().MaxTimes(1)
	unitOfWorkMock.EXPECT().CredentialRepository().Return(credentialRepositoryMock).AnyTimes().MaxTimes(1)
	unitOfWorkMock.EXPECT().ContactRepository().Return(contactRepositoryMock).AnyTimes().MaxTimes(1)

	return unitOfWorkMock, userRepositoryMock, credentialRepositoryMock, contactRepositoryMock
}

func featureCreateUser(unitOfWorkMock *uowmock.MockIUnitOfWork) *CreateUserFeature {
	return &CreateUserFeature{
		uow:               unitOfWorkMock,
		userFactory:       &factory.UserFactory{},
		credentialFactory: &factory.CredentialFactory{},
		contactFactory:    &factory.ContactFactory{},
	}
}

var fake = faker.New()

type generateUserRequestDTO struct {
	userDTO dto.CreateUserRequestDTO
}

func NewGenerateCreateUserRequestDTO() *generateUserRequestDTO {
	return &generateUserRequestDTO{
		userDTO: dto.CreateUserRequestDTO{
			Name:      fake.Person().FirstName(),
			Gender:    fake.Person().Gender(),
			BirthDate: "1989-05-25",
			Password:  "WithUpper1945Letters##",
			Email:     fake.Person().Contact().Email,
			Phone:     "(21) 4002-8922",
		},
	}
}

func (g *generateUserRequestDTO) UseInvalidPassword() *generateUserRequestDTO {
	g.userDTO.Password = "123abc"
	return g
}

func (g *generateUserRequestDTO) UseInvalidName() *generateUserRequestDTO {
	g.userDTO.Name = "R0M4N0!"
	return g
}

func (g *generateUserRequestDTO) UseInvalidEmail() *generateUserRequestDTO {
	g.userDTO.Email = g.userDTO.Email + "!"
	return g
}

func (g *generateUserRequestDTO) Build() *dto.CreateUserRequestDTO {
	return &g.userDTO
}

func Test_CreateUser(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	unitofworkMock, userRepositoryMock, credentialRepositoryMock, contactRepositoryMock := setupRepositoriesMock(ctrl)

	unitofworkMock.EXPECT().Begin().Return(nil)
	unitofworkMock.EXPECT().Commit().Return(nil)

	userRepositoryMock.EXPECT().Save(gomock.Any()).Return(1, nil)
	credentialRepositoryMock.EXPECT().Save(gomock.Any()).Return(nil)
	contactRepositoryMock.EXPECT().Save(gomock.Any()).Return(nil)

	feature := featureCreateUser(unitofworkMock)

	userDTO := NewGenerateCreateUserRequestDTO().Build()

	// Act
	result, err := feature.CreateUser(*userDTO)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func Test_CreateUser_InvalidParamsForCreateUserEntity(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	unitofworkMock, _, _, _ := setupRepositoriesMock(ctrl)

	feature := featureCreateUser(unitofworkMock)
	userDTO := NewGenerateCreateUserRequestDTO().UseInvalidName().Build()

	// Act
	response, err := feature.CreateUser(*userDTO)

	// Assert
	assert.Nil(t, response)
	assert.Error(t, err)
}

func Test_CreateUser_InvalidParamsForCreateCredentialEntity(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	unitofworkMock, userRepositoryMock, _, _ := setupRepositoriesMock(ctrl)
	unitofworkMock.EXPECT().Begin().Return(nil)
	unitofworkMock.EXPECT().Rollback().Return(nil)

	userRepositoryMock.EXPECT().Save(gomock.Any()).Return(1, nil)

	feature := featureCreateUser(unitofworkMock)
	userDTO := NewGenerateCreateUserRequestDTO().UseInvalidPassword().Build()

	// Act
	response, err := feature.CreateUser(*userDTO)

	// Assert
	assert.Nil(t, response)
	assert.Error(t, err)
}

func Test_CreateUser_InvalidParamsForCreateContactEntity(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	unitofworkMock, userRepositoryMock, credentialRepositoryMock, _ := setupRepositoriesMock(ctrl)
	unitofworkMock.EXPECT().Begin().Return(nil)
	unitofworkMock.EXPECT().Rollback().Return(nil)

	userRepositoryMock.EXPECT().Save(gomock.Any()).Return(1, nil)
	credentialRepositoryMock.EXPECT().Save(gomock.Any()).Return(nil)

	feature := featureCreateUser(unitofworkMock)
	userDTO := NewGenerateCreateUserRequestDTO().UseInvalidEmail().Build()

	// Act
	response, err := feature.CreateUser(*userDTO)

	// Assert
	assert.Nil(t, response)
	assert.Error(t, err)
}
