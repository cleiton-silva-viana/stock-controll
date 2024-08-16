package delete_contact

import (
	"testing"

	"stock-controll/tests/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_DeleteContact(t *testing.T) {
	// Arrange
	id := 1
	contactRepositoryMock := new(mocks.ContactRepositoryMock)

	// Act
	deleteContactFeature := DeleteContact(contactRepositoryMock)
	contactRepositoryMock.On("Delete", id).Return(nil)
	err := deleteContactFeature.Execute(id)

	// Assert
	assert.Nil(t, err)
}
