package discount

import (
	"stock-controll/internal/domain/entity/product"
	"stock-controll/internal/domain/services/discount"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func New() *Mock {
	return &Mock{}
}

func (dm *Mock) Apply(item product.IProduct, purchasedQuantity int) (discount.DiscountSummary, error) {
	args := dm.Called(item, purchasedQuantity)
	return args.Get(0).(discount.DiscountSummary), args.Error(1)
}

func (dm *Mock) IsProductValid(item product.IProduct) bool {
	args := dm.Called(item)
	return args.Bool(0)
}
