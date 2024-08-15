package common_errors

import "fmt"

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
