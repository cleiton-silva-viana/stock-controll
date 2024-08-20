package entity
import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

type user_contact struct {
	email string
	phone string
}

func Test_NewContact_MustReturnAContact(t *testing.T) {
	// Arrange
	email := faker.New().Internet().Email()
	phone := "(21) 98322-4321"

	// Act
	result, err := NewContact(email, phone)

	// Assert
	assert.Nil(t, err)
	assert.IsType(t, Contact{}, *result)
}

func Test_NewContact_MustReturnError(t *testing.T) {
	// Arrange
	tests := []struct {
		description string
		user_contact
	}{
		{description: "user with invalid email",
			user_contact: user_contact{
				email: "johan.email.gmail",
				phone: "(21) 98441-5432"}},
		{description: "user with invalid phone",
			user_contact: user_contact{
				email: faker.New().Internet().Email(),
				phone: "21984415432"}},
	}

	// Act
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result, err := NewContact(test.email, test.phone)

			// Assert
			assert.Nil(t, result)
			assert.NotEmpty(t, err)
		})
	}
}
