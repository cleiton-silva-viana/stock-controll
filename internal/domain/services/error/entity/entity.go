package entity

import (
	"fmt"
	"stock-controll/internal/domain/services/error/field"
)

type EntityError struct {
	entityName    string
	errors        []field.FieldError // , FieldErrors, ListErrors
	unknownErrors []error
}

func Error(entityName string) *EntityError {
	return &EntityError{
		entityName: entityName,
		errors:     make([]field.FieldError, 0),
	}
}

func (ee *EntityError) Error() string {
	var message string

	for _, err := range ee.errors {
		message += fmt.Sprintf("%s\n", err.Error())
	}

	return message
}

func (ee *EntityError) AddValidationError(err error) *EntityError {
	if err == nil {
		return ee
	}
	switch err := err.(type) {
	case *field.FieldError:
		ee.errors = append(ee.errors, *err)
	default:
		ee.unknownErrors = append(ee.unknownErrors, err)
	}
	return ee
}

func (ee *EntityError) Errors() []field.FieldError {
	var errorsCopy = make([]field.FieldError, len(ee.errors))
	copy(errorsCopy, ee.errors)
	return errorsCopy
}

func (ee *EntityError) HasError() bool {
	return len(ee.errors) > 0 || len(ee.unknownErrors) > 0
}

/*
{
	"error_summary":
	{
		"code": "ERR_USER_CREATION_FAILED",
		"error_in": "User", <<<<<<<<<< nome da entidade aqui
		"operation": "Create a new user", <<<<<<<<<< nome do use case aqui
		"message": "Falha ao criar o usuário devido a erros de validação e autenticação.",
		"timestamp": "2023-10-01T12:34:56Z",
		"context":
		{
			"user_uuid": "123456",
			"session_uuid": "abcde=12345"
		},
		"error_id": "ERR-20231001-001"
	},

	"error_details":
	{
		"error_validations": <<<<<<<< field error & field errors
		{
			"first_name": <<<<<<<<<< nome do atributo com dados incosistentes
			{
				"code": "ERR_FIELD_TOO_SHORT",
				"message": "O nome informado é muito curto.",
				"solution": "Certifique-se de que o nome tenha pelo menos 2 caracteres.",
				"details": "Um nome válido deve ter no mínimo 2 letras. Nomes muito curtos podem não ser reconhecidos.",
				"documentation": "https://docs.ecommerce/error/err_field_too_short.com"
			},
			"gender":
			{
				"code": "ERR_FIELD_CANNOT_BE_EMPTY",
				"message": "O gênero informado não pode ser vazio.",
				"solution": "Por favor, insira um gênero válido.",
				"details": "Um gênero deve ser fornecido. Gêneros vazios não são aceitos.",
				"documentation": "https://docs.ecommerce/error/err_field_cannot_be_empty.com"
			},
			"roles":
			{
				""
			}
		},

		"authentication_errors":
		{
			"password":
			{
				"Code": "ERR_PASSWORD_TOO_WEAK",
				"Message": "",
				"Details": "",
				"Documentation": "url"
			}
		},

		"authorization_errors":
		{
			"access":
			{
				"Code": "ERR_ACCESS_DENIED",
				"Message": "",
				"Details": "",
				"Documentation": "url"
			}
		}
	}
}

*/
