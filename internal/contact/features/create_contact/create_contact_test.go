package create_contact

import (
	"testing"

	contact_DTO "stock-controll/internal/contact/DTO"
	"stock-controll/tests/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_CreateNewContactUseCase_AllFieldsAreValid(t *testing.T) {
	// Arrange
	var person = contact_DTO.CreateContactDTO{
		Email: "cleitonviana@outlook.com.br",
		Phone: "(21) 99183-9222",
	}
	contactRepositoryMock := new(mocks.ContactRepositoryMock)

	// Act
	createContactFeature := CreateContact(contactRepositoryMock)
	contactRepositoryMock.On("Save").Return(nil)
	err := createContactFeature.Execute(person)

	// Assert
	assert.Nil(t, err)
}
