package value_object

import (
	"regexp"
	"stock-controll/internal/contact/errors"
)

type Email struct {
	email string
}

func NewEmail(email string) (*Email, error) {
	err := validateEmailFormat(email)
	if err != nil {
		return nil, errors.InvalidEmailFormat
	}
	return &Email{
		email: email,
	}, nil
}

func validateEmailFormat(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	isInvalid := !re.MatchString(email)
	if isInvalid {
		return errors.InvalidEmailFormat
	}
	return nil
}
