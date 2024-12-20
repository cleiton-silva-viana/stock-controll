package cart

import (
	"testing"
	"time"

	cartproduct "stock-controll/internal/domain/entity/cart_product"
	"stock-controll/internal/domain/entity/coupon"
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/valueobject/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNoError(t *testing.T) {
	// Arrange
	var userUUID = uuid.New()

	// Act
	cart, err := New(userUUID)

	// Assert
	assert.NoError(t, err)
	require.NotNil(t, cart)
	assert.NoError(t, uuid.IsValid("cart_uuid", cart.UUID()))
	assert.Equal(t, cart.UserUUID(), userUUID)
	assert.Equal(t, cart.Status(), Active)

	assert.Equal(t, time.Now().Minute(), cart.CreatedAt()) // Melhorar
	assert.Equal(t, time.Now().Minute(), cart.UpdatedAt()) // Melhorar

	require.NotNil(t, cart.appliedCoupons)
	assert.Empty(t, cart.AppliedCoupons())

	require.NotNil(t, cart.products)
	assert.Empty(t, cart.ProductsInCart())

	require.NotNil(t, cart.total)
	assert.Equal(t, cart.Subtotal(), 0)
	assert.Equal(t, cart.Discounts(), 0)
	assert.Equal(t, cart.TotalFinal(), 0)
}

func TestNewWithError(t *testing.T) {
	// Arrange
	userUUID := "cfew6few4fg96ew4gwe6g4ewr65g4ert8"

	// Act
	cart, err := New(userUUID)

	// Arrange
	assert.Nil(t, cart)
	require.Error(t, err)
	assert.ErrorAs(t, err, &entity.EntityError{})
	assert.Len(t, err.(*entity.EntityError).Errors(), 1)
}

var data = Cart{
	cartUUID:       uuid.New(),
	userUUID:       uuid.New(),
	createdAt:      time.Now(),
	updatedAt:      time.Now(),
	status:         Active,
	products:       make(map[string]cartproduct.ProductCart),
	appliedCoupons: make([]coupon.ICoupon, 0),
	total: struct {
		subtotal   float64
		discounts  float64
		totalFinal float64
	}{
		subtotal:   0,
		discounts:  0,
		totalFinal: 0,
	},
}

func TestAddProductInCartNoError(t *testing.T) {
	// Arrange
	copy := data
	product := product.Product{}

	// Act
	err := copy.AddProductInCart(&product)

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, copy.ProductsInCart(), product)
}
