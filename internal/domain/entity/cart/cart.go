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

	cartproduct "stock-controll/internal/domain/entity/cart_product"
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/entity/coupon"
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type CartStatus string

const (
	Active    CartStatus = "active"
	Finished  CartStatus = "finished"
	Abandoned CartStatus = "abandoned"
)

type Cart struct {
	cartUUID       string
	userUUID       string
	createdAt      time.Time
	updatedAt      time.Time
	status         CartStatus
	products       map[string]cartproduct.ProductCart
	appliedCoupons []coupon.ICoupon
	total          struct {
		subtotal   float64
		discounts  float64
		totalFinal float64
	}
}

func New(userUUID string) (*Cart, error) {
	err := uuid.IsValid("user_uuid", userUUID)
	if err != nil {
		cartError := validationerrors.New("cart").AddValidationError(err)
		return nil, cartError
	}

	return &Cart{
		cartUUID:       uuid.New(),
		userUUID:       userUUID,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
		status:         Active,
		products:       make(map[string]cartproduct.ProductCart, 0),
		appliedCoupons: make([]coupon.ICoupon, 0),
	}, nil
}

func (c *Cart) UUID() string {
	return c.cartUUID
}

func (c *Cart) UserUUID() string {
	return c.userUUID
}

func (c *Cart) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Cart) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Cart) Status() CartStatus {
	return c.status
}

func (c *Cart) Subtotal() float64 {
	return c.total.subtotal
}

func (c *Cart) TotalFinal() float64 {
	return c.total.totalFinal
}

func (c *Cart) Discounts() float64 {
	return c.total.discounts
}

func (c *Cart) ProductsInCart() []cartproduct.ProductCart {
	itemsInCartCopy := make([]cartproduct.ProductCart, 0, len(c.products))
	for _, item := range c.products {
		itemsInCartCopy = append(itemsInCartCopy, item)
	}
	return itemsInCartCopy
}

const ErrCartInactiveOrAbandoned = "ERR_CART_INACTIVE_OR_ABANDONED"

func (c *Cart) AddProductInCart(item product.IProduct) error {
	if c.status != Active {
		return &validate.FieldError{
			FieldName: "status",
			CodeError: ErrCartInactiveOrAbandoned,
		}
	}

	err := validate.New[any]("item", item, validate.IsNil(item))
	if err != nil {
		return err
	}

	// Se o produto já existe no carrinho, retornamos um erro
	_, exists := c.products[item.UUID()]
	if exists {
		return nil
	}

	productCart, _ := cartproduct.New(item)

	// realizar um loop sobre os descontos e cupons adicionados ao carrinho
	for _, coupon := range c.appliedCoupons {
		discountApplicable := coupon.IsProductValid(productCart)
		// Se o cupom for válido para o produto, adicionamos o cupom ao produto
		if discountApplicable {
			// No ato de adicionar um cupom a um produto, tal produto deve recalcular seus preços e descontos automáticamente
			productCart.AddCoupon(coupon)
		}
	}

	// Com os devidos descontos aplicados
	// Criamos o produto no map e atualizamos o subtotal, discounto aplicado e total final
	c.products[productCart.UUID()] = *productCart
	c.updatedAt = time.Now()
	c.calculate()
	return nil
}

func (c *Cart) RemoveItemInCart(productUUID string) error {
	// Itens não podem ser subtraídos de carrinhos finalizados ou abandonados
	if c.status != Active {
		return &validate.FieldError{
			CodeError: ErrCartInactiveOrAbandoned,
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
	c.total.subtotal -= product.Subtotal()
	c.total.discounts -= product.TotalDiscountApplied()
	c.total.totalFinal -= product.FinalTotal()

	delete(c.products, productUUID)
	c.updatedAt = time.Now()
	return nil
}

// É uma cópia superficial ou profunda?
// Temos risco de algum ponteiro ser alterado e quebrar o sistema posteriormente?
func (c *Cart) AppliedCoupons() []coupon.ICoupon {
	copiedCoupons := make([]coupon.ICoupon, len(c.appliedCoupons))
	copy(copiedCoupons, c.appliedCoupons)
	return copiedCoupons
}

func (c *Cart) AddCoupon(coupon coupon.ICoupon) error {
	if c.status != Active {
		return &validate.FieldError{
			CodeError: ErrCartInactiveOrAbandoned,
		}
	}

	err := validate.New[any]("coupon", coupon, validate.IsNil(coupon))
	if err != nil {
		return err
	}

	const ErrCouponAlreadyApplied = "ERR_COUPON_ALREADY_APPLIED"
	// "ERR_COUPON_ALREADY_APPLIED": {
	// 	"Message": "The discount coupon has already been applied to the cart.",
	// 	"Solution": "Please use a different coupon or remove the existing one before applying a new discount."
	// }
	for _, currentCoupon := range c.appliedCoupons {
		if currentCoupon.UUID() == coupon.UUID() {
			return &validate.FieldError{
				FieldName: "coupon",
				CodeError: ErrCouponAlreadyApplied,
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

// TODO: adicionar erro
func (c *Cart) RemoveCoupon(couponUUID string) error {
	if c.status != Active {
		return &validate.FieldError{
			CodeError: ErrCartInactiveOrAbandoned,
		}
	}

	// verificar se o cupon está presente no carrinho
	var indice int
	for i, currentCoupon := range c.appliedCoupons {
		if currentCoupon.UUID() == couponUUID {
			indice = i
		}
	}

	// <<< refatorar - código em ordem errada! >>>

	// Iterar sobre os produtos e remover o cupom dos respectivos produtos
	for _, product := range c.products {
		product.RemoveCoupon(couponUUID)
	}

	c.appliedCoupons = append(c.appliedCoupons[:indice], c.appliedCoupons[indice+1:]...)

	return &validate.FieldError{
		FieldName: "coupon_uuid",
		CodeError: "",
	}
}

// TODO: a implementação de como os descontos são aplicados e priorizados ainda precisa ser feita.
func (c *Cart) calculate() {
	c.calculateDiscount()
}

func (c *Cart) calculateCoupons() {}

func (c *Cart) calculateDiscount() {
	for _, product := range c.products {
		c.total.discounts = product.TotalDiscountApplied()
		c.total.totalFinal = product.FinalTotal()
		c.total.subtotal = product.Subtotal()
	}
}

func (c *Cart) calculateTaxes() {}
