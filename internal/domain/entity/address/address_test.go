package address

import (
	"strings"
	"testing"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var addressData = address{
	street:     unitary.Fake.Address().StreetName(),
	city:       unitary.Fake.Address().City(),
	state:      unitary.Fake.Address().State(),
	complement: "apto",
	postalCode: "21000-220",
	number:     unitary.Fake.Address().Faker.Currency().Number(),
}

func Test_NewAddress_NoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[address]{
		{
			TestDescription: "street name length with compoust name",
			Handler:         func(a address) { a.street = "rua ibiapuera" },
		},
		{
			TestDescription: "street name length is equal than minimum allowed",
			Handler:         func(a address) { a.street = strings.Repeat("a", minStreetNameLength) },
		},
		{
			TestDescription: "street name length is equal than maximum allowed",
			Handler:         func(a address) { a.street = strings.Repeat("b", maxStreetNameLength) },
		},
		{
			TestDescription: "city name length with compoust name",
			Handler:         func(a address) { a.city = "rio de janeiro" },
		},
		{
			TestDescription: "city name length is equal than minimum allowed",
			Handler:         func(a address) { a.city = strings.Repeat("a", minCityNameLength) },
		},
		{
			TestDescription: "city name length is equal than maximum allowed",
			Handler:         func(a address) { a.city = strings.Repeat("a", maxCityNameLength) },
		},
		{
			TestDescription: "state name length with compoust name",
			Handler:         func(a address) { a.state = "Rio de Janeiro" },
		},
		{
			TestDescription: "state name length is equal than minimum allowed",
			Handler:         func(a address) { a.state = strings.Repeat("a", minStateNameLength) },
		},
		{
			TestDescription: "state name length is equal than maximum allowed",
			Handler:         func(a address) { a.state = strings.Repeat("a", maxStateNameLength) },
		},
		{
			TestDescription: "number name length is equal than minimum allowed",
			Handler:         func(a address) { a.number = minNumberHome },
		},
		{
			TestDescription: "number name length is equal than maximum allowed",
			Handler:         func(a address) { a.number = maxNumberHome },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {

			// Act
			addressInstance, err := NewAddress(
				addressData.street,
				addressData.city,
				addressData.state,
				addressData.postalCode,
				addressData.complement,
				addressData.number)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, addressInstance)
			assert.Equal(t, addressData.street, addressInstance.GetStreet())
			assert.Equal(t, addressData.city, addressInstance.GetCity())
			assert.Equal(t, addressData.state, addressInstance.GetState())
			assert.Equal(t, addressData.complement, addressInstance.GetComplement())
			assert.Equal(t, addressData.postalCode, addressInstance.GetPostalCode())
			assert.Equal(t, addressData.number, addressInstance.GetNumber())
		})
	}
}

func Test_NewAddress_WithError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[address]{
		{
			TestDescription: "street name length is short than allowed",
			Handler:         func(a address) { a.street = strings.Repeat("a", minStreetNameLength-1) },
		},
		{
			TestDescription: "street name length is greater than allowed",
			Handler:         func(a address) { a.street = strings.Repeat("b", maxStreetNameLength+1) },
		},
		{
			TestDescription: "street name contain special characters",
			Handler:         func(a address) { a.street = "dom pedro 2º" },
		},
		{
			TestDescription: "city name length is short than allowed",
			Handler:         func(a address) { a.city = strings.Repeat("a", minCityNameLength-1) },
		},
		{
			TestDescription: "city name length is greater than allowed",
			Handler:         func(a address) { a.city = strings.Repeat("b", maxCityNameLength+1) },
		},
		{
			TestDescription: "city name contain special characters",
			Handler:         func(a address) { a.city = "S@o Paulo" },
		},
		{
			TestDescription: "city name contain number",
			Handler:         func(a address) { a.city = "R1o Grande do Sul" },
		},
		{
			TestDescription: "state name length is short than allowed",
			Handler:         func(a address) { a.state = strings.Repeat("a", minStateNameLength-1) },
		},
		{
			TestDescription: "state name length is greater than allowed",
			Handler:         func(a address) { a.state = strings.Repeat("b", maxStateNameLength+1) },
		},
		{
			TestDescription: "state name contain special characters",
			Handler:         func(a address) { a.state = "S@o Paulo" },
		},
		{
			TestDescription: "state name contain number",
			Handler:         func(a address) { a.state = "R1o Grande do Sul" },
		},
		{
			TestDescription: "postal code with invalid format",
			Handler:         func(a address) { a.postalCode = "215300.300" },
		},
		{
			TestDescription: "home number is short than allowed",
			Handler:         func(a address) { a.state = strings.Repeat("a", minNumberHome-1) },
		},
		{
			TestDescription: "home number is greater than allowed",
			Handler:         func(a address) { a.state = strings.Repeat("b", maxNumberHome+1) },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			var address = addressData
			tt.Handler(address)

			// Act
			addressInstance, err := NewAddress(
				address.street,
				address.city,
				address.state,
				address.postalCode,
				address.complement,
				address.number)

			// Assert
			assert.Nil(t, addressInstance)
			assert.NotNil(t, err)
		})
	}
}
