package valueobject

import (
	"errors"
	"strings"

	"stock-controll/internal/domain/failure"
	"stock-controll/internal/domain/validation"
)


func NameDefaultValidation(fieldName, fieldValue string) (*Name, error) {
	builder := NewNameBuilder()
	nameValidated, err := builder.Field(fieldName, fieldValue).
								Length(3, 25).
								Digts(false).
								SpecialCharacters(false).
								Build()
	if err != nil {
		return nil, err
	}
	return nameValidated, nil
}

type Name struct {
	name string
}

type NameBuilder struct {
	field_name    string
	field_value   string
	min_length    int
	max_length    int
	digits        bool
	special_chars bool
	err           error
}

func NewNameBuilder() *NameBuilder {
	return &NameBuilder{}
}

func (n *NameBuilder) Field(field_name, field_value string) *NameBuilder {
	n.field_name = field_name
	n.field_value = field_value
	return n
}

func (n *NameBuilder) Length(min, max int) *NameBuilder {
	n.min_length = min
	n.max_length = max
	return n
}

func (n *NameBuilder) SpecialCharacters(permitted bool) *NameBuilder {
	n.special_chars = permitted
	return n
}

func (n *NameBuilder) Digts(permitted bool) *NameBuilder {
	n.digits = permitted
	return n
}

func (n *NameBuilder) Build() (*Name, error) {
	if n.field_name == "" || n.field_value == "" ||
		n.min_length == 0 || n.max_length == 0 {
		return nil, errors.New("the configs to NameBuilder are empyts!!!")
	}

	nameParsed := strings.Trim(n.field_value, " ")

	if nameParsed == "" {
		return nil, failure.FieldIsEmpty(n.field_name)
	}

	if len(nameParsed) < n.min_length {
		return nil, failure.FieldIsShort(n.field_name, n.min_length)
	}

	if len(nameParsed) > n.max_length {
		return nil, failure.FieldIsLong(n.field_name, n.max_length)
	}

	nameWithoutSpaces := strings.ReplaceAll(nameParsed, " ", "")
	if !n.special_chars {
		if validation.ContainsSpecialChars(nameWithoutSpaces) {
			return nil, failure.FieldWithSpecialChars(n.field_name)
		}
	}

	if !n.digits {
		if validation.ContainsNumbers(nameParsed) {
			return nil, failure.FieldWithNumber(n.field_name)
		}
	}

	return &Name{
		name: nameParsed,
	}, nil
}
