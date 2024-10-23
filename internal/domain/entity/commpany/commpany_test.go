package entity

import (
	addressEntity "stock-controll/internal/domain/entity/address"
	"stock-controll/test/unitary"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type companyData struct {
	name          string
	cnpj          string
	billingEmail  string
	billingPhone  string
	purchaseEmail string
	purchasePhone string
}

var hyundai = companyData{
	name:          unitary.Fake.Company().Name(),
	cnpj:          "46.318.856/0001-00",
	billingEmail:  unitary.Fake.Internet().CompanyEmail(),
	billingPhone:  "(49)2524-2218",
	purchaseEmail: unitary.Fake.Internet().CompanyEmail(),
	purchasePhone: "(74)3017-4666",
}

func newCompany() company {
	return company{}
}


var addr, _ = addressEntity.NewAddress("rua canudos", "rio de janeiro", "rio de janeiro", "21500-300", "apartamaento 202", 252)

func setCompany(handler func(cd companyData) ) companyData {
	var data = hyundai
	handler(data)
	return data
}

func Test_company_NoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[companyData]{
		{
			TestDescription: "name length is minimum allowed",
			Handler:         func(cd companyData) { cd.name = strings.Repeat("a", companyNameMinLength) },
		},
		{
			TestDescription: "name length is maximum allowed",
			Handler:         func(cd companyData) { cd.name = strings.Repeat("b", companyNameMaxLength) },
		},
		{
			TestDescription: "name contain special characters",
			Handler:         func(cd companyData) { cd.name = "b&b Hammer" },
		},
		{
			TestDescription: "name contain number",
			Handler:         func(cd companyData) { cd.name = "1st price" },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			var companyInstance = company{}
			var data = hyundai
			tt.Handler(data)

			// Assert
			assert.Nil(t, companyInstance.SetName(data.name))
			assert.Nil(t, companyInstance.SetCNPJ(data.cnpj))
			assert.Nil(t, companyInstance.SetBillingEmail(data.billingEmail))
			assert.Nil(t, companyInstance.SetBillingPhone(data.billingPhone))
			assert.Nil(t, companyInstance.SetPurchaseEmail(data.purchaseEmail))
			assert.Nil(t, companyInstance.SetPurchasePhone(data.purchasePhone))
			assert.Nil(t, companyInstance.SetBillingContact(data.billingEmail, data.billingPhone))
			assert.Nil(t, companyInstance.SetPurchaseContact(data.purchaseEmail, data.purchasePhone))
			assert.Equal(t, data.name, companyInstance.GetName())
			assert.Equal(t, data.cnpj, companyInstance.GetCNPJ())
			assert.Equal(t, data.billingEmail, companyInstance.GetBillingEmail())
			assert.Equal(t, data.billingPhone, companyInstance.GetBillingPhone())
			assert.Equal(t, data.purchaseEmail, companyInstance.GetPurchaseEmail())
			assert.Equal(t, data.purchasePhone, companyInstance.GetPurchasePhone())
		})
	}
}

func Test_company_WithError(t *testing.T) {
	t.Parallel()

	// Arrange
	var testNameCases = []unitary.TestField[companyData]{
		{
			TestDescription: "name with length is short than allowed",
			Handler:         func(cd companyData) { cd.name = strings.Repeat("a", companyNameMinLength-1) },
		},
		{
			TestDescription: "name with length is greater than allowed",
			Handler:         func(cd companyData) { cd.name = strings.Repeat("a", companyNameMaxLength+1) },
		},
	}
	for _, tt := range testNameCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			companyInstance := newCompany()
			data := setCompany(tt.Handler)

			// Assert
			assert.NotNil(t, companyInstance.SetName(data.name))

		})
	}

	var testCNPJCases = []unitary.TestField[companyData]{
		{
			TestDescription: "CNPJ with invalid format",
			Handler:         func(cd companyData) { cd.cnpj = "00.000.000.0001.22" },
		},
		{
			TestDescription: "CNPJ length is short than allowed",
			Handler:         func(cd companyData) { cd.cnpj = "00.000.000/0001-2" },
		},
		{
			TestDescription: "CNPJ length is greater than allowed",
			Handler:         func(cd companyData) { cd.cnpj = "00.000.000/0001-222" },
		},
	}
	for _, tt := range testCNPJCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			companyInstance := newCompany()
			data := setCompany(tt.Handler)

			// Assert
			assert.NotNil(t, companyInstance.SetCNPJ(data.name))
		})
	}

	var testBillingContactCases = []unitary.TestField[companyData]{
		{
			TestDescription: "billing email invalid",
			Handler:         func(cd companyData) { cd.billingEmail = "    " },
		},
		{
			TestDescription: "billing phone invalid",
			Handler:         func(cd companyData) { cd.billingPhone = "2188331456" },
		},
	}
	for _, tt := range testBillingContactCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			companyInstance := newCompany()
			data := setCompany(tt.Handler)

			// Assert
			assert.NotNil(t, companyInstance.SetBillingContact(data.billingEmail, data.billingPhone))
		})
	}

	var testPurchaseContactCases = []unitary.TestField[companyData]{
		{
			TestDescription: "purchase email invalid",
			Handler:         func(cd companyData) { cd.purchaseEmail = "invalid@mail" },
		},
		{
			TestDescription: "purchase phone invalid",
			Handler:         func(cd companyData) { cd.purchasePhone = "##########" },
		},
	}
	for _, tt := range testPurchaseContactCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			companyInstance := newCompany()
			data := setCompany(tt.Handler)

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
