package list

import (
	"testing"

	"stock-controll/internal/domain/services/error/field"

	"github.com/stretchr/testify/assert"
)

var err1 = field.FieldError{
	FieldName:    "keyword",
	CodeError:    "ERR_FIELD_WIDTH_SPECIAL_CHARS",
	InvalidValue: "$$$$",
}

var err2 = field.FieldError{
	FieldName:    "keyword",
	CodeError:    "ERR_FIELD_WIDTH_SPECIAL_CHARS",
	InvalidValue: "____",
}

// Neste teste vou verificar se dois itens de uma lista
// com o mesmo tipo de erro não geram duplicidade nos CodeErrors
func TestNewListErrorNoError(t *testing.T) {
	// Arrange
	fieldName := "keywords"
	err := Error(fieldName)

	// Act
	err.AddError(&err1)
	err.AddError(&err2)

	// Assert
	assert.Len(t, err.ErrorIn, 2)
	assert.Len(t, err.CodeErrors, 1)
}
