package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type manufacturer struct {
	name                    string
	cnpj                    string
	category                string
	email                   string
	phone                   string
	error_quantity_expected int
}

func Test_NewManufacturer_MustReturnAManufacturer(t *testing.T) {
	// Arrange
	name := "pepsico"
	cnpj := "24.341.272/0001-13"
	category := "foods"
	email := "contact@pepsico.com"
	phone := "(21) 40028922"

	// Act
	manufacutrerBuilder := NewManufacturerBuilder()
	manufacturer, err := manufacutrerBuilder.
		SetName(name).
		SetCategory(category).
		SetCNPJ(cnpj).
		SetContact(email, phone).
		Build()

	// Assert
	assert.Nil(t, err)
	assert.IsType(t, Manufacturer{}, *manufacturer)
}

func Test_NewManufacturer_WithError(t *testing.T) {
	// Arrange
	var tests = []struct {
		descritipn string
		manufacturer
	}{
		{
			descritipn: "manufacuterer with 2 errors",
			manufacturer: manufacturer{
				name:                    "amazon",
				cnpj:                    "12.345.678/0001-01",
				category:                "all",
				email:                   "amazon.contact.com",
				phone:                   "(21) 40028922",
				error_quantity_expected: 2,
			},
		},
		{
			descritipn: "manufacuterer with 3 errors",
			manufacturer: manufacturer{
				name:                    "$amazon$",
				cnpj:                    "12.345.678/0001-01",
				category:                "",
				email:                   "amazon@contact.com",
				phone:                   "(21) 400289228",
				error_quantity_expected: 3,
			},
		},
	}

	// Act
	for _, test := range tests {
		t.Run(test.descritipn, func(t *testing.T) {
			manufacutrerBuilder := NewManufacturerBuilder()
			manufacturer, err := manufacutrerBuilder.
				SetName(test.name).
				SetCategory(test.category).
				SetCNPJ(test.cnpj).
				SetContact(test.email, test.phone).
				Build()
			
			// Assert
			assert.Nil(t, manufacturer)
			assert.Len(t, err, test.error_quantity_expected)
		})
	}
}
