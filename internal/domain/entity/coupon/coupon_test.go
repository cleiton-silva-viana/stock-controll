package coupon

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"stock-controll/test/unitary"

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

func Setup() *Config {
	return &Config{
		Name:              "PROMOTION",
		MinPurchaseAmount: 100.00,
		UsageLimit:        -1,
		Exclusivity:       "all_customers",
		ExpirationDate:    time.Now().Add(time.Hour * 24),
	}
}

func TestNewCouponNoError(t *testing.T) {
	t.Parallel()

	tests := []unitary.TestField[Config]{
		{
			Description: "coupon with min name length allowed",
			Handler:     func(c *Config) { c.Name = strings.Repeat("a", minNameLength) },
		},
		{
			Description: "coupon with max name length allowed",
			Handler:     func(c *Config) { c.Name = strings.Repeat("b", maxNameLength) },
		},
		{
			Description: "coupon with letters and numbers in name",
			Handler:     func(c *Config) { c.Name = "PROMOTION2024" },
		},
		{
			Description: "coupon with min purchase amout allowed",
			Handler:     func(c *Config) { c.MinPurchaseAmount = minAmount },
		},
		{
			Description: "coupon with min expiration date allowed",
			Handler:     func(c *Config) { c.ExpirationDate = time.Now().Add(minValidityPeriod) },
		},
		{
			Description: "coupon with max expiration date allowed",
			Handler:     func(c *Config) { c.ExpirationDate = time.Now().Add(maxValidityPeriod) },
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint("Test function New of package coupon, %s", test.Description), func(t *testing.T) {
			c := Setup()
			test.Handler(c)

			// Act
			result, err := New(*c)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, result)
			assert.Equal(t, c.Name, result.Name())
			assert.Equal(t, c.MinPurchaseAmount, result.MinPurchaseAmount())
			assert.Equal(t, c.UsageLimit, result.UsageLimit())
			assert.Equal(t, c.Exclusivity, result.Exclusivity())
			assert.Equal(t, c.ExpirationDate, result.ExpirationDate())

			// Como checar se um map contém elementos contidos em um slice?
			require.NotNil(t, result.discountScope)
			assert.Equal(t, result.currentUsage, 0)
		})
	}
}

func TestNewCouponWithError(t *testing.T) {
	t.Parallel()

	tests := []unitary.TestField[Config]{
		{
			Description: "coupon with empty name",
			Handler:     func(c *Config) { c.Name = strings.Repeat(" ", minNameLength) },
		},
		{
			Description: "coupon with name shorter than minimum length",
			Handler:     func(c *Config) { c.Name = strings.Repeat("a", minNameLength-1) },
		},
		{
			Description: "coupon with name longer than maximum length",
			Handler:     func(c *Config) { c.Name = strings.Repeat("a", maxNameLength+1) },
		},
		{
			Description: "coupon with special characters in name",
			Handler:     func(c *Config) { c.Name = "Promo@2024!" },
		},
		{
			Description: "coupon with purchase amount less than minimum allowed",
			Handler:     func(c *Config) { c.MinPurchaseAmount = minAmount - 1 },
		},
		{
			Description: "coupon with expiration date before start date",
			Handler: func(c *Config) {
				c.StartDate = time.Now().Add(2 * time.Hour)
				c.ExpirationDate = time.Now()
			},
		},
		{
			Description: "coupon with expiration date exceeding maximum validity period",
			Handler: func(c *Config) {
				c.StartDate = time.Now()
				c.ExpirationDate = time.Now().Add(maxValidityPeriod + 1*time.Hour)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			c := Setup()
			test.Handler(c)

			// Act
			result, err := New(*c)

			// Assert
			require.Nil(t, result)
			assert.Error(t, err)
		})
	}
}

func TestRedeemNoError(t *testing.T) {}

func TestRedeemWithError(t *testing.T) {}
