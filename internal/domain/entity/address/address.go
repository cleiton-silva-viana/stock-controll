package address

import (
	"regexp"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type IAddress interface {
	Street() string
	Number() int
	Complement() string
	City() string
	State() string
	PostalCode() string
}

type Address struct {
	uuid.UUID
	street     string
	number     int
	complement string
	city       string
	state      string
	postalCode string
}

type Config struct {
	Street     string
	Number     int
	Complement string
	City       string
	State      string
	PostalCode string
}

func New(config Config) (IAddress, error) {

	entityError := entity.Error("address").
		AddValidationError(validateStreet(config.Street)).
		AddValidationError(validateCity(config.City)).
		AddValidationError(validateState(config.State)).
		AddValidationError(validatePostalCode(config.PostalCode)).
		AddValidationError(validateComplement(config.Complement)).
		AddValidationError(validateNumber(config.Number))

	if entityError.HasError() {
		return nil, entityError
	}

	return &Address{
		UUID:       *uuid.New(),
		street:     config.Street,
		number:     config.Number,
		complement: config.Complement,
		city:       config.City,
		state:      config.State,
		postalCode: config.PostalCode,
	}, nil
}

func (a *Address) Number() int {
	return a.number
}

func (a *Address) Street() string {
	return a.street
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) State() string {
	return a.state
}

func (a *Address) PostalCode() string {
	return a.postalCode
}

func (a *Address) Complement() string {
	return a.complement
}

const (
	minStreetNameLength = 2
	maxStreetNameLength = 40
)

func validateStreet(street string) error {
	return validate.New[string](
		"street", street,
		validate.IsBlank(),
		validate.IsLengthInRange(minStreetNameLength, maxStreetNameLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
}

const (
	minNumberHome = 1
	maxNumberHome = 100000
)

func validateNumber(number int) error {
	return validate.New[int](
		"number", number,
		validate.IsInRange(minNumberHome, maxNumberHome),
	)
}

const (
	minComplementLength = 0
	maxComplementLength = 150
)

func validateComplement(complement string) error {
	return validate.New[string](
		"complement", complement,
		validate.IsBlank(),
		validate.IsLengthInRange(minComplementLength, maxComplementLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
}

const (
	minCityNameLength = 3
	maxCityNameLength = 50
)

func validateCity(city string) error {
	return validate.New[string](
		"city", city,
		validate.IsBlank(),
		validate.CheckSpecialChars(validate.Disallow),
		validate.CheckNumbers(validate.Disallow),
		validate.IsLengthInRange(minCityNameLength, maxCityNameLength),
	)
}

const (
	minStateNameLength = 3
	maxStateNameLength = 60
)

func validateState(state string) error {
	return validate.New[string](
		"state", state,
		validate.IsBlank(),
		validate.IsLengthInRange(minStateNameLength, maxStateNameLength),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
}

const ErrAddressWithInvalidZipCodeFormat = "ERR_ADDRESS_WITH_INVALID_ZIP_CODE_FORMAT"

func validatePostalCode(code string) error {
	re := `^\d{5}\-\d{3}$`
	return validate.New[string](
		"postal_code", code,
		validate.IsBlank(),
		validate.IsFormatValid(regexp.MustCompile(re), ErrAddressWithInvalidZipCodeFormat),
	)
}
