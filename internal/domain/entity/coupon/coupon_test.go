package coupon

/*
	name string,
	minPurchaseAmount float64,
	discount discount.IDiscountStrategy,
	expirationDate time.Time,
	usageLimit CouponUseLimit,
	exclusivity CouponExclusivity,
	productsUUIDs []string,

*/

/* var copounFake = coupon{
	name:              "PROMOTION",
	minPurchaseAmount: 100.00,
	usageLimit:        -1,
	currentUsage:      0,
	exclusivity:       AllCustomers,
	ExpirationDate:    time.Now().Add(time.Hour * 24),
	products:          nil,
	discount:          nil,
}

var products = make([]string, 0, 1)
var discountType ,_ = discount.CreatePercentageDiscount(1)
var discountStrategy = discount.NewDiscountForAllProducts(discountType)

func Test_NewCoupon_NoError(t *testing.T) {
	t.Parallel()

	testCases := []unitary.TestField[coupon]{
		{
			TestDescription: "create coupon with min name length allowed",
			Handler:         func(c coupon) { c.name = strings.Repeat("a", minCouponNameLength) },
		},
		{
			TestDescription: "create coupon with max name length allowed",
			Handler:         func(c coupon) { c.name = strings.Repeat("b", maxCouponNameLength) },
		},
		{
			TestDescription: "create coupon with letters and numbers in name",
			Handler:         func(c coupon) { c.name = "PROMOTION2024" },
		},
		{
			TestDescription: "coupoun with min purchase amout allowed",
			Handler:         func(c coupon) { c.minPurchaseAmount = minAmount },
		},
		{
			TestDescription: "coupon with min expiration date allowed",
			Handler:         func(c coupon) { c.ExpirationDate = time.Now().Add(minTimeForCoupon) },
		},
		{
			TestDescription: "coupon with max expiration date allowed",
			Handler:         func(c coupon) { c.ExpirationDate = time.Now().Add(maxTimeForCoupon) },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			var data = copounFake
			tt.Handler(data)

			// Act
			couponInstance, err := NewCoupon(
				data.name,
				data.minPurchaseAmount,
				data.discount, data.ExpirationDate, data.usageLimit, data.exclusivity, products)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, couponInstance)
			assert.Equal(t, data.name, couponInstance.GetName())
			assert.Equal(t, data.minPurchaseAmount, couponInstance.GetMinPurchaseAmount())
			assert.Equal(t, data.usageLimit, couponInstance.GetUsageLimit())
			assert.Equal(t, data.exclusivity, couponInstance.GetExclusivity())
			assert.Equal(t, data.ExpirationDate, couponInstance.GetExpirationDate())
			
			// Como checar se um map contém elementos contidos em um slice?
			require.NotNil(t, )
			
			assert.Contains(t, data.products, couponInstance.HasProduct())
			assert.Contains(t, )
		})
	}
}
 */