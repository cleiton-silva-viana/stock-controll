package validationerrors

import (
	"errors"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/test/unitary"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var fieldError = &validate.FieldError{
	FieldName: "name",
	CodeError: "ERR_FILE_NOT_FOUND",
}

func TestAddErrorCodeNoError(t *testing.T) {
	t.Run("add a error valid", func(t *testing.T) {
		// Arrange
		var err = New("user")

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
		err := New("user").AddValidationError(fieldError)

		// Assert
		require.False(t, err.HasError())
		require.Len(t, err.Errors(), 0)
	})

	t.Run("add a generic erro, the erro not implements FieldError", func(t *testing.T) {
		// Arrange
		genericError := errors.New("a generic error")

		// Act
		err := New("User").AddValidationError(genericError)

		// Assert
		require.False(t, err.HasError())
		require.Len(t, err.Errors(), 0)
	})

	testCases := []unitary.TestField[validate.FieldError]{
		{
			Description: "add field error with empty field name",
			Handler:         func(fe *validate.FieldError) { fe.FieldName = "    " },
		},
		{
			Description: "add field error with empty invalid error code format", // The format for field error code string should be formatted as follows: ERR_
			Handler:         func(fe *validate.FieldError) { fe.CodeError = "ErrFieldCannotBeEmpty" },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			fieldErrCopy := fieldError
			test.Handler(fieldErrCopy)

			// Act
			err := New("User").AddValidationError(fieldErrCopy)

			// Assert
			require.False(t, err.HasError())
			require.Len(t, err.Errors(), 0)
		})
	}
}
