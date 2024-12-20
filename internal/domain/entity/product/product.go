package product

// TODO: Adicionar mutex para condição concorrência
// TODO: Adicionar tags aos produtos
// TODO: implementar no próprio produto, as métricas para definir:
// o valor a ser considerado como baixa quantidade
// o limite de tempo até considerar o produto com data de vencimento próxima

import (
	"image"

	"stock-controll/internal/domain/entity/tag"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
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
	uuid             uuid.UUID
	name             string
	cost             float64
	price            float64
	quantity         int
	description      string
	barcode          string
	brandUUID        uuid.UUID
	manufacturerUUID uuid.UUID
	categoryUUID     uuid.UUID
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

func New(config Config) (*Product, error) {
	productError := entity.Error("product")

	nameErr := validateName(config.Name)
	descriptionErr := validateDescription(config.Description)
	barcodeErr := validateBarcode(config.Barcode)
	brandVO, uuidBrandErr := uuid.Parse("brand_uuid", config.BrandUUID)
	manufacturerVO, uuidManufacturerErr := uuid.Parse("manufacturer_uuid", config.ManufacturerUUID)
	categoryVO, uuidCategoryErr := uuid.Parse("category_uuid", config.CategoryUUID)

	productError.
		AddValidationError(nameErr).
		AddValidationError(descriptionErr).
		AddValidationError(barcodeErr).
		AddValidationError(uuidBrandErr).
		AddValidationError(uuidManufacturerErr).
		AddValidationError(uuidCategoryErr)

	if productError.HasError() {
		return nil, productError
	}
	return &Product{
		uuid:             *uuid.New(),
		name:             config.Name,
		description:      config.Description,
		barcode:          config.Barcode,
		brandUUID:        *brandVO,
		manufacturerUUID: *manufacturerVO,
		categoryUUID:     *categoryVO,
	}, nil
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Cost() float64 {
	return p.cost
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Tags() []tag.Tag {
	tags := make([]tag.Tag, len(p.tags))
	copy(tags, p.tags)
	return tags
}

func (p *Product) Quantity() int {
	return p.quantity
}

func (p *Product) Barcode() string {
	return p.barcode
}

func (p *Product) BrandUUID() string {
	return p.brandUUID.String()
}

func (p *Product) ManufacturerUUID() string {
	return p.manufacturerUUID.String()
}

func (p *Product) CategoryUUID() string {
	return p.categoryUUID.String()
}

const MaxTagsForProduct = 7
const ErrProductTagLimitExceeded = "ERR_PRODUCT_TAG_LIMIT_EXCEEDED"

// Verificar se a tag é correspondente ao produto
// Tags que fazem sentido para o produto...
func (p *Product) SetTag(newTag tag.Tag) error {
	if len(p.tags) >= MaxTagsForProduct {
		return &field.FieldError{
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

const ErrQuantityExceedsStock = "ERR_QUANTITY_EXCEEDS_STOCK"

// TODO: refletir!!!
// Este método deve ser chamado de fora ou devemos te rum método que respeito o tell dont ask?
func (p *Product) SetQuantity(quantity int) error {
	if p.quantity-quantity < 0 {
		return &field.FieldError{
			FieldName: "quantity",
			CodeError: ErrQuantityExceedsStock,
		}
	}
	p.quantity = quantity
	return nil
}

const (
	minPrice                       = 0.10
	maxPrice                       = 1000.00
	ErrProductSalePriceOutOfBounds = "ERR_PRODUCT_SALE_PRICE_OUT_OF_BOUNDS"
)

func (p *Product) SetPrice(price float64) error {
	err := validate.New(
		"price", price,
		validate.IsInRangeFloat64(minPrice, maxPrice),
	)
	if err != nil {
		return &field.FieldError{
			FieldName: "price",
			CodeError: ErrProductSalePriceOutOfBounds,
		}
	}
	return nil
}

/*
Possíveis Vulnerabilidades:
  - Race Conditions em ambiente concorrente
  - Precisão de ponto flutuante em cálculos financeiros
  - Overflow/Underflow em operações matemáticas
*/
func (p *Product) SetCost(cost float64) error {
	err := validate.New(
		"cost", cost,
		validate.IsInRangeFloat64(minCost, maxCost),
	)
	if err != nil {
		return &field.FieldError{
			FieldName: "cost",
			CodeError: ErrProductPurchaseCostOutOfBounds,
		}
	}
	return nil
}

const (
	minNameLength = 10
	maxNameLength = 50
)

func validateName(name string) error {
	return validate.New(
		"name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLength, maxNameLength),
	)
}

const (
	minCost                           = 0.1
	maxCost                           = 900.00
	ErrProductPurchaseCostOutOfBounds = "ERR_PRODUCT_PURCHASE_COST_OUT_OF_BOUNDS"
)

const (
	minDescriptionLength = 10
	maxDescriptionLength = 500
)

func validateDescription(description string) error {
	return validate.New(
		"description", description,
		validate.IsBlank(),
		validate.IsLengthInRange(minDescriptionLength, maxDescriptionLength),
	)
}

const (
	minBarcodeLength = 6
	maxBarcodeLength = 24
)

func validateBarcode(barcode string) error {
	return validate.New(
		"barcode", barcode,
		validate.IsBlank(),
		validate.IsLengthInRange(minBarcodeLength, maxBarcodeLength),
		validate.CheckLetters(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
}

// validações de UUID deve ser feitas no método principal
// remover estes métodos
func validateBrandUUID(BrandUUID string) error {
	return uuid.IsValid("brand_uuid", BrandUUID)
}

func validateManufacturerUUID(ManufacturerUUID string) error {
	return uuid.IsValid("manufacturer_uuid", ManufacturerUUID)
}

func validateCategoryUUID(CategoryUUID string) error {
	return uuid.IsValid("category_uuid", CategoryUUID)
}

/*

// Produto representa um produto genérico em um sistema de e-commerce.
type Produto struct {
    ID         	 	string    	`json:"id"`
    Nome        	string    	`json:"nome"`
    Descricao   	string    	`json:"descricao"`
    Preco       	float64   	`json:"preco"`
	tags			TagProduct

	Quantidade  	int       	`json:"quantidade"`
    DataValidade 	time.Time 	`json:"data_validade,omitempty"`
	Lote 			string

	Peso        	float64   	`json:"peso,omitempty"` // Peso em gramas
    Volume      	float64   	`json:"volume,omitempty"` // Volume em mililitros
    Dimensoes   	Dimensoes 	`json:"dimensoes,omitempty"`

	Imagens     	[]string  	`json:"imagens,omitempty"` // URLs das imagens do produto
    Avaliacoes  	[]Avaliacao `json:"avaliacoes,omitempty"`
}

// Dimensoes representa as dimensões de um produto.
type Dimensoes struct {
    Comprimento float64 `json:"comprimento,omitempty"`
    Largura     float64 `json:"largura,omitempty"`
    Altura      float64 `json:"altura,omitempty"`
}

// Avaliacao representa uma avaliação de um produto.
type Avaliacao struct {
    Usuario    string `json:"usuario"`
    Nota       int    `json:"nota"` // Nota de 1 a 5
    Comentario string `json:"comentario,omitempty"`
}

// Um produto deve possuir
//	- Uma tag de categoria (obrigatório)
// 	- Uma tag de subtacegoria (obrigatório)
// 	- Ao menos 2 tags de atributo (obrigatório)
// 	- Uma tag da marca do produto (obrigatório)
type TagProduct struct {
	category  TagType //TagType = category
	subcategory TagType
	attribute map[string]TagType
}
*/
