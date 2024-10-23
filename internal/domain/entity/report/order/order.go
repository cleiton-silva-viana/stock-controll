package order

import (
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
	"time"
)

type Product struct {
	UUID     string
	quantity int
	price    int
}

type orderStatus string

const (
	pending    orderStatus = "pending"
	completed  orderStatus = "completed"
	inProgress orderStatus = "in_progress"
	canceled   orderStatus = "canceled"
)

type order struct {
	uuid string
	buyerUUID          string
	supplierUUID       string
	requestMadeOn      time.Time
	expectedDeliveryOn time.Time
	products           []Product
	status             orderStatus
	// qrCode ---> implementar
	// profPayment ---> implementar
}

func NewOrder(buyerUUID, supplierUUID string, expectedDelivery time.Time, products []Product) (*order, validationError.IValidationError) {
	var orderError = validationError.NewValidationError("order")
	var orderInstance = &order{
		uuid: common.GenerateUUID(),
		buyerUUID:          buyerUUID,
		supplierUUID:       supplierUUID,
		products:           products,
		expectedDeliveryOn: expectedDelivery,
		requestMadeOn:      time.Now(),
		status:             pending,
	}

	orderError.
		AddValidationError(orderInstance.ValidateUUID(buyerUUID)).
		AddValidationError(orderInstance.ValidateUUID(supplierUUID)).
		AddValidationError(orderInstance.setProducts(products)).
		AddValidationError(orderInstance.UpdateExpectedDelivery(expectedDelivery))

	if orderError.HasError() {
		return nil, orderError
	}

	return orderInstance, nil
}

func (o *order) GetUUID() string {
	return o.uuid
}

// TODO: tornar verificações mais robustas
/*
	Validação de Produtos: O método setProducts atualmente apenas verifica se
	a lista de produtos está vazia. Você pode querer adicionar validações adicionais, 
	como verificar se cada produto tem um preço e uma quantidade válidos.
*/
func (o *order) setProducts(products []Product) *validation.FieldError {
	if len(products) == 0 {
		return &validation.FieldError{
			FieldName:  "products",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	return nil
}

func (o *order) GetBuyerUUID() string {
	return o.buyerUUID
}

func (o *order) GetSupplierUUID() string {
	return o.supplierUUID
}

func (o *order) GetExpectedDelivery() time.Time {
	return o.expectedDeliveryOn
}

func (o *order) UpdateExpectedDelivery(expectedDelivery time.Time) *validation.FieldError {
	var err = validation.Validate[time.Time](
		"expected_delivery",
		expectedDelivery,
		validation.IsBeforeThan(o.requestMadeOn, validation.ErrUnknown),
	)
	if err == nil {
		o.expectedDeliveryOn = expectedDelivery
	}
	return err
}

func (o *order) GetRequestOn() time.Time {
	return o.requestMadeOn
}

func (o *order) GetStatus() orderStatus {
	return o.status
}

func (o *order) SetStatus(newStatus orderStatus) *validation.FieldError {
	if o.status == newStatus {
		return &validation.FieldError{
			FieldName:  "status",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	if o.status == completed || o.status == canceled {
		return &validation.FieldError{
			FieldName:  "status",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	o.status = newStatus
	return nil
}
