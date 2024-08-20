package entity

import (
	"stock-controll/internal/domain/failure"
	"stock-controll/internal/domain/value_object"
	"strings"
)

type Product struct {
	ID             int
	name           string
	description    string
	barcode        string
	manufacturerID int
	categoryID     int
}

type ProductBuilder struct {
	ID             int
	name           string
	description    string
	barcode        string
	manufacturerID int
	categoryID     int
	errorList      []error
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

func (p *ProductBuilder) SetName(name string) *ProductBuilder {
	nameParsed := strings.Trim(name, " ")
	if nameParsed == "" {
		p.errorList = append(p.errorList, failure.NameIsEmpty("name"))
		return p
	}

	const MIN_LENGTH = 2
	if len(nameParsed) < MIN_LENGTH {
		p.errorList = append(p.errorList, failure.NameIsShort("name", MIN_LENGTH))
		return p
	}

	const MAX_LENGTH = 25
	if len(nameParsed) > MAX_LENGTH {
		p.errorList = append(p.errorList, failure.NameIsLong("name", MAX_LENGTH))
		return p
	}

	nameWithoutSpaces := strings.ReplaceAll(name, " ", "")
	if value_object.ContainsSpecialChars(nameWithoutSpaces) {
		p.errorList = append(p.errorList, failure.NameWithInvalidChars("name"))
		return p
	}

	p.name = nameParsed
	return p
}

func (p *ProductBuilder) SetDescription(description string) *ProductBuilder {
	descriptionParsed := strings.Trim(description, " ")
	if descriptionParsed == "" {
		p.errorList = append(p.errorList, failure.NameIsEmpty("description"))
		return p
	}

	const MIN_LENGTH = 10
	if len(descriptionParsed) < MIN_LENGTH {
		p.errorList = append(p.errorList, failure.NameIsShort("description", MIN_LENGTH))
		return p
	}

	const MAX_LENGTH = 100
	if len(descriptionParsed) > MAX_LENGTH {
		p.errorList = append(p.errorList, failure.NameIsLong("description", MAX_LENGTH))
		return p
	}

	p.name = descriptionParsed
	return p
}

func (p *ProductBuilder) SetBarcode(barcode string) *ProductBuilder {
	// Validações

	p.barcode = barcode
	return p
}

func (p *ProductBuilder) SetManufacturerID(manufacturerID int) *ProductBuilder {
	// Validações

	p.manufacturerID = manufacturerID
	return p
}

func (p *ProductBuilder) SetCategoryID(categoryID int) *ProductBuilder {
	// Validações

	p.categoryID = categoryID
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
		manufacturerID: p.manufacturerID,
		categoryID:     p.categoryID,
	}, nil
}
