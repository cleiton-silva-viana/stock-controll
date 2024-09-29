package userfeature

import (
	"stock-controll/internal/application/dto"
	repositorymock "stock-controll/test/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_UpdateUser(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userNewDatas := dto.UserDTO{
		Name:      "cleiton",
		Gender:    "Male",
		BirthDate: "1999-05-01",
	}

	userRepositoryMock := repositorymock.NewMockIUserRepository(ctrl)
	userRepositoryMock.EXPECT().Update(gomock.Any()).Return(nil)

	feature := NewUserRepository(userRepositoryMock)

	// Act
	err := feature.Update(userNewDatas)

	// Assert
	assert.Nil(t, err)
}
