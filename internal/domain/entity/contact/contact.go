package contact

import (
	"regexp"
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
)

type IContact interface {
	GetEmail() string
	GetPhone() string
	SetEmail(email string) *validation.FieldError
	SetPhone(phone string) *validation.FieldError
	GetUUID() string
}

type contact struct {
	uuid string
	phone string
	email string
}

func NewContact(email, phone string) (IContact, validationError.IValidationError) {
	var err = validationError.NewValidationError("contact")
	var c = contact{
		uuid: common.GenerateUUID(),
	}

	err.
		AddValidationError(c.SetEmail(email)).
		AddValidationError(c.SetPhone(phone))

	if err.HasError() {
		return nil, err
	}

	return &c, nil
}

func (c *contact) GetUUID() string {
	return c.uuid
}

func (c *contact) GetEmail() string {
	return c.email
}

const (
	emailMinLength = 12
	emailMaxLength = 255
)

func (c *contact) SetEmail(email string) *validation.FieldError {
	const re = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	var err = validation.Validate("email", email,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(emailMinLength, emailMaxLength, validation.ErrUnknown),
		validation.IsFormatValid(regexp.MustCompile(re), validation.ErrUnknown),
	)

	if err == nil {
		c.email = email
	}
	return err
}

func (c *contact) GetPhone() string {
	return c.phone
}

const (
	PhoneMinLength = 13
	PhoneMaxLength = 14
)

func (c *contact) SetPhone(phone string) *validation.FieldError {
	const re = `^[.(](\d{2})[.)](\d{4,5})[.-](\d{4}$)`

	var err = validation.Validate("phone", phone,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(PhoneMinLength, PhoneMaxLength, validation.ErrUnknown),
		validation.CheckLetters(validation.Disallow, validation.ErrUnknown),
		validation.IsFormatValid(regexp.MustCompile(re), validation.ErrUnknown),
	)

	if err == nil {
		c.phone = phone
	}
	return err
}
