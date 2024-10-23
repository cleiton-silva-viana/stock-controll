package brand

import (
	"stock-controll/internal/domain/entity/common"
	errorentity "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
)

type statusBrand bool
const (
	Active = true
	Inactive = false
)

type Brand struct {
	uuid string
	name             string
	description      string
	logo             string
	manufacturerUUID string
	status           statusBrand
	// salesHistory []Sales
}

func NewBrand(name, description, logo, manufacturerUUID string) (*Brand, errorentity.IValidationError) {
	var b = Brand{
		uuid: common.GenerateUUID(),
	}

	var err = errorentity.
		NewValidationError("Brand").
		AddValidationError(b.SetName(name)).
		AddValidationError(b.SetDescription(description)).
		AddValidationError(b.SetLogo(logo)).
		AddValidationError(b.SetManufacturerUUID(manufacturerUUID))

	if err.HasError() {
		return nil, err
	}

	return &b, nil
}

func (b *Brand) GetUUID() string {
	return b.uuid
}

func (b *Brand) GetName() string {
	return b.name
}

const (
	minNameLengthForBrand = 3
	maxNameLengthForBrand = 30
)

func (b *Brand) SetName(name string) *validation.FieldError {
	var err = validation.Validate[string]("name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minNameLengthForBrand, maxNameLengthForBrand, validation.ErrUnknown),
	)

	if err == nil {
		b.name = name
	}
	return err
}

func (b *Brand) GetDescription() string {
	return b.description
}

const (
	minDescriptionLength = 10
	maxDescriptionLength = 250
)

func (b *Brand) SetDescription(description string) *validation.FieldError {
	var err = validation.Validate[string]("description", description,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minDescriptionLength, maxDescriptionLength, validation.ErrUnknown),
	)

	if err == nil {
		b.description = description
	}
	return err
}

func (b *Brand) GetLogo() string {
	return b.logo
}

func (b *Brand) SetLogo(logo string) *validation.FieldError {
	// Quais validaçõe spor aqui???
	// Verificar se é uma URL ou um caminho de arquivo... implementação futura...
	b.logo = logo
	return nil
}

func (b *Brand) GetStatus() bool {
	return bool(b.status)
}

func (b *Brand) SetStatus(status statusBrand) {
	b.status = status
}

func (b *Brand) GetManufacturerUUID() string {
	return b.manufacturerUUID
}

func (b *Brand) SetManufacturerUUID(uuid string) *validation.FieldError {
	var isValid = common.IsValidUUUID(uuid)

	if isValid {
		b.manufacturerUUID = uuid
	}
	return &validation.FieldError{
	FieldName: "manufacturer_uuid",
	CodeErrors: []string{string(validation.ErrUnknown)},
	}
}
