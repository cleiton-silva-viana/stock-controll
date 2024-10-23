package userfeature

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
	factorymock "stock-controll/test/mock/factory"
	repositorymock "stock-controll/test/mock/repository"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/mock"
)

var fake = faker.New()

var userRequestDTO = &dto.CreateUserRequestDTO{
	Name:      fake.Person().FirstName(),
	Gender:    "male",
	BirthDate: "1989-05-25",
	Password:  "WithUpper1945Letters##",
	Email:     fake.Person().Contact().Email,
	Phone:     "(21) 4002-8922",
}

var UID = "123 456 789"
var user, _ = entity.NewUser(userRequestDTO.Name, userRequestDTO.Gender, userRequestDTO.BirthDate)
var credential, _ = entity.NewCredential(UID, userRequestDTO.Password)
var contact, _ = entity.NewContact(UID, userRequestDTO.Email, userRequestDTO.Phone)

func setupRepositoriesMocks() (
	*repositorymock.UserRepositoryMock,
	*repositorymock.CredentialRepositoryMock,
	*repositorymock.ContactRepositoryMock,
) {
	userRepositoryMock := repositorymock.NewUserRepositoryMock()
	credentialRepositoryMock := repositorymock.NewCredentialRepositoryMock()
	contactRepositoryMock := repositorymock.NewContactRepositoryMock()

	userRepositoryMock.On("Save", mock.AnythingOfType("entity.User")).Return(nil)
	credentialRepositoryMock.On("Save", mock.AnythingOfType("entity.Credential")).Return(nil)
	contactRepositoryMock.On("Save", mock.AnythingOfType("entity.Contact")).Return(nil)

	return userRepositoryMock, credentialRepositoryMock, contactRepositoryMock
}

func setupFactoriesMocks() (
	*factorymock.UserFactoryMock,
	*factorymock.CredentialFactoryMock,
	*factorymock.ContactFactoryMock,
) {
	userFactoryMock := factorymock.NewUserFactoryMock()
	credentialFactoryMock := factorymock.NewCredentialFactoryMock()
	contactFactoryMock := factorymock.NewContactFactoryMock()

	userFactoryMock.On("Create", mock.AnythingOfType("dto.UserDTO")).Return(user, nil)
	credentialFactoryMock.On("Create", mock.AnythingOfType("dto.CreateCredentialDTO")).Return(credential, nil)
	contactFactoryMock.On("Create", mock.AnythingOfType("dto.CreateContactDTO")).Return(contact, nil)

	return userFactoryMock, credentialFactoryMock, contactFactoryMock
}
