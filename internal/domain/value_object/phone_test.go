package valueobject

import (
	"testing"

	"stock-controll/internal/domain/failure"

	"github.com/stretchr/testify/assert"
)

func Test_NewPhone(t *testing.T) {
	// Assert
	tests := []test{
		{
			test_description: "Valid number - landline",
			field:            "(21) 4002-8922",
			want_error:       false,
		},
		{test_description: "Valid number - cellphone",
			field:      "(21) 99314-3214",
			want_error: false,
		},
		{
			test_description: "Invalid number - empty field", 
			field: "     ", 
			want_error: true, 
			expected_error: failure.PhoneWithInvalidLength,
		},
		{
			test_description: "Invalid number - without area code", 
			field: "219269-82225", 
			want_error: true, 
			expected_error: failure.PhoneWithoutAreaCode,
		},
		{
			test_description: "Invalid number - contains letters", 
			field: "(21) 4002-8922a", 
			want_error: true, 
			expected_error: failure.PhoneWithLetters,
		},
		{
			test_description: "Invalid number - contains special characters", 
			field: "(21) 4002-8922#", 
			want_error: true, 
			expected_error: failure.FieldWithSpecialChars("phone"),
		},
		{
			test_description: "Invalid number - less than 10 digits", 
			field: "(21) 4002-892", 
			want_error: true, 
			expected_error: failure.PhoneWithInvalidLength,
		},
		{
			test_description: "Invalid number - long than 11 digits", 
			field: "(21) 40028-23842",
			want_error: true, 
			expected_error: failure.PhoneWithInvalidLength,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.test_description, func(t *testing.T) {
			phone, err := NewPhone(tt.field)

			// Arrange
			if tt.want_error {
				assert.Nil(t, phone)
				assert.Equal(t, tt.expected_error, err)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Phone{}, *phone)
			}
		})
	}
}
