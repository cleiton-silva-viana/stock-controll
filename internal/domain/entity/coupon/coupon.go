package coupon

import (
	"time"

	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/entity/role"
	"stock-controll/internal/domain/services/discount"
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type ICoupon interface {
	UUID() string
	Name() string
	MinPurchaseAmount() float64
	ExpirationDate() time.Time
	UsageLimit() CouponUseLimit
	Exclusivity() CouponExclusivity
	IsUsable() bool
	Redeem(purchasedProducts []product.IProduct) (discount.DiscountSummary, error)
	IsProductValid(product product.IProduct) bool
}

type CouponExclusivity string

const (
	AllCustomers        CouponExclusivity = "all_customer"
	RegisteredCustomers CouponExclusivity = "registered_customer"
	Employees           CouponExclusivity = "employee"
)

type CouponUseLimit int

const (
	Unlimited CouponUseLimit = -1
)

type Coupon struct {
	uuid              string
	name              string
	minPurchaseAmount float64
	discountScope     discount.IProductSpecificDiscountStrategy
	startDate         time.Time
	expirationDate    time.Time
	usageLimit        CouponUseLimit
	currentUsage      int
	exclusivity       CouponExclusivity
}

type Config struct {
	Name              string
	MinPurchaseAmount float64
	DiscountScope     discount.IProductSpecificDiscountStrategy
	StartDate         time.Time
	ExpirationDate    time.Time
	UsageLimit        int
	Exclusivity       string
}

func New(config Config) (*Coupon, error) {

	CouponInstance := &Coupon{
		uuid:              uuid.New(),
		name:              config.Name,
		minPurchaseAmount: config.MinPurchaseAmount,
		discountScope:     config.DiscountScope,
		expirationDate:    config.ExpirationDate,
		usageLimit:        CouponUseLimit(config.UsageLimit), // TODO: validar
		currentUsage:      0,
		exclusivity:       CouponExclusivity(config.Exclusivity), // TODO: validar
	}

	err := CouponInstance.validate()
	if err != nil {
		return nil, err
	}

	return CouponInstance, nil
}

func (c *Coupon) UUID() string {
	return c.uuid
}

func (c *Coupon) Name() string {
	return c.name
}

func (c *Coupon) MinPurchaseAmount() float64 {
	return c.minPurchaseAmount
}

func (c *Coupon) StartDate() time.Time {
	return c.startDate
}

func (c *Coupon) ExpirationDate() time.Time {
	return c.expirationDate
}

func (c *Coupon) UsageLimit() CouponUseLimit {
	return c.usageLimit
}

func (c *Coupon) Exclusivity() CouponExclusivity {
	return c.exclusivity
}

func (c *Coupon) IsUsable() bool {
	return c.usageLimit < 0 || c.currentUsage < int(c.usageLimit)
}

const (
	ErrCouponUsageLimitExceeded  = "ERR_COUPON_USAGE_LIMIT_EXCEEDED"
	ErrUserNotAllowedToUseCoupon = "ERR_USER_NOT_ALLOWED_TO_USE_COUPON"
)

func (c *Coupon) Redeem(userRole role.Role, item product.IProduct, quantity int) (discount.DiscountSummary, error) {
	if !c.IsUsable() {
		return discount.DiscountSummary{}, &validate.FieldError{
			FieldName: "coupon",
			CodeError: ErrCouponUsageLimitExceeded,
		}
	}

	if isUserAllowed(c.exclusivity, userRole) {
		return discount.DiscountSummary{}, &validate.FieldError{
			FieldName: "user_uuid",
			CodeError: ErrUserNotAllowedToUseCoupon,
		}
	}

	// TODO: implementar erro
	if item.Price()*float64(quantity) < c.minPurchaseAmount {
		// compra abaixo do mínimo para aplicabilidade do desconto
		return discount.DiscountSummary{}, &validate.FieldError{
			CodeError: "",
		}
	}

	var discountedAmount, err = c.discountScope.Apply(item, quantity)
	if err == nil {
		c.currentUsage++
	}
	return discountedAmount, err
}

func (c *Coupon) IsProductValid(item product.IProduct) bool {
	return c.discountScope.IsProductValid(item)
}

func (c *Coupon) validate() error {
	couponError := validationerrors.New("Coupon")
	couponError.
		AddValidationError(validateName(c.name)).
		AddValidationError(validateMinPurchaseAllowed(c.minPurchaseAmount)).
		AddValidationError(validateStartDate(c.startDate, c.expirationDate)).
		AddValidationError(validateExpirationDate(c.startDate, c.expirationDate)).
		AddValidationError(validateDiscoutInstance(c.discountScope))

	if couponError.HasError() {
		return couponError
	}
	return nil
}

const (
	minCouponNameLength = 4
	maxCouponNameLength = 20
)

func validateName(name string) error {
	return validate.New("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minCouponNameLength, maxCouponNameLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
}

const minAmount = 0

// TODO: implementar erro
func validateMinPurchaseAllowed(minPurchaseAmount float64) error {
	if minPurchaseAmount < minAmount {
		return &validate.FieldError{
			FieldName: "discount",
			CodeError: "",
		}
	}
	return nil
}

const ErrCouponInvalidDateRange = "ERR_COUPON_INVALID_DATE_RANGE"

func validateStartDate(startDate, expirationDate time.Time) error {
	return validate.New[time.Time](
		"start_date", startDate,
		validate.IsAfterThan(expirationDate, ErrCouponInvalidDateRange),
	)
}

const (
	minCouponValidityPeriod = 7 * 24 * time.Hour
	maxCouponValidityPeriod = 365 * 24 * time.Hour
)

// TODO: Implementar erro
/*
	A data de expiração deve ser no mínimo 7 dias a partir da data de inicio de validade do cupom
	Já a data de expiração deve ser de no máximo 1 ano
*/
func validateExpirationDate(startDate, expirationDate time.Time) error {
	return validate.New[time.Time](
		"expiration_date", expirationDate,
		validate.IsBeforeThan(startDate, ""),
		validate.IsBeforeThan(time.Now().Add(minCouponValidityPeriod), ""),
		validate.IsAfterThan(time.Now().Add(maxCouponValidityPeriod), ""),
	)
}

func validateDiscoutInstance(discountInstance discount.IProductSpecificDiscountStrategy) error {
	return validate.New[any](
		"discout", discountInstance,
		validate.IsNil(discountInstance),
	)
}

// Refatorar
// Code smells
func isUserAllowed(exclusivity CouponExclusivity, userRole role.Role) bool {
	switch exclusivity {
	case AllCustomers:
		return true
	// case RegisteredCustomers:
	// 	return userType == string(RegisteredCustomers)
	// case Employees:
	// 	return userType == string(Employees)
	default:
		return false
	}
}

func calculateTotalAmount(purchasedProducts []product.IProduct) float64 {
	var total float64
	for _, p := range purchasedProducts {
		if p != nil {
			total += p.Price()
		}
	}
	return total
}
