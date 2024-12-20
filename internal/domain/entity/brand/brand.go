package brand

import (
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type Status bool

const (
	Active   Status = true
	Inactive Status = false
)

type Brand struct {
	uuid        uuid.UUID
	name        string
	description string
	// logo             string // url
	manufacturerUUID uuid.UUID
	status           Status
}

const prefix = "BRA"

func New(name, description, manufacturerUUID string) (*Brand, error) {
	var brandError = entity.Error("Brand")

	manufacturerVO, uuidManufacturerErr := uuid.Parse("manufacturer_uuid", manufacturerUUID)
	nameErr := validateName(name)
	descriptionErr := validateDescription(description)
	// logoErr := validateLogo(logo)
	id, idErr := uuid.New(prefix)

	brandError.
		AddValidationError(uuidManufacturerErr).
		AddValidationError(nameErr).
		//	AddValidationError(logoErr)
		AddValidationError(descriptionErr).
		AddValidationError(idErr)

	if brandError.HasError() {
		return nil, brandError
	}

	return &Brand{
		uuid:             *id,
		manufacturerUUID: *manufacturerVO,
		name:             name,
		description:      description,
		status:           Active,
	}, nil
}

func (b *Brand) UUID() string {
	return b.uuid.String()
}

func (b *Brand) ManufacturerUUID() string {
	return b.manufacturerUUID.String()
}

func (b *Brand) Name() string {
	return b.name
}

func (b *Brand) Description() string {
	return b.description
}

func (b *Brand) Status() bool {
	return bool(b.status)
}

const (
	minNameLengthForBrand = 3
	maxNameLengthForBrand = 30
)

func validateName(name string) error {
	return validate.New[string](
		"name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLengthForBrand, maxNameLengthForBrand),
	)
}

const (
	minDescriptionLength = 10
	maxDescriptionLength = 250
)

func validateDescription(description string) error {
	return validate.New[string](
		"description", description,
		validate.IsBlank(),
		validate.IsLengthInRange(minDescriptionLength, maxDescriptionLength),
	)
}
