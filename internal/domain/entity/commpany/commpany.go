package commpany

import (
	"fmt"
	"regexp"

	"stock-controll/internal/domain/entity/address"
	"stock-controll/internal/domain/entity/contact"
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/validate"
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

type CompanyStatus string

const (
	Active   = "active"
	Inactive = "inactive"
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
	uuid            string
	name            string
	cnpj            string
	billingContact  contact.IContact
	purchaseContact contact.IContact
	address         address.IAddress
	status          CompanyStatus
}

func New(config Config) (*Company, error) {
	var instance Company

	errs := validationerrors.New("company").
		AddValidationError(instance.SetName(config.Name)).
		AddValidationError(instance.SetCNPJ(config.CNPJ)).
		AddValidationError(instance.SetStatus(config.Status)).
		AddValidationError(instance.SetBillingContact(
			config.BillingContact.Email(),
			config.BillingContact.Phone())).
		AddValidationError(instance.SetPurchaseContact(
			config.PurchaseContact.Email(),
			config.PurchaseContact.Phone(),
		)).
		AddValidationError(instance.SetAddress(config.Address))

	if errs.HasError() {
		return nil, errs
	}
	return &instance, nil
}

func (c *Company) UUID() string {
	return c.uuid
}

func (c *Company) Name() string {
	return c.name
}

const (
	companyNameMinLength = 3
	companyNameMaxLength = 100
)

func (c *Company) SetName(name string) error {
	err := validate.New("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(companyNameMinLength, companyNameMaxLength),
	)
	if err == nil {
		c.name = name
	}
	return err
}

func (c *Company) CNPJ() string {
	return c.cnpj
}

const CNPJLength = 18
const ErrCNPJWithInvalidFormat = "ERR_CNPJ_WITH_INVALID_FORMAT"

func (c *Company) SetCNPJ(cnpj string) error {
	re := `^(\d{2})(?:\.(\d{3}))(?:\.(\d{3}))(?:\/(\d{4}))(?:\-(\d{2})$)`

	err := validate.New("cnpj", cnpj,
		validate.IsBlank(),
		validate.IsLengthEqualTo(CNPJLength),
		validate.IsFormatValid(regexp.MustCompile(re), ErrCNPJWithInvalidFormat),
	)
	if err == nil {
		c.cnpj = cnpj
	}
	return err
}

const ErrInvalidCompanyStatus = "ERR_INVALID_COMPANY_STATUS"

func (c *Company) SetStatus(newStatus string) error {
	statusParsed := CompanyStatus(newStatus)
	
	if statusParsed == Active || statusParsed == Inactive {
		c.status = statusParsed
		return nil
	}

	return &validate.FieldError{
		FieldName: "status",
		CodeError: ErrInvalidCompanyStatus,
	}
}

func (c *Company) BillingContact() string {
	err := validate.New[any]("billing_contact", c.billingContact,
		validate.IsNil(c.billingContact),
	)
	if err.HasError() {
		return "no has billing contact"
	}
	return fmt.Sprintf("Email: %s\nPhone: %s", c.billingContact, c.billingContact.Phone())
}

func (c *Company) SetBillingContact(email, phone string) error {
	contact, err := contact.New(email, phone)
	if err == nil {
		c.billingContact = contact
	}
	return err
}

func (c *Company) BillingEmail() string {
	return c.billingContact.Email()
}

func (c *Company) SetBillingEmail(email string) error {
	return c.billingContact.SetEmail(email)
}

func (c *Company) BillingPhone() string {
	return c.billingContact.Phone()
}

func (c *Company) SetBillingPhone(phone string) error {
	return c.billingContact.SetPhone(phone)
}

func (c *Company) PurchaseContact() string {
	if c.billingContact == nil {
		return "purchase contact not available"
	}
	return fmt.Sprintf("Email: %s\nPhone: %s", c.purchaseContact.Email(), c.purchaseContact.Phone())
}

func (c *Company) SetPurchaseContact(email, phone string) error {
	contact, err := contact.New(email, phone)
	if err == nil {
		c.purchaseContact = contact
	}
	return err
}

func (c *Company) PurchaseEmail() string {
	return c.purchaseContact.Email()
}

func (c *Company) SetPurchaseEmail(email string) error {
	return c.purchaseContact.SetEmail(email)
}

func (c *Company) PurchasePhone() string {
	return c.purchaseContact.Phone()
}

func (c *Company) SetPurchasePhone(phone string) error {
	return c.purchaseContact.SetPhone(phone)
}

func (c *Company) Address() address.IAddress {
	return c.address
}

func (c *Company) SetAddress(newAddress address.IAddress) error {
	err := validate.New[any]("address", newAddress,
		validate.IsNil(newAddress),
	)
	if err == nil {
		c.address = newAddress
	}
	return err
}
