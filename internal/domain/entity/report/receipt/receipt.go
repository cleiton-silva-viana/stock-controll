package receipt

import (
	"time"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type ReceiptStatus string

const (
	inProgress ReceiptStatus = "in_progress"
	accepted   ReceiptStatus = "accepted"
	rejected   ReceiptStatus = "rejected"
)

type Receipt struct {
	uuid.UUID
	orderUUID        uuid.UUID
	lecturerUUID     uuid.UUID
	manufacturerUUID uuid.UUID
	observations     string
	startedIn        time.Time
	finishedOn       time.Time
	status           ReceiptStatus
}

func New(orderUUID, lecturerUUID, manufacturerUUID string) (*Receipt, error) {
	var receptError = entity.Error("recept")

	oderVO, uuidOrderErr := uuid.Parse("order_uuid", orderUUID)
	lecturerVO, uuidLecturerErr := uuid.Parse("lecturer_uuid", lecturerUUID)
	manufacturerVO, uuidManufacturerErr := uuid.Parse("manufacturer_uuid", manufacturerUUID)

	receptError.
		AddValidationError(uuidOrderErr).
		AddValidationError(uuidLecturerErr).
		AddValidationError(uuidManufacturerErr)

	if receptError.HasError() {
		return nil, receptError
	}

	return &Receipt{
		UUID:             *uuid.New(),
		orderUUID:        *oderVO,
		lecturerUUID:     *lecturerVO,
		manufacturerUUID: *manufacturerVO,
	}, nil
}

const (
	minLengthForObservation = 10
	maxLengthForObservation = 200
)

func (r *Receipt) SetObservations(note string) error {
	var err = validate.New(
		"observation", note,
		validate.IsBlank(),
		validate.IsLengthInRange(minLengthForObservation, maxLengthForObservation),
	)
	if err == nil {
		r.observations = note
	}
	return err
}

func (r *Receipt) Status() ReceiptStatus {
	return r.status
}

// TODO: adicionar erro
func (r *Receipt) AcceptOrder() error {
	if r.status != inProgress {
		return &field.FieldError{
			FieldName: "status",
			CodeError: "",
		}
	}
	r.finishedOn = time.Now()
	r.status = accepted
	return nil
}

// TODO: adicionar error
func (r *Receipt) RejectOrder(cause string) error {
	if r.status != inProgress {
		return &field.FieldError{
			FieldName: "status",
			CodeError: "",
		}
	}
	var err = r.SetObservations(cause)
	if err == nil {
		r.finishedOn = time.Now()
		r.status = rejected
	}
	return err
}
