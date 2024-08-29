package entity

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func Test_NewContact(t *testing.T) {
	// Arrange
	type contact struct {
		email string
		phone string
	}

	tests := []test[contact]{
		{
			description: "Valid contact - must return a contact entity",
			fields: contact{
				phone: "(21) 98322-4321",
				email: faker.New().Internet().Email(),
			},
			wantError: false,
		},
		{
			description: "Invalid contact - phone invalid, expected 1 error in errorList",
			fields: contact{
				phone: "(21) 98322-4321a",
				email: faker.New().Internet().Email(),
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid contact - email invalid, expected 1 error in errorList",
			fields: contact{
				phone: "(21) 98322-4321",
				email: "donald.trump@contact.com.",
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid contact - email and phone invalid, expected 2 errors in errorList",
			fields: contact{
				phone: "(21) 98322-4321a",
				email: "joe.biden.io",
			},
			wantError:              true,
			errorQuantityExpected: 2,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			contact, err := NewContact(1, tt.fields.email, tt.fields.phone)

			// Assert
			if tt.wantError {
				assert.Nil(t, contact)
				assert.Len(t, err, tt.errorQuantityExpected)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Contact{}, *contact)
			}
		})
	}
}
