package entity

import (
	"stock-controll/internal/domain/failure"
	vo "stock-controll/internal/domain/value_object"
)

type SellerSupplier struct {
	id         int
	supplierID int
	firstName  vo.Name
	lastName   vo.Name
	contact    Contact
}

type SellerSupplierBuilder struct {
	id         int
	supplierID int
	firstName  vo.Name
	lastName   vo.Name
	email      vo.Email
	phone      vo.Phone
	errorList  []error
}

func NewSellerSupplierBuilder() *SellerSupplierBuilder {
	return &SellerSupplierBuilder{}
}

func (s *SellerSupplierBuilder) Name(firstName, lastName string) *SellerSupplierBuilder {
	firstNameValidated, err := vo.NameDefaultValidation("firstName", firstName)
	if err != nil {
		s.errorList = append(s.errorList, err)
	}

	lastNameValidated, err := vo.NameDefaultValidation("lastName", lastName)
	if err != nil {
		s.errorList = append(s.errorList, err)
	}

	if firstNameValidated != nil && lastNameValidated != nil {
		s.firstName = *firstNameValidated
		s.lastName = *lastNameValidated
	}

	return s
}

func (s *SellerSupplierBuilder) SupplierID(ID int) *SellerSupplierBuilder {
	if ID < 0 || ID > 1000 {
		s.errorList = append(s.errorList, failure.FieldNotRangeValues("supplierID", []string{"0", "1000"}))
		return s
	}
	s.supplierID = ID
	return s
}

func (s *SellerSupplierBuilder) Contact(email, phone string) *SellerSupplierBuilder {
	emailValidated, err := vo.NewEmail(email)
	if err != nil {
		s.errorList = append(s.errorList, err)
	}

	phoneValidated, err := vo.NewPhone(phone)
	if err != nil {
		s.errorList = append(s.errorList, err)
	}

	if emailValidated != nil && phoneValidated != nil {
		s.email = *emailValidated
		s.phone = *phoneValidated
	}
	return s
}

func (s *SellerSupplierBuilder) Build() (*SellerSupplier, []error) {
	if len(s.errorList) > 0 {
		return nil, s.errorList
	} 
	
	return &SellerSupplier{
		supplierID: s.supplierID,
		firstName: s.firstName,
		lastName: s.lastName,
		contact: Contact{ 
			phone: s.phone, 
			email: s.email,
		}}, nil
}
