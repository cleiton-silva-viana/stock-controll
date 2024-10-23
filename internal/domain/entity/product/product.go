package product

// TODO: Adicionar mutex para condição concorrência
// TODO: Adicionar tags aos produtos

import (
	"image"
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
)

type IProduct interface {
	GetUUID() string
	GetName() string
	GetPrice() float64
	GetQuantity() int
	GetDescription() string
	GetBarcode() string
	GetBrandUUID() string
	GetTags() []Tag
	GetManufacturerUUID() string
	GetCategoryUUID() string
}

type product struct {
	image.Image
	uuid             string
	name             string
	cost             float64
	price            float64
	quantity         int
	description      string
	barcode          string
	brandUUID        string
	manufacturerUUID string
	categoryUUID     string
	tags             []Tag
	status           bool
}

func NewProduct(name, description, barcode, brandUUID, manufacturerUUID, categoryUUID string) (IProduct, validationError.IValidationError) {
	var productError = validationError.NewValidationError("product")
	var productInstance = product{
		uuid:   common.GenerateUUID(),
		status: true,
		tags:   []Tag{},
	}

	productError.
		AddValidationError(productInstance.SetName(name)).
		AddValidationError(productInstance.SetDescription(description)).
		AddValidationError(productInstance.SetBarcode(barcode)).
		AddValidationError(productInstance.SetBrandUUID(brandUUID)).
		AddValidationError(productInstance.SetManufacturerUUID(manufacturerUUID)).
		AddValidationError(productInstance.SetCategoryUUID(categoryUUID))

	if productError.HasError() {
		return nil, productError
	}
	return &productInstance, nil
}

func (p *product) GetUUID() string {
	return p.uuid
}

func (p *product) GetName() string {
	return p.name
}

const (
	productNameMinLength = 10
	productNameMaxLength = 50
)

func (p *product) SetName(name string) *validation.FieldError {
	err := validation.Validate("name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(productNameMinLength, productNameMaxLength, validation.ErrUnknown),
	)
	if err == nil {
		p.name = name
	}
	return err
}

func (p *product) GetCost() float64 {
	return p.cost
}

const (
	minCost = 0.1
	maxCost = 1000.00
)

func (p *product) SetCost(cost float64) *validation.FieldError {
	if cost < minCost || cost > maxCost {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	// Verifique se o custo é um número válido
	if cost != cost { // Verifica se é NaN
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	p.cost = cost
	return nil
}

func (p *product) GetPrice() float64 {
	return p.price
}

const (
	minPrice = 0.10
	maxPrice = 1000.00
)
func (p *product) SetPrice(price float64) *validation.FieldError {
	if price < minPrice || price > maxPrice {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	p.price = price
	return nil
}

func (p *product) GetDescription() string {
	return p.description
}

const (
	productDescriptionMinLength = 10
	productDescriptionMaxLength = 500
)

func (p *product) SetDescription(description string) *validation.FieldError {
	err := validation.Validate("description", description,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(productDescriptionMinLength, productDescriptionMaxLength, validation.ErrUnknown),
	)
	if err == nil {
		p.description = description
	}
	return err
}

func (p *product) GetTags() []Tag {
	tags := make([]Tag, len(p.tags))
	copy(tags, p.tags)
	return tags
}

const MaxTagsForProduct = 7

// Verificar se a tag é correspondente ao produto
// Tags que fazem sentido para o produto...
func (p *product) SetTag(tag Tag) *validation.FieldError {
	if len(p.tags) >= MaxTagsForProduct {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	if len(p.tags) == MaxTagsForProduct {
		newTags := make([]Tag, len(p.tags), MaxTagsForProduct)
		copy(newTags, p.tags)
		p.tags = newTags
	}

	p.tags = append(p.tags, tag)
	return nil
}

func (p *product) GetQuantity() int {
	return p.quantity
}

// não posso subtrair uma quantidade maior que a em estoque
// Ou seja, não podemos ter estoque negativo
func (p *product) UpdateQuantity(quantity int) *validation.FieldError {
	// 10 - 11 = -1 < 0
	if p.quantity-quantity < 0 {
		// quantidade a ser decrementada é menor que a quantidade em estoque
		return &validation.FieldError{
			FieldName:  "quantity",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	p.quantity = quantity
	return nil
}

func (p *product) GetBarcode() string {
	return p.barcode
}

const (
	productBarcodeMinLength = 6
	productBarcodeMaxLength = 24
)

func (p *product) SetBarcode(barcode string) *validation.FieldError {
	var err = validation.Validate("barcode", barcode,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(productBarcodeMinLength, productBarcodeMaxLength, validation.ErrUnknown),
		validation.CheckLetters(validation.Disallow, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
	if err == nil {
		p.barcode = barcode
	}
	return err
}

func (p *product) GetBrandUUID() string {
	return p.brandUUID
}

func (p *product) SetBrandUUID(uuid string) *validation.FieldError {
	isValid := common.IsValidUUUID(uuid)
	if isValid {
		p.brandUUID = uuid
		return nil
	}
	return &validation.FieldError{
		FieldName:  "brand_uuid",
		CodeErrors: []string{string(validation.ErrUnknown)},
	}
}

func (p *product) GetManufacturerUUID() string {
	return p.manufacturerUUID
}

func (p *product) SetManufacturerUUID(uuid string) *validation.FieldError {
	isValid := common.IsValidUUUID(uuid)
	if isValid {
		p.manufacturerUUID = uuid
		return nil
	}
	return &validation.FieldError{
		FieldName:  "manufacturer_uuid",
		CodeErrors: []string{string(validation.ErrUnknown)},
	}
}

func (p *product) GetCategoryUUID() string {
	return p.categoryUUID
}

func (p *product) SetCategoryUUID(uuid string) *validation.FieldError {
	isValid := common.IsValidUUUID(uuid)
	if isValid {
		p.categoryUUID = uuid
		return nil
	}
	return &validation.FieldError{
		FieldName:  "category_uuid",
		CodeErrors: []string{string(validation.ErrUnknown)},
	}
}
