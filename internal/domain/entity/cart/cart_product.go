package cart

import (
	"stock-controll/internal/domain/entity/coupon"
	"stock-controll/internal/domain/entity/discount"
	validationerrors "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/entity/promotion"
	"stock-controll/internal/domain/validation"
)

/* {
    "product": { ... },
    "quantityPurchased": 10,
    "appliedDiscounts": {
        "retailDiscount": {
            "minimumQuantityRequired": 6,
            "discountStrategy": "fixedValue",
            "discountAmount": 0.30,
            "discountValue": 3.00,
            "isApplied": true
        },
        "bulkDiscount": {
            "type": "buy_X_Get_Y",
            "buyQuantity": 2,
            "getQuantity": 3,
            "discountValue": 10.99,
            "description": "Pague 3, leve 5",
            "isActive": true
        },
    }
    "couponDiscounts": [
            {
                "couponCode": "DISCOUNT10",
                "discountStrategy": "fixedValue",
                "discountAmount": 0.30,
                "discountValue": 3.00,
                "isApplied": true
            },
            {
                "couponCode": "BIRTHDATE2024",
                "discountStrategy": "percentage",
                "discountAmount": 0.1,
                "discountValue": 10.30,
                "isApplied": true
            }
    ],
    "totalAmount": {
        "subtotal": 109.00,
        "discountSummary": {
            "totalDiscountApplied": 16.30,
            "finalTotal": 92.70
        }
    }
} */

type TotalAmount struct {
	subtotal float64
	discount.DiscountSummary
}

type ProductCart struct {
	purchasedQuantity int
	appliedPromotions map[string]promotion.IPromotion
	appliedCoupons    map[string]coupon.ICoupon
	product.IProduct
	TotalAmount
}

func NewProductCart(product product.IProduct) (*ProductCart, validationerrors.IValidationError) {
	var productCartErrors = validationerrors.NewValidationError("product_cart")

	productCartErrors.AddValidationError(
		validation.Validate[any]("cart_product", product,
			validation.IsNil(product, validation.ErrUnknown),
		),
	)

	if productCartErrors.HasError() {
		return nil, productCartErrors
	}

	return &ProductCart{
		IProduct:          product,
		purchasedQuantity: 1,
		appliedCoupons:    make(map[string]coupon.ICoupon, 0),
		appliedPromotions: make(map[string]promotion.IPromotion, 0),
		TotalAmount: TotalAmount{
			subtotal:        product.GetPrice(),
			DiscountSummary: discount.DiscountSummary{},
		},
	}, nil
}

func (pc *ProductCart) GetPurchasedQuantity() int {
	return pc.purchasedQuantity
}

// Sempre que alterar a quantidade
// recalcular os descontos e cupons
func (pc *ProductCart) SetPurchasedQuantity(newQuantity int) {
	// Se a nova quantidade for menor que 0, usamos 0
	if pc.purchasedQuantity-newQuantity < 0 {
		pc.purchasedQuantity = 0
	}
	pc.purchasedQuantity += newQuantity
	pc.Calculate()
}

func (pc *ProductCart) GetAppliedDiscount() []coupon.ICoupon {
	var coupons = make([]coupon.ICoupon, 0, len(pc.appliedCoupons))

	for _, coupon := range pc.appliedCoupons {
		coupons = append(coupons, coupon)
	}

	return coupons
}

// Refatorar !!!
func (pc *ProductCart) AddPromotion(promo promotion.IPromotion) *validation.FieldError {
	if promo == nil {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	_, exists := pc.appliedPromotions[promo.GetUUID()]
	if exists {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	// Não podemos adicionar promoções:
	//	 	finalizadas
	//		canceladas
	//		Inválidas para este produto
	status := promo.GetStatus()
	if status != promotion.Scheduled && status != promotion.InProgress {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	return nil
}

func (pc *ProductCart) AddCoupon(coupon coupon.ICoupon) *validation.FieldError {
	if coupon == nil {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	// verificar se o cupon já está no map
	_, exists := pc.appliedCoupons[coupon.GetUUID()]
	if exists {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	isValidCoupon := coupon.IsProductValid(pc)
	if !isValidCoupon {
		return nil
	}

	// adicionando o cupom ao map do produto que contém todos os cupons aplicados
	pc.appliedCoupons[coupon.GetUUID()] = coupon

	// Recalcular os descontos
	pc.Calculate()
	return nil
}

func (pc *ProductCart) RemoveCoupon(couponUUID string) {
	_, exists := pc.appliedCoupons[couponUUID]
	if exists {
		delete(pc.appliedCoupons, couponUUID)
	}
}

func (pc *ProductCart) GetSubtotal() float64 {
	return pc.subtotal
}

func (pc *ProductCart) GetTotalDiscountApplied() float64 {
	return pc.DiscountSummary.GetTotalDiscountApplied()
}

func (pc *ProductCart) GetFinalTotal() float64 {
	return pc.DiscountSummary.GetFinalTotal()
}

func (pc *ProductCart) Calculate() ProductCart {
	// verificar pela flag se já foi realizado o calculo
	pc.CalculateDiscount()
	pc.CalculateCoupons()

	// calcular impostos unitários do produto
	pc.CalculateTaxes()

	return *pc
}

func (pc *ProductCart) CalculateDiscount() {
	/* 	if len(pc.appliedCoupons) < 1 {
	   		return nil
	   	}

	   	var errorsToApplyDiscount = validationerrors.NewValidationError("cart_product")

	   	pc.TotalAmount.subtotal = pc.GetPrice() * float64(pc.purchasedQuantity)
	   	for _, discount := range pc.appliedDiscounts {
	   		discountSummary, err := dis.Apply(pc.IProduct, pc.purchasedQuantity)
	   		pc.DiscountSummary.Update(discountSummary)
	   		errorsToApplyDiscount.AddValidationError(err)
	   	}

	   	if errorsToApplyDiscount.HasError() {
	   		return errorsToApplyDiscount
	   	}
	   	return nil */
}

func (pc *ProductCart) CalculateCoupons() {}

func (pc *ProductCart) CalculateTaxes() {}
