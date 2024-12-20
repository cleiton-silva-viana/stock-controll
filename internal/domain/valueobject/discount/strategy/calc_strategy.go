package strategy

import (
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/valueobject/summary"
)

type IDiscountCalculationStrategy interface {
	Apply(price float64, quantityPurchased int) (*summary.Summary, error)
	Value() float64
}

type PercentageDiscount struct {
	percentage float64
}

func NewPercentageDiscount(discountPercentage int) (*PercentageDiscount, error) {
	err := validatePercentageDiscountValue(discountPercentage)
	if err != nil {
		return nil, err
	}
	return &PercentageDiscount{
		percentage: float64(discountPercentage),
	}, nil
}

func (pd *PercentageDiscount) Value() float64 {
	return pd.percentage
}

func (pd *PercentageDiscount) Apply(price float64, quantityPurchased int) (*summary.Summary, error) {
	subtotal := price * float64(quantityPurchased)
	discounted := subtotal * (pd.percentage / 100)
	total := subtotal - discounted
	return summary.New(subtotal, total)
}

type FixedValueDiscount struct {
	discountValue float64
}

func NewFixedValueDiscount(discountValue float64) (*FixedValueDiscount, error) {
	err := validateFixedDiscountValue(discountValue)
	if err != nil {
		return nil, err
	}
	return &FixedValueDiscount{
		discountValue: discountValue,
	}, nil
}

func (fvd *FixedValueDiscount) Value() float64 {
	return fvd.discountValue
}

const (
	ErrSubtotalMustBeGreaterThanDiscount = "ERR_SUBTOTAL_MUST_BE_GREATER_THAN_DISCOUNT"
)

func (fvd *FixedValueDiscount) Apply(price float64, quantityPurchased int) (*summary.Summary, error) {
	subtotal := price * float64(quantityPurchased)
	if subtotal < fvd.discountValue {
		discounted, _ := summary.New(subtotal, subtotal)
		return discounted, &field.FieldError{
			CodeError: ErrSubtotalMustBeGreaterThanDiscount,
		}
	}
	total := subtotal - fvd.discountValue
	return summary.New(subtotal, total)
}

type MaxDiscountStrategy struct {
	fixedValue FixedValueDiscount
	percentage PercentageDiscount
}

func NewMaxDiscountStrategy(fixedValue float64, percentageValue int) (*MaxDiscountStrategy, error) {
	fixedValueDiscount, err := NewFixedValueDiscount(fixedValue)
	if err != nil {
		return nil, err
	}
	percentageValueDiscount, err := NewPercentageDiscount(percentageValue)
	if err != nil {
		return nil, err
	}
	return &MaxDiscountStrategy{
		fixedValue: *fixedValueDiscount,
		percentage: *percentageValueDiscount,
	}, nil
}

func (md *MaxDiscountStrategy) FixedValue() float64 {
	return md.fixedValue.Value()
}

func (md *MaxDiscountStrategy) PercentageValue() float64 {
	return md.percentage.Value()
}

func (md *MaxDiscountStrategy) Apply(price float64, quantityPurchased int) (*summary.Summary, error) {
	fixedSummary, err := md.fixedValue.Apply(price, quantityPurchased)
	percentageSummary, _ := md.percentage.Apply(price, quantityPurchased)
	if err != nil {
		return percentageSummary, err
	}
	if percentageSummary.Total() < fixedSummary.Total() {
		return percentageSummary, nil
	}
	return fixedSummary, nil
}

const (
	minFixedDiscountValue          = 0.01
	maxFixedDiscountValue          = 100
	ErrDiscountValueBelowMinimum   = "ERR_DISCOUNT_VALUE_BELOW_MINIMUM"
	ErrDiscountValueExceedsMaximum = "ERR_DISCOUNT_VALUE_EXCEEDS_MAXIMUM"
)

func validateFixedDiscountValue(discountValue float64) error {
	if discountValue < minFixedDiscountValue {
		return &field.FieldError{
			FieldName: "value",
			CodeError: ErrDiscountValueBelowMinimum,
		}
	}
	if discountValue > maxFixedDiscountValue {
		return &field.FieldError{
			FieldName: "value",
			CodeError: ErrDiscountValueExceedsMaximum,
		}
	}
	return nil
}

const (
	minPercentageForDiscount            = 1
	ErrDiscountPercentageBelowMinimum   = "ERR_DISCOUNT_PERCENTAGE_BELOW_MINIMUM"
	maxPercentageForDiscount            = 75
	ErrDiscountPercentageExceedsMaximum = "ERR_DISCOUNT_PERCENTAGE_EXCEEDS_MAXIMUM"
)

func validatePercentageDiscountValue(discountPercentage int) error {
	if discountPercentage < minPercentageForDiscount {
		return &field.FieldError{
			FieldName: "percentage",
			CodeError: ErrDiscountPercentageBelowMinimum,
		}
	}
	if discountPercentage > maxPercentageForDiscount {
		return &field.FieldError{
			FieldName: "percentage",
			CodeError: ErrDiscountPercentageExceedsMaximum,
		}
	}
	return nil
}
