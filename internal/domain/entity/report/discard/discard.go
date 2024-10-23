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
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
	"time"
)

type discardStatus string

const (
	done discardStatus = "done"
	toDo discardStatus = "to_do"
)

type discard struct {
	uuid string
	checkerUUID string
	productUUID string
	batchUUID   string
	quantity    int
	createdAt   time.Time
	heldIn      time.Time
	status      discardStatus
}

const (
	MinQuantityForDiscard  = 1
	MaxQuantityForDiscard  = 1000
)

func NewDiscard(checkerUUID, productUUID, batchUUID string, quantity int) (*discard, validationError.IValidationError) {
	var discardError = validationError.NewValidationError("discard")
	var discardInstance = discard{
		uuid: common.GenerateUUID(),
		checkerUUID: checkerUUID,
		productUUID: productUUID,
		batchUUID:   batchUUID,
		quantity:    quantity,
		status:      toDo,
		createdAt:   time.Now(),
	}

	discardError.
		AddValidationError(discardInstance.ValidateUUID(checkerUUID)).
		AddValidationError(discardInstance.ValidateUUID(productUUID)).
		AddValidationError(discardInstance.ValidateUUID(batchUUID)).
		AddValidationError(
			validation.Validate(
				"quantity",
				quantity,
				validation.IsInRange(MinQuantityForDiscard, MaxQuantityForDiscard, validation.ErrUnknown)))

	if discardError.HasError() {
		return nil, discardError
	}

	return &discardInstance, nil
}

func (d *discard) GetUUID() string {
	return d.uuid
}

func (d *discard) GetCheckerUUID() string {
	return d.checkerUUID
}

func (d *discard) GetProductUUID() string {
	return d.productUUID
}

func (d *discard) GetBatchUUID() string {
	return d.batchUUID
}

func (d *discard) GetQuantity() int {
	return d.quantity
}

func (d *discard) GetCreatedAt() time.Time {
	return d.createdAt
}

func (d *discard) GetHeldIn() time.Time {
	return d.heldIn
}

func (d *discard) GetStatus() discardStatus {
	return d.status
}

func (d *discard) SetStatus(newStatus discardStatus) *validation.FieldError {
	if d.status == done {
		return &validation.FieldError{
			FieldName:  "status",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	if newStatus == done {
		d.status = newStatus
		d.heldIn = time.Now()
		return nil
	}
	
	return &validation.FieldError{
		FieldName:  "status",
		CodeErrors: []string{string(validation.ErrUnknown)},
	}
}
