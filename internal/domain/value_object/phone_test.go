package value_object

import (
	"testing"
	
	"stock-controll/internal/domain/failure"

	"github.com/stretchr/testify/assert"
)

func Test_NewPhone_WithValidPhone(t *testing.T) {
	// arrange
	tests := []struct {
		name   string
		device string
		number string
	}{
		{name: "landline number", device: "landline", number: "(21) 4002-8922"},
		{name: "cellphone number", device: "cellphone", number: "(21) 99314-3214"},
	}

	// Act
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := NewPhone(test.number)

			// Assert
			assert.Nil(t, err)
			assert.Equal(t, test.device, result.device)
		})
	}
}

func Test_NewPhone_PhoneWithoutAreaCode_MustReturnAreacodeError(t *testing.T) {
	// Arrange
	phone := "219269-82225"

	// Act
	result, err := NewPhone(phone)

	// Assert
	assert.Nil(t, result)
	assert.ErrorIs(t, err, failure.PhoneWithoutAreaCode)
}

func Test_NewPhone_PhoneWithLetters_MustReturnErrorInvalidFormatError(t *testing.T) {
	// Arrange
	phone := "(21) 4002-8922a"

	// Act
	result, err := NewPhone(phone)

	// Assert
	assert.Nil(t, result)
	assert.ErrorIs(t, err, failure.PhoneWithLetters)
}

func Test_NewPhone_PhoneWithSpecialChars_MustReturnErrorSpecialCharsError(t *testing.T) {
	// Arrange
	phone := "(21) 4002-8922#"

	// Act
	result, err := NewPhone(phone)

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, failure.PhoneWithSpecialCharacters)
}

func Test_NewPhone_PhoneWithInvalidLength_MustReturnErrorLength(t *testing.T) {
	// Arrange
	phone := "(21) 40028-23842"

	// Act
	result, err := NewPhone(phone)

	// Assert
	assert.Nil(t, result)
	assert.ErrorIs(t, err, failure.PhoneWithInvalidLength)
}
