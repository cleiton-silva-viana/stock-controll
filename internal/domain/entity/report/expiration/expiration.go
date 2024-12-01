package expiration

import (
	"time"

	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type Expiration struct {
	uuid         string
	reporterUUID string
	productUUID  string
	quantity     int
	expireDate   time.Time
}

const (
	minQuantity                    = 1
	maxQuantity                    = 100
	ErrExpiredProductReported      = "ERR_EXPIRED_PRODUCT_REPORTED"
	ErrProductExpirationDateTooFar = "ERR_PRODUCT_EXPIRATION_DATE_TOO_FAR"
	ErrLowQuantityProductReported  = "ERR_LOW_QUANTITY_PRODUCT_REPORTED"
)

/*/ TODO: implementar no próprio produto, as métricas para definir:
		o valor a ser considerado como baixa quantidade
		o limite de tempo até considerar o produto com data de vencimento próxima 
*/
func New(reporterUUID, productUUID string, quantity int, expireAs time.Time) (*Expiration, error) {
	var expirationError = validationerrors.New("expiration").
		AddValidationError(uuid.IsValid("reporte_uuid", reporterUUID)).
		AddValidationError(uuid.IsValid("product_uuid", productUUID)).
		AddValidationError(validate.New(
			"quantity",
			quantity,
			validate.IsInRange(minQuantity, maxQuantity),
		)).
		AddValidationError(validate.New[time.Time](
			"expire_date",
			expireAs,
			validate.IsBeforeThan(time.Now(), ErrExpiredProductReported),
		))

	if expirationError.HasError() {
		return nil, expirationError
	}

	return &Expiration{
		uuid:         uuid.New(),
		reporterUUID: reporterUUID,
		productUUID:  productUUID,
		quantity:     quantity,
		expireDate:   expireAs,
	}, nil
}

func (e *Expiration) UUID() string {
	return e.uuid
}

func (e *Expiration) ReporterUUID() string {
	return e.reporterUUID
}

func (e *Expiration) ProductUUID() string {
	return e.productUUID
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
