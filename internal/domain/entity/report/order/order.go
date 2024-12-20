package order

import (
	"time"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type Product struct {
	uuid.UUID
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
	uuid.UUID
	buyerUUID          uuid.UUID
	supplierUUID       uuid.UUID
	requestMadeOn      time.Time
	expectedDeliveryOn time.Time
	products           []Product
	status             OrderStatus
	// qrCode ---> implementar
	// profPayment ---> implementar
}

func New(buyerUUID, supplierUUID string, expectedDelivery time.Time, products []Product) (*Order, error) {
	var orderError = entity.Error("order")

	buyerVO, uuidBuyerErr := uuid.Parse("buyer_uuid", buyerUUID)
	supplierVO, uuidSupplierErr := uuid.Parse("supplier_uuid", supplierUUID)

	orderError.
		AddValidationError(uuidBuyerErr).
		AddValidationError(uuidSupplierErr)

		// AddValidationError(orderInstance.setProducts(products)).
		// AddValidationError(orderInstance.UpdateExpectedDelivery(expectedDelivery))

	if orderError.HasError() {
		return nil, orderError
	}

	return &Order{
		UUID:               *uuid.New(),
		buyerUUID:          *buyerVO,
		supplierUUID:       *supplierVO,
		products:           products,
		expectedDeliveryOn: expectedDelivery,
		requestMadeOn:      time.Now(),
		status:             pending,
	}, nil
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
		return &field.FieldError{
			FieldName: "products",
			CodeError: ErrMinimOneProductRequiredForRestock,
		}
	}
	return nil
}

func (o *Order) ExpectedDelivery() time.Time {
	return o.expectedDeliveryOn
}

// TODO: Este erro abaixo deve ser genérico, ou seja, qualquer operação análoga, deve usar este error code
const ErrDeliveryDateCannotBePast = "ERR_DELIVERY_DATE_CANNOT_BE_PAST"

func (o *Order) UpdateExpectedDelivery(expectedDelivery time.Time) error {
	var err = validate.New[time.Time](
		"expected_delivery", expectedDelivery,
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
		return &field.FieldError{
			FieldName: "status",
			CodeError: ErrStatusAlreadyAssigned,
		}
	}

	if o.status == completed || o.status == canceled {
		return &field.FieldError{
			FieldName: "status",
			CodeError: ErrReportCannotBeModified,
		}
	}

	o.status = newStatus
	return nil
}
