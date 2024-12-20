package name

import (
	"stock-controll/test/unitary"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Setup() *Name {
	return &Name{
		firstName: "Vladmir",
		lastName:  "Putin",
	}
}

func TestNewNameNoError(t *testing.T) {
	tests := []unitary.TestField[Name]{
		{
			Description: "first name with min allowed",
			Handler:     func(n *Name) { n.firstName = strings.Repeat("a", minLength) },
		},
		{
			Description: "first name with max allowed",
			Handler:     func(n *Name) { n.firstName = strings.Repeat("a", maxLength) },
		},
		{
			Description: "compost first name",
			Handler:     func(n *Name) { n.firstName = "Soleiman Rameneim" },
		},
		{
			Description: "last name is empty",
			Handler:     func(n *Name) { n.lastName = "" },
		},
		{
			Description: "compost last name",
			Handler:     func(n *Name) { n.lastName = "trien tie" },
		},
		{
			Description: "last name width min length allowed",
			Handler:     func(n *Name) { n.lastName = strings.Repeat("a", minLength) },
		},
		{
			Description: "last name width max length allowed",
			Handler:     func(n *Name) { n.lastName = strings.Repeat("a", maxLength) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			n := Setup()
			tt.Handler(n)

			// Act
			result, err := New(n.firstName, n.lastName)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, result)
			assert.Equal(t, n.firstName, result.FirstName())
			assert.Equal(t, n.lastName, result.LastName())
		})
	}
}

func TestNewNameWithError(t *testing.T) {
	tests := []unitary.TestField[Name]{
		{
			Description: "first name is empty",
			Handler:     func(n *Name) { n.lastName = "         " },
		},
		{
			Description: "first name is less than min allowed",
			Handler:     func(n *Name) { n.firstName = strings.Repeat("a", minLength-1) },
		},
		{
			Description: "first name is greater than max allowed",
			Handler:     func(n *Name) { n.firstName = strings.Repeat("a", maxLength+1) },
		},
		{
			Description: "first name contains number",
			Handler:     func(n *Name) { n.firstName = "thrird 3" },
		},
		{
			Description: "first name with special chars",
			Handler:     func(n *Name) { n.lastName = "$p$cial" },
		},
		{
			Description: "first & last name are invalids",
			Handler:     func(n *Name) { n.firstName = "     "; n.lastName = "123" },
		},
		{
			Description: "last name length is less than minimum allowed",
			Handler:     func(n *Name) { n.lastName = strings.Repeat("a", minLength-1) },
		},
		{
			Description: "last name length is greater than maximum allowed",
			Handler:     func(n *Name) { n.lastName = strings.Repeat("a", maxLength+1) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			n := Setup()
			tt.Handler(n)

			// Act
			result, err := New(n.firstName, n.lastName)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, result)
			assert.Equal(t, n.firstName, result.FirstName())
			assert.Equal(t, n.lastName, result.LastName())
		})
	}
}
