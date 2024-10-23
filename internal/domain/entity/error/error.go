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
	"sync"

	"stock-controll/internal/domain/validation"
)

type IValidationError interface {
	AddValidationError(error *validation.FieldError) *validationError
	GetErrors() []validation.FieldError
	HasError() bool
}

type validationError struct {
	entity string
	errors []validation.FieldError
	mu     sync.Mutex
}

// Valida ro nome da entidade
func NewValidationError(entityName string) IValidationError {
	return &validationError{
		entity: entityName,
		errors: make([]validation.FieldError, 0),
	}
}

func (e *validationError) AddValidationError(error *validation.FieldError) *validationError {
	if error == nil {
		return e
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.errors = append(e.errors, *error)
	return e
}

func (e *validationError) GetErrors() []validation.FieldError {
	e.mu.Lock()
	defer e.mu.Unlock()

	var errorsCopy = make([]validation.FieldError, len(e.errors))
	copy(errorsCopy, e.errors)
	return errorsCopy
}

func (e *validationError) HasError() bool {
	e.mu.Lock()
	defer e.mu.Unlock()

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

