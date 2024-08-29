package valueobject

import (
	"stock-controll/internal/domain/failure"
	"stock-controll/internal/domain/validation"
	"time"
)

type Date struct {
	birthDate time.Time
}

func NewDate(date string) (*Date, error) {
	isValid, dateTime := validation.DateFormatIsValid(date)
	if !isValid {
		return nil, failure.FieldWithInvalidFormat("date", "yyyy-mm-dd")
	}

	return &Date{
		birthDate: *dateTime,
	}, nil
}

func (b *Date) IsOlderThan(minimunAge int) bool {
	currentDate := time.Now()
	limitDate := currentDate.AddDate(-minimunAge, 0, 0)
	isOlder := b.birthDate.Before(limitDate)
	return isOlder
}
