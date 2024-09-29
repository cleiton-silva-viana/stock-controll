package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewProduct(t *testing.T) {
	// Arrange
	type product struct {
		name           string
		description    string
		barcode        string
		brandID        int
		categoryID     int
		manufacturerID int
	}

	tests := []test[product]{
		{
			description: "Valid product - must return a product entity",
			fields: product{
				name:           "cheetos",
				description:    "A cookie cheezze",
				barcode:        "265463113564",
				brandID:        1,
				categoryID:     1,
				manufacturerID: 1,
			},
			wantError:    false,
		},
		{
			description: "Invalid product - invalid product name, 1 error expected",
			fields: product{
				name:           "Wireless Noise-Canceling Headphones",
				description:    "The SmartTech Pro Watch is the ultimate companion for you.",
				barcode:        "4006381333931",
				brandID:        1,
				categoryID:     1,
				manufacturerID: 1,
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid product - invalid product description, 1 error expected",
			fields: product{
				name:           "Eco-Friendly Water Bottle",
				description:    "                                 ",
				barcode:        "4006381333931",
				brandID:        1,
				categoryID:     1,
				manufacturerID: 1,
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid product - invalid product barcode, 1 error expected",
			fields: product{
				name:           "Eco-Friendly Water Bottle",
				description:    "Experience immersive sound with our Wireless Noise-Canceling Headphones",
				barcode:        "4006381333931a",
				brandID:        1,
				categoryID:     1,
				manufacturerID: 1,
			},
			wantError:              true,
			errorQuantityExpected: 1,
		},
		{
			description: "Invalid product - invalid manufacturer id & brand id, 2 errors are expected",
			fields: product{
				name:           "Eco-Friendly Water Bottle",
				description:    "Experience immersive sound with our Wireless Noise-Canceling Headphones",
				barcode:        "4006381333931",
				brandID:        -1,
				categoryID:     2,
				manufacturerID: 1000000,
			},
			wantError:              true,
			errorQuantityExpected: 2,
		},
		{
			description: "Invalid product - invalid manufacturer, category and brand id, 3 errors are expected",
			fields: product{
				name:           "Eco-Friendly Water Bottle",
				description:    "Experience immersive sound with our Wireless Noise-Canceling Headphones",
				barcode:        "4006381333931",
				brandID:        -1111,
				categoryID:     -200000,
				manufacturerID: 10000000,
			},
			wantError:              true,
			errorQuantityExpected: 3,
		},
		{
			description: "Invalid product - invalid name, description and barcode, 3 errors are expected",
			fields: product{
				name:           "Eco-Friendly Water Bottle$$$$",
				description:    "Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones ",
				barcode:        "400638133393&8",
				brandID:        1021,
				categoryID:     20,
				manufacturerID: 54,
			},
			wantError:              true,
			errorQuantityExpected: 3,
		},
		{
			description: "Invalid product - all fields are invalids",
			fields: product{
				name:           "Eco-Friendly Water Bottle$$$$",
				description:    "Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones Experience immersive sound with our Wireless Noise-Canceling Headphones ",
				barcode:        "400638133393&8",
				brandID:        11111,
				categoryID:     -40,
				manufacturerID: 3458641,
			},
			wantError:              true,
			errorQuantityExpected: 6,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			productBuilder := NewProductBuilder()
			product, err := productBuilder.Name(tt.fields.name).
				Description(tt.fields.description).
				Barcode(tt.fields.barcode).
				CategoryID(tt.fields.categoryID).
				BrandID(tt.fields.brandID).
				ManufacturerID(tt.fields.manufacturerID).
				Build()

			// Assert
			if tt.wantError {
				assert.Nil(t, product)
				assert.Len(t, err.ErrList, tt.errorQuantityExpected)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Product{}, *product)
			}
		})
	}
}
