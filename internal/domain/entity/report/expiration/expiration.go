package expiration

import (
	"time"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type Expiration struct {
	quantity     int
	uuid         uuid.UUID
	reporterUUID uuid.UUID
	productUUID  uuid.UUID
	expireDate   time.Time // e se a data de expiração for igual ou posterior a data d evencimento?
}

const (
	ErrExpiredProductReported      = "ERR_EXPIRED_PRODUCT_REPORTED"
	ErrProductExpirationDateTooFar = "ERR_PRODUCT_EXPIRATION_DATE_TOO_FAR"
	ErrLowQuantityProductReported  = "ERR_LOW_QUANTITY_PRODUCT_REPORTED"
)

func New(reporterUUID, productUUID string, quantity int, expireAs time.Time) (*Expiration, error) {
	var expirationError = entity.Error("expiration")

	uuidReporter, uuidReporterErr := uuid.Parse("reporte_uuid", reporterUUID)
	uuidProduct, uuidProductErr := uuid.Parse("product_uuid", productUUID)

	quantityErr := validateQuantity(quantity)
	expireErr := validateExpire(expireAs)

	expirationError.
		AddValidationError(uuidReporterErr).
		AddValidationError(uuidProductErr).
		AddValidationError(quantityErr).
		AddValidationError(expireErr)

	if expirationError.HasError() {
		return nil, expirationError
	}

	return &Expiration{
		uuid:         *uuid.New(),
		reporterUUID: *uuidReporter,
		productUUID:  *uuidProduct,
		quantity:     quantity,
		expireDate:   expireAs,
	}, nil
}

func (e *Expiration) Quantity() int {
	return e.quantity
}

func (e *Expiration) ExpirationDate() time.Time {
	return e.expireDate
}

func (e *Expiration) IsExpired() bool {
	return time.Now().Unix() > e.expireDate.Unix()
}

const (
	minQuantity = 1
	maxQuantity = 100
)

func validateQuantity(quantity int) error {
	return validate.New(
		"quantity", quantity,
		validate.IsInRange(minQuantity, maxQuantity),
	)
}

func validateExpire(expire time.Time) error {
	return validate.New[time.Time](
		"expire_date", expire,
		validate.IsBeforeThan(time.Now(), ErrExpiredProductReported),
	)
}
