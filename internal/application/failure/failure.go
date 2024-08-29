package failure

import "fmt"

type Error struct {
	Status      int
	Message     string
	ErrorList []error
}

func (e *Error) Error() string {
	return fmt.Sprintf("status: %s\nmessage: %s\nerrors_list: %s", e.Status, e.Message, e.ErrorList)
}

func NewFeatureError(statusHTTP int,message string, errors []error) error {
	return &Error{
		Status: statusHTTP,
		Message: message,
		ErrorList: errors,
	}
}
