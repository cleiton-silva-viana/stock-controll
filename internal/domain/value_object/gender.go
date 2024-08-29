package valueobject

import (
	"strings"

	"stock-controll/internal/domain/failure"
)

type Gender struct {
	gender string
}

var genders = []string{"female", "male"}

func NewGender(gender string) (*Gender, error) {
	genderParsed := strings.ReplaceAll(gender, " ", "")
	genderParsed = strings.ToLower(genderParsed)

	if genderParsed == "" {
		return nil, failure.FieldIsEmpty("gender")
	}
	
	for _, gen := range genders {
		if gen == genderParsed {
			return &Gender{
				gender: gender,
			}, nil
		}
	}
	return nil, failure.FieldNotRangeValues("gender", genders)
}
