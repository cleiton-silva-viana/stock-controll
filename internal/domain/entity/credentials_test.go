package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCredentialst(t *testing.T) {
	//Arrange
	type credential struct {
		password string
		user_ID  int
	}

	tests := []test[credential]{
		{
			description: "Valid credential - password is safe",
			fields: credential{
				password: "!@#123ABCdef",
				user_ID:  1,
			},
			wantError: false,
		},
		{
			description: "Invalid credential - Empty user ID (default 0 are equals empty), 1 error expected",
			fields: credential{
				password: "!@#123ABCdef",
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid credential - password invalid, 1 error expected",
			fields: credential{
				password: "a1d2t3",
				user_ID:  1,
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid credential - password & user id invalids, 2 errors expected",
			fields: credential{
				password: "Oh..No!",
			},
			wantError:              true,
			errorQuantityExpected: 2,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			credential, err := NewCredential(tt.fields.user_ID, tt.fields.password, )

			// Assert
			if tt.wantError {
				assert.Nil(t, credential)
				assert.Len(t, err, tt.errorQuantityExpected)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Credential{}, *credential)
			}
		})
	}
}
