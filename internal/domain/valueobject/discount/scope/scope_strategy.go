package scope

import (
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/services/error/domain"
	"stock-controll/internal/domain/services/error/field"
	discount2 "stock-controll/internal/domain/valueobject/discount/strategy"
	"stock-controll/internal/domain/valueobject/summary"
	"stock-controll/internal/domain/valueobject/uuid"
)

type IProductSpecificDiscountStrategy interface {
	Apply(item product.IProduct, purchasedQuantity int) (*summary.Summary, error)
	IsProductValid(item product.IProduct) bool
}

type TemplateDiscountStrategy struct {
	discount discount2.IDiscountCalculationStrategy
}

const (
	// TODO: corrigir o nome do erro abaixo lá na infra
	ErrProductNotEligibleForApplyThisDiscount = "ERR_PRODUCT_NOT_ELIGIBLE_FOR_THIS_DISCOUNT"
	ErrProductRequiredForApplyDiscount        = "ERR_PRODUCT_REQUIRED_FOR_APPLY_DISCOUNT"
)

func (tds *TemplateDiscountStrategy) Apply(item product.IProduct, quantityPurchased int) (*summary.Summary, error) {
	if item == nil {
		return nil, &domain.DomainError{
			ErrorCodes: []string{ErrProductRequiredForApplyDiscount},
		}
	}
	if !tds.IsProductValid(item) {
		subtotal := item.Price() * float64(quantityPurchased)
		sum, _ := summary.New(subtotal, subtotal)
		return sum, &domain.DomainError{
			ErrorCodes: []string{ErrProductNotEligibleForApplyThisDiscount},
		}
	}
	return tds.discount.Apply(item.Price(), quantityPurchased)
}

func (tds *TemplateDiscountStrategy) IsProductValid(item product.IProduct) bool {
	return item != nil
}

type GlobalDiscountStrategy struct {
	TemplateDiscountStrategy
}

// TODO: Adicionar validações
const ErrDiscountStrategyNotFound = ""

var errDiscountStrategyNotFound = &field.FieldError{
	FieldName: "discount",
	CodeError: ErrDiscountStrategyNotFound,
}

func NewGlobalDiscountStrategy(discount discount2.IDiscountCalculationStrategy) (*GlobalDiscountStrategy, error) {
	if discount == nil {
		return nil, errDiscountStrategyNotFound
	}
	return &GlobalDiscountStrategy{
		TemplateDiscountStrategy: TemplateDiscountStrategy{discount: discount},
	}, nil
}

func (a *GlobalDiscountStrategy) IsProductValid(item product.IProduct) bool {
	return item != nil
}

type SpecificCategoryDiscountStrategy struct {
	TemplateDiscountStrategy
	categories map[string]struct{}
}

func NewSpecificCategoryDiscountStrategy(categoriesUUIDs map[string]struct{}, discount discount2.IDiscountCalculationStrategy) (*SpecificCategoryDiscountStrategy, error) {
	if discount == nil {
		return nil, errDiscountStrategyNotFound
	}
	uuids, err := validateUUIDS("categories_uuid", categoriesUUIDs)
	if err != nil {
		return nil, err
	}
	return &SpecificCategoryDiscountStrategy{
		TemplateDiscountStrategy: TemplateDiscountStrategy{discount: discount},
		categories:               uuids,
	}, nil
}

func (scd *SpecificCategoryDiscountStrategy) IsProductValid(item product.IProduct) bool {
	if item == nil {
		return false
	}
	_, exists := scd.categories[item.CategoryUUID()]
	return exists
}

// TODO: Remover o erro abaixo lá na camada de infra
const ErrDiscountCategoryNotFound = "ERR_DISCOUNT_CATEGORY_NOT_FOUND"

type SpecificBrandDiscountStrategy struct {
	TemplateDiscountStrategy
	brands map[string]struct{}
}

// TODO: implementar error
func NewSpecificBrandDiscountStrategy(brandsUUIDs map[string]struct{}, discount discount2.IDiscountCalculationStrategy) (*SpecificBrandDiscountStrategy, error) {
	if discount == nil {
		return nil, errDiscountStrategyNotFound
	}
	uuids, err := validateUUIDS("brands_uuid", brandsUUIDs)
	if err != nil {
		return nil, err
	}
	return &SpecificBrandDiscountStrategy{
		TemplateDiscountStrategy: TemplateDiscountStrategy{discount: discount},
		brands:                   uuids,
	}, nil
}

func (b *SpecificBrandDiscountStrategy) IsProductValid(item product.IProduct) bool {
	_, exists := b.brands[item.BrandUUID()]
	return exists
}

type SpecificProductsDiscount struct {
	TemplateDiscountStrategy
	productsUUID map[string]struct{}
}

func NewSpecificProductsDiscount(productsUUID map[string]struct{}, discount discount2.IDiscountCalculationStrategy) (*SpecificProductsDiscount, error) {
	if discount == nil {
		return nil, errDiscountStrategyNotFound
	}
	uuids, err := validateUUIDS("products_uuid", productsUUID)
	if err != nil {
		return nil, err
	}
	return &SpecificProductsDiscount{
		TemplateDiscountStrategy: TemplateDiscountStrategy{discount: discount},
		productsUUID:             uuids,
	}, nil
}

func (s *SpecificProductsDiscount) IsProductValid(item product.IProduct) bool {
	_, exists := s.productsUUID[item.UUID()]
	return exists
}

func validateUUIDS(resourceName string, uuids map[string]struct{}) (map[string]struct{}, error) {
	if uuids == nil {
		return nil, &field.FieldError{
			FieldName: resourceName,
			CodeError: "",
		}
	}
	var validUUIDs = make(map[string]struct{})
	for u := range uuids {
		if err := uuid.IsValid("", u); err != nil {
			validUUIDs[u] = struct{}{}
		}
	}
	if len(uuids) == 0 {
		// TODO: criar erro
		return nil, &field.FieldError{
			FieldName: resourceName,
			CodeError: "deve haver ao menos um x para criar o desconto...",
		}
	}
	return validUUIDs, nil
}
