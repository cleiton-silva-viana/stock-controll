package userfeature

import (
	"testing"

	"stock-controll/internal/application/dto"
	factorymock "stock-controll/test/mock/factory"
	repositorymock "stock-controll/test/mock/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UpdateUser(t *testing.T) {
	// Arrange
	newDatas := dto.UserDTO{
		Name:      fake.Person().FirstName(),
		Gender:    "Male",
		BirthDate: "1999-05-01",
	}

	userFactoryMock := factorymock.NewUserFactoryMock()
	userRepository := repositorymock.NewUserRepositoryMock()

	userFactoryMock.On("Create", mock.AnythingOfType("dto.UserDTO")).Return(user, nil)
	userRepository.On("Update", mock.AnythingOfType("entity.User")).Return(nil)

	featureBuilder := NewUserFeatureBuilder()
	featureBuilder.userFactory = userFactoryMock
	featureBuilder.userRepository = userRepository
	feature := featureBuilder.Build()

	// Act
	err := feature.Update(newDatas)

	// Assert
	assert.Nil(t, err)
}
