package uuid

import (
	"regexp"
	"stock-controll/internal/domain/services/validate"
	"strings"

	"github.com/google/uuid" // TODO: problema de dependência externa
)

type UUID struct {
	value  string
	prefix string
}

func New(prefix string) (*UUID, error) {
	err := validatePrefix(prefix)
	if err != nil {
		return nil, err
	}
	id, _ := uuid.NewV7()
	return &UUID{
		prefix: strings.ToLower(prefix),
		value:  id.String(),
	}, nil
}

// aqui recebemos uma injeção de dependência
// recebemos uma função que parseia com base num prefixo
func Parse(fieldName, uuid string) (*UUID, error) {
	err := IsValid(fieldName, uuid)
	if err != nil {
		return nil, err
	}
	return &UUID{
		value: uuid,
	}, nil
}

func (u *UUID) String() string {
	return u.value
}

const (
	uuidLength           = 36
	ErrInvalidUUIDFormat = "ERR_INVALID_UUID_FORMAT"
	re                   = `^(\w{8})(\-)(\w{4})(\-)(\w{4})(\-)(\w{4})(\-)(\w{12})`
)

func IsValid(fieldName, uuid string) error {
	return validate.New(fieldName, uuid,
		validate.IsBlank(),
		validate.IsLengthEqualTo(uuidLength),
		validate.IsFormatValid(regexp.MustCompile(re), ErrInvalidUUIDFormat),
	)
}

const (
	prefixLength = 3
	prefixRegExp = `(^[a-zA-Z]{3}$)$`
)

func validatePrefix(prefix string) error {
	validate.New("prefix", prefix,
		validate.IsBlank(),
		validate.IsLengthEqualTo(prefixLength),
		validate.IsFormatValid(prefixRegExp /* CODE */),
	)
	return nil
}
