package feature

import "fmt"

type Error struct {
	Status      string
	Message     string
	Errors_list []error
}

func (e *Error) Error() string {
	return fmt.Sprintf("status: %s\nmessage: %s\nerrors_list: %s", e.Status, e.Message, e.Errors_list)
}

func NewFeatureError(message string, errors []error) error {
	return &Error{
		Status: "error",
		Message: message,
		Errors_list: errors,
	}
}
