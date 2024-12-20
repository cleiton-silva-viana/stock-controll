package address

import (
	"strings"
	"testing"

	"stock-controll/internal/domain/valueobject/uuid"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Setup() (*Config, *Address) {
	return &Config{
			Street:     unitary.Fake.Address().StreetName(),
			City:       unitary.Fake.Address().City(),
			State:      unitary.Fake.Address().State(),
			Complement: "apto",
			PostalCode: "21000-220",
			Number:     unitary.Fake.Address().Faker.Currency().Number(),
		}, &Address{
			UUID:       *uuid.New(),
			street:     unitary.Fake.Address().StreetName(),
			city:       unitary.Fake.Address().City(),
			state:      unitary.Fake.Address().State(),
			complement: "apto",
			postalCode: "21000-220",
			number:     unitary.Fake.Address().Faker.Currency().Number(),
		}
}

func TestNewAddressNoError(t *testing.T) {
	t.Parallel()

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
			// Arrange
			c, _ := Setup()
			test.Handler(c)

			// Act
			addressInstance, err := New(*c)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, addressInstance)
			assert.Equal(t, c.Street, addressInstance.Street())
			assert.Equal(t, c.City, addressInstance.City())
			assert.Equal(t, c.State, addressInstance.State())
			assert.Equal(t, c.Complement, addressInstance.Complement())
			assert.Equal(t, c.PostalCode, addressInstance.PostalCode())
			assert.Equal(t, c.Number, addressInstance.Number())
		})
	}
}

func TestNewAddressWithError(t *testing.T) {
	t.Parallel()

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
			// Arrange
			c, _ := Setup()
			test.Handler(c)

			// Act
			result, err := New(*c)

			// Assert
			assert.Nil(t, result)
			assert.NotNil(t, err)
		})
	}
}
