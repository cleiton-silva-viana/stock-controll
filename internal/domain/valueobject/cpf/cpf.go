package cpf

import (
	"regexp"
	"stock-controll/internal/domain/services/validate"
)

type CPF struct {
	cpf string
}

func New(cpf string) (*CPF, error) {
	err := validateCPF(cpf)
	if err != nil {
		return nil, err
	}
	return &CPF{
		cpf: cpf,
	}, nil
}

func (c *CPF) CPF() string {
	return c.cpf
}

const (
	length                  = 14
	ErrCPFWithInvalidFormat = "ERR_CPF_WITH_INVALID_FORMAR"
	re                      = `^\d{3}\.\d{3}\.\d{3}\-\d{2}$`
)

func validateCPF(cpf string) error {
	return validate.New[string]("cpf", cpf,
		validate.IsBlank(),
		validate.IsLengthEqualTo(length),
		validate.IsFormatValid(regexp.MustCompile(re), ErrCPFWithInvalidFormat),
	)
}
