package userfeature

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/factory"
	persistencemock "stock-controll/test/mock/persistence"
	uowmock "stock-controll/test/mock/uow"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CreateUser(t *testing.T) {
	// Arrange

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	unitOfWorkMock := uowmock.NewMockIUnitOfWork(ctrl)

	userRepositoryMock := persistencemock.NewMockISQLUser(ctrl)
	credentialRepositoryMock := persistencemock.NewMockISQLCredential(ctrl)
	contactRepositoryMock := persistencemock.NewMockISQLContact(ctrl)

	unitOfWorkMock.EXPECT().Begin().Return(nil)

	unitOfWorkMock.
		EXPECT().UserRepository().Return(userRepositoryMock)
	unitOfWorkMock.
		EXPECT().CredentialRepository().Return(credentialRepositoryMock)
	unitOfWorkMock.
		EXPECT().ContactRepository().Return(contactRepositoryMock)

	userRepositoryMock.EXPECT().Save(gomock.Any()).Return(1, nil)
	credentialRepositoryMock.EXPECT().Save(gomock.Any()).Return(nil)
	contactRepositoryMock.EXPECT().Save(gomock.Any()).Return(nil)

	fake := faker.New()
	userDTO := dto.CreateUserRequestDTO{
		Name:      fake.Person().FirstName(),
		Gender:    fake.Person().Gender(),
		BirthDate: "1989-05-25",
		Password:  "WithUpper1945Letters##",
		Email:     fake.Person().Contact().Email,
		Phone:     "(21) 4002-8922",
	}
	feature := NewUserFeature(
		unitOfWorkMock,
		&factory.UserFactory{},
		&factory.CredentialFactory{},
		&factory.ContactFactory{})

	// Act
	response, err := feature.CreateUser(userDTO)

	// Arrange
	assert.Nil(t, err)
	assert.IsType(t, dto.CreateUserResponseDTO{}, *response)
}

