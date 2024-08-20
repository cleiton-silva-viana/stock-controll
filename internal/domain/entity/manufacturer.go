package entity

import "stock-controll/internal/domain/value_object"

type Manufacturer struct {
	name     *value_object.Name
	category *value_object.Name
	cnpj     *value_object.CNPJ
	contact  *Contact
}

type ManufacturerBuilder struct {
	name     *value_object.Name
	category *value_object.Name
	cnpj     *value_object.CNPJ
	contact  *Contact
	errorList []error
}

func NewManufacturerBuilder() *ManufacturerBuilder {
	return &ManufacturerBuilder{}
}

func (m *ManufacturerBuilder) SetName(name string) *ManufacturerBuilder {
	nameParsed, err := value_object.NewName(name, "name", 50)
	if err != nil {
		m.errorList = append(m.errorList, err)
		return m
	}
	m.name = nameParsed
	return m
}

func (m *ManufacturerBuilder) SetCategory(category string) *ManufacturerBuilder {
	categoryParsed, err := value_object.NewName(category, "category", 30)
	if err != nil {
		m.errorList = append(m.errorList, err)
		return m
	}
	m.category = categoryParsed
	return m
}

func (m *ManufacturerBuilder) SetCNPJ(cnpj string) *ManufacturerBuilder {
	cpnjParsed, err := value_object.NewCNPJ(cnpj)
	if err != nil {
		m.errorList = append(m.errorList, err)
		return m
	}
	m.cnpj = cpnjParsed
	return m
}

func (m *ManufacturerBuilder) SetContact(email, phone string) *ManufacturerBuilder {
	contactParsed, errList := NewContact(email, phone)
	if errList != nil {
		m.errorList = append(m.errorList, errList...)
		return m
	}
	m.contact = contactParsed
	return m
}

func (m *ManufacturerBuilder) Build() (*Manufacturer, []error) {
	if len(m.errorList) > 0 {
		return nil, m.errorList

	} 
	return &Manufacturer{
		name:     m.name,
		category: m.category,
		cnpj: m.cnpj,
		contact:  m.contact,
	}, nil
}
