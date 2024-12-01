package product

import (
	"stock-controll/internal/domain/entity/tag"

	"github.com/stretchr/testify/mock"
)

// ProductMock é um mock da interface IProduct
type ProductMock struct {
	mock.Mock
}

func Mock() *ProductMock {
	return &ProductMock{}
}

// UUID implementa o método UUID da interface IProduct
func (m *ProductMock) UUID() string {
	args := m.Called()
	return args.String(0)
}

// Name implementa o método Name da interface IProduct
func (m *ProductMock) Name() string {
	args := m.Called()
	return args.String(0)
}

// Price implementa o método Price da interface IProduct
func (m *ProductMock) Price() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

// Quantity implementa o método Quantity da interface IProduct
func (m *ProductMock) Quantity() int {
	args := m.Called()
	return args.Int(0)
}

// Description implementa o método Description da interface IProduct
func (m *ProductMock) Description() string {
	args := m.Called()
	return args.String(0)
}

// Barcode implementa o método Barcode da interface IProduct
func (m *ProductMock) Barcode() string {
	args := m.Called()
	return args.String(0)
}

// BrandUUID implementa o método BrandUUID da interface IProduct
func (m *ProductMock) BrandUUID() string {
	args := m.Called()
	return args.String(0)
}

// Tags implementa o método Tags da interface IProduct
func (m *ProductMock) Tags() []tag.Tag {
	args := m.Called()
	return args.Get(0).([]tag.Tag)
}

// ManufacturerUUID implementa o método ManufacturerUUID da interface IProduct
func (m *ProductMock) ManufacturerUUID() string {
	args := m.Called()
	return args.String(0)
}

// CategoryUUID implementa o método CategoryUUID da interface IProduct
func (m *ProductMock) CategoryUUID() string {
	args := m.Called()
	return args.String(0)
}
