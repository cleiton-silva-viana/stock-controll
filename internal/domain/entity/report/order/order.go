package order

import (
	"time"

	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type Product struct {
	UUID     string
	quantity int
	price    int
}

type OrderStatus string

const (
	pending    OrderStatus = "pending"
	completed  OrderStatus = "completed"
	inProgress OrderStatus = "in_progress"
	canceled   OrderStatus = "canceled"
)

type Order struct {
	uuid               string
	buyerUUID          string
	supplierUUID       string
	requestMadeOn      time.Time
	expectedDeliveryOn time.Time
	products           []Product
	status             OrderStatus
	// qrCode ---> implementar
	// profPayment ---> implementar
}

func NewOrder(buyerUUID, supplierUUID string, expectedDelivery time.Time, products []Product) (*Order, error) {
	var orderError = validationerrors.New("order")
	var orderInstance = &Order{
		uuid:               uuid.New(),
		buyerUUID:          buyerUUID,
		supplierUUID:       supplierUUID,
		products:           products,
		expectedDeliveryOn: expectedDelivery,
		requestMadeOn:      time.Now(),
		status:             pending,
	}

	orderError.
		AddValidationError(uuid.IsValid("buyer_uuid", buyerUUID)).
		AddValidationError(uuid.IsValid("supplier_uuid", supplierUUID)).
		AddValidationError(orderInstance.setProducts(products)).
		AddValidationError(orderInstance.UpdateExpectedDelivery(expectedDelivery))

	if orderError.HasError() {
		return nil, orderError
	}

	return orderInstance, nil
}

func (o *Order) UUID() string {
	return o.uuid
}

// TODO: tornar verificações mais robustas
/*
	Validação de Produtos: O método setProducts atualmente apenas verifica se
	a lista de produtos está vazia. Você pode querer adicionar validações adicionais,
	como verificar se cada produto tem um preço e uma quantidade válidos.
*/

const ErrMinimOneProductRequiredForRestock = "ERR_MINIMUM_ONE_PRODUCT_REQUIRED_FOR_RESTOCK"

func (o *Order) setProducts(products []Product) error {
	if len(products) == 0 {
		return &validate.FieldError{
			FieldName: "products",
			CodeError: ErrMinimOneProductRequiredForRestock,
		}
	}
	return nil
}

func (o *Order) BuyerUUID() string {
	return o.buyerUUID
}

func (o *Order) SupplierUUID() string {
	return o.supplierUUID
}

func (o *Order) ExpectedDelivery() time.Time {
	return o.expectedDeliveryOn
}

// TODO: Este erro abaixo deve ser genérico, ou seja, qualquer operação análoga, deve usar este error code
const ErrDeliveryDateCannotBePast = "ERR_DELIVERY_DATE_CANNOT_BE_PAST"

func (o *Order) UpdateExpectedDelivery(expectedDelivery time.Time) error {
	var err = validate.New[time.Time](
		"expected_delivery",
		expectedDelivery,
		validate.IsBeforeThan(o.requestMadeOn, ErrDeliveryDateCannotBePast),
	)
	if err == nil {
		o.expectedDeliveryOn = expectedDelivery
	}
	return err
}

func (o *Order) RequestOn() time.Time {
	return o.requestMadeOn
}

func (o *Order) Status() OrderStatus {
	return o.status
}

// TODO: os erros devem ser genéricos, aplicados a todos os tipos de relatórios
const (
	ErrStatusAlreadyAssigned  = "ERR_STATUS_ALREADY_ASSIGNED"
	ErrReportCannotBeModified = "ERR_REPORT_CANNOT_BE_MODIFIED"
)

func (o *Order) SetStatus(newStatus OrderStatus) error {
	if o.status == newStatus {
		return &validate.FieldError{
			FieldName: "status",
			CodeError: ErrStatusAlreadyAssigned,
		}
	}

	if o.status == completed || o.status == canceled {
		return &validate.FieldError{
			FieldName: "status",
			CodeError: ErrReportCannotBeModified,
		}
	}

	o.status = newStatus
	return nil
}
