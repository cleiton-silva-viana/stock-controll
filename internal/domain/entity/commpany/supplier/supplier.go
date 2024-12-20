package supplier

import (
	"stock-controll/internal/domain/entity/payment"
	"time"

	"stock-controll/internal/domain/entity/address"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/valueobject/uuid"

	"stock-controll/internal/domain/entity/user"
)

type Supplier struct {
	company Company
	sellers map[string]user.IUser
	pay payment.IPayment
}

func New(name, cnpj, billingEmail, billingPhone, purchaseEmail, purchasePhone string, addr address.IAddress) (*Supplier, error) {
	var supplierErrors = entity.Error("supplier")
	var supplierInstance = Supplier{
		company: Company{
			uuid:   uuid.New(),
			status: Active,
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

func (s *Supplier) AddSaller(saller user.IUser) {
	s.sellers[saller.UUID()] = saller
}

func (s *Supplier) RemoveSaller(uuid string) {
	_, exist := s.sellers[uuid]
	if exist {
		delete(s.sellers, uuid)
	}
}

func (s *Supplier) Sallers() []user.IUser {
	var sellers = make([]user.IUser, 0, len(s.sellers))
	for _, seller := range s.sellers {
		sellers = append(sellers, seller)
	}
	return sellers
}

type paymentCondition string

const ()

// TODO: imnplementar
func (s *Supplier) SetPaymentCondition() {}
