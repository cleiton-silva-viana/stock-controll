package valueobject

import (
	"strconv"
	"testing"
	"time"

	"stock-controll/internal/domain/failure"

	"github.com/stretchr/testify/assert"
)

func Test_NewDate(t *testing.T) {
	// Arrange
	tests := []test{
		{
			test_description: "Valid date - format is valid",
			field:            "1999-01-05",
			want_error:       false,
		},
		{
			test_description: "Invalid date - format invalid",
			field:            "01-05-1999",
			want_error:       true,
			expected_error:   failure.FieldWithInvalidFormat("date", "yyyy-mm-dd"),
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.test_description, func(t *testing.T) {
			date, err := NewDate(tt.field)

			// Assert
			if tt.want_error {
				assert.Nil(t, date)
				assert.Equal(t, tt.expected_error, err)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Date{}, *date)
			}
		})
	}
}

func Test_IsOlderThan(t *testing.T) {
	// Arrange
	tests := []test{
		{
			test_description: "Age valid - are equals to minimal age",
			field:            "18",
			want_error:       false,
		},
		{
			test_description: "Age valid - is older than minimal age",
			field:            "24",
			want_error:       false,
		},
		{
			test_description: "Age invalid - is minor than minimal age",
			field:            "16",
			want_error:       true,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.test_description, func(t *testing.T) {
			const MINIMUN_AGE = 18
			now := time.Now()
			age, _ := strconv.Atoi(tt.field)
			birthDate := now.AddDate(-age, 0, 0)
			formattedDate := birthDate.Format("2006-01-02")

			date, _ := NewDate(formattedDate)
			isOlder := date.IsOlderThan(MINIMUN_AGE)

			// Assert
			assert.Equal(t, isOlder, !tt.want_error)
		})
	}
}
