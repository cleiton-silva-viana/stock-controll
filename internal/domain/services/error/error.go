package validationerrors

/*
	"validationsErrors": {
		"first_name": {
			"Code": "ERR_NAME_TOO_SHORT"
			"Message": "O nome informado é muito curto.",
			"Solution": "Certifique-se de que o nome tenha pelo menos 2 caracteres.",
			"Details": "Um nome válido deve ter no mínimo 2 letras. Nomes muito curtos podem não ser reconhecidos.",
			"Documentation": https://...
		},
		"gender": {
			"Code": "ERR_GENDER_CANNOT_BE_EMPTY"
			"Message": "O gênero informado não pode ser vazio.",
			"Solution": "Por favor, insira um gênero válido.",
			"Details": "Um gênero deve ser fornecido. Gêneros vazios não são aceitos.",
			"Documentation": https://...
		},
	},
*/

import (
	"fmt"
	"stock-controll/internal/domain/services/validate"
)

type ValidationError struct {
	entityName string
	errors     []validate.FieldError
}

// TODO: Validar o nome da entidade
func New(entityName string) *ValidationError {
	return &ValidationError{
		entityName: entityName,
		errors:     make([]validate.FieldError, 0),
	}
}

func (ve *ValidationError) Error() string {
	var errs string

	for _, err := range ve.errors {
		errs += fmt.Sprintf("%s\n", err.CodeError)
	}

	return errs
}

// DUVIDA: devemos fazer o que caso o erro enviado seja do tipo nil?
func (ve *ValidationError) AddValidationError(err error) *ValidationError {
	if err == nil {
		return ve
	}
	if fieldError, ok := err.(*validate.FieldError); ok {
		ve.errors = append(ve.errors, *fieldError)
	}
	return ve
}

func (e *ValidationError) Errors() []validate.FieldError {
	var errorsCopy = make([]validate.FieldError, len(e.errors))
	copy(errorsCopy, e.errors)
	return errorsCopy
}

func (e *ValidationError) HasError() bool {
	return len(e.errors) > 0
}

/*
{
	"ErrorSummary": {
		"Code": "ERR_USER_CREATION_FAILED",
		"ErrorIn": "User",
		"Operation": "Create a new user",
		"Message": "Falha ao criar o usuário devido a erros de validação e autenticação.",
		"TimesTamp": "2023-10-01T12:34:56Z"
		"Context": {
			"UserUUID": "123456"
			"SessionUUID": "abcde=12345"
		}
		"ErrorID": "ERR-20231001-001"
	}

	"ErrorDetails": {

		"AuthenticationErrors": {
			"password": {
				"Code": "ERR_PASSWORD_TOO_WEAK",
				"Message": "",
				"Details": "",
			},
		}

		"AuthorizationErrors": {
			"access": {
				"Code": "ERR_ACCESS_DENIED",
				"Message": "",
				"Details": "",
			},
		},
	},
}

*/
