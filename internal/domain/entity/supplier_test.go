package entity

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func Test_NewSupplier(t *testing.T) {
	// Arrange
	type supplier struct {
		ID             int
		name           string
		CNPJ           string
		contactBilling struct {
			email string
			phone string
		}
		contactPurschase struct {
			email string
			phone string
		}
		// paymentsConditions string
		sallers []SellerSupplier
	}

	tests := []test[supplier]{
		{
			description: "Valid supplier - must retun a supply entity",
			fields: supplier{
				ID:   1,
				name: "pepsiCo",
				CNPJ: "01.535.423/0001-79",
				contactPurschase: struct {
					email string
					phone string
				}{
					email: faker.New().Internet().Email(),
					phone: "(11) 9311-4887",
				},
				contactBilling: struct {
					email string
					phone string
				}{
					email: faker.New().Internet().Email(),
					phone: "(11) 9333-4887",
				},
				sallers: nil,
			},
			wantError: false,
		},
		{
			description: "Invalid supplier - all fields are invalids",
			fields: supplier{
				ID:   -1,
				name: "pepsiCo#",
				CNPJ: "01.535.423/0001-19",
				contactPurschase: struct {
					email string
					phone string
				}{
					email: "pespico@purschase",
					phone: "(11) 9311-488447",
				},
				contactBilling: struct {
					email string
					phone string
				}{
					email: "@purschase.com",
					phone: "1549333-4887",
				},
				sallers: nil,
			},
			wantError: true,
			errorQuantityExpected: 6,
		},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			builder := NewSupplierBuilder()
			supplier, err := builder.
				Indentity(tt.fields.name, tt.fields.CNPJ).
				ContactPurschase(1, tt.fields.contactPurschase.email, tt.fields.contactPurschase.phone).
				ContactBilling(1, tt.fields.contactBilling.email, tt.fields.contactBilling.phone).
				Sellers(tt.fields.sallers).
				Build()

			// Assert
			if tt.wantError {
				assert.Nil(t, supplier)
			} else {
				assert.Nil(t, err)
				assert.IsType(t, Supplier{}, *supplier)
			}
		})
	}
}
