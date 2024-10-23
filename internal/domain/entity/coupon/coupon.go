package coupon

import (
	"time"

	"stock-controll/internal/domain/entity/common"
	"stock-controll/internal/domain/entity/discount"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/entity/role"
	"stock-controll/internal/domain/validation"
)

type ICoupon interface {
	GetUUID() string
	GetName() string
	GetMinPurchaseAmount() float64
	GetExpirationDate() time.Time
	GetUsageLimit() CouponUseLimit
	GetExclusivity() CouponExclusivity
	IsUsable() bool
	Redeem(purchasedProducts []product.IProduct) (discount.DiscountSummary, *validation.FieldError)
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

type coupon struct {
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

func NewCoupon(
	name string,
	minPurchaseAmount float64,
	discountScope discount.IProductSpecificDiscountStrategy,
	startDate time.Time, // validar
	expirationDate time.Time, // validar
	usageLimit CouponUseLimit,
	exclusivity CouponExclusivity,
) (*coupon, validationError.IValidationError) {

	CouponInstance := &coupon{
		uuid:              common.GenerateUUID(),
		name:              name,
		minPurchaseAmount: minPurchaseAmount,
		discountScope:     discountScope,
		expirationDate:    expirationDate,
		usageLimit:        usageLimit,
		currentUsage:      0,
		exclusivity:       exclusivity,
	}

	err := CouponInstance.validate()
	if err != nil {
		return nil, err
	}

	return CouponInstance, nil
}

func (c *coupon) GetUUID() string {
	return c.uuid
}

func (c *coupon) GetName() string {
	return c.name
}

func (c *coupon) GetMinPurchaseAmount() float64 {
	return c.minPurchaseAmount
}

func (c *coupon) GetstartDate() time.Time {
	return c.startDate
}

func (c *coupon) GetExpirationDate() time.Time {
	return c.expirationDate
}

func (c *coupon) GetUsageLimit() CouponUseLimit {
	return c.usageLimit
}

func (c *coupon) GetExclusivity() CouponExclusivity {
	return c.exclusivity
}

func (c *coupon) IsUsable() bool {
	return c.usageLimit < 0 || c.currentUsage < int(c.usageLimit)
}

/*
	A função Redeem pode ser melhorada para lidar com erros de forma mais clara.
*/
func (c *coupon) Redeem(userRole role.Role, item product.IProduct, quantity int) (discount.DiscountSummary, *validation.FieldError) {
	if !c.IsUsable() {
		// excedeu o número de uso do cupom
		return discount.DiscountSummary{}, &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	if isUserAllowed(c.exclusivity, userRole) {
		// Usuário não permitido a usar este cupom
		return discount.DiscountSummary{}, &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	if item.GetPrice()*float64(quantity) < c.minPurchaseAmount {
		// compra abaixo do mínimo para aplicabilidade do desconto
		return discount.DiscountSummary{}, &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	var discountedAmount, err = c.discountScope.Apply(item, quantity)
	if err == nil {
		c.currentUsage++
	}
	return discountedAmount, err
}

func (c *coupon) IsProductValid(item product.IProduct) bool {
	return c.discountScope.IsProductValid(item)
}


// TODO: validar start date
// Observe que, operação de validação de datas, já está sendo algo corriqueir, ou seja, está havendo repetição de código
// Reoslver!!!
func (c *coupon) validate() validationError.IValidationError {
	couponError := validationError.NewValidationError("Coupon")
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

func validateName(name string) *validation.FieldError {
	return validation.Validate("name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minCouponNameLength, maxCouponNameLength, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
}

const minAmount = 0

func validateMinPurchaseAllowed(minPurchaseAmount float64) *validation.FieldError {
	if minPurchaseAmount < minAmount {
		return &validation.FieldError{
			FieldName:  "discount",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	return nil
}


/*
	A repetição de lógica de validação de datas pode ser refatorada para evitar duplicação.
*/
func validateStartDate(startDate, expirationDate time.Time) *validation.FieldError {
	return validation.Validate[time.Time](
		"start_date", startDate,
		validation.IsAfterThan(expirationDate, validation.ErrUnknown),
	) 
}

const (
    minCouponValidityPeriod = 7 * 24 * time.Hour  // 7 dias
    maxCouponValidityPeriod = 365 * 24 * time.Hour // 1 ano
)

// Melhorar 
/*
	A data de expiração deve ser no mínimo 7 dias a partir da data de inicio de validade do cupom
	Já a data de expiração deve ser de no máximo 1 ano 
*/
func validateExpirationDate(startDate, expirationDate time.Time)  *validation.FieldError {
	return validation.Validate[time.Time](
		"expiration_date", expirationDate,
		validation.IsBeforeThan(startDate, validation.ErrUnknown),
		validation.IsBeforeThan(time.Now().Add(minCouponValidityPeriod), validation.ErrUnknown),
		validation.IsAfterThan(time.Now().Add(maxCouponValidityPeriod), validation.ErrUnknown),
	)
}

func validateDiscoutInstance(discountInstance discount.IProductSpecificDiscountStrategy) *validation.FieldError {
	return validation.Validate[any](
		"discout", discountInstance,
		validation.IsNil(discountInstance, validation.ErrUnknown),
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
			total += p.GetPrice()
		}
	}
	return total
}
