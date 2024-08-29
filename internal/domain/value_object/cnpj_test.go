package valueobject

import (
	"stock-controll/internal/domain/failure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCPNJ(t *testing.T) {
	// Arrange
	tests := []test{
		{
			test_description: "CNPJ valid - correct requierements",
			field:            "11.222.333/0001-81",
			want_error:       false,
		},
		{
			test_description: "CPNJ invalid - format invalid",
			field:            "00.000.000.0001.22",
			want_error:       true,
			expected_error:   failure.FieldWithInvalidFormat("cnpj", "XX.XXX.XXX/0001-XX"),
		},
		{
			test_description: "CPNJ invalid - contains special characters invalids",
			field:            "00 000 000 0001 22",
			want_error:       true,
			expected_error:   failure.FieldWithInvalidFormat("cnpj", "XX.XXX.XXX/0001-XX"),
		},
		{
			test_description: "CPNJ invalid - digits quantity minor than expected",
			field:            "00.000.000/0001-2",
			want_error:       true,
			expected_error:   failure.FieldWithInvalidFormat("cnpj", "XX.XXX.XXX/0001-XX"),
		},
		{
			test_description: "CPNJ invalid - digits quantity greater than expected",
			field:            "00.000.000/0001-222",
			want_error:       true,
			expected_error:   failure.FieldWithInvalidFormat("cnpj", "XX.XXX.XXX/0001-XX"),
		},
		{
			test_description: "CPNJ invalid - contains checker digits inconsistents",
			field:            "12.345.678/0001-99",
			want_error:       true,
			expected_error:   failure.CPNJWithInvalidCheckerDigits,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.test_description, func(t *testing.T) {
			cnpj, err := NewCNPJ(tt.field)

			// Assert
			if tt.want_error {
				assert.Nil(t, cnpj)
				assert.Equal(t, tt.expected_error, err)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, CNPJ{}, *cnpj)
			}

		})
	}
}
