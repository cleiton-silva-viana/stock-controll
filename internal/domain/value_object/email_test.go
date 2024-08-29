package valueobject

import (
	"testing"

	"stock-controll/internal/domain/failure"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var fake = faker.New()

func Test_NewEmail(t *testing.T) {
	// Arrange
	tests := []test{
		{
			test_description: "Valid email - genenated with faker package",
			field:            fake.Internet().Email(),
			want_error:       false,
		},
		{
			test_description: "Invalid email - without character @",
			field:            "user.example.com",
			want_error:       true,
			expected_error: failure.FieldWithInvalidFormat("email", "example@domain.com"),
		},
		{
			test_description: "Invalid email - without domain",
			field:            "user.example@",
			want_error:       true,
			expected_error: failure.FieldWithInvalidFormat("email", "example@domain.com"),
		},
		{
			test_description: "Invalid email - without user name",
			field:            "@example.com",
			want_error:       true,
			expected_error: failure.FieldWithInvalidFormat("email", "example@domain.com"),
		},
		{
			test_description: "Invalid email - invalid domain",
			field:            "@example..com",
			want_error:       true,
			expected_error: failure.FieldWithInvalidFormat("email", "example@domain.com"),
		},
		{
			test_description: "Invalid email - with white spaces",
			field:            "user @example.com",
			want_error:       true,
			expected_error: failure.FieldWithInvalidFormat("email", "example@domain.com"),
		},
		{
			test_description: "Invalid email - domain without TDL",
			field:            "user@example",
			want_error:       true,
			expected_error: failure.FieldWithInvalidFormat("email", "example@domain.com"),
		},
		{
			test_description: "Invalid email - initializing with characters specials",
			field:            "@user@example.com",
			want_error:       true,
			expected_error: failure.FieldWithInvalidFormat("email", "example@domain.com"),
		},
		{
			test_description: "Invalid email - ending with special characters",
			field:            "user@example.com.",
			want_error:       true,
			expected_error: failure.FieldWithInvalidFormat("email", "example@domain.com"),
		},
		{
			test_description: "Invalid email - format incorrect",
			field:            "user@example,com.",
			want_error:       true,
			expected_error: failure.FieldWithInvalidFormat("email", "example@domain.com"),
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.test_description, func(t *testing.T) {
			email, err := NewEmail(tt.field)
			
			// Assert
			if tt.want_error {
				assert.Nil(t, email)
				assert.Equal(t, tt.expected_error, err)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Email{}, *email)
			}
		})
	}
}
