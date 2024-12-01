package uuid

// TODO: Remover da camada de domínio dependencia externa

import (
	"regexp"
	"stock-controll/internal/domain/services/validate"

	"github.com/google/uuid" // <<<< problema de dependência externa
)

func New() string {
	uuid, _ := uuid.NewV7()
	return uuid.String()
}

const (
	re                       = `^(\w{8})(\-)(\w{4})(\-)(\w{4})(\-)(\w{4})(\-)(\w{12})`
	uuidLength               = 36
	ErrUUIDWithInvalidFormat = "ERR_UUID_WITH_INVALID_FORMAT"
)

func IsValid(fieldName, uuid string) error {
	return validate.New(fieldName, uuid,
		validate.IsBlank(),
		validate.IsLengthEqualTo(uuidLength),
		validate.IsFormatValid(regexp.MustCompile(re), ErrUUIDWithInvalidFormat),
	)
}
