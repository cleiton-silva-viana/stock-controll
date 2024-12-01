package receipt

import (
	"time"

	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type ReceiptStatus string

const (
	inProgress ReceiptStatus = "in_progress"
	accepted   ReceiptStatus = "accepted"
	rejected   ReceiptStatus = "rejected"
)

type Receipt struct {
	uuid             string
	orderUUID        string
	lecturerUUID     string
	manufacturerUUID string
	observations     string
	startedIn        time.Time
	finishedOn       time.Time
	status           ReceiptStatus
}

func NewReceipt(orderUUID, lecturerUUID, manufacturerUUID string) (*Receipt, error) {
	var receptError = validationerrors.New("recept")
	var receptInstance = Receipt{
		uuid:      uuid.New(),
		status:    inProgress,
		startedIn: time.Now(),
	}

	receptError.
		AddValidationError(uuid.IsValid("order_uuid", orderUUID)).
		AddValidationError(uuid.IsValid("lecturer_uuid", lecturerUUID)).
		AddValidationError(uuid.IsValid("manufacturer_uuid", manufacturerUUID))

	if receptError.HasError() {
		return nil, receptError
	}

	return &receptInstance, nil
}

func (r *Receipt) UUID() string {
	return r.uuid
}

func (r *Receipt) OrderUUID() string {
	return r.orderUUID
}

func (r *Receipt) LecturerUUID() string {
	return r.lecturerUUID
}

func (r *Receipt) ManufacturerUUID() string {
	return r.manufacturerUUID
}

const (
	minLengthForObservation = 10
	maxLengthForObservation = 200
)

func (r *Receipt) SetObservations(note string) error {
	var err = validate.New("observation", note,
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
		return &validate.FieldError{
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
		return &validate.FieldError{
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
