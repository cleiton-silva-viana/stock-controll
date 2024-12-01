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

	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type DiscardStatus string

const (
	done DiscardStatus = "done"
	toDo DiscardStatus = "to_do"
)

type Discard struct {
	uuid        string
	checkerUUID string
	productUUID string
	batchUUID   string
	quantity    int
	createdAt   time.Time
	heldIn      time.Time
	status      DiscardStatus
}

const (
	MinQuantity = 1
	MaxQuantity = 1000
)

func NewDiscard(checkerUUID, productUUID, batchUUID string, quantity int) (*Discard, error) {
	var discardError = validationerrors.New("discard").
		AddValidationError(uuid.IsValid("checker_uuid", checkerUUID)).
		AddValidationError(uuid.IsValid("product_uuid", productUUID)).
		AddValidationError(uuid.IsValid("batch_uuid", batchUUID)).
		AddValidationError(
			validate.New[int](
				"quantity",
				quantity,
				validate.IsInRange(MinQuantity, MaxQuantity),
			),
		)

	if discardError.HasError() {
		return nil, discardError
	}

	return &Discard{
		uuid:        uuid.New(),
		checkerUUID: checkerUUID,
		productUUID: productUUID,
		batchUUID:   batchUUID,
		quantity:    quantity,
		status:      toDo,
		createdAt:   time.Now(),
	}, nil
}

func (d *Discard) UUID() string {
	return d.uuid
}

func (d *Discard) CheckerUUID() string {
	return d.checkerUUID
}

func (d *Discard) ProductUUID() string {
	return d.productUUID
}

func (d *Discard) BatchUUID() string {
	return d.batchUUID
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
	ErrStatusChangeNotAllowed = "ERR_STATUS_CHANGE_NOT_ALLOWED"
	ErrInvalidDiscardOperationStatus = "ERR_INVALID_DISCARD_OPERATION_STATUS"
)

func (d *Discard) SetStatus(newStatus DiscardStatus) error {
	if d.status == done {
		return &validate.FieldError{
			FieldName:  "status",
			CodeError: ErrStatusChangeNotAllowed,
		}
	}

	if newStatus == done {
		d.status = newStatus
		d.heldIn = time.Now()
		return nil
	}

	return &validate.FieldError{
		FieldName:  "status",
		CodeError: ErrInvalidDiscardOperationStatus, 
	}
}
