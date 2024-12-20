package sale

import (
	"testing"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/valueobject/uuid"
	
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var data = Config{
	SellerUUID:    uuid.New().String(),
	ClientUUID:    uuid.New().String(),
	Discount:      Discount{}, // testar
	Products:      nil,
	PaymentMethod: Cash, // testar
	Status:        InProcessing,
}

func TestNewNoError(t *testing.T) {
	// Arrange
	dataCopy := data

	// Act
	instance, err := New(dataCopy)

	// Assert
	assert.NoError(t, err)
	require.NotNil(t, instance)
	assert.Equal(t, data.SellerUUID, instance.SellerUUID())
	assert.Equal(t, data.ClientUUID, instance.ClientUUID())
	require.NotEmpty(t, instance.Discount())
	assert.Equal(t, data.Discount, instance.Discount())
	assert.Equal(t, string(data.PaymentMethod), instance.PaymentMethod())
	require.NotNil(t, instance.Products())
	assert.Equal(t, data.Products, instance.Products())
	assert.Equal(t, string(data.Status), instance.Status())
}

func TestNewWithError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Config]{
		{
			Description: "seller uuid is invalid",
			Handler:     func(c *Config) { c.SellerUUID = "" },
		},
		{
			Description: "client uuid is invalid",
			Handler:     func(c *Config) { c.ClientUUID = "a65ewfd65ewqf4ew6f4ew65f4ew6few4f" },
		},
		{
			Description: "Products is nil",
			Handler:     func(c *Config) { c.Products = nil },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			instance, err := New(dataCopy)

			// Assert
			assert.Nil(t, instance)
			require.Error(t, err)
			assert.ErrorIs(t, err, &entity.EntityError{})
		})
	}
}
