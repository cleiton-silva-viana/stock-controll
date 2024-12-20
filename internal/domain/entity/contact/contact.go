package contact

import (
	"regexp"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type IContact interface {
	UUID() string
	Email() string
	SetEmail(email string) error
	Phone() string
	SetPhone(phone string) error
}

type Contact struct {
	uuid uuid.UUID
	phone string
	email string
}

func New(email, phone string) (*Contact, error) {

	var err = entity.Error("contact").
		AddValidationError(validateEmail(email)).
		AddValidationError(validatePhone(phone))

	if err.HasError() {
		return nil, err
	}

	return &Contact{
		uuid:  *uuid.New(),
		phone: phone,
		email: email,
	}, nil
}

func (c *Contact) UUID() string {
	return c.uuid.String()
}

func (c *Contact) Email() string {
	return c.email
}

func (c *Contact) Phone() string {
	return c.phone
}

const (
	emailMinLength          = 8
	emailMaxLength          = 255
	ErrCPFWithInvalidFormat = "ERR_CPF_WITH_INVALID_FORMAT"
)

func validateEmail(email string) error {
	const re = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	return validate.New(
		"email", email,
		validate.IsBlank(),
		validate.IsLengthInRange(emailMinLength, emailMaxLength),
		validate.IsFormatValid(regexp.MustCompile(re), ErrCPFWithInvalidFormat),
	)
}

const (
	PhoneMinLength            = 13
	PhoneMaxLength            = 14
	ErrPhoneWithInvalidFormat = "ERR_PHONE_WITH_INVALID_FORMAT"
)

func validatePhone(phone string) error {
	const re = `^[.(](\d{2})[.)](\d{4,5})[.-](\d{4}$)`

	return validate.New(
		"phone", phone,
		validate.IsBlank(),
		validate.IsLengthInRange(PhoneMinLength, PhoneMaxLength),
		validate.CheckLetters(validate.Disallow),
		validate.IsFormatValid(regexp.MustCompile(re), ErrPhoneWithInvalidFormat),
	)
}
