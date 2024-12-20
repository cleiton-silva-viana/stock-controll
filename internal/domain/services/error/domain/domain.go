package domain

import (
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/error/list"
)

type DomainError struct {
	TypedErrors  []error // Field error, field errors ou List error
	UnknowErrors []error
	ErrorCodes []string // somente erros de domínio
}

func Error() *DomainError {
	return &DomainError{
		TypedErrors:  make([]error, 0, 2),
		UnknowErrors: make([]error, 0, 0),
	}
}

// TODO: implementar
func (de *DomainError) Error() string {
	return ""
}

func (de *DomainError) AddError(err error) *DomainError {
	if err == nil {
		return de
	}

	switch err := err.(type) {
	case *field.FieldError:
	// case *field.FieldErrors:
	case *list.ListError:
		de.TypedErrors = append(de.TypedErrors, err)
	default:
		de.UnknowErrors = append(de.UnknowErrors, err)
	}
	return de
}

func (de *DomainError) HasError() bool {
	return len(de.TypedErrors) > 0 || len(de.UnknowErrors) > 0
}

/*
	domain_error
	{
		keywords: <<<< list error
		{
		code_error: "ERR_INVALID_KEYWORDS":
		{
			message: ""
			details:
			{
				"cOC#":
				{
					code_error:	"ERR_CHARACTERS_SPECIAL_CANNOT_BE_ONLY"
					message: ""
				},
				"pepsi":
				{
					code_error: "ERR_DUPLICATED_KEYWORD"
					message: ""
				}
			}
		}
		name:
		{
			code_error: ERR_INVALID_USER_NAME
			message: ""
		}
	}
*/
