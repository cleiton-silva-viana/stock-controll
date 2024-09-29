package userfeature

import (
	repositorymock "stock-controll/test/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_DeleteUser(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userID := 1

	userRepositoryMock := repositorymock.NewMockIUserRepository(ctrl)
	userRepositoryMock.EXPECT().Delete(userID).Return(nil)

	feature := NewUserRepository(userRepositoryMock)

	// Act
	err := feature.Delete(userID)

	// Assert
	assert.Nil(t, err)
}

