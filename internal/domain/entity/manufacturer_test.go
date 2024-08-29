package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewManufacturer(t *testing.T) {
	// Arrange
	type manufacturer struct {
		name     string
		cnpj     string
		category string
	}

	tests := []test[manufacturer]{
		{
			description: "Valid manufacturer - all fields are valids, must return a manufacturer entity",
			fields: manufacturer{
				name:     "pepsico",
				cnpj:     "24.341.272/0001-13",
				category: "foods",
			},
			wantError:    false,
		},
		{
			description: "Invalid manufacturer - field 'name' invalid, 1 error espected",
			fields: manufacturer{
				name:     "#@$%!",
				cnpj:     "24.341.272/0001-13",
				category: "foods",
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid manufacturer - field 'cnpj' invalid, 1 error espected",
			fields: manufacturer{
				name:     "coca cola",
				cnpj:     "24.341.272/0001-14",
				category: "drinks",
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid manufacturer - field 'category' invalid, 1 error espected",
			fields: manufacturer{
				name:     "bauduco",
				cnpj:     "24.341.272/0001-13",
				category: " *_* ",
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid manufacturer - field 'cnpj' & 'category' invalid, 2 errors espected",
			fields: manufacturer{
				name:     "coca cola",
				cnpj:     "24.341.272/0001-12",
				category: " =) ",
			},
			wantError:              true,
			errorQuantityExpected: 2,
		},
		{
			description: "Invalid manufacturer - al fields is invalid, 3 errors espected",
			fields: manufacturer{
				name:     "coca-cola",
				cnpj:     "24.341.272/0001-12",
				category: " =) ",
			},
			wantError:              true,
			errorQuantityExpected: 3,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			manufacturerBuilder := NewManufacturerBuilder()
			manufacturer, err := manufacturerBuilder.
				Name(tt.fields.name).
				Category(tt.fields.category).
				CNPJ(tt.fields.cnpj).
				Build()

			//Assert
			if tt.wantError {
				assert.Nil(t, manufacturer)
				assert.Len(t, err, tt.errorQuantityExpected)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Manufacturer{}, *manufacturer)
			}
		})
	}
}
