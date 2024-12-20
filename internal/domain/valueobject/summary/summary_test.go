package summary

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"stock-controll/test/unitary"
	"testing"
)

func Setup() *Summary {
	return &Summary{
		discount: 10,
		subtotal: 100,
		total:    90,
	}
}

func TestNewSummaryNoError(t *testing.T) {
	tests := []unitary.TestField[Summary]{
		{
			Description: "subtotal & total with min value allowed",
			Handler: func(s *Summary) {
				s.discount = 0
				s.subtotal = 0
				s.total = 0
			},
		},
		{
			Description: "total with min value allowed",
			Handler: func(s *Summary) {
				s.discount = 100
				s.subtotal = 100
				s.total = 0
			},
		},
		{
			Description: "subtotal & total with max value allowed",
			Handler: func(s *Summary) {
				s.discount = 0
				s.subtotal = 10000.00
				s.total = 10000.00
			},
		},
		{
			Description: "subtotal & total with max value allowed",
			Handler: func(s *Summary) {
				s.discount = 10000.00
				s.subtotal = 10000.00
				s.total = 0
			},
		},
		{
			Description: "subtotal is greater to total",
			Handler: func(s *Summary) {
				s.discount = 10
				s.subtotal = 100
				s.total = 90
			},
		},
		{
			Description: "subtotal is equal to total",
			Handler: func(s *Summary) {
				s.total = s.subtotal
				s.discount = 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			s := Setup()
			tt.Handler(s)

			// Act
			result, err := New(s.subtotal, s.total)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, result)
			assert.Equal(t, s.subtotal, result.Subtotal())
			assert.Equal(t, s.total, result.Total())
			assert.Equal(t, s.discount, result.Discount())
		})
	}
}

func TestNewSummaryWithError(t *testing.T) {
	tests := []unitary.TestField[Summary]{
		{
			Description: "subtotal is less to total",
			Handler: func(s *Summary) {
				s.total = 10
				s.subtotal = 9
			},
		},
		{
			Description: "subtotal with negative value",
			Handler: func(s *Summary) {
				s.total = 0
				s.subtotal = -9
			},
		},
		{
			Description: "subtotal with negative value",
			Handler: func(s *Summary) {
				s.total = -1
				s.subtotal = 99
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			s := Setup()
			tt.Handler(s)

			// Act
			result, err := New(s.subtotal, s.total)

			// Assert
			assert.Nil(t, result)
			assert.Error(t, err)
		})
	}
}
