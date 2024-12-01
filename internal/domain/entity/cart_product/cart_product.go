package cartproduct

import (
	"stock-controll/internal/domain/entity/coupon"
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/entity/promotion"
	"stock-controll/internal/domain/services/discount"
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/validate"
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
            "type": "buy_X__Y",
            "buyQuantity": 2,
            "Quantity": 3,
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

func New(product product.IProduct) (*ProductCart, error) {
	var productCartErrors = validationerrors.New("product_cart")

	productCartErrors.AddValidationError(
		validate.New[any]("cart_product", product,
			validate.IsNil(product),
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
			subtotal:        product.Price(),
			DiscountSummary: discount.DiscountSummary{},
		},
	}, nil
}

func (pc *ProductCart) PurchasedQuantity() int {
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

func (pc *ProductCart) AppliedDiscount() []coupon.ICoupon {
	var coupons = make([]coupon.ICoupon, 0, len(pc.appliedCoupons))

	for _, coupon := range pc.appliedCoupons {
		coupons = append(coupons, coupon)
	}

	return coupons
}

// TODO: Adicionar erro
func (pc *ProductCart) AddPromotion(promo promotion.IPromotion) error {
	if promo == nil {
		return &validate.FieldError{
			CodeError: "",
		}
	}

	_, exists := pc.appliedPromotions[promo.UUID()]
	if exists {
		return &validate.FieldError{
			CodeError: "",
		}
	}

	// Não podemos adicionar promoções:
	//	 	finalizadas
	//		canceladas
	//		Inválidas para este produto
	status := promo.Status()
	if status != promotion.Scheduled && status != promotion.InProgress {
		return &validate.FieldError{
			CodeError: "",
		}
	}

	return nil
}

func (pc *ProductCart) AddCoupon(coupon coupon.ICoupon) error {
	if coupon == nil {
		return &validate.FieldError{
			CodeError: "",
		}
	}

	// verificar se o cupon já está no map
	_, exists := pc.appliedCoupons[coupon.UUID()]
	if exists {
		return &validate.FieldError{
			CodeError: "",
		}
	}

	isValidCoupon := coupon.IsProductValid(pc)
	if !isValidCoupon {
		return nil
	}

	// adicionando o cupom ao map do produto que contém todos os cupons aplicados
	pc.appliedCoupons[coupon.UUID()] = coupon

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

func (pc *ProductCart) Subtotal() float64 {
	return pc.subtotal
}

func (pc *ProductCart) TotalDiscountApplied() float64 {
	return pc.DiscountSummary.TotalDiscountApplied()
}

func (pc *ProductCart) FinalTotal() float64 {
	return pc.DiscountSummary.FinalTotal()
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

	   	pc.TotalAmount.subtotal = pc.Price() * float64(pc.purchasedQuantity)
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
