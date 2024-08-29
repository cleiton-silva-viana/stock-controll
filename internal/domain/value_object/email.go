package valueobject

import (
	"regexp"
	
	"stock-controll/internal/domain/failure"
)

type Email struct {
	email string
}

func NewEmail(email string) (*Email, error) {
	err := validateEmailFormat(email)
	if err != nil {
		return nil, err
	}
	return &Email{
		email: email,
	}, nil
}

func (e *Email) GetEmail() string {
	return e.email
}

func validateEmailFormat(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	isInvalid := !re.MatchString(email)
	if isInvalid {
		return failure.FieldWithInvalidFormat("email", "example@domain.com")
	}
	return nil
}
