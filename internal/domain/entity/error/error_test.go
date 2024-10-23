package validationerrors

import (
	"stock-controll/internal/domain/validation"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


var fieldError = &validation.FieldError{
	FieldName: "name",
	CodeErrors: []string{"ERR_FILE_NOT_FOUND"},
}

func Test_AddErrorCode_NoError(t *testing.T) {
	t.Run("add a error valid", func(t *testing.T) {	
		// Arrange
		var err = NewValidationError("user")

		// Act
		err.AddValidationError(fieldError)
		
		// Assert
		require.True(t, err.HasError())
		assert.Len(t, err.GetErrors(), 1)
	})

	t.Run("add a nil reference for field error", func(t *testing.T) {
		// Arrange
		var err = NewValidationError("user")
		fieldError = nil

		// Act
		err.AddValidationError(fieldError)

		// Assert
		require.Len(t, err.GetErrors(), 0)
	})
}
