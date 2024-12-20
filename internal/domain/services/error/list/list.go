package list

import (
	"stock-controll/internal/domain/services/error/field"
)

type ItemError struct {
	ErrorIn    map[interface{}]string
	CodeErrors map[string]field.FieldError
}

func itemError() *ItemError {
	return &ItemError{
		ErrorIn:    make(map[interface{}]string),
		CodeErrors: make(map[string]field.FieldError),
	}
}

func (ie *ItemError) AddError(err error) *ItemError {
	if err == nil {
		return ie
	}
	switch err := err.(type) {
	case *field.FieldError:
		ie.ErrorIn[err.InvalidValue] = err.CodeError
		ie.CodeErrors[err.CodeError] = *err
	}
	return ie
}

type ListError struct {
	FieldName string
	ItemError
}

func Error(fieldName string) *ListError {
	return &ListError{
		FieldName: fieldName,
		ItemError: *itemError(),
	}
}

// Melhorar implementação
func (le *ListError) Error() string {
	return le.FieldName
}

func (le *ListError) HasError() bool {
	return len(le.ErrorIn) > 0
}

/*
Example:
	"keywords": {
		"error_in": {
			"manga": {
				"error_code": "ERR_DUPLICATED_KEYWORD"
			},
			"a": {
				"error_code": "ERR_LENGTH_TOO_SHORT"
			},
			"pesico%": {
				"error_code": "ERR_SPECIAL_CHARACTES"
			},
			"@#%$#$": {
				"error_code": "ERR_SPECIAL_CHARACTES"
			}
		},
		"code_errros": {
			"ERR_DUPLICATED_KEYWORD": {
				"message": "duplicate keyword in the list",
				"tip": "remove duplicated keywork of the list",
				"documentation": "https://docs.ecommerce/error/err_field_cannot_be_empty.com"
			},
			"ERR_LENGTH_TOO_SHORT": {
				"message": "the length of keyword is too short than minimum allowed",
				"tip": "use a keywork with length greater than 2 characters",
				"documentation": "https://docs.ecommerce/error/err_field_cannot_be_empty.com"
			},
			"ERR_SPECIAL_CHARACTERS": {
				"message": "the keyword dont have special characters",
				"tip": "use a keyword with letter & number only",
				"documentation": "https://docs.ecommerce/error/err_field_cannot_be_empty.com"
			}
			
		}
	}
*/
