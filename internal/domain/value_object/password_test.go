package valueobject

import (
	"testing"
	
	"stock-controll/internal/domain/failure"

	"github.com/stretchr/testify/assert"
)

func Test_NewPassword(t *testing.T) {
	//Arrange
	tests := []test{
		{
			test_description: "Valid password - with 8 characters", 
			field: "AbCd24@$", 
			want_error: false,
		},
		{
			test_description: "Valid password - with 16 characters", 
			field: "AbCd24@$A&56dv$1", 
			want_error: false,
		},
		
		{
			test_description: "Valid password - with 24 characters", 
			field: "Abcd$%@$AbCd24@$A65d24@$", 
			want_error: false,
		},
		{
			test_description: "Invalid password - less than 8 characters", 
			field: "Halo$12", 
			want_error: true, 
			expected_error: failure.PasswordIsShort(8),
		},
		{
			test_description: "Invalid password - more than 24 characters", 
			field: "Abcd$%@$AbCd24@$A65d24@$cf$12", 
			want_error: true, 
			expected_error: failure.PasswordIsLong(24),
		},
		{
			test_description: "Invalid password - without lower cases", 
			field: "HALO123%$", 
			want_error: true, 
			expected_error: failure.PasswordNotContainsLowerCases,
		},
		{
			test_description: "Invalid password - without upper cases", 
			field: "halo123%$", 
			want_error: true, 
			expected_error: failure.PasswordNotContainsUpperCases,
		},
		{
			test_description: "Invalid password - without special characters", 
			field: "PaloAlto123", 
			want_error: true, 
			expected_error: failure.FieldWithoutSpecialChars("password"),
		},
		{
			test_description: "Invalid password - without number", 
			field: "PaloAlto@#$", 
			want_error: true, 
			expected_error: failure.FieldWithoutNumber("password"),
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.test_description, func(t *testing.T) {
			password, err := NewPassword(tt.field)

			// Assert
			if tt.want_error {
				assert.Nil(t, password)
				assert.Equal(t, tt.expected_error, err)
				} else {
				assert.Nil(t, err)
				assert.IsType(t, Password{}, *password)
				assert.NotEmpty(t, password.password_hashed)
				assert.NotEmpty(t, password.password_salt)
			}
		})
	}
}
