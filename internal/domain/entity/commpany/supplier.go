package entity

import (
	"time"

	"stock-controll/internal/domain/entity/address"
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/entity/user"
)

type supplier struct {
	company
	sellers            map[string]user.IUser
	paymnetsConditions string
}

func NewSupplier(name, cnpj, billingEmail, billingPhone, purchaseEmail, purchasePhone string, addr address.IAddress) (*supplier, validationError.IValidationError) {
	var supplierErrors = validationError.NewValidationError("supplier")
	var supplierInstance = supplier{
		company: company{
			uuid:   common.GenerateUUID(),
			status: active,
		},
		sellers: make(map[string]user.IUser),
	}

	supplierErrors.
		AddValidationError(supplierInstance.SetName(name)).
		AddValidationError(supplierInstance.SetCNPJ(cnpj)).
		AddValidationError(supplierInstance.SetBillingContact(billingEmail, billingPhone)).
		AddValidationError(supplierInstance.SetPurchaseContact(purchaseEmail, purchasePhone)).
		AddValidationError(supplierInstance.SetAddress(addr))

	if supplierErrors.HasError() {
		return nil, supplierErrors
	}

	return &supplierInstance, nil

}

func (s *supplier) AddSaller(saller user.IUser) {
	s.sellers[saller.GetUUID()] = saller
}

func (s *supplier) RemoveSaller(uuid string) {
	_, exist := s.sellers[uuid]
	if exist {
		delete(s.sellers, uuid)
	}
}

func (s *supplier) GetSallers() []user.IUser {
	var sellers = make([]user.IUser, 0, len(s.sellers))
	for _, seller := range s.sellers {
		sellers = append(sellers, seller)
	}
	return sellers
}

type paymentCondition string

const ()

// TODO: imnplementar
func (s *supplier) SetPaymentCondition() {

}

type supplierBuilder struct {
	ICompany
	sellers            []user.IUser
	paymentsConditions string
	minimunOrder       int
	deliveryTime       time.Time
}
