package failure

import "fmt"

func NameIsShort(fieldName string, minimunLength int) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the %s cannot contain less than %d letters", fieldName, minimunLength),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}

func NameIsLong(fieldName string, maxLength int) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the %s cannot contain more than %d chars", fieldName, maxLength),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}

func NameIsEmpty(fieldName string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the %s is empty", fieldName),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}

func NameWithInvalidChars(fieldName string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the %s cannot contains special characters", fieldName),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}

func NameWithNumber(fieldName string) error {
	return &Field{
		Field: fieldName,
		Description: Description{
			Message:  fmt.Sprintf("the %s cannot contains numbers", fieldName),
			Solution: fmt.Sprintf("Check if %s is correct", fieldName),
		},
	}
}
