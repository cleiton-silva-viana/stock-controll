package contact

import (
	"fmt"
	"strings"
	"testing"

	"stock-controll/internal/domain/services/uuid"
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var person = Contact{
	email: unitary.Fake.Person().Contact().Email,
	phone: "(61)3517-3828",
}

func TestNewNoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[Contact]{
		{
			Description: "the email contains minimum length allowed",
			Handler:     func(c *Contact) { c.email = "gl@g.com" },
		},
		{
			Description: "the email contains maximum length allowed",
			Handler:     func(c *Contact) { c.email = fmt.Sprintf("%s@gmail.com", strings.Repeat("a", emailMaxLength-10)) },
		},
		{
			Description: "create a landline",
			Handler:     func(c *Contact) { c.phone = "(21)4002-3214" },
		},
		{
			Description: "create a cellphone",
			Handler:     func(c *Contact) { c.phone = "(21)94852-3254" },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			var data = person
			test.Handler(&data)

			// Act
			contact, err := New(data.email, data.phone)

			// Arrange
			assert.NoError(t, err)
			require.NotNil(t, contact)
			assert.Equal(t, data.phone, contact.Phone())
			assert.Equal(t, data.email, contact.Email())
			assert.NoError(t, uuid.IsValid("contact_uuid", contact.UUID()))
		})
	}
}

func TestNewWithError(t *testing.T) {
	t.Parallel()

	// Arrange
	var testCases = []unitary.TestField[Contact]{
		{
			Description: "the email is empty",
			Handler:     func(c *Contact) { c.email = "       " },
		},
		{
			Description: "the email is short tha allowed",
			Handler:     func(c *Contact) { c.email = "a@b.org" },
		},
		{
			Description: "the email is greater than allowed",
			Handler:     func(c *Contact) { c.email = fmt.Sprintf("%s@email.com", strings.Repeat("a", 251)) },
		},
		{
			Description: "the email not contains arroba(@) character",
			Handler:     func(c *Contact) { c.email = "user.example.com" },
		},
		{
			Description: "the email not contains the domain (ex: outlook.com)",
			Handler:     func(c *Contact) { c.email = "user.example@" },
		},
		{
			Description: "the email not contais user name",
			Handler:     func(c *Contact) { c.email = "@example.com" },
		},
		{
			Description: "the email contains invalid domain format",
			Handler:     func(c *Contact) { c.email = "@example..com" },
		},
		{
			Description: "the email contain empty character",
			Handler:     func(c *Contact) { c.email = "user @example.com" },
		},
		{
			Description: "the email not contai TDL domain",
			Handler:     func(c *Contact) { c.email = "user@example" },
		},
		{
			Description: "the email initializing with characters specials",
			Handler:     func(c *Contact) { c.email = "@user@example.com" },
		},
		{
			Description: "the email ending with special characters",
			Handler:     func(c *Contact) { c.email = "user@example.com." },
		},
		{
			Description: "the email contains invalid format",
			Handler:     func(c *Contact) { c.email = "user@example,com." },
		},
		{

			Description: "phone is empty",
			Handler:     func(c *Contact) { c.phone = "" },
		},
		{

			Description: "phone is filled with spaces ",
			Handler:     func(c *Contact) { c.phone = "              " },
		},
		{

			Description: "phone with invalid format",
			Handler:     func(c *Contact) { c.phone = "219269-82225" },
		},
		{

			Description: "phone contains letesters",
			Handler:     func(c *Contact) { c.phone = "(21)4002-8922a" },
		},
		{

			Description: "phone contains special characters",
			Handler:     func(c *Contact) { c.phone = "(21)4002-8922#" },
		},
		{

			Description: "phone length is less than minimum allowed",
			Handler:     func(c *Contact) { c.phone = "(21)4002-892" },
		},
		{

			Description: "phone length is long than maximum allowed",
			Handler:     func(c *Contact) { c.phone = "(21)40028-23842" },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			var data = person
			test.Handler(&data)

			// Act
			contact, err := New(data.email, data.phone)

			// Assert
			assert.Nil(t, contact)
			assert.Error(t, err)
		})
	}
}

func TestStateConsistencyAfterInvalidSet(t *testing.T) {
	testCases := []unitary.Consistence{
		{
			T:            t,
			Description:  "test consistence of email",
			Getter:       func() interface{} { return person.Email },
			Setter:       func(value interface{}) error { return person.SetEmail(value.(string)) },
			InvalidValue: "emailWithoutDomain@",
		},
		{
			T:            t,
			Description:  "test consistence of phone",
			Getter:       func() interface{} { return person.Phone },
			Setter:       func(value interface{}) error { return person.SetPhone(value.(string)) },
			InvalidValue: "215554228578",
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {

			unitary.ConsistenceTest(test)
		})
	}
}
