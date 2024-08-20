package value_object

import (
	"testing"

	"stock-controll/internal/domain/failure"

	"github.com/stretchr/testify/assert"
)

type test_name struct {
	description    string
	name           string
	fieldName      string
	error_expected error
}

func Test_CreateName_WithValidName(t *testing.T) {
	// Arrange
	var tests = []test_name{
		{description: "name with 3 letters", fieldName: "name", name: "Ana"},
		{description: "name with 20 chars", fieldName: "name", name: "Alexandre de Silveir"},
	}

	// Act
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result, err := NewName(test.name, test.fieldName, 20)

			// Assert
			assert.Nil(t, err)
			assert.IsType(t, &Name{}, result)
		})
	}
}

func Test_CreateName_WithInvalidName(t *testing.T) {
	// Arrange
	var tests = []test_name{
		{description: "name with minor than 3 chars", name: "an", fieldName: "name", error_expected: failure.NameIsShort("name", 3)},
		{description: "name with big than 20 chars", name: "Maximilianoshompenhaurcasatelar", fieldName: "name", error_expected: failure.NameIsLong("name", 20)},
		{description: "name with number", name: "R0B3rT0", fieldName: "name", error_expected: failure.NameWithNumber("name")},
		{description: "name with special characters", name: "Ezilabeth$", fieldName: "name", error_expected: failure.NameWithInvalidChars("name")},
		{description: "name with space (empyt)", name: "   ", fieldName: "name", error_expected: failure.NameIsEmpty("name")},
	}

	// Act
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result, err := NewName(test.name, test.fieldName, 20)

			// Assert
			assert.Nil(t, result)
			assert.Equal(t, err.Error(), test.error_expected.Error())
		})
	}
}
