package discount

import (
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type IProductSpecificDiscountStrategy interface {
	Apply(item product.IProduct, purchasedQuantity int) (DiscountSummary, error)
	IsProductValid(item product.IProduct) bool
}

type TemplateDiscountStrategy struct {
	discount IDiscountCalculationStrategy
}

// TODO: implementar error
func (b *TemplateDiscountStrategy) Apply(item product.IProduct, quantityPurchased int) (DiscountSummary, error) {
	if b.IsProductValid(item) {
		return b.discount.Apply(item.Price(), quantityPurchased)
	}
	return DiscountSummary{
			totalDiscountApplied: 0,
			finalTotal:           item.Price() * float64(quantityPurchased),
		}, &validate.FieldError{
			CodeError: "",
		}
}

func (b *TemplateDiscountStrategy) IsProductValid(item product.IProduct) bool {
	return false
}

type GlobalDiscountStrategy struct {
	TemplateDiscountStrategy
}

// TODO: Adicionar validações
// TODO: implementar error
func NewGlobalDiscountStrategy(discount IDiscountCalculationStrategy) (*GlobalDiscountStrategy, error) {
	if discount == nil {
		return nil, &validate.FieldError{
			CodeError: "",
		}
	}

	return &GlobalDiscountStrategy{
		TemplateDiscountStrategy: TemplateDiscountStrategy{discount: discount},
	}, nil
}

func (a *GlobalDiscountStrategy) IsProductValid(item product.IProduct) bool {
	return true
}

type SpecificCategoryDiscountStrategy struct {
	TemplateDiscountStrategy
	categories map[string]struct{}
}

// TODO: implementar error
func NewSpecificCategoryDiscountStrategy(categoriesUUIDs map[string]struct{}, discount IDiscountCalculationStrategy) (*SpecificCategoryDiscountStrategy, error) {
	if discount == nil {
		return nil, &validate.FieldError{
			CodeError: "",
		}
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

func (c *SpecificCategoryDiscountStrategy) IsProductValid(item product.IProduct) bool {
	_, exists := c.categories[item.CategoryUUID()]
	return exists
}

type SpecificBrandDiscountStrategy struct {
	TemplateDiscountStrategy
	brands map[string]struct{}
}

// TODO: implementar error
func NewSpecificBrandDiscountStrategy(brandsUUIDs map[string]struct{}, discount IDiscountCalculationStrategy) (*SpecificBrandDiscountStrategy, error) {
	if discount == nil {
		return nil, &validate.FieldError{
			CodeError: "",
		}
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

// TODO: implementar error
func NewSpecificProductsDiscount(productsUUID map[string]struct{}, discount IDiscountCalculationStrategy) (*SpecificProductsDiscount, error) {
	if discount == nil {
		return nil, &validate.FieldError{
			CodeError: "",
		}
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

// TODO: implementar error
func validateUUIDS(resouceName string, uuids map[string]struct{}) (map[string]struct{}, error) {
	if uuids == nil {
		return nil, &validate.FieldError{
			FieldName: resouceName,
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
		return nil, &validate.FieldError{
			FieldName: resouceName,
			CodeError: "deve haver ao menos um x para criar o desconto...",
		}
	}

	return validUUIDs, nil
}
