package entity

import (
	"stock-controll/internal/domain/failure"
	"testing"

	"github.com/stretchr/testify/assert"
)

type product_datas struct {
	name           string
	description    string
	barcode        string
	brandID        int
	categoryID     int
	manufacturerID int
}

type tests_product_field struct {
	test_descrition string
	product_datas
	new_value_field string
	error_expected          error
	error_quantity_expected int
}

var product = product_datas{
	name:           "Coca Cola",
	description:    "Refrigerante de cola, 2 litros",
	barcode:        "7891234567890",
	brandID:        1,
	categoryID:     2,
	manufacturerID: 3,
}

func Test_NewProduct_MustReturnAProduct(t *testing.T) {
	// Arrange
	name := "cheetos"
	description := "A cookie cheezze"
	barcode := "265463113564"
	categoryID := 3
	manufacturerID := 1

	// Act
	productBuilder := NewProductBuilder()
	product, err := productBuilder.SetName(name).
		SetDescription(description).
		SetBarcode(barcode).
		SetCategoryID(categoryID).
		SetManufacturerID(manufacturerID).
		Build()

	// Assert
	assert.Nil(t, err)
	assert.IsType(t, Product{}, *product)
}

func Test_NewProduct_TestErrorsFieldName(t *testing.T) {
	// Arrange
	tests := []tests_product_field{
		{
			test_descrition: "name with longer than 25 characters",
			new_value_field: "Mindful Moments: Your Daily Oasis for Calm and Focus",
			product_datas:  product,
			error_expected: failure.NameIsLong("name", 25),
		},
		{
			test_descrition: "name with less than 2 characters",
			new_value_field: "a", 
			product_datas: product, 
			error_expected: failure.NameIsShort("name", 2),
		},
		{
			test_descrition: "empyt name",
			new_value_field: "", 
			product_datas: product, 
			error_expected: failure.NameIsEmpty("name"),
		},
		{
			test_descrition: "name with special chars",
			new_value_field: "Behavio@$", 
			product_datas: product, error_expected: 
			failure.NameWithInvalidChars("name"),
		},
	}

	// Act
	for _, test := range tests {
		t.Run(test.test_descrition, func(t *testing.T) {
			test.product_datas.name = test.new_value_field

			productBuilder := NewProductBuilder()
			product, err := productBuilder.SetName(test.name).
				SetDescription(test.description).
				SetBarcode(test.barcode).
				SetCategoryID(test.categoryID).
				SetManufacturerID(test.manufacturerID).
				Build()

			// Assert
			assert.Nil(t, product)
			assert.Equal(t, err[0].Error(), test.error_expected.Error())
		})
	}
}

func Test_NewProduct_TestErrorsFieldDescrition(t *testing.T) {
	// Arrange
	tests := []tests_product_field{
		{
			test_descrition: "descrition with longer than 100 characters",
			new_value_field: "Ultimate All-in-One Multi-Purpose Kitchen Appliance with Smart Technology for Effortless Cooking, Baking, and Food Preparation",
			product_datas:  product,
			error_expected: failure.NameIsLong("description", 100),
			error_quantity_expected: 1,
		},
		{
			test_descrition: "descrition with less than 10 characters",
			new_value_field: "Lamp", 
			product_datas: product, 
			error_expected: failure.NameIsShort("description", 10),
			error_quantity_expected: 1,
		},
		{
			test_descrition: "empyt descrition",
			new_value_field: "", 
			product_datas: product, 
			error_expected: failure.NameIsEmpty("description"),
			error_quantity_expected: 1,
		},
	}

	// Act
	for _, test := range tests {
		t.Run(test.test_descrition, func(t *testing.T) {
			test.product_datas.description = test.new_value_field

			productBuilder := NewProductBuilder()
			product, err := productBuilder.SetName(test.name).
				SetDescription(test.description).
				SetBarcode(test.barcode).
				SetCategoryID(test.categoryID).
				SetManufacturerID(test.manufacturerID).
				Build()

			// Assert
			assert.Nil(t, product)
			assert.Len(t, err, test.error_quantity_expected)
			assert.Equal(t, err[0], test.error_expected)
		})
	}
}
