package value_object

import (
	"stock-controll/internal/domain/failure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCNPJ_MustReturnACNPJ(t *testing.T) {
	// Arrange
	cnpj := "11.222.333/0001-81"

	// Act
	result, err := NewCNPJ(cnpj)

	// Assert
	assert.Nil(t, err)
	assert.IsType(t, CNPJ{}, *result)
}

func Test_NewCNPJ_MustReturnErrorFormat(t *testing.T) {
	// Arrange
	tests := []struct {
		description    string
		cnpj           string
		error_expected error
	}{
		{description: "CPNJ with invalid quantity digits", cnpj: "00.000.000/0001-222", error_expected: failure.CPNJWithInvalidFormat},
		{description: "CPNJ with invalid format", cnpj: "00.000.000.0001.222", error_expected: failure.CPNJWithInvalidFormat},
		{description: "CPNJ with invalid chars", cnpj: "00 000 000 0001 222", error_expected: failure.CPNJWithInvalidFormat},
		{description: "CPNJ with invalid checker digits", cnpj: "12.345.678/0001-99", error_expected: failure.CPNJWithInvalidCheckerDigits},
	}

	// Act
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result, err := NewCNPJ(test.cnpj)

			// Assert
			assert.Nil(t, result)
			assert.ErrorIs(t, err, test.error_expected)
		})
	}
}
