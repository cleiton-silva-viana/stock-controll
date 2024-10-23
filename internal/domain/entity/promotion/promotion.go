package promotion

// TODO: A fazer

// Aplicar descontos a compras no atacado
// A partir de 6 unidades, p exemplo
// Implementar posteriormente...

// Descontos Diretos: O preço do produto é reduzido diretamente na etiqueta ou no caixa. Por exemplo, um produto que custa R$ 10 pode ter um desconto de R$ 2, sendo vendido por R$ 8.
// Promoções de "Leve Mais, Pague Menos": Ofertas que incentivam a compra de múltiplos itens, como "compre 2 e leve 3" ou "leve 3 e pague 2". O desconto é aplicado no total da compra.
// Descontos em Cartão de Fidelidade: Clientes que possuem um cartão de fidelidade podem receber descontos exclusivos ou acumular pontos que podem ser trocados por descontos em compras futuras.
// Descontos em Produtos Específicos: Produtos em promoção podem ter preços reduzidos por tempo limitado, como "oferta do dia" ou "promoção de fim de semana".
// Descontos em Compras em Grande Quantidade: Algumas vezes, supermercados oferecem descontos para compras em maior volume, como "preço especial para atacado".
// Descontos Sazonais: Promoções que ocorrem em datas específicas, como Black Friday, Natal ou liquidações de fim de estação.
// Descontos em Produtos Próximos da Data de Validade: Produtos que estão prestes a vencer podem ter descontos significativos para incentivar a venda antes do prazo.

import (
	"time"

	"stock-controll/internal/domain/entity/common"
	"stock-controll/internal/domain/entity/discount"
	validationerrors "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/validation"
)

type IPromotion interface {
	GetUUID() string
	GetName() string
	GetDescription() string
	GetStartDate() time.Time
	GetEndDate() time.Time
	GetIsActive() bool
	GetStatus() promotionStatus
	CancelPromotion() *validation.FieldError
	ExtendPromotionDate(date time.Time) *validation.FieldError
	InitPromotion()
	FinalizePromotion()
}

type promotionStatus string

const (
	Scheduled  promotionStatus = "scheduled"   // promoção agendada
	InProgress promotionStatus = "in_progress" // promoção em progresso
	Ended      promotionStatus = "ended"
	Canceled   promotionStatus = "canceled"
)

type promotion struct {
	uuid              string
	name              string
	description       string
	startDate         time.Time
	endDate           time.Time
	status            promotionStatus
	extendedPromotion int
	discount.IProductSpecificDiscountStrategy
}

func NewPromotion(name, description string, discount discount.IProductSpecificDiscountStrategy, startDate, endDate time.Time) (*promotion, validationerrors.IValidationError) {
	promotionErrors := validationerrors.NewValidationError("promotion")

	promotionErrors.
		AddValidationError(validateName(name)).
		AddValidationError(validateDescription(description)).
		AddValidationError(validateStartDate(startDate, endDate)).
		AddValidationError(validateEndDate(endDate)).
		AddValidationError(validateDiscount(discount))

	if promotionErrors.HasError() {
		return nil, promotionErrors
	}

	return &promotion{
		uuid:                             common.GenerateUUID(),
		name:                             name,
		description:                      description,
		startDate:                        startDate,
		endDate:                          endDate,
		IProductSpecificDiscountStrategy: discount,
	}, nil
}

func (p *promotion) GetUUID() string {
	return p.uuid
}

func (p *promotion) GetName() string {
	return p.name
}

func (p *promotion) GetDescription() string {
	return p.description
}

func (p *promotion) GetStartDate() time.Time {
	return p.startDate
}

func (p *promotion) GetEndDate() time.Time {
	return p.endDate
}

func (p *promotion) GetStatus() promotionStatus {
	return p.status
}

func (p *promotion) CancelPromotion() *validation.FieldError {
	if p.status == Ended || p.status == Canceled {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	p.status = Canceled
	return nil
}

const maxDateExtension = time.Hour*24 - 7

// Notificações: Se a promoção for estendida, pode ser útil notificar os clientes ou usuários do sistema sobre a nova data de término.
// Limites de Extensão: Você pode querer definir limites sobre quantas vezes uma promoção pode ser estendida ou por quanto tempo. Isso pode ajudar a evitar abusos e garantir que as promoções permaneçam relevantes.
func (p *promotion) ExtendPromotionDate(newEndDate time.Time) *validation.FieldError {
	if p.extendedPromotion != 0 {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	if newEndDate.After(p.endDate.Add(maxDateExtension)) {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	p.extendedPromotion += 1
	p.endDate = newEndDate
	return nil
}

func (p *promotion) FinalizePromotion() {
	if p.status != InProgress {
		return
	}

	if time.Now().After(p.endDate) {
		p.status = Ended
	}
}

func (p *promotion) InitPromotion() {
	if p.status != Scheduled {
		return
	}

	if time.Now().Before(p.startDate) {
		p.status = InProgress
	}
}

func (p *promotion) IsActive() bool {
	return p.status == InProgress
}

func (p *promotion) PromotionIsValidFor(item product.IProduct) bool {
	return p.IProductSpecificDiscountStrategy.IsProductValid(item)
}

func (p *promotion) StartPromotionScheduler() {
	go func() {
		for {
			time.Sleep(time.Hour * 24)
			p.CheckPromotion()
		}
	}()
}

func (p *promotion) CheckPromotion() {

}

const (
	minLengthForPromotionName = 6
	maxLengthForPromotionName = 24
)

func validateName(name string) *validation.FieldError {
	return validation.Validate[string]("name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minLengthForPromotionName, maxLengthForPromotionName, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
}

const (
	minDescriptionLength = 12
	maxDescriptionLength = 1000
)

func validateDescription(description string) *validation.FieldError {
	return validation.Validate[string]("description", description,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minDescriptionLength, maxDescriptionLength, validation.ErrUnknown),
	)
}

const (
	maxStartDateOffset   = time.Hour * 24 * 30
	maxPromotionDuration = time.Hour * 24 * 180
)

/*
Validação: A validação de promoções é abrangente, mas, assim como no pacote de cupons, a validação de datas pode ser refatorada para evitar repetição.
*/

// Tanto inicio quanto fim da promoção devem ocorrer na virada de um dia para outro, ou seja, nos configuramos as promoções para iniciarem a partir da 00:00 de um dia específico e terminar a partir das 00:00 de outro dia posterior
func validateStartDate(startDate, endDate time.Time) *validation.FieldError {
	return validation.Validate[time.Time]("start_date", startDate,
		validation.IsAfterThan(endDate, validation.ErrUnknown),
		validation.IsAfterThan(time.Now().Add(maxStartDateOffset), validation.ErrUnknown),
	)
}

func validateEndDate(endDate time.Time) *validation.FieldError {
	return validation.Validate[time.Time]("end_date", endDate,
		validation.IsBeforeThan(time.Now(), validation.ErrUnknown),
		validation.IsAfterThan(time.Now().Add(maxPromotionDuration), validation.ErrUnknown),
	)
}

func validateDate(date time.Time) *validation.FieldError {
	// Data não pode ter minutos e horas setados
	// Queremos uma data do tipo: 2024-01-01TM00:00:00
	if date.Hour() != 0 || date.Minute() != 0 || date.Second() != 0 {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	return nil
}

func validateDiscount(discount discount.IProductSpecificDiscountStrategy) *validation.FieldError {
	return validation.Validate[any]("discount", discount,
		validation.IsNil(discount, validation.ErrUnknown),
	)
}
