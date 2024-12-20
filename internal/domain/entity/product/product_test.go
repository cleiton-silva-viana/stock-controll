package product

import (
	"strings"
	"testing"

	"stock-controll/internal/domain/entity/tag"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/valueobject/uuid"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var data = Config{
	Name:             "metal ice",
	Description:      "a metal with ice for your pratice",
	Barcode:          "12346631234988",
	BrandUUID:        uuid.New().String(),
	ManufacturerUUID: uuid.New().String(),
	CategoryUUID:     uuid.New().String(),
}

func TestNewNoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []unitary.TestField[Config]{
		{
			Description: "Product with minimum name length",
			Handler:     func(c *Config) { c.Name = strings.Repeat("a", minNameLength) },
		},
		{
			Description: "Product with maximum name length",
			Handler:     func(c *Config) { c.Name = strings.Repeat("b", maxNameLength) },
		},
		{
			Description: "Product with min description length",
			Handler:     func(c *Config) { c.Description = strings.Repeat("s", minDescriptionLength) },
		},
		{
			Description: "Product with maximum description length",
			Handler:     func(c *Config) { c.Description = strings.Repeat("z", maxDescriptionLength) },
		},
	}

	for _, test := range testsCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			result, err := New(dataCopy)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, result)
			assert.Equal(t, strings.ToLower(dataCopy.Name), strings.ToLower(result.Name()))
			assert.Equal(t, strings.ToLower(dataCopy.Description), strings.ToLower(result.Description()))
			assert.Equal(t, dataCopy.BrandUUID, result.BrandUUID())
			assert.Equal(t, dataCopy.CategoryUUID, result.CategoryUUID())
			assert.Equal(t, dataCopy.ManufacturerUUID, result.ManufacturerUUID())
		})
	}
}

func TestNewWithError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[Config]{
		{
			Description: "empty product name",
			Handler:     func(c *Config) { c.Name = "         " },
		},
		{
			Description: "product name is short than minimum allowed",
			Handler:     func(c *Config) { c.Name = strings.Repeat("a", minNameLength-1) },
		},
		{
			Description: "product name is greater than maximum allowed",
			Handler:     func(c *Config) { c.Name = strings.Repeat("a", maxNameLength+1) },
		},
		{
			Description: "product description is short than minimum allowed",
			Handler:     func(c *Config) { c.Description = strings.Repeat("d", minDescriptionLength-1) },
		},
		{
			Description: "product description is greater than maximum allowed",
			Handler:     func(c *Config) { c.Description = strings.Repeat("A", maxDescriptionLength+1) },
		},
		{
			Description: "barcode is empty",
			Handler:     func(c *Config) { c.Barcode = "         " },
		},
		{
			Description: "barcode contain letter",
			Handler:     func(c *Config) { c.Barcode = "1234565789a" },
		},
		{
			Description: "barcode contain special chars",
			Handler:     func(c *Config) { c.Barcode = "#$#@54556546R%$#" },
		},
		{
			Description: "invalid brand uuid",
			Handler:     func(c *Config) { c.BrandUUID = "0192589b-33df-7408/b68e-e95b7b1cf94e" },
		},
		{
			Description: "invalid manufacturer uuid",
			Handler:     func(c *Config) { c.ManufacturerUUID = "092589b-a7c7-79d3-842d-9318f5f45961" },
		},
		{
			Description: "invlaid category uuid",
			Handler:     func(c *Config) { c.CategoryUUID = "0192589c-174b-77ce-8395" },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			result, err := New(dataCopy)

			// Assert
			assert.Nil(t, result)
			require.Error(t, err)
			assert.ErrorAs(t, err, &entity.EntityError{})
		})
	}
}

var productInstance = Product{
	uuid:             *uuid.New(),
	name:             "cookier ice",
	description:      "a metal with ice for your pratice",
	barcode:          "12346631234988",
	brandUUID:        *uuid.New(),
	manufacturerUUID: *uuid.New(),
	categoryUUID:     *uuid.New(),
	cost:             2.99,
	price:            10,
	quantity:         39,
	tags:             nil,
	Image:            nil,
}

func TestSetCostNoError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Product]{
		{
			Description: "cost equal than minimum allowed",
			Handler:     func(p *Product) { p.cost = minCost },
		},
		{
			Description: "const equal than maximum allowed",
			Handler:     func(p *Product) { p.cost = maxCost },
		},
		{
			Description: "cost is mid range value allowed",
			Handler:     func(p *Product) { p.cost = 100.00 },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := productInstance
			test.Handler(&dataCopy)

			// Act
			err := productInstance.SetCost(dataCopy.cost)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, productInstance.Cost(), dataCopy.cost)
		})
	}
}

func TestSetCostWithError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Product]{
		{
			Description: "zero cost",
			Handler:     func(p *Product) { p.cost = 0 }, // TODO: Criar um erro específico para este cenário
		},
		{
			Description: "cost is negative",
			Handler:     func(p *Product) { p.cost = -10.00 }, // TODO: Criar erro específico para este cenário
		},
		{
			Description: "cost is less than min allowed",
			Handler:     func(p *Product) { p.cost = minCost - 0.1 },
		},
		{
			Description: "cost is greater than max allowed",
			Handler:     func(p *Product) { p.cost = maxCost + 0.1 },
		},
		{
			Description: "cost with many decimal places", // TODO: criar erro específico para este cenário
			Handler:     func(p *Product) { p.cost = 500.123 },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := productInstance
			test.Handler(&dataCopy)

			// Act
			err := productInstance.SetCost(dataCopy.cost)

			// Assert
			assert.Error(t, err)
			assert.NotEqual(t, dataCopy.cost, productInstance.Cost())
		})
	}
}

func TestSetPriceNoError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Product]{
		{
			Description: "price equal than min allowed",
			Handler:     func(p *Product) { p.price = minCost },
		},
		{
			Description: "price equal than max allowed",
			Handler:     func(p *Product) { p.price = maxCost },
		},
		{
			Description: "price is mid range value allowed",
			Handler:     func(p *Product) { p.price = 250.99 },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := productInstance
			test.Handler(&dataCopy)

			// Act
			err := productInstance.SetPrice(dataCopy.price)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, productInstance.Price(), dataCopy.price)
		})
	}
}

func TestSetPriceWithError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Product]{
		{
			Description: "zero price",
			Handler:     func(p *Product) { p.price = 0 }, // TODO: Criar um erro específico para este cenário
		},
		{
			Description: "price is negative",
			Handler:     func(p *Product) { p.price = -10.00 }, // TODO: Criar erro específico para este cenário
		},
		{
			Description: "price is less than min allowed",
			Handler:     func(p *Product) { p.price = minCost - 0.1 },
		},
		{
			Description: "price is greater than max allowed",
			Handler:     func(p *Product) { p.price = maxCost + 0.1 },
		},
		{
			Description: "price with many decimal places", // TODO: criar erro específico para este cenário
			Handler:     func(p *Product) { p.price = 500.123 },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := productInstance
			test.Handler(&dataCopy)

			// Act
			err := productInstance.SetPrice(dataCopy.price)

			// Assert
			assert.Error(t, err)
			assert.NotEqual(t, dataCopy.price, productInstance.Price())
		})
	}
}

func TestStateConsistencyAfterInvalidSet(t *testing.T) {
	dataCopy := productInstance
	// Arrange
	testCases := []unitary.Consistence{
		{
			T:            t,
			Description:  "test consistence of cost value",
			Getter:       func() interface{} { return dataCopy.Cost() },
			Setter:       func(value interface{}) error { return dataCopy.SetCost(value.(float64)) },
			InvalidValue: -0.01,
		},
		{
			T:            t,
			Description:  "test consistence of price value",
			Getter:       func() interface{} { return dataCopy.Price },
			Setter:       func(value interface{}) error { return dataCopy.SetPrice(value.(float64)) },
			InvalidValue: 1000000,
		},
		{
			T:            t,
			Description:  "test consistence of product tag",
			Getter:       func() interface{} { return dataCopy.Name },
			Setter:       func(value interface{}) error { return dataCopy.SetTag(value.(tag.Tag)) },
			InvalidValue: nil,
		},
		{
			T:            t,
			Description:  "test consistence quantity update",
			Getter:       func() interface{} { return dataCopy.Name },
			Setter:       func(value interface{}) error { return dataCopy.SetQuantity(value.(int)) },
			InvalidValue: -1,
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {

			// Act & Assert
			unitary.ConsistenceTest(test)
		})
	}
}
