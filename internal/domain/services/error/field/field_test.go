package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = &FieldError{
	FieldName:    "image",
	CodeError:    "ERR_INVALID_FILE_TYPE",
	InvalidValue: "pdf",
}

func TestErrorNoError(t *testing.T) {
	// Act
	err := Error(example.FieldName, example.CodeError, example.InvalidValue)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, err.FieldName, example.FieldName)
	assert.Equal(t, err.CodeError, example.CodeError)
	assert.Equal(t, err.InvalidValue, example.InvalidValue)
}
