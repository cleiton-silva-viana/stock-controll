package receipt

import (
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
	"time"
)

type receiptStatus string

const (
	inProgress receiptStatus = "in_progress"
	accepted   receiptStatus = "accepted"
	rejected   receiptStatus = "rejected"
)

type receipt struct {
	uuid             string
	orderUUID        string
	lecturerUUID     string
	manufacturerUUID string
	observations     string
	startedIn        time.Time
	finishedOn       time.Time
	status           receiptStatus
}

func NewReceipt(orderUUID, lecturerUUID, manufacturerUUID string) (*receipt, validationError.IValidationError) {
	var receptError = validationError.NewValidationError("recept")
	var receptInstance = receipt{
		uuid: common.GenerateUUID(),
		status:    inProgress,
		startedIn: time.Now(),
	}

	receptError.
		AddValidationError(receptInstance.ValidateUUID(orderUUID)).
		AddValidationError(receptInstance.ValidateUUID(lecturerUUID)).
		AddValidationError(receptInstance.ValidateUUID(manufacturerUUID))

	if receptError.HasError() {
		return nil, receptError
	}

	return &receptInstance, nil
}

func (r *receipt) GetUUID() string {
	return r.uuid
}

func (r *receipt) GetOrderUUID() string {
	return r.orderUUID
}

func (r *receipt) GetLecturerUUID() string {
	return r.lecturerUUID
}

func (r *receipt) GetManufacturerUUID() string {
	return r.manufacturerUUID
}

const (
	minLengthForObservation = 10
	maxLengthForObservation = 200
)

func (r *receipt) SetObservations(note string) *validation.FieldError {
	var err = validation.Validate("observation", note,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minLengthForObservation, maxLengthForObservation, validation.ErrUnknown),
	)
	if err == nil {
		r.observations = note
	}
	return err
}

func (r *receipt) GetStatus() receiptStatus {
	return r.status
}

func (r *receipt) AcceptOrder() *validation.FieldError {
	if r.status != inProgress {
		return &validation.FieldError{
			FieldName:  "status",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	r.finishedOn = time.Now()
	r.status = accepted
	return nil
}

func (r *receipt) RejectOrder(cause string) *validation.FieldError {
	if r.status != inProgress {
		return &validation.FieldError{
			FieldName:  "status",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	var err = r.SetObservations(cause)
	if err == nil {
		r.finishedOn = time.Now()
		r.status = rejected
	}
	return err
}
