package entity

import (
	"stock-controll/internal/domain/failure"
	"stock-controll/internal/domain/validation"
	vo "stock-controll/internal/domain/value_object"
	"strings"
)

type Product struct {
	id             int
	name           vo.Name
	description    string
	barcode        string
	brandID        int
	manufacturerID int
	categoryID     int
}

type ProductBuilder struct {
	id             int
	name           vo.Name
	description    string
	barcode        string
	manufacturerID int
	brandID        int
	categoryID     int
	errorList      []error
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

func (p *ProductBuilder) Name(name string) *ProductBuilder {
	builder := vo.NewNameBuilder()
	nameValidated, err := builder.Field("name", name).
		Length(2, 25).
		Digts(true).
		SpecialCharacters(true).
		Build()
	if err != nil {
		p.errorList = append(p.errorList, err)
		return p
	}
	p.name = *nameValidated
	return p
}

func (p *ProductBuilder) Description(description string) *ProductBuilder {
	if strings.ReplaceAll(description, " ", "") == "" {
		p.errorList = append(p.errorList, failure.FieldIsEmpty("description"))
	}

	if len(description) < 10 || len(description) > 100 {
		p.errorList = append(p.errorList, failure.FieldNotRangeValues("description", []string{"10", "100"}))
		return p
	}
	p.description = description
	return p
}

func (p *ProductBuilder) Barcode(barcode string) *ProductBuilder {
	builder := vo.NewNameBuilder()
	_, err := builder.Field("barcode", barcode).
		Digts(true).
		Length(6, 24).
		SpecialCharacters(false).
		Build()
	if err != nil {
		p.errorList = append(p.errorList, err)
		return p
	}

	if validation.ContainsLetters(barcode) {
		p.errorList = append(p.errorList, failure.FieldWithLetters("barcode"))
		return p
	}

	p.barcode = barcode
	return p
}

func (p *ProductBuilder) ManufacturerID(manufacturerID int) *ProductBuilder {
	if manufacturerID < 0 || manufacturerID > 100000 {
		p.errorList = append(p.errorList, failure.FieldNotRangeValues("manufacturer", []string{"0", "100000"}))
		return p
	}
	p.manufacturerID = manufacturerID
	return p
}

func (p *ProductBuilder) CategoryID(categoryID int) *ProductBuilder {
	if categoryID < 0 || categoryID > 100 {
		p.errorList = append(p.errorList, failure.FieldNotRangeValues("category", []string{"0", "100"}))
		return p
	}
	p.categoryID = categoryID
	return p
}

func (p *ProductBuilder) BrandID(brandID int) *ProductBuilder {
	if brandID < 0 || brandID > 10000 {
		p.errorList = append(p.errorList, failure.FieldNotRangeValues("brand", []string{"0", "100"}))
		return p
	}
	p.brandID = brandID
	return p
}

func (p *ProductBuilder) Build() (*Product, []error) {
	if len(p.errorList) > 0 {
		return nil, p.errorList
	}
	return &Product{
		name:           p.name,
		description:    p.description,
		barcode:        p.barcode,
		brandID:        p.brandID,
		manufacturerID: p.manufacturerID,
		categoryID:     p.categoryID,
	}, nil
}
