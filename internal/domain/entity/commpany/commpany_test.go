package commpany

import (
	"strings"
	"testing"
	
	"stock-controll/internal/domain/entity/address"
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
)

var data = Config{
	Name:          unitary.Fake.Company().Name(),
	CNPJ:          "46.318.856/0001-00",
	// BillingEmail:  unitary.Fake.Internet().CompanyEmail(),
	// BillingPhone:  "(49)2524-2218",
	// PurchaseEmail: unitary.Fake.Internet().CompanyEmail(),
	// PurchasePhone: "(74)3017-4666",
}


var addrConfig = address.Config{
	Number: 1,
	Street: "rua canudos",
	City: "rio de janeiro",
	State: "rio de janeiro",
	PostalCode: "21500-300",
	Complement: "apartamaento 202",
}

func TestCompanyNoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[Config]{
		{
			Description: "name length is minimum allowed",
			Handler:         func(c *Config) { c.Name = strings.Repeat("a", companyNameMinLength) },
		},
		{
			Description: "name length is maximum allowed",
			Handler:         func(c *Config) { c.Name = strings.Repeat("b", companyNameMaxLength) },
		},
		{
			Description: "name contain special characters",
			Handler:         func(c *Config) { c.Name = "b&b Hammer" },
		},
		{
			Description: "name contain number",
			Handler:         func(c *Config) { c.Name = "1st price" },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			var data = hyundai
			test.Handler(&data)

			// Assert
			assert.Nil(t, companyInstance.SetName(data.name))
			assert.Nil(t, companyInstance.SetCNPJ(data.cnpj))
			assert.Nil(t, companyInstance.SetBillingEmail(data.billingEmail))
			assert.Nil(t, companyInstance.SetBillingPhone(data.billingPhone))
			assert.Nil(t, companyInstance.SetPurchaseEmail(data.purchaseEmail))
			assert.Nil(t, companyInstance.SetPurchasePhone(data.purchasePhone))
			assert.Nil(t, companyInstance.SetBillingContact(data.billingEmail, data.billingPhone))
			assert.Nil(t, companyInstance.SetPurchaseContact(data.purchaseEmail, data.purchasePhone))
			assert.Equal(t, data.name, companyInstance.Name())
			assert.Equal(t, data.cnpj, companyInstance.CNPJ())
			assert.Equal(t, data.billingEmail, companyInstance.BillingEmail())
			assert.Equal(t, data.billingPhone, companyInstance.BillingPhone())
			assert.Equal(t, data.purchaseEmail, companyInstance.PurchaseEmail())
			assert.Equal(t, data.purchasePhone, companyInstance.PurchasePhone())
		})
	}
}

func TestCompanyWithError(t *testing.T) {
	t.Parallel()

	// Arrange
	var testNameCases = []unitary.TestField[Config]{
		{
			Description: "name with length is short than allowed",
			Handler:         func(c *Config) { c.Name = strings.Repeat("a", companyNameMinLength-1) },
		},
		{
			Description: "name with length is greater than allowed",
			Handler:         func(c *Config) { c.Name = strings.Repeat("a", companyNameMaxLength+1) },
		},
	}
	for _, test := range testNameCases {
		t.Run(test.Description, func(t *testing.T) {

			companyInstance := newCompany()
			data := setCompany(test.Handler)

			// Assert
			assert.NotNil(t, companyInstance.SetName(data.name))

		})
	}

	var testCNPJCases = []unitary.TestField[Config]{
		{
			Description: "CNPJ with invalid format",
			Handler:         func(c *Config) { c.CNPJ = "00.000.000.0001.22" },
		},
		{
			Description: "CNPJ length is short than allowed",
			Handler:         func(c *Config) { c.CNPJ = "00.000.000/0001-2" },
		},
		{
			Description: "CNPJ length is greater than allowed",
			Handler:         func(c *Config) { c.CNPJ = "00.000.000/0001-222" },
		},
	}
	for _, test := range testCNPJCases {
		t.Run(test.Description, func(t *testing.T) {
			companyInstance := newCompany()
			data := setCompany(test.Handler)

			// Assert
			assert.NotNil(t, companyInstance.SetCNPJ(data.name))
		})
	}

	var testBillingContactCases = []unitary.TestField[Config]{
		{
			Description: "billing email invalid",
			Handler:         func(c *Config) { c.billingEmail = "    " },
		},
		{
			Description: "billing phone invalid",
			Handler:         func(c *Config) { c.billingPhone = "2188331456" },
		},
	}
	for _, test := range testBillingContactCases {
		t.Run(test.Description, func(t *testing.T) {
			companyInstance := newCompany()
			data := setCompany(test.Handler)

			// Assert
			assert.NotNil(t, companyInstance.SetBillingContact(data.billingEmail, data.billingPhone))
		})
	}

	var testPurchaseContactCases = []unitary.TestField[companyData]{
		{
			Description: "purchase email invalid",
			Handler:         func(c *Config) { cd.purchaseEmail = "invalid@mail" },
		},
		{
			Description: "purchase phone invalid",
			Handler:         func(c *Config) { cd.purchasePhone = "##########" },
		},
	}
	for _, test := range testPurchaseContactCases {
		t.Run(test.Description, func(t *testing.T) {
			companyInstance := newCompany()
			data := setCompany(test.Handler)

			// Assert
			assert.NotNil(t, companyInstance.SetBillingContact(data.purchaseEmail, data.purchasePhone))
		})
	}

	t.Run("purchase email invalid", func(t *testing.T) {
		// Arrage
		companyInstance := newCompany()
		var invalidEmail = "invalid@mail.."

		// Assert
		assert.NotNil(t, companyInstance.SetBillingEmail(invalidEmail))
		assert.NotNil(t, companyInstance.SetPurchaseEmail(invalidEmail))
	})

	t.Run("purchase phone invalid", func(t *testing.T) {
		companyInstance := newCompany()
		var invalidPhone = "ab-a123-1566"

		// Assert
		assert.NotNil(t, companyInstance.SetBillingPhone(invalidPhone))
		assert.NotNil(t, companyInstance.SetPurchasePhone(invalidPhone))
	})
}
