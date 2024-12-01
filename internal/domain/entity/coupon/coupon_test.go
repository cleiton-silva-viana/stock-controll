package coupon

import (
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/test/unitary"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
	name string,
	minPurchaseAmount float64,
	discount discount.IDiscountStrategy,
	expirationDate time.Time,
	usageLimit CouponUseLimit,
	exclusivity CouponExclusivity,
	productsUUIDs []string,

*/

var data = Config{
	Name:              "PROMOTION",
	MinPurchaseAmount: 100.00,
	UsageLimit:        -1,
	Exclusivity:       "all_customers",
	ExpirationDate:    time.Now().Add(time.Hour * 24),
}

func TestNewnNoError(t *testing.T) {
	t.Parallel()

	testCases := []unitary.TestField[Config]{
		{
			Description: "create coupon with min name length allowed",
			Handler:     func(c *Config) { c.Name = strings.Repeat("a", minCouponNameLength) },
		},
		{
			Description: "create coupon with max name length allowed",
			Handler:     func(c *Config) { c.Name = strings.Repeat("b", maxCouponNameLength) },
		},
		{
			Description: "create coupon with letters and numbers in name",
			Handler:     func(c *Config) { c.Name = "PROMOTION2024" },
		},
		{
			Description: "coupoun with min purchase amout allowed",
			Handler:     func(c *Config) { c.MinPurchaseAmount = minAmount },
		},
		{
			Description: "coupon with min expiration date allowed",
			Handler:     func(c *Config) { c.ExpirationDate = time.Now().Add(minCouponValidityPeriod) },
		},
		{
			Description: "coupon with max expiration date allowed",
			Handler:     func(c *Config) { c.ExpirationDate = time.Now().Add(maxCouponValidityPeriod) },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			copy := data
			test.Handler(&copy)

			// Act
			couponInstance, err := New(copy)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, couponInstance)
			require.NoError(t, uuid.IsValid("coupon_uuid", couponInstance.uuid))
			assert.Equal(t, copy.Name, couponInstance.Name())
			assert.Equal(t, copy.MinPurchaseAmount, couponInstance.MinPurchaseAmount())
			assert.Equal(t, copy.UsageLimit, couponInstance.UsageLimit())
			assert.Equal(t, copy.Exclusivity, couponInstance.Exclusivity())
			assert.Equal(t, copy.ExpirationDate, couponInstance.ExpirationDate())

			// Como checar se um map contém elementos contidos em um slice?
			require.NotNil(t, couponInstance.discountScope)
			assert.Equal(t, couponInstance.currentUsage, 0)
		})
	}
}

func TestRedeemNoError(t *testing.T) {}

func TestRedeemWithError(t *testing.T) {}
