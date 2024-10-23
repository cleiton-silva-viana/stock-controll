package expiration

import (
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
	"time"
)

type expiration struct {
	uuid         string
	reporterUUID string
	productUUID  string
	quantity     int
	expireDate   time.Time
}

func NewShortExpirationDate(reporterUUID, productUUID string, quantity int, expireAs time.Time) (*expiration, validationError.IValidationError) {
	var expirationError = validationError.NewValidationError("expiration")
	var expirationInstance = expiration{}

	expirationError.
		AddValidationError(expirationInstance.ValidateUUID(reporterUUID)).
		AddValidationError(expirationInstance.ValidateUUID(productUUID)).
		AddValidationError(
			validation.Validate[int](
				"quantity",
				quantity,
				validation.IsInRange(
					1, 1000, validation.ErrUnknown))).
		AddValidationError(
			validation.Validate[time.Time](
				"expire_date",
				expireAs,
				validation.IsAfterThan(time.Now(), validation.ErrUnknown),
				validation.IsBeforeThan(time.Now().AddDate(0, 0, 45), validation.ErrUnknown),
			))
	if expirationError.HasError() {
		return nil, expirationError
	}

	return &expiration{
		uuid:         common.GenerateUUID(),
		reporterUUID: reporterUUID,
		productUUID:  productUUID,
		quantity:     quantity,
		expireDate:   expireAs,
	}, nil
}

func (e *expiration) GetUUID() string {
	return e.uuid
}

func (e *expiration) GetReporterUUID() string {
	return e.reporterUUID
}

func (e *expiration) GetProductUUID() string {
	return e.productUUID
}

func (e *expiration) GetQuantity() int {
	return e.quantity
}

func (e *expiration) GetExpirationDate() time.Time {
	return e.expireDate
}

func (e *expiration) IsExpired() bool {
	return time.Now().Unix() > e.expireDate.Unix()
}
