package position

import (
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type Position struct {
	uuid string
	name string
}

const (
	minPositionNameLength = 6
	maxPositionNameLength = 18
)

func New(name string) (*Position, error) {
	positionInstance := &Position{
		uuid: uuid.New(),
		name: "",
	}

	errs := validationerrors.
		New("position").
		AddValidationError(positionInstance.SetName(name))

	if errs.HasError() {
		return nil, errs
	}

	return positionInstance, nil
}

func (p *Position) UUID() string {
	return p.uuid
}

func (p *Position) Name() string {
	return p.name
}

func (p *Position) SetName(name string) error {
	err := validate.New("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minPositionNameLength, maxPositionNameLength),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow))
	if err == nil {
		p.name = name
	}
	return err
}
