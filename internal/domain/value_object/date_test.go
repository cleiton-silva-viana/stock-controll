package value_object

import (
	"stock-controll/internal/domain/failure"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewDate_WithValidDates(t *testing.T) {
	// Arrange
	date := "1999-01-05"
	expected_type := Date{}

	// Act
	result, err := NewDate(date)

	// Assert
	assert.Nil(t, err)
	assert.IsType(t, &expected_type, result)
}

func Test_NewDate_WithInvalidDateFormat(t *testing.T) {
	// Arrange
	date := "01-05-1999"

	// Act
	result, err := NewDate(date)

	// Assert
	assert.Nil(t, result)
	assert.ErrorIs(t, err, failure.DateWithInvalidFormat)
}

func Test_IsOlderThan_AreAdults(t *testing.T) {
	// Arrange
	now := time.Now()
	minimunAge := 18
	tests := []struct {
		description string
		age         int
		throwError  bool
	}{
		{description: "user age and minimal age are equals", age: 18, throwError: true},
		{description: "user age is older than minimal age", age: 24, throwError: false},
	}

	// Act
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			birthDate := now.AddDate(-test.age, 0, 0)
			formattedDate := birthDate.Format("2006-01-02")

			date, _ := NewDate(formattedDate)
			result, err := date.IsOlderThan(minimunAge, test.throwError)

			// Assert
			assert.Nil(t, err)
			assert.True(t, result)
		})
	}
}

func Test_IsOlderThan_AreMinor(t *testing.T) {
	// Arrange
	userAge := 16
	now := time.Now()
	birthDate := now.AddDate(-userAge, 0, 0)
	formattedDate := birthDate.Format("2006-01-02")
	minimunAge := 18

	// Act
	date, _ := NewDate(formattedDate)
	result, err := date.IsOlderThan(minimunAge, true)

	// Assert
	assert.ErrorIs(t, err, failure.InsufficientAge)
	assert.False(t, result)
}
