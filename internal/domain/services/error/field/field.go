package field

const NoError = "NO_ERROR"

// Tenho que ter um método que verifica se a string de ero possui um %s para que nós coloquemos o valor inválido no erro

type FieldError struct {
	FieldName    string
	CodeError    string
	InvalidValue interface{}
}

func Error(fieldName, CodeError string, invalidValue interface{}) *FieldError {

	return &FieldError{
		FieldName:    fieldName,
		CodeError:    CodeError,
		InvalidValue: invalidValue,
	}
}

func (fe *FieldError) Error() string {
	return fe.CodeError
}

func (fe *FieldError) AddErrorCode(code string) *FieldError {
	if code == NoError {
		return fe
	}
	fe.CodeError = code
	return fe
}

func (fe *FieldError) HasError() bool {
	return fe.CodeError != NoError
}

type FieldErrors struct {
	FieldName     string
	FieldErrors   map[string]struct{}
	UnknownErrors []error
	invalidValue  interface{}
}

func Errors(fieldName string) *FieldErrors {
	return &FieldErrors{
		FieldName:     fieldName,
		FieldErrors:   make(map[string]struct{}, 2),
		UnknownErrors: make([]error, 0, 1),
	}
}

func (fe *FieldErrors) HasError() bool {
	return len(fe.FieldErrors) > 0 || len(fe.UnknownErrors) > 0
}

// TODO: testar
func (fe *FieldErrors) AddError(err error) *FieldErrors {
	if err == nil {
		return fe
	}
	switch err := err.(type) {
	case *FieldError:
		fe.FieldErrors[err.CodeError] = struct{}{}
	default:
		fe.UnknownErrors = append(fe.UnknownErrors, err)
	}
	return fe
}

/*
	<<< field error >>>

	FieldName: value_discount
	Error:
	{
		ErrInvalidFixedValueDiscount:
		{
			Message:
		}
	}
*/

/*
	<<< field errors >>>

	FieldName: discount_strategy
	Errors:
	{
		ErrInvalidFixedValueDiscount:
		{
			Message:
		},
		ErrInvalidPercentageValueDiscount:
		{
			Message:
		}
	}
*/
