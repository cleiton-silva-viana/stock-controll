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

	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/services/discount"
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type IPromotion interface {
	UUID() string
	Name() string
	Description() string
	StartDate() time.Time
	EndDate() time.Time
	IsActive() bool
	Status() PromotionStatus
	CancelPromotion() *validate.FieldError
	ExtendPromotionDate(date time.Time) *validate.FieldError
	InitPromotion()
	FinalizePromotion()
}

type PromotionStatus string

const (
	Scheduled  PromotionStatus = "scheduled"   // promoção agendada
	InProgress PromotionStatus = "in_progress" // promoção em progresso
	Ended      PromotionStatus = "ended"
	Canceled   PromotionStatus = "canceled"
)

type Promotion struct {
	uuid              string
	name              string
	description       string
	startDate         time.Time
	endDate           time.Time
	status            PromotionStatus
	extendedPromotion int
	discount.IProductSpecificDiscountStrategy
}

type ConfigPromotion struct {
	UUID              string
	Name              string
	Description       string
	StartDate         time.Time
	EndDate           time.Time
	Status            PromotionStatus
	ExtendedPromotion int
	Discount          discount.IProductSpecificDiscountStrategy // TODO: criar um método Get para etsa propriedade
}

func New(config ConfigPromotion) (*Promotion, error) {
	// temos que verificar se a pormoção está sendo recuperada do banco de dados ou se está sendo criada

	promotionErrors := validationerrors.New("promotion").
		AddValidationError(validateName(config.Name)).
		AddValidationError(validateDescription(config.Description)).
		AddValidationError(validateStartDate(config.Status, config.StartDate, config.EndDate)).
		AddValidationError(validateEndDate(config.Status, config.EndDate)).
		AddValidationError(validateDiscount(config.Discount))

	if promotionErrors.HasError() {
		return nil, promotionErrors
	}

	if config.UUID == "" {
		config.UUID = uuid.New()
	}

	return &Promotion{
		uuid:                             config.UUID,
		name:                             config.Name,
		description:                      config.Description,
		startDate:                        config.StartDate,
		endDate:                          config.EndDate,
		IProductSpecificDiscountStrategy: config.Discount,
		status:                           config.Status,
		extendedPromotion:                config.ExtendedPromotion,
	}, nil
}

func (p *Promotion) UUID() string {
	return p.uuid
}

func (p *Promotion) Name() string {
	return p.name
}

func (p *Promotion) Description() string {
	return p.description
}

func (p *Promotion) StartDate() time.Time {
	return p.startDate
}

func (p *Promotion) EndDate() time.Time {
	return p.endDate
}

func (p *Promotion) Status() PromotionStatus {
	return p.status
}

func (p *Promotion) ExtendedPromotion() int {
	return p.extendedPromotion
}

const ErrPromotionStatusChangeNotAllowed = "ERR_PROMOTION_STATUS_CHANGE_NOT_ALLOWED"

func (p *Promotion) CancelPromotion() error {
	if p.status == Ended || p.status == Canceled {
		return &validate.FieldError{
			FieldName: "status",
			CodeError: ErrPromotionStatusChangeNotAllowed,
		}
	}
	p.status = Canceled
	return nil
}

const maxDateExtension = time.Hour*24 - 7
const ErrPromotionExtensionLimitReached = "ERR_PROMOTION_EXTENSION_LIMIT_REACHED"
const ErrPromotionExtensionExceedsMaximumDays = "ERR_PROMOTION_EXTENSION_EXCEEDS_MAXIMUM_DAYS"

// Notificações: Se a promoção for estendida, pode ser útil notificar os clientes ou usuários do sistema sobre a nova data de término.
// Limites de Extensão: Você pode querer definir limites sobre quantas vezes uma promoção pode ser estendida ou por quanto tempo. Isso pode ajudar a evitar abusos e garantir que as promoções permaneçam relevantes.
func (p *Promotion) ExtendPromotionDate(endDate time.Time) error {
	if p.extendedPromotion != 0 {
		return &validate.FieldError{
			FieldName: "extension_promotion",
			CodeError: ErrPromotionExtensionLimitReached,
		}
	}

	if endDate.After(p.endDate.Add(maxDateExtension)) {
		return &validate.FieldError{
			FieldName: "extension_promotion",
			CodeError: ErrPromotionExtensionExceedsMaximumDays,
		}
	}

	p.extendedPromotion += 1
	p.endDate = endDate
	return nil
}

func (p *Promotion) InitPromotion() {
	if p.status != Scheduled {
		return
	}

	if time.Now().Before(p.startDate) {
		p.status = InProgress
	}
}

func (p *Promotion) IsActive() bool {
	return p.status == InProgress
}

func (p *Promotion) PromotionIsValidFor(item product.IProduct) bool {
	return p.IProductSpecificDiscountStrategy.IsProductValid(item)
}

func (p *Promotion) StartPromotionScheduler() {
	go func() {
		for {
			time.Sleep(time.Hour * 24)
			p.CheckPromotion()
		}
	}()
}

func (p *Promotion) CheckPromotion() {}

func (p *Promotion) finalizePromotion() {
	if p.status != InProgress {
		return
	}

	if time.Now().After(p.endDate) {
		p.status = Ended
	}
}

const (
	minLengthForPromotionName = 6
	maxLengthForPromotionName = 24
)

func validateName(name string) error {
	return validate.New[string]("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minLengthForPromotionName, maxLengthForPromotionName),
		validate.CheckSpecialChars(validate.Disallow),
	)
}

const (
	minDescriptionLength = 12
	maxDescriptionLength = 1000
)

func validateDescription(description string) error {
	return validate.New[string]("description", description,
		validate.IsBlank(),
		validate.IsLengthInRange(minDescriptionLength, maxDescriptionLength),
	)
}

const (
	maxStartDateOffset   = time.Hour * 24 * 30
	maxPromotionDuration = time.Hour * 24 * 180
)

const (
	ErrPromotionStartDateBeforeCurrent = "ERR_PROMOTION_START_DATE_BEFORE_CURRENT"
	ErrPromotionStartDateAfterEndDate  = "ERR_PROMOTION_START_DATE_AFTER_END_DATE"
)

// TODO: melhorar implementação da primeira condicional, veja que é repetida no método abaixo
func validateStartDate(promotionStatus PromotionStatus, startDate, endDate time.Time) error {
	//
	if promotionStatus != InProgress && promotionStatus != Scheduled {
		return nil
	}

	if err := validateDate("start_date", endDate); err != nil {
		return err
	}
	return validate.New[time.Time]("start_date", startDate,
		validate.IsAfterThan(time.Now().Add(maxStartDateOffset), ErrPromotionStartDateBeforeCurrent),
		validate.IsAfterThan(endDate, ErrPromotionStartDateAfterEndDate),
	)
}

const (
	ErrPromotionExpirationDateAfterLimit   = "ERR_PROMOTION_EXPIRATION_DATE_AFTER_LIMIT"
	ErrPromotionExpirationDateBeforeCurret = "ERR_PROMOTION_EXPIRATION_DATE_BEFORE_CURRENT"
)

func validateEndDate(promotionStatus PromotionStatus, endDate time.Time) error {
	if err := validateDate("end_date", endDate); err != nil {
		return err
	}
	return validate.New[time.Time]("end_date", endDate,
		validate.IsBeforeThan(time.Now(), ErrPromotionExpirationDateBeforeCurret),
		validate.IsAfterThan(time.Now().Add(maxPromotionDuration), ErrPromotionExpirationDateAfterLimit),
	)
}

const ErrTimeFieldsMustBeZeroed = "ERR_TIME_FIELDS_MUST_BE_ZEROED"

func validateDate(fieldName string, date time.Time) error {
	if date.Hour() != 0 || date.Minute() != 0 || date.Second() != 0 {
		return &validate.FieldError{
			FieldName: fieldName,
			CodeError: ErrTimeFieldsMustBeZeroed,
		}
	}
	return nil
}

func validateDiscount(discount discount.IProductSpecificDiscountStrategy) error {
	return validate.New[any]("discount", discount,
		validate.IsNil(discount),
	)
}
