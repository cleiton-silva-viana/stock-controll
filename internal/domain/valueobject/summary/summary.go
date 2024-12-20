package summary

import "stock-controll/internal/domain/services/error/field"

type Summary struct {
	discount float64 // Valor do desconto
	total    float64 // Valor total do produto - desconto
	subtotal float64 // Valor total sem descontos
}

func New(subtotal, total float64) (*Summary, error) {

	err := validateValues(subtotal, total)
	if err != nil {
		return nil, err
	}

	return &Summary{
		discount: subtotal - total,
		total:    total,
		subtotal: subtotal,
	}, nil
}

func (s *Summary) Subtotal() float64 { return s.subtotal }
func (s *Summary) Discount() float64 {
	return s.discount
}
func (s *Summary) Total() float64 {
	return s.total
}

const (
	minValue         = 0
	maxValue         = 100000.00
	ErrNegativeValue = "ERR_NEGATIVE_VALUE"
	ErrValueTooHigh  = "ERR_VALUE_TOO_HIGH"
)

func validateValue(fieldName string, value float64) error {
	if value < minValue {
		return field.Error(fieldName, ErrNegativeValue, value)
	}

	if value > maxValue {
		return field.Error(fieldName, ErrValueTooHigh, value)
	}
	return nil
}

const ErrSubtotalLessThanTotal = "ERR_SUBTOTAL_LESS_THAN_TOTAL"

func validateValues(subtotal, total float64) error {

	err := validateValue("subtotal", subtotal)
	if err != nil {
		return err
	}

	err = validateValue("total", total)
	if err != nil {
		return err
	}

	if subtotal < total {
		return field.Error("subtotal", ErrSubtotalLessThanTotal, subtotal)
	}
	return nil
}
