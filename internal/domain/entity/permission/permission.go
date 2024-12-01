package permission

// TODO: Usar entityError

import (
	validationError "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/validate"
)

type Permission struct {
	uuid string
	name string
}

func New(name string) (*Permission, error) {
	var permissionError = validationError.New("permission")
	var permissionInstance = Permission{}

	permissionError.
		AddValidationError(permissionInstance.SetName(name))

	if permissionError.HasError() {
		return nil, permissionError
	}
	return &permissionInstance, nil
}

func (p *Permission) UUID() string {
	return p.uuid
}

func (p *Permission) Name() string {
	return p.name
}

const (
	minNameLength = 6
	maxNameLength = 24
)

func (p *Permission) SetName(name string) error {
	err := validate.New[string]("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLength, maxNameLength),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err == nil {
		p.name = name
	}
	return err
}

func (p *Permission) Equals(other *Permission) bool {
	if other == nil {
		return false
	}
	return p.uuid == other.uuid && p.name == other.name
}
