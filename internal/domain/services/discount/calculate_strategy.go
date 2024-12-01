package discount

import (
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/validate"
)

type DiscountSummary struct {
	totalDiscountApplied float64 // valor do desconto
	finalTotal           float64 // valor total do produto - desconto
}

func (ds *DiscountSummary) TotalDiscountApplied() float64 {
	return ds.totalDiscountApplied
}

func (ds *DiscountSummary) FinalTotal() float64 {
	return ds.finalTotal
}

func (ds *DiscountSummary) Update(discountSummary DiscountSummary) {
	ds.totalDiscountApplied += discountSummary.totalDiscountApplied
	ds.finalTotal += discountSummary.finalTotal
}

type IDiscountCalculationStrategy interface {
	Apply(price float64, quantityPurchased int) (DiscountSummary, *validate.FieldError)
}

type PercentageDiscount struct {
	percentage float64
}

const (
	minPercentageForDiscount          = 1
	ErrDiscountPercentageBelowMinimum = "ERR_DISCOUNT_PERCENTAGE_BELOW_MINIMUM"
	/*
		ERR_DISCOUNT_PERCENTAGE_BELOW_MINIMUM: {
			"Message": "A porcentagem de desconto é inferior ao mínimo permitido.",
			"Solution": "Por favor, verifique a porcentagem de desconto e tente novamente."
		}
	*/

	maxPercentageForDiscount            = 100
	ErrDiscountPercentageExceedsMaximum = "ERR_DISCOUNT_PERCENTAGE_EXCEEDS_MAXIMUM"
	/*
		"ERR_DISCOUNT_PERCENTAGE_EXCEEDS_MAXIMUM": {
			"Message": "A porcentagem do desconto é maior que o máximo permitido.",
			"Solution": "Por favor, verifique a porcentagem do desconto e tente novamente."
		}
	*/
)

func NewPercentageDiscount(discountPercentage int) (*PercentageDiscount, *validationerrors.ValidationError) {

	if discountPercentage < minPercentageForDiscount {
		return nil, validationerrors.New("percentage_discount").
			AddValidationError(&validate.FieldError{
				FieldName: "percentage",
				CodeError: ErrDiscountPercentageBelowMinimum,
			})
	}

	if discountPercentage > maxPercentageForDiscount {
		return nil, validationerrors.New("percentage_discount").
			AddValidationError(&validate.FieldError{
				FieldName: "percentage",
				CodeError: ErrDiscountPercentageExceedsMaximum,
			})
	}

	return &PercentageDiscount{
		percentage: float64(discountPercentage),
	}, nil
}

func (d *PercentageDiscount) Apply(price float64, quantityPurchased int) (DiscountSummary, error) {
	subtotal := price * float64(quantityPurchased)
	finalAmountAfterDiscount := subtotal * (1 - d.percentage/100)

	return DiscountSummary{
		finalTotal:           finalAmountAfterDiscount,
		totalDiscountApplied: subtotal - finalAmountAfterDiscount,
	}, nil
}

type FixedValueDiscount struct {
	discountValue float64
}

const (
	minFixedDiscountValue        = 0
	ErrDiscountValueBelowMinimum = "ERR_DISCOUNT_VALUE_BELOW_MINIMUM"
	/*
		"ERR_DISCOUNT_VALUE_BELOW_MINIMUM": {
			"Message": "O valor do desconto é menor que o mínimo permitido.",
			"Solution": "Por favor, verifique o valor do desconto e tente novamente."
		}
	*/

	maxFixedDiscountValue          = 100
	ErrDiscountValueExceedsMaximum = "ERR_DISCOUNT_VALUE_EXCEEDS_MAXIMUM"
	/*
		"ERR_DISCOUNT_VALUE_EXCEEDS_MAXIMUM": {
			"Message": "O valor do desconto é maior que o máximo permitido.",
			"Solution": "Por favor, verifique o valor do desconto e tente novamente."
		}
	*/

)

func NewFixedValueDiscount(discountValue float64) (*FixedValueDiscount, *validationerrors.ValidationError) {
	if discountValue < minFixedDiscountValue {
		return nil, validationerrors.New("value_discount").
			AddValidationError(&validate.FieldError{
				FieldName: "value",
				CodeError: ErrDiscountValueBelowMinimum,
			})
	}

	if discountValue > maxFixedDiscountValue {
		return nil, validationerrors.New("value_discount").
			AddValidationError(&validate.FieldError{
				FieldName: "value",
				CodeError: ErrDiscountValueExceedsMaximum,
			})
	}

	return &FixedValueDiscount{
		discountValue: discountValue,
	}, nil
}

const ErrSubtotalMustBeGreaterThanDiscount = "ERR_SUBTOTAL_MUST_BE_GREATER_THAN_DISCOUNT"
/*
	"ERR_SUBTOTAL_MUST_BE_GREATER_THAN_DISCOUNT": {
		"Message": "O valor do subtotal deve ser maior que o valor do desconto para que ele possa ser aplicado.",
		"Solution": "Por favor, ajuste o subtotal para que seja maior que o valor do desconto."
	}
*/
func (d *FixedValueDiscount) Apply(price float64, quantityPurchased int) (DiscountSummary, *validate.FieldError) {
	subtotal := price * float64(quantityPurchased)

	var discounted = DiscountSummary{
		totalDiscountApplied: 0,
		finalTotal:           subtotal,
	}

	if subtotal < d.discountValue {
		return discounted, &validate.FieldError{
			CodeError: ErrSubtotalMustBeGreaterThanDiscount,
		}
	}

	discounted.finalTotal = d.discountValue
	discounted.totalDiscountApplied = subtotal - d.discountValue

	return discounted, nil
}

type MaxDiscountStrategy struct {
	fixedValue FixedValueDiscount
	percentage PercentageDiscount
}

func NewMaxDiscountStrategy(fixedValue FixedValueDiscount, percentage PercentageDiscount) *MaxDiscountStrategy {
	return &MaxDiscountStrategy{
		fixedValue: fixedValue,
		percentage: percentage,
	}
}

func (md *MaxDiscountStrategy) Apply(price float64, quantityPurchased int) (DiscountSummary, *validate.FieldError) {
	fixedSummary, fixedErr := md.fixedValue.Apply(price, quantityPurchased)
	percentageSummary, _ := md.percentage.Apply(price, quantityPurchased)

	if fixedErr != nil {
		return percentageSummary, fixedErr
	}

	if percentageSummary.totalDiscountApplied > fixedSummary.totalDiscountApplied {
		return percentageSummary, nil
	}
	return fixedSummary, nil
}
