package failure

import (
	"fmt"
	"strings"
)

type Field struct {
	Field       string
	Description Description
}

type Description struct {
	Message  string
	Solution string
}

func (e *Field) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Description.Message)
}

func NewFieldError(field, message, solution string) error {
	return &Field{
		Field: field,
		Description: Description{
			Message:  message,
			Solution: solution,
		},
	}
}

func FieldIsShort(fieldName string, minimunLength int) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the '%s' field cannot contain less than %d characters", fieldName, minimunLength),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}

func FieldIsLong(fieldName string, maxLength int) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the '%s' field cannot contain more than %d chars", fieldName, maxLength),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}

func FieldIsEmpty(fieldName string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the '%s' field is empty", fieldName),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}

func FieldWithSpecialChars(fieldName string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the '%s' field cannot contains special characters", fieldName),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}

func FieldWithLetters(fieldName string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the '%s' field must have contains letters", fieldName),
			Solution: fmt.Sprintf("check the format of %s field", fieldName),
		},
	}
}

func FieldWithoutSpecialChars(fieldName string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the '%s' field must have contains special characters", fieldName),
			Solution: fmt.Sprintf("add special characters on %s field", fieldName),
		},
	}
}

func FieldWithNumber(fieldName string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the '%s' field cannot contains numbers", fieldName),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}

func FieldWithoutNumber(fieldName string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the '%s' field must have contains numbers", fieldName),
			Solution: fmt.Sprintf("add numbers in %s field", fieldName),
		},
	}
}

func FieldWithInvalidFormat(fieldName, validFormat string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the '%s' field is on invalid format", fieldName),
			Solution: fmt.Sprintf("check if %s field contains the format %s", fieldName, validFormat),
		},
	}
}

func FieldNotRangeValues(fieldName string, ranger []string) error {
	rangerStr := strings.Join(ranger, ", ")
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the %s field not contains a valid range values", fieldName),
			Solution: fmt.Sprintf("select values between (%s)s", rangerStr),
		},
	}
}
