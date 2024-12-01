package brand

import (
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type StatusBrand bool

const (
	Active   = true
	Inactive = false
)

type Brand struct {
	uuid             string
	name             string
	description      string
	logo             string
	manufacturerUUID string
	status           StatusBrand
	// salesHistory []Sales
}

func New(name, description, logo, manufacturerUUID string) (*Brand, *validationerrors.ValidationError) {
	var brandInstance = Brand{
		uuid: uuid.New(),
	}

	var err = validationerrors.
		New("Brand").
		AddValidationError(brandInstance.SetName(name)).
		AddValidationError(brandInstance.SetDescription(description)).
		AddValidationError(brandInstance.SetLogo(logo)).
		AddValidationError(brandInstance.SetManufacturerUUID(manufacturerUUID))

	if err.HasError() {
		return nil, err
	}

	return &brandInstance, nil
}

func (b *Brand) UUID() string {
	return b.uuid
}

func (b *Brand) Name() string {
	return b.name
}

const (
	minNameLengthForBrand = 3
	maxNameLengthForBrand = 30
)

func (b *Brand) SetName(name string) error {
	var err = validate.New[string]("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLengthForBrand, maxNameLengthForBrand),
	)

	if err == nil {
		b.name = name
	}
	return err
}

func (b *Brand) Description() string {
	return b.description
}

const (
	minDescriptionLength = 10
	maxDescriptionLength = 250
)

func (b *Brand) SetDescription(description string) error {
	var err = validate.New[string]("description", description,
		validate.IsBlank(),
		validate.IsLengthInRange(minDescriptionLength, maxDescriptionLength),
	)

	if err == nil {
		b.description = description
	}
	return err
}

func (b *Brand) Logo() string {
	return b.logo
}

// TODO: Logo deve ser um svg
// TODO: implementar teste de consistência
func (b *Brand) SetLogo(logo string) error {
	b.logo = logo
	return nil
}

func (b *Brand) Status() bool {
	return bool(b.status)
}

func (b *Brand) SetStatus(status StatusBrand) {
	b.status = status
}

func (b *Brand) ManufacturerUUID() string {
	return b.manufacturerUUID
}

func (b *Brand) SetManufacturerUUID(manufacturerUUID string) error {
	err := uuid.IsValid("manufacturer_uuid", manufacturerUUID)
	if err == nil {
		b.manufacturerUUID = manufacturerUUID
	}
	return err
}
