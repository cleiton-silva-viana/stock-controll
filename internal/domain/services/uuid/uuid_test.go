package uuid

import (
	"regexp"
	"testing"

	"stock-controll/internal/domain/services/validate"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	// Act
	uuid := New()

	// Assert
	require.Len(t, uuid, uuidLength)
	assert.Regexp(t, regexp.MustCompile(re), uuid)
}

func TestIsValidNoError(t *testing.T) {
	// Arrange
	uuid := New()

	// Act
	err := IsValid("uuid", uuid)

	// Assert
	assert.NoError(t, err)
}

// TODO: usar o pacote unitary...
func TestIsValidWithError(t *testing.T) {
	// Arrange
	var tests = map[string]struct {
		uuid string
	}{
		"uuid is empty": {
			uuid: "",
		},
		"uuid is filled with empty chars": {
			uuid: "                                    ",
		},
		"uuid with invalid length": {
			uuid: "12345678-1234-1234-1234-12345678",
		},

		"uuid has invalid format": {
			uuid: "019252a1kd83a-7700-be9e-e7d3c8c89b68",
		},
	}

	for name, field := range tests {
		t.Run(name, func(t *testing.T) {

			// Act
			err := IsValid("uuid", field.uuid)

			// Assert
			require.Error(t, err)
			assert.ErrorAs(t, &validate.FieldError{}, &err)
		})
	}
}
