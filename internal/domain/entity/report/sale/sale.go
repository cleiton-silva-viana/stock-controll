package sale

/*
	A partir de R$2000 em compras, o cliente pessoa físicas deve ser obrigado a fornecer CPF
*/

import (
	"time"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/valueobject/uuid"
)

type PaymentMethod string

const (
	CreditCard PaymentMethod = "credit_card"
	DebitCard  PaymentMethod = "debit_card"
	Boleto     PaymentMethod = "boleto"
	Cash       PaymentMethod = "cash"
	Other      PaymentMethod = "other"
)

type salesStatus string

const (
	Sold         salesStatus = "sold"
	InProcessing salesStatus = "in_processing"
	Canceled     salesStatus = "canceled"
)

type Discount struct {
	Amount        float32
	PromotionCode string
}

type Product struct {
	productUUID string
	price       float32
	quantity    int
}

type ISale interface {
	SaleUUID() string
	TimesTamp() time.Time
	ClientUUID() string
	SalerUUID() string
	Products() []Product
	PaymentMethod() string
	SaleStatus() string
	Discount() string
	Amout() float32
	UpdateStatus() error
	CalculateTotal() error
}

type Sale struct {
	uuid          uuid.UUID
	sellerUUID    uuid.UUID
	clientUUID    uuid.UUID
	discount      Discount
	amount        float32
	products      []Product
	paymentMethod PaymentMethod
	status        salesStatus
	timestamp     time.Time
}

type Config struct {
	SellerUUID    string
	ClientUUID    string
	Discount      Discount
	Products      []Product
	PaymentMethod PaymentMethod
	Status        salesStatus
}

/*
TODO:  Adicionar validações para payment, discount, status e products

Ao invés de paymentMethod, devemos receber um objeto do tipo payment
E com base no status de pagamento, gerar o relatório
Ou seja, se payment metho retornar um processing, já podemos gerar um relatório
Tal abordagem evita que criemos um relatório para um evento nunca ocorrido
ou um suposto evento futuro
*/
func New(config Config) (*Sale, error) {
	var saleError = entity.Error("sale")

	// validar produtos
	// validar método de pagamento
	// validar discount
	// validar status
	clientVO, uuidClientErr := uuid.Parse("client_uuid", config.ClientUUID)
	sellerVO, uuidSellerErr := uuid.Parse("seller_uuid", config.SellerUUID)

	saleError.
		AddValidationError(uuidClientErr).
		AddValidationError(uuidSellerErr)

	if saleError.HasError() {
		return nil, saleError
	}

	// Pensar em como gerar descontos

	return &Sale{
		uuid:          *uuid.New(),
		clientUUID:    *clientVO,
		sellerUUID:    *sellerVO,
		products:      config.Products,
		paymentMethod: config.PaymentMethod,
		discount:      config.Discount,
		status:        config.Status,
		timestamp:     time.Now(),
	}, nil
}

func (s *Sale) SellerUUID() string {
	return s.sellerUUID.String()
}

func (s *Sale) ClientUUID() string {
	return s.clientUUID.String()
}

func (s *Sale) Discount() Discount {
	return s.discount
}

// TODO: implementar
func (s *Sale) CalculateDiscount(product Product) {}

func (s *Sale) Amout() float32 {
	return s.amount
}

// deve ficar aqui????
/* func (s *Sale) calculateAmount() {
	var amount int

	for _, product := range s.products {
		(product.price * float32(product.quantity)) - s.discount
	}

} */

func (s *Sale) CalculateTotal() error {
	return nil
}

func (s *Sale) Products() []Product {
	return s.products
}

func (s *Sale) PaymentMethod() string {
	return string(s.paymentMethod)
}

func (s *Sale) Status() string {
	return string(s.status)
}

// TODO: usar erro genérico para tratamento de erros de atribuíção de status
func (s *Sale) SetStatus(newStatus salesStatus) error {
	if s.status == newStatus {
		return &field.FieldError{
			FieldName: "status",
			CodeError: "",
		}
	}
	s.status = newStatus
	return nil
}

func (s *Sale) Timestamp() time.Time {
	return time.Now()
}

const ErrNoProductsAssociated = "ERR_NO_PRODUCTS_ASSOCIATED"

// TODO: melhorar checagem...
func (s *Sale) validateProducts(products []Product) error {
	if len(products) == 0 {
		return &field.FieldError{
			FieldName: "products",
			CodeError: ErrNoProductsAssociated,
		}
	}
	return nil
}
