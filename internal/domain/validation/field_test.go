package validation

import (
	"testing"
	"time"


	"github.com/stretchr/testify/assert"
)

func Test_ContainsSpecialChars(t *testing.T) {
	// Arrange
	tests := []test{
		{
			description: "Must return true - the string have special characters",
			field:            "Maçâ$$$",
			wantError:       true,
		},
		{
			description: "Must return false - the string not have special characters",
			field:            "orange",
			wantError:       false,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ok := ContainsSpecialChars(tt.field)

			// Assert
			assert.Equal(t, tt.wantError, ok)
		})
	}
}

func Test_ContainsNumbers(t *testing.T) {
	// Arrange
	tests := []test{
		{
			description: "Must return true - the string have numbers",
			field:            "World War 2",
			wantError:       true,
		},
		{
			description: "Must return false - the string not have number",
			field:            "World War II",
			wantError:       false,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ok := ContainsNumbers(tt.field)

			// Assert
			assert.Equal(t, tt.wantError, ok)
		})
	}
}

func Test_ContainsLetters(t *testing.T) {
	// Arrange
	tests := []test{
		{
			description: "Must return true - the string have letters",
			field:            "pink",
			wantError:       true,
		},
		{
			description: "Must return false - the string not have letters",
			field:            "#$¨%#$46874  ",
			wantError:       false,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ok := ContainsLetters(tt.field)

			// Assert
			assert.Equal(t, tt.wantError, ok)
		})
	}
}
func Test_ContainsLowerCaseLetters(t *testing.T) {
	// Arrange
	tests := []test{
		{
			description: "Must return true - the string have lowercase letters",
			field:            "pink",
			wantError:       true,
		},
		{
			description: "Must return false - the string not have lowercase letters",
			field:            "EMMA",
			wantError:       false,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ok := ContainsLowerCaseLetters(tt.field)

			// Assert
			assert.Equal(t, tt.wantError, ok)
		})
	}
}

func Test_ContainsUpperCaseLetters(t *testing.T) {
	// Arrange
	tests := []test{
		{
			description: "Must return true - the string have uppercase letters",
			field:            "MINDFLOW",
			wantError:       true,
		},
		{
			description: "Must return false - the string not have uppercase letters",
			field:            "killswitch engage",
			wantError:       false,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ok := ContainsUpperCaseLetters(tt.field)

			// Assert
			assert.Equal(t, tt.wantError, ok)
		})
	}
}

func Test_DateFormatIsValid(t *testing.T) {
	// Arrange
	tests := []test{
		{
			description: "Must return true - the date format is valid",
			field:            "2024-09-21",
			wantError:       true,
		},
		{
			description: "Must return false - the date format is not valid",
			field:            "1945/11/25",
			wantError:       false,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ok, date := DateFormatIsValid(tt.field)

			// Assert
			if tt.wantError {
				assert.IsType(t, time.Time{}, *date)
				assert.True(t, ok)
			} else {
				assert.Nil(t, date)
				assert.False(t, ok)
			}
		})
	}
}
