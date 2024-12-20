package entity

import (
	"testing"

	"stock-controll/internal/domain/services/error/field"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var fieldError = &field.FieldError{
	FieldName: "name",
	CodeError: "ERR_FILE_NOT_FOUND",
}

func TestAddErrorCodeNoError(t *testing.T) {
	t.Run("add a error valid", func(t *testing.T) {
		// Arrange
		var err = Error("user")

		// Act
		err.AddValidationError(fieldError)

		// Assert
		require.True(t, err.HasError())
		assert.Len(t, err.Errors(), 1)
	})

	t.Run("add a nil reference for field error", func(t *testing.T) {
		// Arrange
		fieldError = nil

		// Act
		err := Error("user").AddValidationError(fieldError)

		// Assert
		require.False(t, err.HasError())
		require.Len(t, err.Errors(), 0)
	})
}
