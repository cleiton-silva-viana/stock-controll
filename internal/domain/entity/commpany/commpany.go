package commpany

import (
	"fmt"
	"regexp"

	"stock-controll/internal/domain/entity/address"
	"stock-controll/internal/domain/entity/contact"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type ICompany interface {
	UUID() string
	Name() string
	CNPJ() string
	BillingContact() string
	BillingEmail() string
	BillingPhone() string
	PurchaseContact() string
	PurchaseEmail() string
	PurchasePhone() string
}

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)

type Config struct {
	UUID            string
	Name            string
	CNPJ            string
	Status          string
	BillingContact  contact.IContact
	PurchaseContact contact.IContact
	Address         address.IAddress
}
type Company struct {
	uuid            uuid.UUID
	name            string
	cnpj            string // TODO: vai virar value object
	billingContact  contact.IContact
	purchaseContact contact.IContact
	address         address.IAddress
	status          Status
}

const prefix = "COM"

// TODO: criar função específica para validações de contato
// Devemos impedir que contato comercial seja o mesmo contato de cobrança?
func New(config Config) (*Company, error) {
	errors := entity.Error("company")

	nameErr := validateName(config.Name)
	cnpjErr := validateCNPJ(config.CNPJ)
	statusErr := validateStatus(config.Status)
	id, idErr := uuid.New(prefix)

	errors.
		AddValidationError(nameErr).
		AddValidationError(cnpjErr).
		AddValidationError(statusErr).
		AddValidationError(idErr)

	// validateBillingContact(
	// 		config.BillingContact.Email(),
	// 		config.BillingContact.Phone())).
	// 	AddValidationError(instance.SetPurchaseContact(
	// 		config.PurchaseContact.Email(),
	// 		config.PurchaseContact.Phone(),
	// 	)).
	// 	AddValidationError(instance.SetAddress(config.Address))

	if errors.HasError() {
		return nil, errors
	}
	return &Company{
		uuid: *id,
		name: config.Name,
		cnpj: config.CNPJ,
	}, nil
}

func (c *Company) UUID() string {
	return c.uuid.String()
}

func (c *Company) Name() string {
	return c.name
}

func (c *Company) Address() address.IAddress {
	return c.address
}

func (c *Company) CNPJ() string {
	return c.cnpj
}

func (c *Company) BillingContact() string {
	err := validate.New[any](
		"billing_contact", c.billingContact,
		validate.IsNil(c.billingContact),
	)
	if err.HasError() {
		return "no has billing contact"
	}
	return fmt.Sprintf("Email: %s\nPhone: %s", c.billingContact, c.billingContact.Phone())
}

func (c *Company) BillingEmail() string {
	return c.billingContact.Email()
}

func (c *Company) BillingPhone() string {
	return c.billingContact.Phone()
}

func (c *Company) PurchaseContact() string {
	if c.billingContact == nil {
		return "purchase contact not available"
	}
	return fmt.Sprintf("Email: %s\nPhone: %s", c.purchaseContact.Email(), c.purchaseContact.Phone())
}

func (c *Company) PurchaseEmail() string {
	return c.purchaseContact.Email()
}

func (c *Company) PurchasePhone() string {
	return c.purchaseContact.Phone()
}

const (
	companyNameMinLength = 3
	companyNameMaxLength = 100
)

func validateName(name string) error {
	return validate.New(
		"name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(companyNameMinLength, companyNameMaxLength),
	)
}

const CNPJLength = 18
const ErrCNPJWithInvalidFormat = "ERR_CNPJ_WITH_INVALID_FORMAT"

func validateCNPJ(cnpj string) error {
	re := `^(\d{2})(?:\.(\d{3}))(?:\.(\d{3}))(?:\/(\d{4}))(?:\-(\d{2})$)`

	return validate.New(
		"cnpj", cnpj,
		validate.IsBlank(),
		validate.IsLengthEqualTo(CNPJLength),
		validate.IsFormatValid(regexp.MustCompile(re), ErrCNPJWithInvalidFormat),
	)
}

const ErrInvalidCompanyStatus = "ERR_INVALID_COMPANY_STATUS"

// TODO: verificar
func validateStatus(newStatus string) error {
	statusParsed := CompanyStatus(newStatus)

	if statusParsed == Active || statusParsed == Inactive {
		c.status = statusParsed
		return nil
	}

	return &field.FieldError{
		FieldName: "status",
		CodeError: ErrInvalidCompanyStatus,
	}
}

// TODO: verificar
func validateBillingContact(email, phone string) error {
	contact, err := contact.New(email, phone)
	if err == nil {
		c.billingContact = contact
	}
	return err
}

// TODO: verificar
func validateBillingEmail(email string) error {
	return c.billingContact.SetEmail(email)
}

// TODO: verificar
func validateBillingPhone(phone string) error {
	return c.billingContact.SetPhone(phone)
}

func validatePurchaseContact(email, phone string) error {
	contact, err := contact.New(email, phone)
	if err == nil {
		c.purchaseContact = contact
	}
	return err
}

func validatePurchaseEmail(email string) error {
	return c.purchaseContact.SetEmail(email)
}

func validatePurchasePhone(phone string) error {
	return purchaseContact.SetPhone(phone)
}

func validateAddress(newAddress address.IAddress) error {
	return validate.New[any](
		"address", newAddress,
		validate.IsNil(newAddress),
	)
}
