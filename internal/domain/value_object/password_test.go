package value_object

import (
	"testing"
	
	"stock-controll/internal/domain/failure"

	"github.com/stretchr/testify/assert"
)

func Test_NewPassword_WithValidPassword(t *testing.T) {
	// Arrange
	var tests = []struct {
		descrition string
		password   string
	}{
		{descrition: "password with 8 characters", password: "AbCd24@$"},
		{descrition: "password with 24 characters", password: "Abcd$%@$AbCd24@$A65d24@$"},
	}

	// Act
	for _, test := range tests {
		t.Run(test.descrition, func(t *testing.T) {
			result, err := NewPassword(test.password)

			// Assert
			assert.Nil(t, err)
			assert.IsType(t, &Password{}, result)
			assert.NotEmpty(t, result.password_hashed)
			assert.NotEmpty(t, result.password_salt)
		})
	}
}

func Test_NewPassword_WithInvalidParams(t *testing.T) {
	// Arrange
	var tests = []struct {
		descritpion    string
		password       string
		expected_error error
	}{
		{descritpion: "password less than 8 characters", password: "Halo$12", expected_error: failure.PasswordIsShort},
		{descritpion: "password with more than 24 characters", password: "Palo123Alto456San!@#Diego$%&", expected_error: failure.PasswordIsLong},
		{descritpion: "password without lower cases", password: "HALO123%$", expected_error: failure.PasswordNotContainsLowerCases},
		{descritpion: "password without upper cases", password: "halo123%$", expected_error: failure.PasswordNotContainsUpperCases},
		{descritpion: "password without special characters", password: "PaloAlto123", expected_error: failure.PasswordNotContainsSpecialChars},
		{descritpion: "password without number", password: "PaloAlto@#$", expected_error: failure.PasswordNotContainsNumbers},
	}

	// Act
	for _, test := range tests {
		t.Run(test.descritpion, func(t *testing.T) {
			result, err := NewPassword(test.password)

			// Assert
			assert.Nil(t, result)
			assert.ErrorIs(t, err, test.expected_error)
		})
	}
}
