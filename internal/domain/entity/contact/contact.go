package contact

import (
	"regexp"

	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type IContact interface {
	UUID() string
	Email() string
	SetEmail(email string) error
	Phone() string
	SetPhone(phone string) error
}

type Contact struct {
	uuid  string
	phone string
	email string
}

func New(email, phone string) (IContact, error) {
	var contactInstance = Contact{
		uuid: uuid.New(),
	}

	var err = validationerrors.New("contact").
		AddValidationError(contactInstance.SetEmail(email)).
		AddValidationError(contactInstance.SetPhone(phone))

	if err.HasError() {
		return nil, err
	}

	return &contactInstance, nil
}

func (c *Contact) UUID() string {
	return c.uuid
}

func (c *Contact) Email() string {
	return c.email
}

const (
	emailMinLength          = 12
	emailMaxLength          = 255
	ErrCPFWithInvalidFormat = "ERR_CPF_WITH_INVALID_FORMAT"
)

func (c *Contact) SetEmail(email string) error {
	const re = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	var err = validate.New("email", email,
		validate.IsBlank(),
		validate.IsLengthInRange(emailMinLength, emailMaxLength),
		validate.IsFormatValid(regexp.MustCompile(re), ErrCPFWithInvalidFormat),
	)
	if err == nil {
		c.email = email
	}
	return err
}

func (c *Contact) Phone() string {
	return c.phone
}

const (
	PhoneMinLength            = 13
	PhoneMaxLength            = 14
	ErrPhoneWithInvalidFormat = "ERR_PHONE_WITH_INVALID_FORMAT"
)

func (c *Contact) SetPhone(phone string) error {
	const re = `^[.(](\d{2})[.)](\d{4,5})[.-](\d{4}$)`

	var err = validate.New("phone", phone,
		validate.IsBlank(),
		validate.IsLengthInRange(PhoneMinLength, PhoneMaxLength),
		validate.CheckLetters(validate.Disallow),
		validate.IsFormatValid(regexp.MustCompile(re), ErrPhoneWithInvalidFormat),
	)

	if err == nil {
		c.phone = phone
	}
	return err
}
