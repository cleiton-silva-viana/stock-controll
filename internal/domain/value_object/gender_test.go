package value_object

import (
	"testing"
	
	"stock-controll/internal/domain/failure"

	"github.com/stretchr/testify/assert"
)

func Test_CreateGender(t *testing.T) {
	// Arrange
	var tests = []struct{
		name string
		gender string
	} {
		{name: "create a gender male", gender: "Male"},
		{name: "create a gender female", gender: "Female"},
	}

	// Act
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {	
			result, err := NewGender(test.gender)
			
			// Assert
			assert.Nil(t, err)
			assert.IsType(t, &Gender{}, result)
		})
	}
}

func Test_CreateGender_InvalidGenders(t *testing.T) {
	// Arrange
	gender := "plants"

	// Act
	result, err := NewGender(gender)

	// Assert
	assert.Nil(t, result)
	assert.ErrorIs(t, err, failure.GenderInvalid)
}
