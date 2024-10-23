package cart

/*
	TODO: Prioridade: Defina a prioridade de aplicação dos descontos, caso mais de um desconto seja aplicável ao mesmo tempo.

	Persistência: O carrinho pode ser armazenado em um banco de dados,
	e pode ser associado a uma sessão de usuário ou a um login,
	dependendo do modelo de autenticação da loja.

	Expiração: Pode ser interessante implementar uma lógica para
	expirar carrinhos que não foram atualizados por um determinado período.

	Desconto de Fidelidade
	Descrição: Descontos oferecidos a clientes que fazem parte de um programa de fidelidade.
	Exemplo: "Clientes do programa de fidelidade ganham 15% de desconto em todas as compras."

	Desconto em Combos
	Descrição: Descontos aplicados quando produtos específicos são comprados juntos.
	Exemplo: "Compre um refrigerante e um lanche e ganhe 20% de desconto no lanche."

	Desconto de Fidelidade
	Descrição: Descontos oferecidos a clientes que fazem parte de um programa de fidelidade.
	Exemplo: "Clientes do programa de fidelidade ganham 15% de desconto em todas as compras."

	Desconto por Data de Aniversário
	Descrição: Descontos especiais oferecidos a clientes em seu mês de aniversário.
	Exemplo: "Ganhe 20% de desconto em qualquer compra durante o mês do seu aniversário."

	Desconto por Pagamento Antecipado
	Descrição: Descontos oferecidos para pagamentos feitos antes do prazo.
	Exemplo: "5% de desconto para pagamentos feitos até 10 dias antes da data de vencimento."

	Desconto por Pagamento no PIX
	Descrição: Descontos oferecidos para pagamentos feitos via PIX.
	Exemplo: "1% de desconto nos pagamentos feitos via pix"
*/

/*
{
    "cart_uuid": "123",
    "user_uuid": "123",
	"created_at": "0000.00.00",
	"updated_at": "0000.00.00",
	"cart_status": "active",
    "products_purchased": {
        "product_uuid": { ... },
		"product_uuid": { ... },
    }
    "total": {
        "subtotal": 300.00,
        "discounts": 10.00,
        "finalTotal": 289.00
    },
    "appliedCoupons": [
		"ANIVERSARIO2024",
		"CUSTOMER",
	],
    "shippingInfo": {
        "shippingAddress": "rua canudos",
        "shippingMethod": "car",
        "shippingCost": 12.00
    }
    "services": [],
    "taxes": []
}
*/

import (
	"time"

	"stock-controll/internal/domain/entity/common"
	"stock-controll/internal/domain/entity/coupon"
	validationerrors "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/validation"
)

type CartStatus string

const (
	Active    CartStatus = "active"
	Finished  CartStatus = "finished"
	Abandoned CartStatus = "abandoned"
)

type cart struct {
	cartUUID       string
	userUUID       string
	createdAt      time.Time
	updatedAt      time.Time
	status         CartStatus
	products       map[string]ProductCart
	appliedCoupons []coupon.ICoupon
	total          struct {
		subtotal   float64
		discounts  float64
		finalTotal float64
	}
}

func NewCart(userUUID string) (*cart, validationerrors.IValidationError) {
	isValid := common.IsValidUUUID(userUUID)
	if !isValid {
		err := validationerrors.NewValidationError("cart")
		err.AddValidationError(&validation.FieldError{
			FieldName:  "user_uuid",
			CodeErrors: []string{string(validation.ErrUnknown)},
		})
		return nil, err
	}
	return &cart{
		cartUUID:       common.GenerateUUID(),
		userUUID:       userUUID,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
		status:         Active,
		products:       make(map[string]ProductCart, 0),
		appliedCoupons: make([]coupon.ICoupon, 0),
	}, nil
}

func (c *cart) GetUUID() string {
	return c.cartUUID
}

func (c *cart) GetUserUUID() string {
	return c.userUUID
}

func (c *cart) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *cart) GetUpdatedAt() time.Time {
	return c.updatedAt
}

func (c *cart) GetStatus() CartStatus {
	return c.status
}

func (c *cart) GetSubtotal() float64 {
	return c.total.subtotal
}

func (c *cart) GetTotalFinal() float64 {
	return c.total.finalTotal
}

func (c *cart) GetDiscounts() float64 {
	return c.total.discounts
}

// Produtos devem ser retornados em ordem alfabética
func (c *cart) GetProductsInCart() []ProductCart {
	itemsInCartCopy := make([]ProductCart, 0, len(c.products))
	for _, item := range c.products {
		itemsInCartCopy = append(itemsInCartCopy, item)
	}
	return itemsInCartCopy
}

func (c *cart) AddProductInCart(item product.IProduct) *validation.FieldError {
	// carrinho inativo ou abandonado não podem ser modificados
	if c.status != Active {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	if item == nil {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	// Se o produto já existe no carrinho, retornamos um erro ou simplismente retornamos nil
	// Checar qual melhor abordagem
	_, exists := c.products[item.GetUUID()]
	if exists {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	productCart, _ := NewProductCart(item)

	// realizar um loop sobre os descontos e cupons adicionados ao carrinho
	for _, coupon := range c.appliedCoupons {
		discountApplicable := coupon.IsProductValid(productCart)
		// Se o cupom for válido para o produto, adiucionamos o cupom ao produto
		if discountApplicable {
			// No ato de adicionar um cupom a um produto, tal produto deve recalcular seus preços e descontos automáticamente
			productCart.AddCoupon(coupon)
		}
	}

	// Com os devidos descontos aplicados
	// Criamos o produto no map e atualizamos o subtotal, discounto aplicado e total final
	c.products[productCart.GetUUID()] = *productCart
	c.updatedAt = time.Now()
	c.calculate()
	return nil
}

func (c *cart) RemoveItemInCart(productUUID string) *validation.FieldError {
	// Itens não podem ser subtraídos de carrinhos finalizados ou abandonados
	if c.status != Active {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	// Se o produto não existir no carrinho
	// Retornamos nil ou um erro? checar!
	product, exists := c.products[c.cartUUID]
	if !exists {
		return nil
	}

	// Vamos recuperar o subtotal, discounto aplicado e total final do produto
	// E iremos subtrair do carrinho
	c.total.subtotal -= product.GetSubtotal()
	c.total.discounts -= product.GetTotalDiscountApplied()
	c.total.finalTotal -= product.GetFinalTotal()

	delete(c.products, productUUID)
	c.updatedAt = time.Now()
	return nil
}

// É uma cópia superficial ou profunda?
// Temos risco de algum ponteiro ser alterado e quebrar o sistema posteriormente?
func (c *cart) GetAppliedCoupons() []coupon.ICoupon {
	copiedCoupons := make([]coupon.ICoupon, len(c.appliedCoupons))
	copy(copiedCoupons, c.appliedCoupons)
	return copiedCoupons
}

func (c *cart) AddCoupon(coupon coupon.ICoupon) *validation.FieldError {
	if c.status != Active {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	if coupon != nil {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	// verificando se o cupom já está aplicado
	for _, currentCoupon := range c.appliedCoupons {
		if currentCoupon.GetUUID() == coupon.GetUUID() {
			return &validation.FieldError{
				FieldName:  "coupon",
				CodeErrors: []string{string(validation.ErrUnknown)},
			}
		}
	}

	for _, product := range c.products {
		// Verificando quais produtos o cupom se aplica
		isApplicable := coupon.IsProductValid(product)
		if isApplicable {
			// Aplicando o cupom de descontos apenas nos produtos elegíveis
			// Ignorando os erros
			// O que fazer s eum cupom não puder ser aplicado a um produto?
			// Devemos manter o erro lá, ou o erro cá?
			// Pensar a respeito
			product.AddCoupon(coupon)
		}
	}

	c.appliedCoupons = append(c.appliedCoupons, coupon)
	c.calculate()
	return nil
}

func (c *cart) RemoveCoupon(couponUUID string) *validation.FieldError {
	if c.status != Active {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	// verificar se o cupon está presente no carrinho
	var indice int
	for i, currentCoupon := range c.appliedCoupons {
		if currentCoupon.GetUUID() == couponUUID {
			indice = i
		}
	}

	// Iterar sobre os produtos e remover o cupom dos respectivos produtos
	for _, product := range c.products {
		product.RemoveCoupon(couponUUID)
	}

	c.appliedCoupons = append(c.appliedCoupons[:indice], c.appliedCoupons[indice+1:]...)

	return &validation.FieldError{
		FieldName:  "coupon_uuid",
		CodeErrors: []string{string(validation.ErrUnknown)},
	}
}


// TODO: a implementação de como os descontos são aplicados e priorizados ainda precisa ser feita.
func (c *cart) calculate() {
	c.calculateDiscount()
}

func (c *cart) calculateCoupons() {}

func (c *cart) calculateDiscount() {
	for _, product := range c.products {
		c.total.discounts = product.GetTotalDiscountApplied()
		c.total.finalTotal = product.GetFinalTotal()
		c.total.subtotal = product.GetSubtotal()
	}
}

func (c *cart) calculateTaxes() {}
