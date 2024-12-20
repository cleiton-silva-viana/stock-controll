package cpf

import (
	"fmt"
	"stock-controll/test/unitary"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCpfNoError(t *testing.T) {
	// Arrange
	c := "177.868.886-97"

	// Act
	result, err := New(c)

	// Assert
	assert.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, c, result.CPF())
}

func TestNewCpfWithError(t *testing.T) {
	tests := []unitary.TestField[CPF]{
		{
			Description: "cpf empty",
			Handler:     func(c *CPF) { c.cpf = "                " },
		},
		{
			Description: "cpf is short",
			Handler:     func(c *CPF) { c.cpf = "177.868-97" },
		},
		{
			Description: "cpf is long",
			Handler:     func(c *CPF) { c.cpf = "177.868.886-978" },
		},
		{
			Description: "cpf with invalid format",
			Handler:     func(c *CPF) { c.cpf = "167.868.886.97" },
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Test New function of CPF package, error expected because %s", tt.Description), func(t *testing.T) {
			// Arrange
			var c CPF
			tt.Handler(&c)

			// Act
			result, err := New(c.cpf)

			// Assert
			assert.Nil(t, result)
			assert.Error(t, err)
		})
	}
}
