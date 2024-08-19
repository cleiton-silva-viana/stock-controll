package value_object

import (
	"strings"
	
	"stock-controll/internal/domain/failure"
)

type Name struct {
	name string
}

func NewName(name string) (*Name, error) {
	nameParsed := strings.Trim(name, " ")

	if nameParsed == "" {
		return nil, failure.NameIsEmpty
	}

	if len(nameParsed) < 3 {
		return nil, failure.NameIsShort
	}

	if len(nameParsed) > 20 {
		return nil, failure.NameIsLong
	}

	if ContainsNumbers(name) {
		return nil, failure.NameWithInvalidChars
	}

	nameWithotSpaces := strings.ReplaceAll(name, " ", "")
	if ContainsSpecialChars(nameWithotSpaces) {
		return nil, failure.NameWithInvalidChars
	}

	return &Name{
		name: name,
	}, nil
}
