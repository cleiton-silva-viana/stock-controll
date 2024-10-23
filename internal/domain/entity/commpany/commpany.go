package entity

// TODO: implementar o método para status

import (
	"fmt"
	"regexp"

	"stock-controll/internal/domain/entity/address"
	"stock-controll/internal/domain/entity/contact"
	"stock-controll/internal/domain/validation"
)

type ICompany interface {
	GetUUID() string
	GetName() string
	GetCNPJ() string
	GetBillingContact() string
	GetBillingEmail() string
	GetBillingPhone() string
	GetPurchaseContact() string
	GetPurchaseEmail() string
	GetPurchasePhone() string
}

type companyStatus string

const (
	active   = "active"
	inactive = "inactive"
)

type company struct {
	uuid            string
	name            string
	cnpj            string
	billingContact  contact.IContact
	purchaseContact contact.IContact
	address         address.IAddress
	status          companyStatus
}

func (c *company) GetUUID() string {
	return c.uuid
}

func (c *company) GetName() string {
	return c.name
}

const (
	companyNameMinLength = 3
	companyNameMaxLength = 100
)

func (c *company) SetName(name string) *validation.FieldError {
	err := validation.Validate("name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(companyNameMinLength, companyNameMaxLength, validation.ErrUnknown),
	)
	if err == nil {
		c.name = name
	}
	return err
}

func (c *company) GetCNPJ() string {
	return c.cnpj
}

const CNPJLength = 18

// Deve ser imutável
func (c *company) SetCNPJ(cnpj string) *validation.FieldError {
	re := `^(\d{2})(?:\.(\d{3}))(?:\.(\d{3}))(?:\/(\d{4}))(?:\-(\d{2})$)`

	err := validation.Validate("cnpj", cnpj,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthEqualTo(CNPJLength, validation.ErrUnknown),
		validation.IsFormatValid(regexp.MustCompile(re), validation.ErrUnknown),
	)
	if err == nil {
		c.cnpj = cnpj
	}
	return err
}

// TODO: Melhorar - verificar um formato mais interessante
func (c *company) GetBillingContact() string {
	if c.billingContact == nil {
		return "contact not available"
	}
	return fmt.Sprintf("Email: %s\nPhone: %s", c.billingContact, c.billingContact.GetPhone())
}

func (c *company) SetBillingContact(email, phone string) *validation.FieldError {
	contact, err := contact.NewContact(email, phone)
	if err == nil {
		c.billingContact = contact
	}
	return nil
	// return err
}

func (c *company) GetBillingEmail() string {
	return c.billingContact.GetEmail()
}

func (c *company) SetBillingEmail(email string) *validation.FieldError {
	return c.billingContact.SetEmail(email)
}

func (c *company) GetBillingPhone() string {
	return c.billingContact.GetPhone()
}

func (c *company) SetBillingPhone(phone string) *validation.FieldError {
	return c.billingContact.SetPhone(phone)
}

func (c *company) GetPurchaseContact() string {
	if c.billingContact == nil {
		return "contact not available"
	}
	return fmt.Sprintf("Email: %s\nPhone: %s", c.purchaseContact.GetEmail(), c.purchaseContact.GetPhone())
}

func (c *company) SetPurchaseContact(email, phone string) *validation.FieldError {
	contact, err := contact.NewContact(email, phone)
	if err == nil {
		c.purchaseContact = contact
	}
	// return err
	return nil
}

func (c *company) GetPurchaseEmail() string {
	return c.purchaseContact.GetEmail()
}

func (c *company) SetPurchaseEmail(email string) *validation.FieldError {
	return c.purchaseContact.SetEmail(email)
}

func (c *company) GetPurchasePhone() string {
	return c.purchaseContact.GetPhone()
}

func (c *company) SetPurchasePhone(phone string) *validation.FieldError {
	return c.purchaseContact.SetPhone(phone)
}

func (c *company) GetAddress() address.IAddress {
	return c.address
}

// Adicionar validações
func (c *company) SetAddress(newAddress address.IAddress) *validation.FieldError {
	c.address = newAddress
	return nil
}
