package uuid

import (
	"fmt"
	"testing"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Setup() *UUID {
	return New()
}

func TestIsValidNoError(t *testing.T) {
	// Arrange
	u := Setup()

	// Act
	err := IsValid("uuid", u.String())

	// Assert
	assert.NoError(t, err)
}

func TestParseNoError(t *testing.T) {
	// Arrange
	u := Setup()

	// Act
	result, err := Parse("uuid", u.String())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, result.String(), u.String())
}

var tests = []unitary.TestField[UUID]{
	{
		Description: "uuid is empty",
		Handler:     func(u *UUID) { u.value = "" },
	},
	{
		Description: "uuid is filled with empty chars",
		Handler:     func(u *UUID) { u.value = "                                    " },
	},
	{
		Description: "uuid with invalid length",
		Handler:     func(u *UUID) { u.value = "12345678-1234-1234-1234-12345678" },
	},
	{
		Description: "uuid has invalid format",
		Handler:     func(u *UUID) { u.value = "019252a1kd83a-7700-be9e-e7d3c8c89b68" },
	},
}

func TestIsValidWithError(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("test function IsValid - expected error because %s", tt.Description), func(t *testing.T) {
			// Arrange
			u := Setup()
			tt.Handler(u)

			// Act
			err := IsValid("uuid", u.value)

			// Assert
			require.Error(t, err)
		})
	}
}

func TestParseWithError(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Test funciton Parse - expected because- %s", tt.Description), func(t *testing.T) {
			// Arrange
			u := Setup()
			tt.Handler(u)

			// Act
			err := IsValid("uuid", u.value)

			// Assert
			require.Error(t, err)
		})
	}
}
