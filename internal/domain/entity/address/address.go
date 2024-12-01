package address

import (
	"regexp"

	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

// TODO: Adicionar os setters
type IAddress interface {
	Street() string
	Number() int
	Complement() string
	City() string
	State() string
	PostalCode() string
}

type Address struct {
	uuid       string
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

func New(config Config) (IAddress, *validationerrors.ValidationError) {
	a := Address{
		uuid: uuid.New(),
	}

	err := validationerrors.New("address").
		AddValidationError(a.SetStreet(config.Street)).
		AddValidationError(a.SetCity(config.City)).
		AddValidationError(a.SetState(config.State)).
		AddValidationError(a.SetPostalCode(config.PostalCode)).
		AddValidationError(a.SetComplement(config.Complement)).
		AddValidationError(a.SetNumber(config.Number))

	if err.HasError() {
		return nil, err
	}
	return &a, nil
}

func (a *Address) Street() string {
	return a.street
}

const (
	minStreetNameLength = 2
	maxStreetNameLength = 40
)

func (a *Address) SetStreet(street string) error {
	err := validate.New[string]("street", street,
		validate.IsBlank(),
		validate.IsLengthInRange(minStreetNameLength, maxStreetNameLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err == nil {
		a.street = street
	}
	return err
}

func (a *Address) Number() int {
	return a.number
}

const (
	minNumberHome = 1
	maxNumberHome = 100000
)

func (a *Address) SetNumber(number int) error {
	err := validate.New[int]("number", number,
		validate.IsInRange(minNumberHome, maxNumberHome),
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

func (a *Address) Complement() string {
	return a.complement
}

func (a *Address) SetComplement(complement string) error {
	err := validate.New[string]("complement", complement,
		validate.IsBlank(),
		validate.IsLengthInRange(minComplementLength, maxComplementLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err == nil {
		a.complement = complement
	}
	return err
}

func (a *Address) City() string {
	return a.city
}

const (
	minCityNameLength = 3
	maxCityNameLength = 50
)

func (a *Address) SetCity(city string) error {
	err := validate.New[string]("city", city,
		validate.IsBlank(),
		validate.CheckSpecialChars(validate.Disallow),
		validate.IsLengthInRange(minCityNameLength, maxCityNameLength),
	)
	if err == nil {
		a.city = city
	}
	return err
}

func (a *Address) State() string {
	return a.state
}

const (
	minStateNameLength = 3
	maxStateNameLength = 60
)

func (a *Address) SetState(state string) error {
	err := validate.New[string]("state", state,
		validate.IsBlank(),
		validate.IsLengthInRange(minStateNameLength, maxStateNameLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err == nil {
		a.state = state
	}
	return err
}

func (a *Address) PostalCode() string {
	return a.postalCode
}

const ErrAddressWithInvalidZipCodeFormat = "ERR_ADDRESS_WITH_INVALID_ZIP_CODE_FORMAT"

func (a *Address) SetPostalCode(code string) error {
	re := `^\d{5}\-\d{3}$`
	err := validate.New[string]("postal_code", code,
		validate.IsBlank(),
		validate.IsFormatValid(regexp.MustCompile(re), ErrAddressWithInvalidZipCodeFormat),
	)
	if err == nil {
		a.postalCode = code
	}
	return err
}
