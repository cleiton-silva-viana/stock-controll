package address

import (
	"regexp"
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
)

type IAddress interface {
	GetStreet() string
	GetNumber() int
	GetComplement() string
	GetCity() string
	GetState() string
	GetPostalCode() string
}

type address struct {
	uuid       string
	street     string
	number     int
	complement string
	city       string
	state      string
	postalCode string
}

func NewAddress(street, city, state, postalCode, complement string, number int) (IAddress, validationError.IValidationError) {
	var err = validationError.NewValidationError("address")
	var a = address{
		uuid: common.GenerateUUID(),
	}

	err.
		AddValidationError(a.SetStreet(street)).
		AddValidationError(a.SetCity(city)).
		AddValidationError(a.SetState(state)).
		AddValidationError(a.SetPostalCode(postalCode)).
		AddValidationError(a.SetComplement(complement)).
		AddValidationError(a.SetNumber(number))

	if err.HasError() {
		return nil, err
	}
	return &a, nil
}

func (a *address) GetStreet() string {
	return a.street
}

const (
	minStreetNameLength = 2
	maxStreetNameLength = 40
)

func (a *address) SetStreet(street string) *validation.FieldError {
	err := validation.Validate[string]("street", street,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minStreetNameLength, maxStreetNameLength, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
	if err == nil {
		a.street = street
	}
	return err
}

func (a *address) GetNumber() int {
	return a.number
}

const (
	minNumberHome = 1
	maxNumberHome = 100000
)

func (a *address) SetNumber(number int) *validation.FieldError {
	err := validation.Validate[int]("number", number,
		validation.IsInRange(minNumberHome, maxNumberHome, validation.ErrUnknown),
	)
	if err == nil {
		a.number = number
	}
	return err
}

const (
	minComplementLength = 0
	maxComplementLength = 150
)

func (a *address) GetComplement() string {
	return a.complement
}

func (a *address) SetComplement(complement string) *validation.FieldError {
	err := validation.Validate[string]("complement", complement,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minComplementLength, maxComplementLength, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
	if err == nil {
		a.complement = complement
	}
	return err
}

func (a *address) GetCity() string {
	return a.city
}

const (
	minCityNameLength = 3
	maxCityNameLength = 50
)

func (a *address) SetCity(city string) *validation.FieldError {
	err := validation.Validate[string]("city", city,
		validation.IsBlank(validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
		validation.IsLengthInRange(minCityNameLength, maxCityNameLength, validation.ErrUnknown),
	)
	if err == nil {
		a.city = city
	}
	return err
}

func (a *address) GetState() string {
	return a.state
}

const (
	minStateNameLength = 3
	maxStateNameLength = 60
)

func (a *address) SetState(state string) *validation.FieldError {
	err := validation.Validate[string]("state", state,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minStateNameLength, maxStateNameLength, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
	if err == nil {
		a.state = state
	}
	return err
}

func (a *address) GetPostalCode() string {
	return a.postalCode
}

// Checar formato de código postal
func (a *address) SetPostalCode(code string) *validation.FieldError {
	re := `^\d{5}\-\d{3}$`
	err := validation.Validate[string]("postal_code", code,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsFormatValid(regexp.MustCompile(re), validation.ErrUnknown),
	)
	if err == nil {
		a.postalCode = code
	}
	return err
}

type addressBuilder struct {
	street     string
	number     int
	complement string
	city       string
	state      string
	postalCode string
}

func NewAddressBuilder() *addressBuilder {
	return &addressBuilder{}
}

func (a *addressBuilder) SetStreet(street string) *addressBuilder {
	a.street = street
	return a
}

func (a *addressBuilder) SetNumber(number int) *addressBuilder {
	a.number = number
	return a
}

func (a *addressBuilder) SetComplement(complement string) *addressBuilder {
	a.complement = complement
	return a
}

func (a *addressBuilder) SetCity(city string) *addressBuilder {
	a.city = city
	return a
}

func (a *addressBuilder) SetState(state string) *addressBuilder {
	a.state = state
	return a
}

func (a *addressBuilder) SetPostalCode(postalCode string) *addressBuilder {
	a.postalCode = postalCode
	return a
}

func (a *addressBuilder) Build() (IAddress, validationError.IValidationError) {
	return NewAddress(a.street, a.city, a.state, a.postalCode, a.complement, a.number)
}
