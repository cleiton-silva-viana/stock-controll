package coupon

import (
	"stock-controll/internal/domain/valueobject/discount/scope"
	"stock-controll/internal/domain/valueobject/summary"
	"time"

	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/entity/role"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type ICoupon interface {
	UUID() string
	Name() string
	MinPurchaseAmount() float64
	ExpirationDate() time.Time
	UsageLimit() UseLimit
	Exclusivity() Exclusivity
	IsUsable() bool
	Redeem(purchasedProducts []product.IProduct) (summary.Summary, error)
	IsProductValid(product product.IProduct) bool
}

type Exclusivity string

const (
	AllCustomers        Exclusivity = "all_customer"
	RegisteredCustomers Exclusivity = "registered_customer"
	Employees           Exclusivity = "employee"
)

type UseLimit int

const (
	Unlimited UseLimit = -1
)

type Coupon struct {
	uuid              uuid.UUID
	name              string
	minPurchaseAmount float64
	discountScope     scope.IProductSpecificDiscountStrategy
	startDate         time.Time
	expirationDate    time.Time
	usageLimit        UseLimit
	currentUsage      int
	exclusivity       Exclusivity
}

type Config struct {
	Name              string
	MinPurchaseAmount float64
	DiscountScope     scope.IProductSpecificDiscountStrategy
	StartDate         time.Time
	ExpirationDate    time.Time
	UsageLimit        int
	Exclusivity       string
}

func New(config Config) (*Coupon, error) {

	couponError := entity.Error("coupon")

	nameErr := validateName(config.Name)
	minPurchasedErr := validateMinPurchaseAllowed(config.MinPurchaseAmount)
	startDateErr := validateStartDate(config.StartDate, config.ExpirationDate)
	expirationDateErr := validateExpirationDate(config.StartDate, config.ExpirationDate)
	discountScopeErr := validateDiscountInstance(config.DiscountScope)

	couponError.
		AddValidationError(nameErr).
		AddValidationError(minPurchasedErr).
		AddValidationError(startDateErr).
		AddValidationError(expirationDateErr).
		AddValidationError(discountScopeErr)

	return &Coupon{
		uuid:              *uuid.New(),
		name:              config.Name,
		minPurchaseAmount: config.MinPurchaseAmount,
		discountScope:     config.DiscountScope,
		expirationDate:    config.ExpirationDate,
		usageLimit:        UseLimit(config.UsageLimit), // TODO: validar
		currentUsage:      0,
		exclusivity:       Exclusivity(config.Exclusivity), // TODO: validar
	}, nil
}

func (c *Coupon) UUID() string { return c.uuid.String() }

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

func (c *Coupon) UsageLimit() UseLimit {
	return c.usageLimit
}

func (c *Coupon) Exclusivity() Exclusivity {
	return c.exclusivity
}

func (c *Coupon) IsUsable() bool {
	return c.usageLimit < 0 || c.currentUsage < int(c.usageLimit)
}

const (
	ErrCouponUsageLimitExceeded  = "ERR_COUPON_USAGE_LIMIT_EXCEEDED"
	ErrUserNotAllowedToUseCoupon = "ERR_USER_NOT_ALLOWED_TO_USE_COUPON"
	ErrPurchasedAmountTooLow     = "ERR_PURCHASE_AMOUNT_TOO_LOW"
)

// Quando eu resgato um cupom, o cupom incide sobre um produto, todos os produtos uma categoria ou etc?
// Como calcular???
// E se produto for nil?
// TODO: ajeitar os summary
func (c *Coupon) Redeem(userRole role.Role, item product.IProduct, quantity int) (summary.Summary, error) {
	if !c.IsUsable() {
		return summary.Summary{}, &field.FieldError{
			FieldName: "coupon",
			CodeError: ErrCouponUsageLimitExceeded,
		}
	}

	if isUserAllowed(c.exclusivity, userRole) {
		return summary.Summary{}, &field.FieldError{
			FieldName: "user_uuid",
			CodeError: ErrUserNotAllowedToUseCoupon,
		}
	}

	if item.Price()*float64(quantity) < c.minPurchaseAmount {
		return summary.Summary{}, &field.FieldError{
			FieldName: "coupon",
			CodeError: "ErrPurchasedAmountTooLow",
		}
	}

	var discountedAmount, err = c.discountScope.Apply(item, quantity)
	if err == nil {
		c.currentUsage++
	}
	return *discountedAmount, err
}

func (c *Coupon) IsProductValid(item product.IProduct) bool {
	return c.discountScope.IsProductValid(item)
}

const (
	minNameLength = 4
	maxNameLength = 20
)

func validateName(name string) error {
	return validate.New(
		"name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLength, maxNameLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
}

const (
	minAmount                  = 0
	ErrMinimumPurchaseRequired = "ERR_MINIMUM_PURCHASE_REQUIRED"
)

func validateMinPurchaseAllowed(minPurchaseAmount float64) error {
	if minPurchaseAmount < minAmount {
		return &field.FieldError{
			FieldName: "discount",
			CodeError: ErrMinimumPurchaseRequired,
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
	minValidityPeriod                  = 7 * 24 * time.Hour
	maxValidityPeriod                  = 365 * 24 * time.Hour
	ErrCouponExpirationBeforeStartDate = "ERR_COUPON_EXPIRATION_BEFORE_START_DATE"
	ErrCouponExpirationExceedsMaxDate  = "ERR_COUPON_EXPIRATION_EXCEEDS_MAX_DATE"
)

// Observar
func validateExpirationDate(startDate, expirationDate time.Time) error {
	return validate.New[time.Time](
		"expiration_date", expirationDate,
		validate.IsBeforeThan(startDate, ErrCouponExpirationBeforeStartDate),
		validate.IsAfterThan(time.Now().Add(maxValidityPeriod), ErrCouponExpirationExceedsMaxDate),
	)
}

func validateDiscountInstance(discountInstance scope.IProductSpecificDiscountStrategy) error {
	return validate.New[any](
		"discount", discountInstance,
		validate.IsNil(discountInstance),
	)
}

// Refatorar
// Code smells
func isUserAllowed(exclusivity Exclusivity, userRole role.Role) bool {
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
