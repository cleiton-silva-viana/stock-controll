package address

import (
	"strings"
	"testing"

	"stock-controll/internal/domain/services/uuid"
	
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var config = Config{
	Street:     unitary.Fake.Address().StreetName(),
	City:       unitary.Fake.Address().City(),
	State:      unitary.Fake.Address().State(),
	Complement: "apto",
	PostalCode: "21000-220",
	Number:     unitary.Fake.Address().Faker.Currency().Number(),
}

func TestNewAddressNoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[Config]{
		{
			Description: "street name length with compoust name",
			Handler:     func(c *Config) { c.Street = "rua ibiapuera" },
		},
		{
			Description: "street name length is equal than minimum allowed",
			Handler:     func(c *Config) { c.Street = strings.Repeat("a", minStreetNameLength) },
		},
		{
			Description: "street name length is equal than maximum allowed",
			Handler:     func(c *Config) { c.Street = strings.Repeat("b", maxStreetNameLength) },
		},
		{
			Description: "city name length with compoust name",
			Handler:     func(c *Config) { c.City = "rio de janeiro" },
		},
		{
			Description: "city name length is equal than minimum allowed",
			Handler:     func(c *Config) { c.City = strings.Repeat("a", minCityNameLength) },
		},
		{
			Description: "city name length is equal than maximum allowed",
			Handler:     func(c *Config) { c.City = strings.Repeat("a", maxCityNameLength) },
		},
		{
			Description: "state name length with compoust name",
			Handler:     func(c *Config) { c.State = "Rio de Janeiro" },
		},
		{
			Description: "state name length is equal than minimum allowed",
			Handler:     func(c *Config) { c.State = strings.Repeat("a", minStateNameLength) },
		},
		{
			Description: "state name length is equal than maximum allowed",
			Handler:     func(c *Config) { c.State = strings.Repeat("a", maxStateNameLength) },
		},
		{
			Description: "number name length is equal than minimum allowed",
			Handler:     func(c *Config) { c.Number = minNumberHome },
		},
		{
			Description: "number name length is equal than maximum allowed",
			Handler:     func(c *Config) { c.Number = maxNumberHome },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			copy := config
			test.Handler(&copy)

			// Act
			addressInstance, err := New(copy)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, addressInstance)
			assert.Equal(t, copy.Street, addressInstance.Street())
			assert.Equal(t, copy.City, addressInstance.City())
			assert.Equal(t, copy.State, addressInstance.State())
			assert.Equal(t, copy.Complement, addressInstance.Complement())
			assert.Equal(t, copy.PostalCode, addressInstance.PostalCode())
			assert.Equal(t, copy.Number, addressInstance.Number())
		})
	}
}

func TestNewAddressWithError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[Config]{
		{
			Description: "street name length is short than allowed",
			Handler:     func(c *Config) { c.Street = strings.Repeat("a", minStreetNameLength-1) },
		},
		{
			Description: "street name length is greater than allowed",
			Handler:     func(c *Config) { c.Street = strings.Repeat("b", maxStreetNameLength+1) },
		},
		{
			Description: "street name contain special characters",
			Handler:     func(c *Config) { c.Street = "dom pedro 2º" },
		},
		{
			Description: "city name length is short than allowed",
			Handler:     func(c *Config) { c.City = strings.Repeat("a", minCityNameLength-1) },
		},
		{
			Description: "city name length is greater than allowed",
			Handler:     func(c *Config) { c.City = strings.Repeat("b", maxCityNameLength+1) },
		},
		{
			Description: "city name contain special characters",
			Handler:     func(c *Config) { c.City = "S@o Paulo" },
		},
		{
			Description: "city name contain number",
			Handler:     func(c *Config) { c.City = "R1o Grande do Sul" },
		},
		{
			Description: "state name length is short than allowed",
			Handler:     func(c *Config) { c.State = strings.Repeat("a", minStateNameLength-1) },
		},
		{
			Description: "state name length is greater than allowed",
			Handler:     func(c *Config) { c.State = strings.Repeat("b", maxStateNameLength+1) },
		},
		{
			Description: "state name contain special characters",
			Handler:     func(c *Config) { c.State = "S@o Paulo" },
		},
		{
			Description: "state name contain number",
			Handler:     func(c *Config) { c.State = "R1o Grande do Sul" },
		},
		{
			Description: "postal code with invalid format",
			Handler:     func(c *Config) { c.PostalCode = "215300.300" },
		},
		{
			Description: "home number is short than allowed",
			Handler:     func(c *Config) { c.State = strings.Repeat("a", minNumberHome-1) },
		},
		{
			Description: "home number is greater than allowed",
			Handler:     func(c *Config) { c.State = strings.Repeat("b", maxNumberHome+1) },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			copy := config
			test.Handler(&copy)

			// Act
			addressInstance, err := New(copy)

			// Assert
			assert.Nil(t, addressInstance)
			assert.NotNil(t, err)
		})
	}
}

var addressInstance = Address{
	uuid:       uuid.New(),
	street:     unitary.Fake.Address().StreetName(),
	city:       unitary.Fake.Address().City(),
	state:      unitary.Fake.Address().State(),
	complement: "apto",
	postalCode: "21000-220",
	number:     unitary.Fake.Address().Faker.Currency().Number(),
}

func TestStateConsistencyAfterInvalidSet(t *testing.T) {
	copy := addressInstance

	// Arrange
	testCases := []unitary.Consistence{
		{
			T:            t,
			Description:  "test consistence of street name",
			Getter:       func() interface{} { return copy.street },
			Setter:       func(value interface{}) error { return copy.SetStreet(value.(string)) },
			InvalidValue: "Street Z$r0",
		},
		{
			T:            t,
			Description:  "test consistence of home number",
			Getter:       func() interface{} { return copy.City },
			Setter:       func(value interface{}) error { return copy.SetNumber(value.(int)) },
			InvalidValue: -1,
		},
		{
			T:            t,
			Description:  "test consistence of complement",
			Getter:       func() interface{} { return copy.City },
			Setter:       func(value interface{}) error { return copy.SetComplement(value.(string)) },
			InvalidValue: strings.Repeat("a", maxComplementLength+1),
		},
		{
			T:            t,
			Description:  "test consistence of city name",
			Getter:       func() interface{} { return copy.City },
			Setter:       func(value interface{}) error { return copy.SetCity(value.(string)) },
			InvalidValue: "#%¨#$5",
		},
		{
			T:            t,
			Description:  "test consistence of state",
			Getter:       func() interface{} { return copy.City },
			Setter:       func(value interface{}) error { return copy.SetCity(value.(string)) },
			InvalidValue: "_______",
		},
		{
			T:            t,
			Description:  "test consistence of postal code",
			Getter:       func() interface{} { return copy.City },
			Setter:       func(value interface{}) error { return copy.SetCity(value.(string)) },
			InvalidValue: "21550-300",
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {

			// Act & Assert
			unitary.ConsistenceTest(test)
		})
	}
}
