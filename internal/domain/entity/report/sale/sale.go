package report

import (
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
	"time"
)

type paymentMethod string

const (
	CreditCard paymentMethod = "credit_card"
	DebitCard  paymentMethod = "debit_card"
	Boleto     paymentMethod = "boleto"
	Cash       paymentMethod = "cash"
	Other      paymentMethod = "other"
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
	GetSaleUUID() string
	GetTimesTamp() time.Time
	GetClientUUID() string
	GetSalerUUID() string
	GetProducts() []Product
	GetPaymentMethod() string
	GetSaleStatus() string
	GetDiscount() string
	GetAmout() float32
	UpdateStatus() error
	CalculateTotal() error
}

type sale struct {
	sellerUUID    string
	clientUUID    string
	discount      Discount
	amount        float32
	products      []Product
	paymentMethod paymentMethod
	status        salesStatus
	timestamp     time.Time
	uuid          string
}

func NewSale(clientUUID, sellerUUID string, products []Product, payment paymentMethod, discountApplyed Discount, status salesStatus) (*sale, validationError.IValidationError) {
	var saleError = validationError.NewValidationError("sale")
	var saleInstance = sale{
		uuid:          common.GenerateUUID(),
		clientUUID:    clientUUID,
		sellerUUID:    sellerUUID,
		products:      products,
		paymentMethod: payment,
		discount:      discountApplyed,
		status:        status,
		timestamp:     time.Now(),
	}

	saleError.
		AddValidationError(common.IsValidUUUID(clientUUID)).
		AddValidationError(common.IsValidUUUID(sellerUUID)).
		AddValidationError(saleInstance.validateProducts(products))

	if saleError.HasError() {
		return nil, saleError
	}

	// Pensar em como gerar descontos

	return &saleInstance, nil
}

func (s *sale) GetSellerUUID() string {
	return s.sellerUUID
}

func (s *sale) GetClientUUID() string {
	return s.clientUUID
}

func (s *sale) GetDiscount() Discount {
	return s.discount
}

// TODO: implementar
func (s *sale) CalculateDiscount(product Product) {

}

func (s *sale) GetAmout() float32 {
	return s.amount
}

// deve ficar aqui????
/* func (s *sale) calculateAmount() {
	var amount int

	for _, product := range s.products {
		(product.price * float32(product.quantity)) - s.discount
	}

} */

func (s *sale) CalculateTotal() error { return nil }

func (s *sale) GetProducts() []Product {
	return s.products
}

func (s *sale) GetPaymentMethod() string {
	return string(s.paymentMethod)
}

func (s *sale) GetStatus() string {
	return string(s.status)
}

func (s *sale) SetStatus(newStatus salesStatus) *validation.FieldError {
	if s.status == newStatus {
		return &validation.FieldError{
			FieldName:  "status",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	s.status = newStatus
	return nil
}

func (s *sale) GetTimestamp() time.Time {
	return time.Now()
}

func (s *sale) validateProducts(products []Product) *validation.FieldError {
	if len(products) == 0 {
		return &validation.FieldError{
			FieldName:  "products",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	return nil
}
