package permission

// TODO: Usar entityError

import (
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type Permission struct {
	uuid uuid.UUID
	name string
}

func New(name string) (*Permission, error) {
	var permissionError = entity.Error("permission").
		AddValidationError(validateName(name))

	if permissionError.HasError() {
		return nil, permissionError
	}
	return &Permission{
		uuid: *uuid.New(),
		name: name,
	}, nil
}

func (p *Permission) UUID() string {
	return p.uuid.String()
}

func (p *Permission) Name() string {
	return p.name
}

func (p *Permission) Equals(other *Permission) bool {
	if other == nil {
		return false
	}
	return p.UUID() == other.UUID() && p.name == other.name
}

const (
	minNameLength = 6
	maxNameLength = 24
)

func validateName(name string) error {
	return validate.New[string](
		"name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLength, maxNameLength),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
}
