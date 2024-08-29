package entity

import (
	vo "stock-controll/internal/domain/value_object"
)

type Manufacturer struct {
	name     *vo.Name
	category *vo.Name
	cnpj     *vo.CNPJ
}

type ManufacturerBuilder struct {
	name      *vo.Name
	category  *vo.Name
	cnpj      *vo.CNPJ
	errorList []error
}

func NewManufacturerBuilder() *ManufacturerBuilder {
	return &ManufacturerBuilder{}
}

func (m *ManufacturerBuilder) Name(name string) *ManufacturerBuilder {
	nameParsed, err :=  vo.NameDefaultValidation("name", name)

	if err != nil {
		m.errorList = append(m.errorList, err)
		return m
	}
	m.name = nameParsed
	return m
}

func (m *ManufacturerBuilder) Category(category string) *ManufacturerBuilder {
	const minLength = 2
	const maxLength = 25

	categoryBuilder := vo.NewNameBuilder()
	categoryParsed, err := categoryBuilder.Field("user name", category).
		Length(minLength, maxLength).
		Build()

	if err != nil {
		m.errorList = append(m.errorList, err)
		return m
	}
	m.category = categoryParsed
	return m
}

func (m *ManufacturerBuilder) CNPJ(cnpj string) *ManufacturerBuilder {
	cpnjParsed, err := vo.NewCNPJ(cnpj)
	if err != nil {
		m.errorList = append(m.errorList, err)
		return m
	}
	m.cnpj = cpnjParsed
	return m
}

func (m *ManufacturerBuilder) Build() (*Manufacturer, []error) {
	if len(m.errorList) > 0 {
		return nil, m.errorList

	}
	return &Manufacturer{
		name:     m.name,
		category: m.category,
		cnpj:     m.cnpj,
	}, nil
}
