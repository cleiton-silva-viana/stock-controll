package feature

import (
	"testing"

	"stock-controll/internal/application/dto"

	"github.com/stretchr/testify/assert"
)

func Test_CreateNewContactUseCase_AllFieldsAreValid(t *testing.T) {
	// Arrange
	var person = dto.CreateContactDTO{
		Email: "cleitonviana@outlook.com.br",
		Phone: "(21) 99183-9222",
	}

	// Act
	createContactFeature := ContactFeature(contactRepositoryMock)
	contactRepositoryMock.On("Save").Return(nil)
	err := createContactFeature.CreateContact(person)

	// Assert
	assert.Nil(t, err)
}
