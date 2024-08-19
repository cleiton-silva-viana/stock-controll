package value_object

import (
	"strings"
	
	"stock-controll/internal/domain/failure"
)

type Gender struct {
	gender string
}

func NewGender(gender string) (*Gender, error) {
	gender = strings.ToLower(gender)
	if gender == "female" || gender == "male" {
		return &Gender{
			gender: gender,
		}, nil
	}
	return nil, failure.GenderInvalid
}
