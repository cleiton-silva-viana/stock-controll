package strategy

import (
	"stock-controll/test/unitary"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type NewDiscountTest struct {
	Description string
	Value       struct {
		Fixed      float64
		Percentage int
	}
}

type ApplyDiscountTest struct {
	Value struct {
		Percentage int
		Fixed      float64
	}
	Price                 float64
	QuantityPurchased     int
	ExpectedFinalTotal    float64
	ExpectedTotalDiscount float64
}

func Setup() *ApplyDiscountTest {
	return &ApplyDiscountTest{
		Value: struct {
			Percentage int
			Fixed      float64
		}{
			Percentage: 10,
			Fixed:      0.1,
		},
		Price:                 100.00,
		QuantityPurchased:     1,
		ExpectedFinalTotal:    90.00,
		ExpectedTotalDiscount: 10.00,
	}
}

func TestPercentageDiscount(t *testing.T) {
	t.Run("Valide percentage discounts", func(t *testing.T) {
		tests := []NewDiscountTest{
			{
				Description: "with min percentage discount allowed",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Percentage: minPercentageForDiscount,
				},
			},
			{
				Description: "with max percentage discount allowed",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Fixed: maxPercentageForDiscount,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Act
				result, err := NewPercentageDiscount(tt.Value.Percentage)

				// Assert
				assert.NoError(t, err)
				require.NotNil(t, result)

				// Erro na asserção abaixo, pois, se mandarmos um inteiro, ele vai converter para float
				// assert.Equal(t, result.Value(), tt.Value)
			})
		}
	})

	t.Run("Invalid percentage discounts", func(t *testing.T) {
		tests := []NewDiscountTest{
			{
				Description: "with percentage discount equal to 0",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Percentage: 0,
				},
			},
			{
				Description: "with percentage discount negative",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Percentage: -1,
				},
			},
			{
				Description: "with percentage discount greater than max allowed",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Percentage: maxPercentageForDiscount + 1,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Act
				result, err := NewPercentageDiscount(tt.Value.Percentage)

				// Assert
				require.Nil(t, result)
				assert.Error(t, err)
			})
		}
	})

	t.Run("Discount successfully applied", func(t *testing.T) {
		tests := []unitary.TestField[ApplyDiscountTest]{
			{
				Description: "to one unit",
				Handler: func(adt *ApplyDiscountTest) {
					adt.QuantityPurchased = 1
				},
			},
			{
				Description: "in multiple units",
				Handler: func(adt *ApplyDiscountTest) {
					adt.QuantityPurchased = 3
					adt.ExpectedFinalTotal = adt.ExpectedFinalTotal * 3
					adt.ExpectedTotalDiscount = adt.ExpectedTotalDiscount * 3
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Arrange
				config := Setup()
				tt.Handler(config)
				d, _ := NewPercentageDiscount(config.Value.Percentage)

				// Act
				result, err := d.Apply(config.Price, config.QuantityPurchased)

				// Assert
				assert.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, result.Total(), config.ExpectedFinalTotal)
				assert.Equal(t, result.Discount(), config.ExpectedTotalDiscount)
			})
		}
	})

	t.Run("Discount applied unsuccessfully", func(t *testing.T) {
		tests := []unitary.TestField[ApplyDiscountTest]{
			{
				Description: "price is less than minimum allowed",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 0.09
					adt.ExpectedFinalTotal = adt.Price * float64(adt.QuantityPurchased)
					adt.ExpectedTotalDiscount = 0
				},
			},
			{
				Description: "price is negative",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = -100
					adt.ExpectedFinalTotal = adt.Price * float64(adt.QuantityPurchased)
					adt.ExpectedTotalDiscount = 0
				},
			},
			{
				Description: "quantity units is less than 1",
				Handler: func(adt *ApplyDiscountTest) {
					adt.QuantityPurchased = 0
					adt.ExpectedFinalTotal = adt.Price * float64(adt.QuantityPurchased)
					adt.ExpectedTotalDiscount = 0
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Arrange
				config := Setup()
				tt.Handler(config)
				d, _ := NewPercentageDiscount(config.Value.Percentage)

				// Act
				result, err := d.Apply(config.Price, config.QuantityPurchased)

				// Assert
				assert.Error(t, err)
				require.NotNil(t, result)
				assert.Equal(t, result.Total(), config.ExpectedFinalTotal)
				assert.Equal(t, result.Discount(), config.ExpectedTotalDiscount)
			})
		}
	})
}

func TestFixedValueDiscount(t *testing.T) {
	t.Run("Valid value for discounts", func(t *testing.T) {
		tests := []NewDiscountTest{
			{
				Description: "with value discount equal than min allowed",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Fixed: minFixedDiscountValue,
				},
			},
			{
				Description: "with value discount less than max allowed",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Fixed: minFixedDiscountValue,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Act
				result, err := NewFixedValueDiscount(tt.Value.Fixed)

				// Assert
				assert.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, result.Value(), tt.Value)
			})
		}
	})

	t.Run("Invalid value for discounts", func(t *testing.T) {
		tests := []NewDiscountTest{
			{
				Description: "with value discount equal to 0",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Fixed: 0,
				},
			},
			{
				Description: "with value discount less than min allowed",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Fixed: minFixedDiscountValue - 0.1,
				},
			},
			{
				Description: "with value discount greater than max allowed",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Fixed: maxFixedDiscountValue + 0.1,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Act
				result, err := NewFixedValueDiscount(tt.Value.Fixed)

				// Assert
				assert.Nil(t, result)
				assert.Error(t, err)
			})
		}
	})

	t.Run("Apply discount with success", func(t *testing.T) {
		tests := []unitary.TestField[ApplyDiscountTest]{
			{
				Description: "unit price equal to the discount amount",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 1.00
					adt.Value.Fixed = 1.00
					adt.QuantityPurchased = 1
					adt.ExpectedTotalDiscount = 1.00
					adt.ExpectedFinalTotal = 0
				},
			},
			{
				Description: "the unit price is greater than the discount amount",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 10.00
					adt.Value.Fixed = 2.00
					adt.QuantityPurchased = 1
					adt.ExpectedTotalDiscount = 2.00
					adt.ExpectedFinalTotal = 8.00
				},
			},
			{
				Description: "the unit price is less than discount amount, but the quantity * unit is equal than the discount amount",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 0.10
					adt.Value.Fixed = 1.00
					adt.QuantityPurchased = 10
					adt.ExpectedTotalDiscount = 1.00
					adt.ExpectedFinalTotal = 0
				},
			},
			{
				Description: "the unit price is less than the discount amount, but the quantity * unit is greater than the discount amount",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 0.25
					adt.Value.Fixed = 1.00
					adt.QuantityPurchased = 5
					adt.ExpectedTotalDiscount = 1.00
					adt.ExpectedFinalTotal = 0.25
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Arrange
				config := Setup()
				tt.Handler(config)
				d, _ := NewFixedValueDiscount(config.Value.Fixed)

				// Act
				result, err := d.Apply(config.Price, config.QuantityPurchased)

				// Assert
				assert.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, config.ExpectedFinalTotal, result.Total())
				assert.Equal(t, config.ExpectedTotalDiscount, result.Discount())
			})
		}
	})

	t.Run("Apply discount unsuccessfully", func(t *testing.T) {
		tests := []unitary.TestField[ApplyDiscountTest]{
			{
				Description: "product price is equals to 0",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 0
					adt.QuantityPurchased = 1
					adt.ExpectedFinalTotal = 0
					adt.ExpectedTotalDiscount = 0
				},
			},
			{
				Description: "the amount of product is 0",
				Handler: func(adt *ApplyDiscountTest) {
					adt.QuantityPurchased = 0
					adt.ExpectedFinalTotal = 0
					adt.ExpectedTotalDiscount = 0
				},
			},
			{
				Description: "The unit price of the product is negative",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = -0.10
					adt.Value.Fixed = 0.99
					adt.QuantityPurchased = 10
					adt.ExpectedFinalTotal = -0.10
					adt.ExpectedTotalDiscount = 0
				},
			},
			{
				Description: "the unit price of the product is less than the discount value",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 0.99
					adt.Value.Fixed = 1.00
					adt.QuantityPurchased = 1
					adt.ExpectedFinalTotal = 1.00
					adt.ExpectedTotalDiscount = 0
				},
			},
			{
				Description: "The unit price of the product with quantity is less than the discount value",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 0.09
					adt.Value.Fixed = 0.99
					adt.QuantityPurchased = 10
					adt.ExpectedFinalTotal = 0.90
					adt.ExpectedTotalDiscount = 0
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Assert
				config := Setup()
				tt.Handler(config)
				d, _ := NewFixedValueDiscount(config.Value.Fixed)

				// Act
				result, err := d.Apply(config.Price, config.QuantityPurchased)

				// Arrange
				assert.Error(t, err)
				require.NotNil(t, result)
				assert.Equal(t, config.ExpectedFinalTotal, result.Total())
				assert.Equal(t, config.ExpectedTotalDiscount, result.Discount())
			})
		}
	})
}

func TestMaxiDiscountStrategy(t *testing.T) {
	t.Run("Valid values for discounts", func(t *testing.T) {
		// Arrange
		fixedValue := 0.7
		percentage := 10

		// Act
		result, err := NewMaxDiscountStrategy(fixedValue, percentage)

		// Assert
		assert.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, result.FixedValue(), fixedValue)
		assert.Equal(t, result.PercentageValue(), percentage)
	})

	t.Run("Invalid values for discounts", func(t *testing.T) {
		tests := []NewDiscountTest{
			{
				Description: "with invalid fixed value ",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Fixed:      100000,
					Percentage: 10,
				},
			},
			{
				Description: "with invalid percentage value",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Fixed:      0.1,
					Percentage: 200,
				},
			},
			{
				Description: "with invalid percentage & fixed value ",
				Value: struct {
					Fixed      float64
					Percentage int
				}{
					Fixed:      -10,
					Percentage: 0,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Act
				result, err := NewMaxDiscountStrategy(tt.Value.Fixed, tt.Value.Percentage)

				// Assert
				assert.Nil(t, result)
				assert.Error(t, err)
			})
		}
	})

	t.Run("Apply discount with success", func(t *testing.T) {
		tests := []unitary.TestField[ApplyDiscountTest]{
			{
				Description: "applied fixed discount",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 0.99
					adt.QuantityPurchased = 10
					adt.Value.Fixed = 3.65
					adt.Value.Percentage = 40
					adt.ExpectedTotalDiscount = 3.65
					adt.ExpectedFinalTotal = 5.35
				},
			},
			{
				Description: "applied percentage discount",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 0.99
					adt.QuantityPurchased = 10
					adt.Value.Fixed = 3.59
					adt.Value.Percentage = 40
					adt.ExpectedTotalDiscount = 3.60
					adt.ExpectedFinalTotal = 5.40
				},
			},
			{
				Description: "percentage & fixed value equals",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 1
					adt.QuantityPurchased = 1
					adt.Value.Fixed = 0.50
					adt.Value.Percentage = 50
					adt.ExpectedTotalDiscount = 0.50
					adt.ExpectedFinalTotal = 0.50
				},
			},
			{
				Description: "max discounts strategy applied",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 250
					adt.QuantityPurchased = 1
					adt.Value.Fixed = 100
					adt.Value.Percentage = 100
					adt.ExpectedTotalDiscount = 250
					adt.ExpectedFinalTotal = 0
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Arrange
				config := Setup()
				tt.Handler(config)
				d, _ := NewMaxDiscountStrategy(config.Value.Fixed, config.Value.Percentage)

				// Act
				result, err := d.Apply(config.Price, config.QuantityPurchased)

				// Assert
				assert.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, config.ExpectedFinalTotal, result.Total())
				assert.Equal(t, config.ExpectedTotalDiscount, result.Discount())
			})
		}
	})

	t.Run("Apply discount unsuccessfully", func(t *testing.T) {
		tests := []unitary.TestField[ApplyDiscountTest]{
			{
				Description: "product price is negative",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = -0.01
					adt.QuantityPurchased = 10
					adt.Value.Fixed = 0.01
					adt.Value.Percentage = 5
					adt.ExpectedTotalDiscount = 0
					adt.ExpectedFinalTotal = 0.1
				},
			},
			{
				Description: "the amount purchased is 0",
				Handler: func(adt *ApplyDiscountTest) {
					adt.Price = 0.99
					adt.QuantityPurchased = 0
					adt.Value.Fixed = 0.1
					adt.Value.Percentage = 15
					adt.ExpectedTotalDiscount = 0
					adt.ExpectedFinalTotal = 0
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Arrange
				config := Setup()
				tt.Handler(config)
				d, _ := NewMaxDiscountStrategy(config.Value.Fixed, config.Value.Percentage)

				// Act
				result, err := d.Apply(config.Price, config.QuantityPurchased)

				// Assert
				assert.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, config.ExpectedFinalTotal, result.Total())
				assert.Equal(t, config.ExpectedTotalDiscount, result.Discount())
			})
		}
	})
}
