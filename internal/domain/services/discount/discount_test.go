package discount
/* 
import (
	"stock-controll/test/unitary"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var dataForTest = DiscountResult{
	totalWithoutDiscount: 0,
	totalWithDiscount:    0,
}

func Test_DiscountResult_AddDiscountValues_NoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[DiscountResult]{
		{
			TestDescription: "adicionando o valor mínimo para o total sem",
			Handler:         func(dr DiscountResult) { dr.totalWithoutDiscount = minValeForTotal },
		},
		{
			TestDescription: "adicionando o valor mínimo para produto com desconto",
			Handler:         func(dr DiscountResult) { dr.totalWithDiscount = minValeForTotal },
		},
		{
			TestDescription: "adicionando valores iguais para ambos os tipos de totais",
			Handler: func(dr DiscountResult) {
				dr.totalWithoutDiscount = 8
				dr.totalWithDiscount = 8
			},
		},
		{
			TestDescription: "adicionando um valor maior para total e um menor para valor com desconto",
			Handler: func(dr DiscountResult) {
				dr.totalWithoutDiscount = 100
				dr.totalWithDiscount = 90
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			var data = dataForTest
			tt.Handler(data)

			// Act
			err := dataForTest.AddDiscountValues(data.totalWithoutDiscount, data.totalWithDiscount)

			// Assert
			require.Nil(t, err)
			assert.Equal(t, dataForTest.totalWithDiscount, data.totalWithDiscount)
			assert.Equal(t, dataForTest.totalWithoutDiscount, data.totalWithoutDiscount)
		})
	}
}

func Test_DiscountResult_AddDiscountValues_WithError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[DiscountResult]{
		{
			TestDescription: "valor do produto com desconto é maior que o do produto sem desconto",
			Handler:         func(dr DiscountResult) { dr.totalWithDiscount = 10; dr.totalWithoutDiscount = 9 },
		},
		{
			TestDescription: "valor do total com desconto é negativo",
			Handler:         func(dr DiscountResult) { dr.totalWithDiscount = -1 },
		},
		{
			TestDescription: "valor do total sem desconto é negativo",
			Handler:         func(dr DiscountResult) { dr.totalWithoutDiscount = -10 },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			var data = dataForTest
			tt.Handler(data)

			// Act
			err := dataForTest.AddDiscountValues(data.totalWithoutDiscount, data.totalWithDiscount)

			// Assert
			assert.NotNil(t, err)
		})
	}
}

func Test_DiscountResult_UpdateWithDiscount(t *testing.T) {
	// Arrange
	var data = dataForTest
	var updater = DiscountResult{
		totalWithoutDiscount: 100.00,
		totalWithDiscount:    90.00,
	}

	// Act
	data.UpdateWithDiscount(updater)

	assert.Equal(t, updater.totalWithoutDiscount, updater.totalWithoutDiscount)
	assert.Equal(t, updater.totalWithDiscount, data.totalWithDiscount)
}

func Test_DiscountResult_CalculateDiscount(t *testing.T) {
	// Arrange
	const total float64 = 1000
	const subtotal float64 = 20
	const discount float64 = 980

	var data = DiscountResult{
		totalWithoutDiscount: total,
		totalWithDiscount:    subtotal,
	}

	// Act
	discountCalculated := data.CalculateDiscount()

	// Assert
	assert.Equal(t, discount, discountCalculated)
}

type TestDiscountStrategy[T any] struct {
	TestDescription string
	Value           T
}

func Test_NewPercentageDiscount_NoError(t *testing.T) {
	// Arrange
	testCases := []TestDiscountStrategy[int]{
		{
			TestDescription: "valor mínimo para desconto",
			Value:           minPercentageForDiscount,
		},
		{
			TestDescription: "valor máximo para desconto",
			Value:           maxPercentageForDiscount,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {

			// Act
			percentageInstance, err := NewPercentageDiscount(tt.Value)

			// Assert
			require.Nil(t, err)
			require.NotNil(t, percentageInstance)
			assert.Equal(t, tt.Value, percentageInstance.percentage)
		})
	}
}

func Test_NewPercentageDiscount_WithError(t *testing.T) {
	// Arrange
	testCases := []TestDiscountStrategy[int]{
		{
			TestDescription: "valor menor que o permitido",
			Value:           minPercentageForDiscount - 1,
		},
		{
			TestDescription: "valor maior que o permitido",
			Value:           maxPercentageForDiscount + 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {

			// Act
			percentageInstance, err := NewPercentageDiscount(tt.Value)

			// Assert
			require.Nil(t, percentageInstance)
			require.NotNil(t, err)
		})
	}
}

// implementar
func Test_PercentageDiscount_Apply(t *testing.T) {}

func Test_NewFixedValueDiscount_NoError(t *testing.T) {
	// Arrange
	testCases := []TestDiscountStrategy[float64]{
		{
			TestDescription: "value is equal than minimum allowed",
			Value:           minFixedDiscountValue,
		},
		{
			TestDescription: "value is equal than maximum allowed",
			Value:           maxFixedDiscountValue,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {

			// Act
			discountInstance, err := NewFixedValueDiscount(tt.Value)

			// Assert
			require.Nil(t, err)
			require.NotNil(t, discountInstance)
			assert.Equal(t, tt.Value, discountInstance.discountValue)
		})
	}
}

func Test_NewFixedValueDiscount_WithError(t *testing.T) {
	// Arrange
	testCases := []TestDiscountStrategy[float64]{
		{
			TestDescription: "value discount is short than minimum allowed",
			Value:           minFixedDiscountValue - 1,
		},
		{
			TestDescription: "value discount is greater than maximum allowed",
			Value:           maxFixedDiscountValue + 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {

			// Act
			discountInstance, err := NewFixedValueDiscount(tt.Value)

			// Assert
			require.Nil(t, discountInstance)
			require.NotNil(t, err)
		})
	}
}

// Implementar
func Test_FixedValueDiscount_Apply(t *testing.T) {}




 */