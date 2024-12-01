package product

// TODO: Adicionar mutex para condição concorrência
// TODO: Adicionar tags aos produtos

import (
	"image"

	"stock-controll/internal/domain/entity/tag"
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type IProduct interface {
	UUID() string
	Name() string
	Price() float64
	Quantity() int
	Description() string
	Barcode() string
	Tags() []tag.Tag
	BrandUUID() string
	ManufacturerUUID() string
	CategoryUUID() string
}

type Product struct {
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
	tags             []tag.Tag
}

type Config struct {
	Name             string
	Description      string
	Barcode          string
	BrandUUID        string
	ManufacturerUUID string
	CategoryUUID     string
}

func New(config Config) (IProduct, error) {
	productInstance := Product{
		uuid: uuid.New(),
		tags: make([]tag.Tag, 0),
	}

	productError := validationerrors.New("product").
		AddValidationError(productInstance.SetName(config.Name)).
		AddValidationError(productInstance.SetDescription(config.Description)).
		AddValidationError(productInstance.SetBarcode(config.Barcode)).
		AddValidationError(productInstance.SetBrandUUID(config.BrandUUID)).
		AddValidationError(productInstance.SetManufacturerUUID(config.ManufacturerUUID)).
		AddValidationError(productInstance.SetCategoryUUID(config.CategoryUUID))

	if productError.HasError() {
		return nil, productError
	}
	return &productInstance, nil
}

func (p *Product) UUID() string {
	return p.uuid
}

func (p *Product) Name() string {
	return p.name
}

const (
	minNameLength = 10
	maxNameLength = 50
)

func (p *Product) SetName(name string) error {
	err := validate.New("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLength, maxNameLength),
	)
	if err == nil {
		p.name = name
	}
	return err
}

func (p *Product) Cost() float64 {
	return p.cost
}

const (
	minCost                           = 0.1
	maxCost                           = 900.00
	ErrProductPurchaseCostOutOfBounds = "ERR_PRODUCT_PURCHASE_COST_OUT_OF_BOUNDS"
)

/*
Possíveis Vulnerabilidades:
  - Race Conditions em ambiente concorrente
  - Precisão de ponto flutuante em cálculos financeiros
  - Overflow/Underflow em operações matemáticas
*/
func (p *Product) SetCost(cost float64) error {
	err := validate.New("cost", cost,
		validate.IsInRangeFloat64(minCost, maxCost),
	)
	if err != nil {
		return &validate.FieldError{
			FieldName: "cost",
			CodeError: ErrProductPurchaseCostOutOfBounds,
		}
	}

	p.cost = cost
	return nil
}

func (p *Product) Price() float64 {
	return p.price
}

const (
	minPrice                       = 0.10
	maxPrice                       = 1000.00
	ErrProductSalePriceOutOfBounds = "ERR_PRODUCT_SALE_PRICE_OUT_OF_BOUNDS"
)

func (p *Product) SetPrice(price float64) error {
	err := validate.New("price", price,
		validate.IsInRangeFloat64(minPrice, maxPrice),
	)
	if err != nil {
		return &validate.FieldError{
			FieldName: "price",
			CodeError: ErrProductSalePriceOutOfBounds,
		}
	}
	p.price = price
	return nil
}

func (p *Product) Description() string {
	return p.description
}

const (
	minDescriptionLength = 10
	maxDescriptionLength = 500
)

func (p *Product) SetDescription(description string) error {
	err := validate.New("description", description,
		validate.IsBlank(),
		validate.IsLengthInRange(minDescriptionLength, maxDescriptionLength),
	)
	if err == nil {
		p.description = description
	}
	return err
}

func (p *Product) Tags() []tag.Tag {
	tags := make([]tag.Tag, len(p.tags))
	copy(tags, p.tags)
	return tags
}

const MaxTagsForProduct = 7
const ErrProductTagLimitExceeded = "ERR_PRODUCT_TAG_LIMIT_EXCEEDED"

// Verificar se a tag é correspondente ao produto
// Tags que fazem sentido para o produto...
func (p *Product) SetTag(newTag tag.Tag) error {
	if len(p.tags) >= MaxTagsForProduct {
		return &validate.FieldError{
			FieldName: "tag",
			CodeError: ErrProductTagLimitExceeded,
		}
	}

	if len(p.tags) == MaxTagsForProduct {
		newTags := make([]tag.Tag, len(p.tags), MaxTagsForProduct)
		copy(newTags, p.tags)
		p.tags = newTags
	}

	p.tags = append(p.tags, newTag)
	return nil
}

func (p *Product) Quantity() int {
	return p.quantity
}

const ErrQuantityExceedsStock = "ERR_QUANTITY_EXCEEDS_STOCK"

func (p *Product) SetQuantity(quantity int) error {
	if p.quantity-quantity < 0 {
		return &validate.FieldError{
			FieldName: "quantity",
			CodeError: ErrQuantityExceedsStock,
		}
	}
	p.quantity = quantity
	return nil
}

func (p *Product) Barcode() string {
	return p.barcode
}

const (
	minBarcodeLength = 6
	maxBarcodeLength = 24
)

func (p *Product) SetBarcode(barcode string) error {
	var err = validate.New("barcode", barcode,
		validate.IsBlank(),
		validate.IsLengthInRange(minBarcodeLength, maxBarcodeLength),
		validate.CheckLetters(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err == nil {
		p.barcode = barcode
	}
	return err
}

func (p *Product) BrandUUID() string {
	return p.brandUUID
}

func (p *Product) SetBrandUUID(BrandUUID string) error {
	err := uuid.IsValid("brand_uuid", BrandUUID)
	if err == nil {
		p.brandUUID = BrandUUID
	}
	return err
}

func (p *Product) ManufacturerUUID() string {
	return p.manufacturerUUID
}

func (p *Product) SetManufacturerUUID(ManufacturerUUID string) error {
	err := uuid.IsValid("manufacturer_uuid", ManufacturerUUID)
	if err == nil {
		p.manufacturerUUID = ManufacturerUUID
	}
	return err
}

func (p *Product) CategoryUUID() string {
	return p.categoryUUID
}

func (p *Product) SetCategoryUUID(CategoryUUID string) error {
	err := uuid.IsValid("category_uuid", CategoryUUID)
	if err == nil {
		p.categoryUUID = CategoryUUID
	}
	return err
}
