package userfeature

import (
	uowmock "stock-controll/test/mock/uow"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DeleteUser(t *testing.T) {
	// Arrange
	uowMock := uowmock.NewUniOfWorkMock()
	userRepositoryMock, credentialRepositoryMock, contactRepositoryMock := setupRepositoriesMocks()

	uowMock.On("Begin").Return(nil)
	uowMock.On("Commit").Return(nil)

	userRepositoryMock.On("Delete", UID).Return(nil)
	credentialRepositoryMock.On("Delete", UID).Return(nil)
	contactRepositoryMock.On("Delete", UID).Return(nil)

	feature := NewUserFeatureBuilder().
		SetUnitOfWork(uowMock).
		SetRepositories(userRepositoryMock, credentialRepositoryMock, contactRepositoryMock).
		Build()

	// Act
	err := feature.Delete(UID)

	// Assert
	assert.Nil(t, err)
}
