package name

import "stock-controll/internal/domain/services/validate"

type Name struct {
	firstName string
	lastName  string
}

func New(firstName, lastName string) (*Name, []error) {
	var errs = make([]error, 0, 2)	

	firstNameError := validateName("first_name", firstName)
	var lastNameError error

	if firstNameError != nil {
		errs = append(errs, firstNameError)
	}
	
	if lastName != " " {
		lastNameError = validateName("last_name", lastName)
	}
	
	if lastNameError != nil {
		errs = append(errs, lastNameError)
	}	

	if len(errs) > 0 {
		return nil, errs
	}

	return &Name{
		firstName: firstName,
		lastName: lastName,
	}, nil
}

func (n *Name) FirstName() string {
	return n.firstName
}

func (n *Name) LastName() string {
	return n.lastName
}

const (
	minLength = 2
	maxLength = 16
)

func validateName(fieldName, name string) error {
	return validate.New[string](fieldName, name,
		validate.IsLengthInRange(minLength, maxLength),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
}
