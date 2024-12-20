package cartproduct

import (
	"testing"

	product "stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/services/error/entity"

	productMock "stock-controll/test/mock/entity/product"
	promotionMock "stock-controll/test/mock/entity/promotion"
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var item = productMock.Mock()

func TestNewNoError(t *testing.T) {
	// Act
	instance, err := New(item)

	// Assert
	assert.NoError(t, err)
	require.NotNil(t, instance)
	require.NotNil(t, instance.appliedCoupons)
	require.NotNil(t, instance.appliedPromotions)
	require.NotNil(t, instance.IProduct)
	require.NotNil(t, instance.TotalAmount)
	assert.Equal(t, instance.PurchasedQuantity(), 1)
	assert.Equal(t, instance.FinalTotal(), instance.Price())
	assert.Equal(t, instance.TotalDiscountApplied(), 0)
	assert.Equal(t, instance.FinalTotal(), instance.Price())
	assert.Empty(t, instance.AppliedDiscount())
}

func TestNewWithError(t *testing.T) {
	testCases := []unitary.TestDependence[product.IProduct]{
		{
			Description: "product is empty",
			Dependency:  &product.Product{},
		},
		{
			Description: "dependency is nil",
			Dependency:  nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {

			// Act
			instance, err := New(test.Dependency)

			// Assert
			assert.Nil(t, instance)
			require.Error(t, err)
			assert.ErrorAs(t, err, &entity.EntityError{})
		})
	}
}

var promo = promotionMock.Mock()

func TestAddPromotionNoError(t *testing.T) {

}

func TestAddPromotionWithError(t *testing.T) {}

func TestAddCouponNoError(t *testing.T) {}

func TestAddCouponWithError(t *testing.T) {}

func TestRemoveCouponNoError(t *testing.T) {}

func TestRemoveCouponWithError(t *testing.T) {}
