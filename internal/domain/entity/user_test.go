package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUser(t *testing.T) {
	// Arrange
	type user struct {
		name      string
		gender    string
		birthDate string
	}

	tests := []test[user]{
		{
			description: "Valid user - create a user of male gender",
			fields: user{
				name:      "johan",
				gender:    "MALE",
				birthDate: "1989-05-11",
			},
			wantError:    false,
		},
		{
			description: "Invalid user - the 'name' is invalid",
			fields: user{
				name:      "C4rl0s Gun'n R0s3",
				gender:    "female",
				birthDate: "2005-12-25",
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid user - the 'gender' field is invalid",
			fields: user{
				name:      "Carolline",
				gender:    "zombie",
				birthDate: "2001-01-05",
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid user - the 'birthDate' is less than minimun age requeriment",
			fields: user{
				name:      "Carolline",
				gender:    "female",
				birthDate: "2009-01-05",
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid user - all fields are invalids",
			fields: user{
				name:      "Donald",
				gender:    "Puri Puri Pr1s1oner",
				birthDate: "2001/01/05",
			},
			wantError:              true,
			errorQuantityExpected: 2,
		},
	}
	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			user, err := NewUser(tt.fields.name, tt.fields.gender, tt.fields.birthDate)
			
			// Assert
			if tt.wantError {
				assert.Nil(t, user)
				assert.Len(t, err, tt.errorQuantityExpected)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, User{}, *user)
			}
		})
	}
}
