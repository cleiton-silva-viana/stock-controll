package value_object

import (
	"testing"
	
	"stock-controll/internal/domain/failure"

	"github.com/stretchr/testify/assert"
)

func Test_CreateName_WithValidName(t *testing.T) {
	// Arrange
	var tests = []struct {
		description    string
		name           string
	}{
		{description: "name with 3 letters", name: "Ana"},
		{description: "name with 20 chars", name: "Alexandre de Silveir"},
	}

	// Act
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result, err := NewName(test.name)

			// Assert
			assert.Nil(t, err)
			assert.IsType(t, &Name{}, result)
		})
	}
}

func Test_CreateName_WithInvalidName(t *testing.T) {
	// Arrange
	var tests = []struct {
		description    string
		name           string
		error_expected error
	}{
		{description: "name with minor than 3 chars", name: "an", error_expected: failure.NameIsShort},
		{description: "name with big than 20 chars", name: "Maximilianoshompenhaurcasatelar", error_expected: failure.NameIsLong},
		{description: "name with number", name: "R0B3rT0", error_expected: failure.NameWithInvalidChars},
		{description: "name with special characters", name: "Ezilabeth$", error_expected: failure.NameWithInvalidChars},
		{description: "name with space (empyt)", name: "   ", error_expected: failure.NameIsEmpty},
	}

	// Act
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result, err := NewName(test.name)

			// Assert
			assert.Nil(t, result)
			assert.ErrorIs(t, err, test.error_expected)
		})
	}
}
