package position

import (
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type Position struct {
	uuid uuid.UUID
	name string
}

const (
	minNameLength = 6
	maxNameLength = 18
)

func New(name string) (*Position, error) {
	errs := entity.Error("position").
		AddValidationError(validateName(name))

	if errs.HasError() {
		return nil, errs
	}

	return &Position{
		uuid: *uuid.New(),
		name: name,
	}, nil
}

func (p *Position) UUID() string {
	return p.uuid.String()
}

func (p *Position) Name() string {
	return p.name
}

func validateName(name string) error {
	return validate.New(
		"name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLength, maxNameLength),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow))
}
