package contact

import (
	"fmt"
	"stock-controll/internal/domain/entity/common"
	"stock-controll/test/unitary"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var person = contact{
	email: unitary.Fake.Person().Contact().Email, 
	phone: "(61)3517-3828",
}

func Test_NewContact_NoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[contact]{
		{
			TestDescription: "the email contains minimum length allowed",
			Handler:         func(c contact) { c.email = "gl@g.com" },
		},
		{
			TestDescription: "the email contains maximum length allowed",
			Handler: func(c contact) { c.email = fmt.Sprintf("%s@gmail.com", strings.Repeat("a", emailMaxLength-10))},
		},
		{
			TestDescription: "create a landline",
			Handler:         func(c contact) { c.phone = "(21)4002-3214" },
		},
		{
			TestDescription: "create a cellphone",
			Handler:  func(c contact) { c.phone = "(21)94852-3254" },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			var data = person
			tt.Handler(data)

			// Act
			contact, err := NewContact(data.email, data.phone)

			// Arrange
			assert.Nil(t, err)
			require.NotNil(t, contact)
			assert.Equal(t, data.phone, contact.GetPhone())
			assert.Equal(t, data.email, contact.GetEmail())
			assert.True(t, common.IsValidUUUID(contact.GetUUID()))
		})
	}
}

func Test_NewContact_WithError(t *testing.T) {
	t.Parallel()

	// Arrange
	var testCases = []unitary.TestField[contact]{
			{
				TestDescription: "the email is empty",
				Handler:   func(c contact) { c.email = "       " },
			},
			{
				TestDescription: "the email is short tha allowed",
				Handler:           func(c contact) { c.email = "a@b.org" } ,
			},
			{
				TestDescription: "the email is greater than allowed",
				Handler:          func(c contact) { c.email = fmt.Sprintf("%s@email.com", strings.Repeat("a", 251)) },
			},
			{
				TestDescription: "the email not contains arroba(@) character",
				Handler:          func(c contact) { c.email = "user.example.com"},
			},
			{
				TestDescription: "the email not contains the domain (ex: outlook.com)",
				Handler:         func(c contact) { c.email = "user.example@"},
			},
			{
				TestDescription: "the email not contais user name",
				Handler:         func(c contact) { c.email = "@example.com"},
			},
			{
				TestDescription: "the email contains invalid domain format",
				Handler:          func(c contact) { c.email = "@example..com"} ,
			},
			{
				TestDescription: "the email contain empty character",
				Handler:         func(c contact) { c.email = "user @example.com"}  ,
			},
			{
				TestDescription: "the email not contai TDL domain",
				Handler:          func(c contact) { c.email = "user@example"} ,
			},
			{
				TestDescription: "the email initializing with characters specials",
				Handler:          func(c contact) { c.email = "@user@example.com" } ,
			},
			{
				TestDescription: "the email ending with special characters",
				Handler:         func(c contact) { c.email = "user@example.com."}  ,
			},
			{
				TestDescription: "the email contains invalid format",
				Handler:          func(c contact) { c.email = "user@example,com." } ,
			},
			{
				
				TestDescription: "phone is empty",
				Handler:  func(c contact) { c.phone = "" },
			},
			{
				
				TestDescription: "phone is filled with spaces ",
				Handler:  func(c contact) { c.phone = "              "  },
			},
			{
				
				TestDescription: "phone with invalid format",
				Handler:  func(c contact) { c.phone = "219269-82225" },
			},
			{
				
				TestDescription: "phone contains letters",
				Handler:  func(c contact) { c.phone = "(21)4002-8922a" },
			},
			{
				
				TestDescription: "phone contains special characters",
				Handler:  func(c contact) { c.phone = "(21)4002-8922#" },
			},
			{
				
				TestDescription: "phone length is less than minimum allowed",
				Handler:  func(c contact) { c.phone = "(21)4002-892"},
			},
			{
				
				TestDescription: "phone length is long than maximum allowed",
				Handler:  func(c contact) { c.phone = "(21)40028-23842"},
			},
		}

		for _, tt := range testCases {
			t.Run(tt.TestDescription, func(t *testing.T) {
				var data = person
				tt.Handler(data)

				// Act
				contact, err := NewContact(data.email, data.phone)

				// Assert
				assert.Nil(t, contact)
				assert.NotNil(t, err)
			})
		}
}
