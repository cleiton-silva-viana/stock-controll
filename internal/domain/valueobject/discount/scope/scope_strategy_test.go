package scope

import (
	"stock-controll/internal/domain/valueobject/discount/strategy"
	"stock-controll/internal/domain/valueobject/uuid"
	"stock-controll/test/unitary"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type UUIDsForTest struct {
	Discount strategy.IDiscountCalculationStrategy
	UUIDS    map[string]struct{}
}

func SetupScope() *UUIDsForTest {
	return &UUIDsForTest{
		Discount: &strategy.FixedValueDiscount{discountValue: 0.10},
		UUIDS:    map[string]struct{}{},
	}
}

func TestGlobalDiscountStrategy(t *testing.T) {
	t.Run("Valid global discount strategy", func(t *testing.T) {
		// Arrange
		_discount, _ := strategy.NewFixedValueDiscount(10)

		// Act
		result, err := NewGlobalDiscountStrategy(_discount)

		// Assert
		assert.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("Invalid global discount strategy", func(t *testing.T) {
		// Arrange
		_discount, _ := strategy.NewFixedValueDiscount(10)

		// Act
		result, err := NewGlobalDiscountStrategy(_discount)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("validate method IsValid of global discount strategy", func(t *testing.T) {
		// Arrange

		// Act

		// Assert
	})
}

func TestSpecificCategoryDiscount(t *testing.T) {
	t.Run("Valid specific category discount", func(t *testing.T) {
		tests := []unitary.TestField[UUIDsForTest]{
			{
				Description: "with min items required",
				Handler: func(n *UUIDsForTest) {
					var uuids map[string]struct{}
					uuids[uuid.New("bra").String()] = struct{}{}
					n.UUIDS = uuids
				},
			},
			{
				Description: "with valid & invalid uuids",
				Handler: func(n *UUIDsForTest) {
					uuids := make(map[string]struct{})
					uuids["invalid_uuid"] = struct{}{}
					uuids[uuid.New().String()] = struct{}{} // valid uuid
					n.UUIDS = uuids
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Arrange
				d := SetupScope()
				tt.Handler(d)

				// Act
				result, err := NewSpecificCategoryDiscountStrategy(d.UUIDS, d.Discount)

				// Assert
				assert.NoError(t, err)
				require.NotNil(t, result)
			})
		}
	})

	t.Run("Invalid specific category discount", func(t *testing.T) {
		tests := []unitary.TestField[UUIDsForTest]{
			{
				Description: "because discount is a null object",
				Handler: func(n *UUIDsForTest) {
					n.Discount = nil
				},
			},
			{
				Description: "because category uuids is a null object",
				Handler: func(n *UUIDsForTest) {
					n.UUIDS = nil
				},
			},
			{
				Description: "because category uuids is a empty object",
				Handler: func(n *UUIDsForTest) {
					n.UUIDS = make(map[string]struct{}, 0)
				},
			},
			{
				Description: "because category uuids is a invalid uuids",
				Handler: func(n *UUIDsForTest) {
					uuids := make(map[string]struct{})
					uuids["132548"] = struct{}{}
					uuids["abcq"] = struct{}{}
					uuids["TR#$T$3654sdvwsdv"] = struct{}{}
					uuids["        "] = struct{}{}
					uuids[""] = struct{}{}
					uuids["invalid uuid format"] = struct{}{}
					n.UUIDS = uuids
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Description, func(t *testing.T) {
				// Arrange
				d := SetupScope()
				tt.Handler(d)

				// Act
				result, err := NewSpecificCategoryDiscountStrategy(d.UUIDS, d.Discount)

				// Assert
				assert.Nil(t, result)
				assert.Error(t, err)
			})
		}
	})

	t.Run("Check function IsProductValid", func(t *testing.T) {})

}

func TestSpecificBrandDiscount(t *testing.T) {
	t.Run("Valid specific brands discount", func(t *testing.T) {

	})

	t.Run("Invalid operation of adding new brands", func(t *testing.T) {

	})

	t.Run("Check function IsProductValid", func(t *testing.T) {

	})
}

func TestSpecificProductsDiscount(t *testing.T) {
	t.Run("Valid specific products discount", func(t *testing.T) {

	})

	t.Run("Invalid operation of adding new brands", func(t *testing.T) {

	})

	t.Run("Check function IsProductValid", func(t *testing.T) {

	})
}
