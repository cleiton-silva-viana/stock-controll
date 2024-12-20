package discard

/*
Tornar a estrutura discard imutável
pode exigir algumas mudanças na forma como você lida com a lógica de negócios,
mas os benefícios em termos de segurança, previsibilidade e facilidade de teste
 geralmente compensam o esforço.
 Avalie as necessidades do seu projeto
 e considere se a imutabilidade é uma abordagem que
 se alinha com seus objetivos de design.
*/

import (
	"time"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type DiscardStatus string

const (
	done DiscardStatus = "done"
	toDo DiscardStatus = "to_do"
)

type Discard struct {
	uuid.UUID
	checkerUUID uuid.UUID
	productUUID uuid.UUID
	batchUUID   uuid.UUID
	quantity    int
	createdAt   time.Time
	heldIn      time.Time
	status      DiscardStatus
}

const (
	minQuantity = 1
	maxQuantity = 1000
)

func New(checkerUUID, productUUID, batchUUID string, quantity int) (*Discard, error) {
	var discardError = entity.Error("discard")

	quantityErr := validateQuantity(quantity)
	uuidChecker, uuidCheckerErr := uuid.Parse("checker_uuid", checkerUUID)
	uuidProduct, uuidProcutErr := uuid.Parse("product_uuid", productUUID)
	uuidBatch, uuidBatchErr := uuid.Parse("batch_uuid", batchUUID)

	discardError.
		AddValidationError(quantityErr).
		AddValidationError(uuidCheckerErr).
		AddValidationError(uuidProcutErr).
		AddValidationError(uuidBatchErr)

	if discardError.HasError() {
		return nil, discardError
	}

	return &Discard{
		UUID:        *uuid.New(),
		checkerUUID: *uuidChecker,
		productUUID: *uuidProduct,
		batchUUID:   *uuidBatch,
		quantity:    quantity,
		status:      toDo,
		createdAt:   time.Now(),
	}, nil
}

func (d *Discard) Quantity() int {
	return d.quantity
}

func (d *Discard) CreatedAt() time.Time {
	return d.createdAt
}

func (d *Discard) HeldIn() time.Time {
	return d.heldIn
}

func (d *Discard) Status() DiscardStatus {
	return d.status
}

const (
	ErrStatusChangeNotAllowed        = "ERR_STATUS_CHANGE_NOT_ALLOWED"
	ErrInvalidDiscardOperationStatus = "ERR_INVALID_DISCARD_OPERATION_STATUS"
)

func (d *Discard) UpdateStatus(newStatus DiscardStatus) error {
	if d.status == done {
		return &field.FieldError{
			FieldName: "status",
			CodeError: ErrStatusChangeNotAllowed,
		}
	}

	if newStatus != done {
		return &field.FieldError{
			FieldName: "status",
			CodeError: ErrInvalidDiscardOperationStatus,
		}
	}

	d.status = newStatus
	d.heldIn = time.Now()
	return nil
}

func validateQuantity(quantity int) error {
	return validate.New[int](
		"quantity", quantity,
		validate.IsInRange(minQuantity, maxQuantity),
	)
}
