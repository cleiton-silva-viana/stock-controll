package value_object

import (
	"strings"

	"stock-controll/internal/domain/failure"
)

type Name struct {
	name string
}

func NewName(name, fieldName string, maxLength int) (*Name, error) {
	nameParsed := strings.Trim(name, " ")

	if nameParsed == "" {
		return nil, failure.NameIsEmpty(fieldName)
	}

	const MIN_VALUE = 3
	if len(nameParsed) < MIN_VALUE {
		return nil, failure.NameIsShort(fieldName, MIN_VALUE)
	}

	if len(nameParsed) > maxLength {
		return nil, failure.NameIsLong(fieldName, maxLength)
	}

	if ContainsNumbers(name) {
		return nil, failure.NameWithNumber(fieldName)
	}

	nameWithOutSpaces := strings.ReplaceAll(name, " ", "")
	if ContainsSpecialChars(nameWithOutSpaces) {
		return nil, failure.NameWithInvalidChars(fieldName)
	}

	return &Name{
		name: name,
	}, nil
}
