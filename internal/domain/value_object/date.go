package value_object

import (
	"stock-controll/internal/domain/failure"
	"time"
)

type Date struct {
	birthDate time.Time
}

func (b *Date) IsOlderThan(minimunAge int, trowError bool) (bool, error) {
	currentDate := time.Now()
	limitDate := currentDate.AddDate(-minimunAge, 0, 0)
	isOlder :=  b.birthDate.Before(limitDate)

	if trowError && !isOlder {
		return false, failure.InsufficientAge
	}

	return isOlder, nil
}

func NewDate(date string) (*Date, error) {
	isValid, dateTime := DateFormatIsValid(date)
	if !isValid {
		return nil, failure.DateWithInvalidFormat
	}

	return &Date{
		birthDate: *dateTime,
	}, nil
}

