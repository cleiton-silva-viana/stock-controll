package discount

import (
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
)

type DiscountSummary struct {
	totalDiscountApplied float64 // valor do desconto
	finalTotal           float64 // valor total do produto - desconto
}

func (ds *DiscountSummary) GetTotalDiscountApplied() float64 {
	return ds.totalDiscountApplied
}

func (ds *DiscountSummary) GetFinalTotal() float64 {
	return ds.finalTotal
}

func (ds *DiscountSummary) Update(discountSummary DiscountSummary) {
	ds.totalDiscountApplied += discountSummary.totalDiscountApplied
	ds.finalTotal += discountSummary.finalTotal
}

type IDiscountCalculationStrategy interface {
	Apply(price float64, quantityPurchased int) (DiscountSummary, *validation.FieldError)
}

type PercentageDiscount struct {
	percentage float64
}

const (
	minPercentageForDiscount = 1
	maxPercentageForDiscount = 100
)

func NewPercentageDiscount(discountPercentage int) (*PercentageDiscount, validationError.IValidationError) {
	if discountPercentage < minPercentageForDiscount || discountPercentage > maxPercentageForDiscount {
		return nil, validationError.NewValidationError("percentage_discount").
			AddValidationError(&validation.FieldError{
				FieldName:  "percentage",
				CodeErrors: []string{string(validation.ErrUnknown)},
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
	minFixedDiscountValue = 0
	maxFixedDiscountValue = 100
)

func NewFixedValueDiscount(discountValue float64) (*FixedValueDiscount, validationError.IValidationError) {
	if discountValue < minFixedDiscountValue || discountValue > maxFixedDiscountValue {
		return nil, validationError.NewValidationError("value_discount").
			AddValidationError(&validation.FieldError{
				FieldName:  "value",
				CodeErrors: []string{string(validation.ErrUnknown)},
			})
	}
	return &FixedValueDiscount{
		discountValue: discountValue,
	}, nil
}

func (d *FixedValueDiscount) Apply(price float64, quantityPurchased int) (DiscountSummary, *validation.FieldError) {
	subtotal := price * float64(quantityPurchased)

	var discounted = DiscountSummary{
		totalDiscountApplied: 0,
		finalTotal:           subtotal,
	}

	if subtotal < d.discountValue {
		return discounted, &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
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

func (md *MaxDiscountStrategy) Apply(price float64, quantityPurchased int) (DiscountSummary, *validation.FieldError) {
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