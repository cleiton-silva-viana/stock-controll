package discount

import (
	"stock-controll/internal/domain/entity/common"
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/validation"
)

type IProductSpecificDiscountStrategy interface {
	Apply(item product.IProduct, purchasedQuantity int) (DiscountSummary, *validation.FieldError)
	IsProductValid(item product.IProduct) bool
}

type TemplateDiscountStrategy struct {
	discount IDiscountCalculationStrategy
}

func (b *TemplateDiscountStrategy) Apply(item product.IProduct, quantityPurchased int) (DiscountSummary, *validation.FieldError) {
	if b.IsProductValid(item) {
		return b.discount.Apply(item.GetPrice(), quantityPurchased)
	}
	return DiscountSummary{
			totalDiscountApplied: 0,
			finalTotal:           item.GetPrice() * float64(quantityPurchased),
		}, &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
}

func (b *TemplateDiscountStrategy) IsProductValid(item product.IProduct) bool {
	return false
}

type GlobalDiscountStrategy struct {
	TemplateDiscountStrategy
}

// TODO: Adicionar validações
func NewGlobalDiscountStrategy(discount IDiscountCalculationStrategy) (*GlobalDiscountStrategy, *validation.FieldError) {
	if discount == nil {
		return nil, &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
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

func NewSpecificCategoryDiscountStrategy(categoriesUUIDs map[string]struct{}, discount IDiscountCalculationStrategy) (*SpecificCategoryDiscountStrategy, *validation.FieldError) {
	if discount == nil {
		return nil, &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
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
	_, exists := c.categories[item.GetCategoryUUID()]
	return exists
}

type SpecificBrandDiscountStrategy struct {
	TemplateDiscountStrategy
	brands map[string]struct{}
}

func NewSpecificBrandDiscountStrategy(brandsUUIDs map[string]struct{}, discount IDiscountCalculationStrategy) (*SpecificBrandDiscountStrategy, *validation.FieldError) {
	if discount == nil {
		return nil, &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
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
	_, exists := b.brands[item.GetBrandUUID()]
	return exists
}

type SpecificProductsDiscount struct {
	TemplateDiscountStrategy
	productsUUID map[string]struct{}
}

func NewSpecificProductsDiscount(productsUUID map[string]struct{}, discount IDiscountCalculationStrategy) (*SpecificProductsDiscount, *validation.FieldError) {
	if discount == nil {
		return nil, &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
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
	_, exists := s.productsUUID[item.GetUUID()]
	return exists
}

func validateUUIDS(resouceName string, uuids map[string]struct{}) (map[string]struct{}, *validation.FieldError) {
	if uuids == nil {
		return nil, &validation.FieldError{
			FieldName:  resouceName,
			CodeErrors: []string{"o recurso esperado não pode ser"},
		}
	}

	var validUUIDs = make(map[string]struct{})

	for uuid := range uuids {
		isValid := common.IsValidUUUID(uuid)
		if isValid {
			validUUIDs[uuid] = struct{}{}
		}
	}

	if len(uuids) == 0 {
		return nil, &validation.FieldError{
			FieldName:  resouceName,
			CodeErrors: []string{"deve haver ao menos um %s para criar o desconto..."},
		}
	}

	return validUUIDs, nil
}
