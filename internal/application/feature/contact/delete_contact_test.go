package feature

import (
	"testing"

	"stock-controll/tests/mocks"

	"github.com/stretchr/testify/assert"
)

var contactRepositoryMock = new(mocks.ContactRepositoryMock)

func Test_DeleteContact(t *testing.T) {
	// Arrange
	id := 1

	// Act
	deleteContactFeature := ContactFeature(contactRepositoryMock)
	contactRepositoryMock.On("Delete", id).Return(nil)
	err := deleteContactFeature.DeleteContact(id)

	// Assert
	assert.Nil(t, err)
}
