package valueobject

import (
	"testing"

	"stock-controll/internal/domain/failure"

	"github.com/stretchr/testify/assert"
)

func Test_NewGender(t *testing.T) {
	// Arrange
	tests := []test{
		{
			test_description: "Valid gender - create a male gender",
			field:            "Male",
			want_error:       false,
		},
		{
			test_description: "Valid gender - create a female gender",
			field:            "Female",
			want_error:       false,
		},
		{
			test_description: "Valid gender - with uppercase letters",
			field:            "FEMALE",
			want_error:       false,
		},
		{
			test_description: "Valid gender - with lowercase letters",
			field:            "female",
			want_error:       false,
		},
		{
			test_description: "Invalid gender - empty field",
			field:            "    ",
			want_error:       true,
			expected_error:   failure.FieldIsEmpty("gender"),
		},
		{
			test_description: "Invalid gender - gender does not belong to the range of allowed values",
			field:            "Plant",
			want_error:       true,
			expected_error:   failure.FieldNotRangeValues("gender", []string{"female", "male"}),
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.test_description, func(t *testing.T) {
			gender, err := NewGender(tt.field)

			// Arrange
			if tt.want_error {
				assert.Nil(t, gender)
				assert.Equal(t, tt.expected_error, err)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Gender{}, *gender)
			}
		})
	}
}
