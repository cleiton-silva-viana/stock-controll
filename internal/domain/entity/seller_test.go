package entity

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func Test_NewSellerSupplier(t *testing.T) {
	type seller struct {
		firstName  string
		lastName   string
		supplierID int
		email      string
		phone      string
	}

	// Arrange
	tests := []test[seller]{
		{
			description: "Valid seller - all fields are valids, must return a seller entity",
			fields: seller{
				firstName:  "alan",
				lastName:   "Wake",
				supplierID: 2,
				email:      faker.New().Internet().Email(),
				phone: "(21) 93189-2233",
			},
			wantError:    false,
		},
		{
			description: "Invalid seller - names are invalid format",
			fields: seller{
				firstName:  "$alan$",
				lastName:   "      ",
				supplierID: 2,
				email:      faker.New().Internet().Email(),
				phone: "(21) 93189-2233",
			},
			wantError:    true,
			errorQuantityExpected: 2,
		},
		{
			description: "Invalid seller - contacts are invalid format and supplier ID is incorrect",
			fields: seller{
				firstName:  "Jonathan",
				lastName:   "Jowstar",
				supplierID: 20597949,
				email:      "jonatha.email@xyz",
				phone: "(21) 93189-2233xyz",
			},
			wantError:    true,
			errorQuantityExpected: 3,
		},
		{
			description: "Invalid seller - all fields are invalid",
			fields: seller{
				firstName:  "Jonathan163",
				lastName:   "Jowstar#r5",
				supplierID: 20597949,
				email:      "jonatha.email@xyz",
				phone: "(21) 93189-2233xyz",
			},
			wantError:    true,
			errorQuantityExpected: 5,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			sellerBuilder := NewSellerSupplierBuilder()
			seller, err := sellerBuilder.Name(tt.fields.firstName, tt.fields.lastName).
										SupplierID(tt.fields.supplierID).
										Contact(tt.fields.email, tt.fields.phone).
										Build()

			//Assert
			if tt.wantError {
				assert.Nil(t, seller)
				assert.Len(t, err, tt.errorQuantityExpected)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, SellerSupplier{}, *seller)
			}
		})
	}
}
