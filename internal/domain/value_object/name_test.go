package valueobject

import (
	"stock-controll/internal/domain/failure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewName(t *testing.T) {
	//Arrange
	type name struct {
		test
		withDigits       bool
		withSpecialChars bool
		maxLength        int
		minLength        int
	}

	t.Run("checking digits setter behavior in name", func(t *testing.T) {
		// Arrange
		tests := []name{
			{
				test: test{
					field:            "Joe v2",
					test_description: "Valid name - with digits permitted",
					want_error:       false,
					
				},
				withDigits:       true,
				withSpecialChars: false,
				maxLength:        50,
				minLength:        3,
			},
			{
				test: test{
					field:            "etec v2",
					test_description: "Invalid name - field contains digits (not permitted)",
					want_error:       true,
					expected_error:   failure.FieldWithNumber("name"),
				},
				withDigits:       false,
				withSpecialChars: false,
				maxLength:        50,
				minLength:        3,
			},
			{
				test: test{
					field:            "Ang*L",
					test_description: "Valid name - with characters specials permiited",
					want_error:       false,
				},
				withDigits:       false,
				withSpecialChars: true,
				maxLength:        50,
				minLength:        3,
			},
			{
				test: test{
					field:            "Ang*L",
					test_description: "Invalid name - contains characters specials not permitted",
					want_error:       true,
					expected_error: failure.FieldWithSpecialChars("name"),
				},
				withDigits:       false,
				withSpecialChars: false,
				maxLength:        50,
				minLength:        3,
			},
			{
				test: test{
					field:            "Ana",
					test_description: "Valid name - length is equal to min length permitted",
					want_error:       false,
				},
				withDigits:       false,
				withSpecialChars: false,
				maxLength:        12,
				minLength:        3,
			},
			{
				test: test{
					field:            "MATHEUS",
					test_description: "Valid name - length is between minimum and maximum length allowed",
					want_error:       false,
				},
				withDigits:       false,
				withSpecialChars: false,
				maxLength:        12,
				minLength:        3,
			},
			{
				test: test{
					field:            "Ana",
					test_description: "Invalid name - length is less than the minimun length allowed",
					want_error:       true,
					expected_error: failure.FieldIsShort("name", 6),
				},
				withDigits:       false,
				withSpecialChars: false,
				maxLength:        12,
				minLength:        6,
			},
			{
				test: test{
					field:            "carlos",
					test_description: "Valid name - length is equal than to the maximun length allowed",
					want_error:       false,
				},
				withDigits:       false,
				withSpecialChars: false,
				maxLength:        6,
				minLength:        3,
			},
			{
				test: test{
					field:            "Bomsounaro",
					test_description: "Invalid name - length is longer than to the maximun length permitted",
					want_error:       true,
					expected_error: failure.FieldIsLong("name", 6),
				},
				withDigits:       false,
				withSpecialChars: false,
				maxLength:        6,
				minLength:        3,
			},
	
		}

		// Act
		for _, tt := range tests {
			nameBuilder := NewNameBuilder()
			name, err := nameBuilder.Field("name", tt.field).
				Length(tt.minLength, tt.maxLength).
				Digts(tt.withDigits).
				SpecialCharacters(tt.withSpecialChars).
				Build()
			// Assert
			if tt.want_error {
				assert.Nil(t, name)
				assert.Equal(t, tt.expected_error, err)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Name{}, *name)
			}

		}
	})
}
