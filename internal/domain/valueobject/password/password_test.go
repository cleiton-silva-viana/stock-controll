package password

import (
	"stock-controll/test/unitary"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type pass struct {
	password string
}

func TestNewPassowrdNoError(t *testing.T) {
	// Arrange
	var testCases = []unitary.TestField[pass]{
		{
			Description: "Password with max length allowed",
			Handler:     func(p *pass) { p.password = "AbCd24@$A&56dv" },
		},
		{
			Description: "Password with min length allowed",
			Handler:     func(p *pass) { p.password = "Abcd$%@$AbCd24@$A65d24@$" },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			var p pass
			test.Handler(&p)

			// Act
			result, err := New(p.password)

			// Arrange
			assert.NoError(t, err)
			require.NotNil(t, result)
			require.NotEmpty(t, result.Hash())
			require.NotEmpty(t, result.Salt())
			assert.Len(t, result.Salt(), passwordSaltLength)
			assert.GreaterOrEqual(t, len(result.Hash()), passwordSaltLength+passwordMinLength)
		})
	}
}

func TestNewPasswordWithError(t *testing.T) {
	var testCases = []unitary.TestField[pass]{
		{
			Description: "the password is empyt",
			Handler:     func(p *pass) { p.password = "" },
		},
		{
			Description: "the password length is equals to min length for valid password, howerer all chars are empty characters",
			Handler:     func(p *pass) { p.password = "        " },
		},
		{
			Description: "the password contains 1 digit less than the minimum acceptable",
			Handler:     func(p *pass) { p.password = "Ha$1a7#" },
		},
		{
			Description: "the password is invalid because it does not contains letesters",
			Handler:     func(p *pass) { p.password = "123@#$123" },
		},
		{
			Description: "the password is invalid for not contains lower cases letesters",
			Handler:     func(p *pass) { p.password = "ABC123@#$JK1" },
		},
		{
			Description: "The password is invalid for not contains upper cases",
			Handler:     func(p *pass) { p.password = "halo123%$baca" },
		},
		{
			Description: "the Password is invalid for not contains numbers",
			Handler:     func(p *pass) { p.password = "PaloAlto@#$" },
		},
		{
			Description: "the password is invalid for not contains special characters",
			Handler:     func(p *pass) { p.password = "PaloAlto123" },
		},
		{
			Description: "the password contains 1 digit more than the minimum acceptable",
			Handler:     func(p *pass) { p.password = "Abcd$%@$AbCd24@$A65d24@$cf$12" },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			var p pass
			test.Handler(&p)

			// Act
			result, err := New(p.password)

			// Assert
			assert.Nil(t, result)
			require.Error(t, err)
		})
	}
}
