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
	"stock-controll/internal/domain/valueobject/discount/scope"
	"time"

	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type IPromotion interface {
	UUID() string
	Name() string
	Description() string
	StartDate() time.Time
	EndDate() time.Time
	IsActive() bool
	Status() Status
	CancelPromotion() field.FieldError
	ExtendPromotionDate(date time.Time) field.FieldError
	InitPromotion()
	FinalizePromotion()
}

type Status string

const (
	Scheduled  Status = "scheduled"   // promoção agendada
	InProgress Status = "in_progress" // promoção em progresso
	Ended      Status = "ended"       // promoção finalizada
	Canceled   Status = "canceled"    // promoção cancelada
)

type Scope string

const (
	Specific     Scope = "specific"
	Category     Scope = "category"
	Brand        Scope = "brand"
	Manufacturer Scope = "manufacturer"
	Distinct     Scope = "distinct"
)

type Promotion struct {
	uuid              uuid.UUID
	name              string
	description       string
	status            Status
	scope             Scope
	discount          scope.IProductSpecificDiscountStrategy // sem método acessor
	startDate         time.Time
	endDate           time.Time
	extendedPromotion int // sem método acessor
}

type ConfigPromotion struct {
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Status      string
	Scope       string
	Discount    scope.IProductSpecificDiscountStrategy
}

func New(config ConfigPromotion) (*Promotion, error) {
	errors := entity.Error("promotion")

	statusParsed, statusErr := parseStatus(config.Status)
	scopeParsed, scopeErr := parseScope(config.Scope)

	errors.
		AddValidationError(statusErr).
		AddValidationError(scopeErr).
		AddValidationError(validateName(config.Name)).
		AddValidationError(validateDescription(config.Description)).
		AddValidationError(validateDiscount(config.Discount))

	if errors.HasError() {
		return nil, errors
	}

	errors.
		AddValidationError(validateStartDate(config.Status, config.StartDate, config.EndDate)).
		AddValidationError(validateEndDate(config.Status, config.EndDate))

	if errors.HasError() {
		return nil, errors
	}

	return &Promotion{
		uuid:        *uuid.New(),
		name:        config.Name,
		description: config.Description,
		startDate:   config.StartDate,
		endDate:     config.EndDate,
		discount:    config.Discount,
		status:      statusParsed,
		scope:       scopeParsed,
	}, nil
}

func (p *Promotion) UUID() string {
	return p.uuid.String()
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

func (p *Promotion) Status() Status {
	return p.status
}

const ErrPromotionStatusChangeNotAllowed = "ERR_PROMOTION_STATUS_CHANGE_NOT_ALLOWED"

func (p *Promotion) CancelPromotion() error {
	if p.status == Ended || p.status == Canceled {
		return &field.FieldError{
			FieldName: "status",
			CodeError: ErrPromotionStatusChangeNotAllowed,
		}
	}
	p.status = Canceled
	return nil
}

const (
	maxDateExtension                        = time.Hour * 24 * 7
	ErrPromotionExtensionLimitReached       = "ERR_PROMOTION_EXTENSION_LIMIT_REACHED"
	ErrPromotionExtensionExceedsMaximumDays = "ERR_PROMOTION_EXTENSION_EXCEEDS_MAXIMUM_DAYS"
)

func (p *Promotion) ExtendPromotionDate(endDate time.Time) error {
	if p.extendedPromotion != 0 {
		return &field.FieldError{
			FieldName: "extension_promotion",
			CodeError: ErrPromotionExtensionLimitReached,
		}
	}

	if endDate.After(p.endDate.Add(maxDateExtension)) {
		return &field.FieldError{
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
	return p.discount.IsProductValid(item)
}

// ????
func (p *Promotion) StartPromotionScheduler() {
	go func() {
		for {
			time.Sleep(time.Hour * 24)
			p.CheckPromotion()
		}
	}()
}

// ????
func (p *Promotion) CheckPromotion() {}

const ErrCannotFinishPromotionNotInProgress = "ERR_CANNOT_FINISH_PROMOTION_NOT_IN_PROGRESS"

func (p *Promotion) finalizePromotion() error {
	if p.status != InProgress {
		return &field.FieldError{
			FieldName: "status",
			CodeError: ErrCannotFinishPromotionNotInProgress,
		}
	}

	if time.Now().After(p.endDate) {
		p.status = Ended
	}
	return nil
}

const (
	minLengthForPromotionName = 6
	maxLengthForPromotionName = 24
)

func validateName(name string) error {
	return validate.New[string](
		"name", name,
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
	return validate.New[string](
		"description", description,
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

func validateStartDate(promotionStatus string, startDate, endDate time.Time) error {
	if Status(promotionStatus) != InProgress && Status(promotionStatus) != Scheduled {
		return nil
	}

	err := validateDate("start_date", endDate)
	if err != nil {
		return err
	}

	return validate.New[time.Time](
		"start_date", startDate,
		validate.IsAfterThan(time.Now().Add(maxStartDateOffset), ErrPromotionStartDateBeforeCurrent),
		validate.IsAfterThan(endDate, ErrPromotionStartDateAfterEndDate),
	)
}

const (
	ErrPromotionExpirationDateAfterLimit   = "ERR_PROMOTION_EXPIRATION_DATE_AFTER_LIMIT"
	ErrPromotionExpirationDateBeforeCurret = "ERR_PROMOTION_EXPIRATION_DATE_BEFORE_CURRENT"
)

func validateEndDate(status string, endDate time.Time) error {
	if err := validateDate("end_date", endDate); err != nil {
		return err
	}
	return validate.New[time.Time](
		"end_date", endDate,
		validate.IsBeforeThan(time.Now(), ErrPromotionExpirationDateBeforeCurret),
		validate.IsAfterThan(time.Now().Add(maxPromotionDuration), ErrPromotionExpirationDateAfterLimit),
	)
}

const ErrTimeFieldsMustBeZeroed = "ERR_TIME_FIELDS_MUST_BE_ZEROED"

func validateDate(fieldName string, date time.Time) error {
	if date.Hour() != 0 || date.Minute() != 0 || date.Second() != 0 {
		return &field.FieldError{
			FieldName: fieldName,
			CodeError: ErrTimeFieldsMustBeZeroed,
		}
	}
	return nil
}

func validateDiscount(discount scope.IProductSpecificDiscountStrategy) error {
	return validate.New[any](
		"discount", discount,
		validate.IsNil(discount),
	)
}

const ErrInvalidPromotionStatus = "ERR_INVALID_PROMOTION_STATUS"

func parseStatus(statusStr string) (Status, error) {
	switch statusStr {
	case string(Scheduled):
		return Scheduled, nil
	case string(InProgress):
		return InProgress, nil
	case string(Ended):
		return Ended, nil
	case string(Canceled):
		return Canceled, nil
	default:
		return "", &field.FieldError{
			FieldName:    "status",
			CodeError:    ErrInvalidPromotionStatus,
			InvalidValue: statusStr,
		}
	}
}

const ErrInvalidPromotionScope = "ERR_INVALID_PROMOTION_SCOPE"

func parseScope(scopeStr string) (Scope, error) {
	switch scopeStr {
	case string(Specific):
		return Specific, nil
	case string(Category):
		return Category, nil
	case string(Brand):
		return Brand, nil
	case string(Manufacturer):
		return Manufacturer, nil
	case string(Distinct):
		return Distinct, nil
	default:
		return "", &field.FieldError{
			FieldName:    "scope",
			CodeError:    ErrInvalidPromotionScope,
			InvalidValue: scopeStr,
		}
	}
}

/*

Promotion representa uma promoção no sistema de e-commerce.

type Promotion struct {
    ID            string        `json:"id"`
    Nome          string        `json:"nome"`
    Descricao     string        `json:"descricao,omitempty"`
    Tipo          TipoPromotion `json:"tipo"`
    ValorDesconto float64       `json:"valor_desconto"` // Valor do desconto em porcentagem (%)
    DataInicio    time.Time     `json:"data_inicio"`
    DataFim       time.Time     `json:"data_fim"`


	Produtos      []string      `json:"produtos,omitempty"`      // IDs dos produtos específicos
    Categoria     string        `json:"categoria,omitempty"`     // Categoria de produtos
    Marca         string        `json:"marca,omitempty"`         // Marca dos produtos
    Fabricante    string        `json:"fabricante,omitempty"`    // Fabricante dos produtos
    ListaProdutos []string      `json:"lista_produtos,omitempty"`// Lista de IDs de produtos distintos
}
*/
