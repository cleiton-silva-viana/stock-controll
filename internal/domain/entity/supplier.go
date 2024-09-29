package entity

import (
	"net/http"
	"stock-controll/internal/domain/failure"
	vo "stock-controll/internal/domain/value_object"
)

type Supplier struct {
	uid int
	name vo.Name
	CNPJ vo.CNPJ
	contactBilling Contact
	contactPurschase Contact
	sellers []SellerSupplier
}

type SupplierBuilder struct {
	id int
	name vo.Name
	CNPJ vo.CNPJ
	contactBilling Contact
	contactPurschase Contact
	sellers []SellerSupplier
	errorList []error
	// adicionar condições de pagamento 
	// adiciona rprazo de entrega
	// adicionar pedido mínimo

}

func NewSupplierBuilder() *SupplierBuilder {
	return &SupplierBuilder{}
}

func (s *SupplierBuilder) Indentity(name, CNPJ string) *SupplierBuilder {
	nameValidated, err := vo.NameDefaultValidation("name", name)
	if err != nil {
		s.errorList = append(s.errorList, err)
	}

	CNPJvalidated, err := vo.NewCNPJ(CNPJ)
	if err != nil {
		s.errorList = append(s.errorList, err)
	}

	if nameValidated == nil || CNPJvalidated == nil {
		return s
	}

	s.name = *nameValidated
	s.CNPJ = *CNPJvalidated
	return s
}

func (s *SupplierBuilder) ContactBilling(ID int, email, phone string) *SupplierBuilder {
	contact, err := NewContact(ID, email, phone)

	if err != nil {
		s.errorList = append(s.errorList, err)
		return s
	}
	s.contactBilling = *contact
	return s
}

func (s *SupplierBuilder) ContactPurschase(ID int, email, phone string) *SupplierBuilder {
	contact, err := NewContact(ID, email, phone)
	if err != nil {
		s.errorList = append(s.errorList, err)
		return s
	}
	s.contactBilling = *contact
	return s
}

func (s *SupplierBuilder) Sellers(sellers []SellerSupplier) *SupplierBuilder {
	return s
}

func (s *SupplierBuilder) Build() (*Supplier, *failure.Fields) {
	if len(s.errorList) > 0 {
		return nil, &failure.Fields{
			Status: http.StatusBadRequest,
			ErrList: s.errorList,
		}
	}
	return &Supplier{
		name: s.name,
		CNPJ: s.CNPJ,
		contactBilling: s.contactBilling,
		contactPurschase: s.contactPurschase,
		sellers: s.sellers,
	}, nil
}
