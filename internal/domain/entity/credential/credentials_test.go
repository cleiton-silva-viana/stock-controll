package credential

import (
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/test/unitary"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type credentialData struct {
	password string
	userUUID string
}

var person = credentialData{
	password: "AbCd24@$Aefg&56dv",
	userUUID: "01928cee-b413-72f3-ad15-a3a297f0a114",
}

func TestNewCredentialPasswordNoError(t *testing.T) {
	// Arrange
	var testCases = []unitary.TestField[credentialData]{
		{
			Description: "Password with minimun characters valids",
			Handler:         func(cd *credentialData) { cd.password = "AbCd24@$A&56dv" },
		},
		{
			Description: "Password with maximun characters valids",
			Handler:         func(cd *credentialData) { cd.password = "Abcd$%@$AbCd24@$A65d24@$" },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			var data = person
			test.Handler(&data)

			// Act
			credential, err := NewCredential(data.userUUID, data.password)

			// Arrange
			assert.NoError(t, err)
			require.NotNil(t, credential)
			assert.Equal(t, person.userUUID, credential.GetUUID())
			assert.LessOrEqual(t, credential.GetCreatedAt(), time.Now())
			assert.LessOrEqual(t, credential.GetUpdatedAt(), time.Now())
			assert.GreaterOrEqual(t, credential.GetCreatedAt().Unix(), time.Now().Add(-time.Minute*5).Unix())
			assert.GreaterOrEqual(t, credential.GetUpdatedAt().Unix(), time.Now().Add(-time.Minute*5).Unix())
			require.NotEmpty(t, credential.GetPasswordHash()) // verificar se contem tamanho em bytes ao esperado
			require.NotEmpty(t, credential.GetPasswordSalt())
			assert.Len(t, credential.GetPasswordSalt(), passwordSaltLength)
			assert.GreaterOrEqual(t, credential.GetPasswordHash(), passwordSaltLength+passwordMinLength)
		})
	}
}

func TestNewCredentialPasswordWithError(t *testing.T) {
	var testCases = []unitary.TestField[credentialData]{
		{
			Description: "the password is empyt",
			Handler:         func(cd *credentialData) { cd.password = "" },
		},
		{
			Description: "the password length is equals to min length for valid password, howerer all chars are empty characters",
			Handler:         func(cd *credentialData) { cd.password = "        " },
		},
		{
			Description: "the password contains 1 digit less than the minimum acceptable",
			Handler:         func(cd *credentialData) { cd.password = "Ha$1a7#" },
		},
		{
			Description: "the password is invalid because it does not contains letesters",
			Handler:         func(cd *credentialData) { cd.password = "123@#$123" },
		},
		{
			Description: "the password is invalid for not contains lower cases letesters",
			Handler:         func(cd *credentialData) { cd.password = "ABC123@#$JK1" },
		},
		{
			Description: "The password is invalid for not contains upper cases",
			Handler:         func(cd *credentialData) { cd.password = "halo123%$baca" },
		},
		{
			Description: "the Password is invalid for not contains numbers",
			Handler:         func(cd *credentialData) { cd.password = "PaloAlto@#$" },
		},
		{
			Description: "the password is invalid for not contains special characters",
			Handler:         func(cd *credentialData) { cd.password = "PaloAlto123" },
		},
		{
			Description: "the password contains 1 digit more than the minimum acceptable",
			Handler:         func(cd *credentialData) { cd.password = "Abcd$%@$AbCd24@$A65d24@$cf$12" },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			var data = person
			test.Handler(&data)

			// Act
			credential, err := NewCredential(data.userUUID, data.password)

			// Assert
			assert.Nil(t, credential)
			require.Error(t, err)
			assert.ErrorAs(t, &validationerrors.ValidationError{}, err)
		})
	}
}
