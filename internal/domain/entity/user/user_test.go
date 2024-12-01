package user

import (
	"strings"
	"testing"
	"time"

	validationerrors "stock-controll/internal/domain/services/error"
	
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var config = UserConfig{
	FirstName: unitary.Fake.Person().FirstName(),
	LastName:  unitary.Fake.Person().LastName(),
	CPF:       "186.887.817-89",
	Gender:    unitary.Fake.Person().Gender(),
	BirthDate: time.Now().AddDate(-55, 0, 0),
}

func TestNewNoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []unitary.TestField[UserConfig]{
		{
			Description: "user with minimum age",
			Handler:         func(uc *UserConfig) { uc.BirthDate = time.Now().AddDate(-18, 0, 0) },
		},
		{
			Description: "user with maximum age",
			Handler:         func(uc *UserConfig) { uc.BirthDate = time.Now().AddDate(-100, 0, 0) },
		},
		{
			Description: "Compound first name",
			Handler:         func(uc *UserConfig) { uc.FirstName = "Maria Clara" },
		},
		{
			Description: "first name with minimum length",
			Handler:         func(uc *UserConfig) { uc.FirstName = strings.Repeat("a", userNameMinLength) },
		},
		{
			Description: "first name with maximum length",
			Handler:         func(uc *UserConfig) { uc.FirstName = strings.Repeat("a", userNameMaxLength) },
		},
		{
			Description: "Compound last name",
			Handler:         func(uc *UserConfig) { uc.LastName = "Maria Clara" },
		},
		{
			Description: "last name with minimum length",
			Handler:         func(uc *UserConfig) { uc.LastName = strings.Repeat("a", userNameMinLength) },
		},
		{
			Description: "last name with maximum length",
			Handler:         func(uc *UserConfig) { uc.LastName = strings.Repeat("a", userNameMaxLength) },
		},
		{
			Description: "create a male gender",
			Handler:         func(uc *UserConfig) { uc.Gender = "Male" },
		},
		{
			Description: "create a female gender",
			Handler:         func(uc *UserConfig) { uc.Gender = "Female" },
		},
		{
			Description: "create gender with uppercases letters",
			Handler:         func(uc *UserConfig) { uc.Gender = "FEMALE" },
		},
		{
			Description: "create gender with lowercases letters",
			Handler:         func(uc *UserConfig) { uc.Gender = "male" },
		},
	}

	for _, test := range testsCases {
		t.Run(test.Description, func(t *testing.T) {
			datasCopy := config
			test.Handler(&datasCopy)

			// Act
			user, err := New(datasCopy)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, user)
			assert.NotEmpty(t, user.UUID())
			assert.Equal(t, strings.ToLower(datasCopy.FirstName), user.FirstName())
			assert.Equal(t, strings.ToLower(datasCopy.LastName), user.LastName())
			assert.Equal(t, strings.ToLower(datasCopy.Gender), user.Gender())
			assert.Equal(t, datasCopy.CPF, user.CPF())
			assert.Equal(t, datasCopy.BirthDate, user.BirthDate())
		})
	}
}

func TestNewWithError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []unitary.TestField[UserConfig]{
		{
			Description: "first name empty",
			Handler:         func(uc *UserConfig) { uc.FirstName = "            " },
		},
		{
			Description: "fisrt name with special characters",
			Handler:         func(uc *UserConfig) { uc.FirstName = "Maria @#" },
		},
		{
			Description: "first name with number",
			Handler:         func(uc *UserConfig) { uc.FirstName = "Otavio 123" },
		},
		{
			Description: "first name is filled with spaces characters",
			Handler:         func(uc *UserConfig) { uc.FirstName = "           " },
		},
		{
			Description: "first name length greater than minimum",
			Handler:         func(uc *UserConfig) { uc.FirstName = strings.Repeat("d", userNameMinLength-1) },
		},
		{
			Description: "first name length greater than maximum",
			Handler:         func(uc *UserConfig) { uc.FirstName = strings.Repeat("c", userNameMaxLength+1) },
		},
		{
			Description: "last name empty",
			Handler:         func(uc *UserConfig) { uc.LastName = "            " },
		},
		{
			Description: "last name with special characters",
			Handler:         func(uc *UserConfig) { uc.LastName = "Maria @#" },
		},
		{
			Description: "last name with number",
			Handler:         func(uc *UserConfig) { uc.LastName = "Otavio 123" },
		},
		{
			Description: "last name is filled with spaces characters",
			Handler:         func(uc *UserConfig) { uc.LastName = "           " },
		},
		{
			Description: "last name length greater than minimum",
			Handler:         func(uc *UserConfig) { uc.LastName = strings.Repeat("d", userNameMinLength-1) },
		},
		{
			Description: "last name length greater than maximum",
			Handler:         func(uc *UserConfig) { uc.LastName = strings.Repeat("c", userNameMaxLength+1) },
		},
		{
			Description: "cpf empty",
			Handler:         func(uc *UserConfig) { uc.CPF = "                " },
		},
		{
			Description: "cpf is short",
			Handler:         func(uc *UserConfig) { uc.CPF = "177.868-97" },
		},
		{
			Description: "cpf is long",
			Handler:         func(uc *UserConfig) { uc.CPF = "177.868.886-978" },
		},
		{
			Description: "cpf with invalid format",
			Handler:         func(uc *UserConfig) { uc.CPF = "167.868.886.97" },
		},
		{
			Description: "cpf with invalid checker digits",
			Handler:         func(uc *UserConfig) { uc.CPF = "123.456.789-01" },
		},
		{
			Description: "age is lower than 18 years old",
			Handler:         func(uc *UserConfig) { uc.BirthDate = time.Now().AddDate(-18, 0, +2) },
		},
		{
			Description: "age over maximum",
			Handler:         func(uc *UserConfig) { uc.BirthDate = time.Now().AddDate(-100, 0, -1) },
		},
		{

			Description: "the gender is empty",
			Handler: func(uc *UserConfig) { uc.Gender = ""},
		},

		{
			Description: "the gender field is filled with is empty characters",
			Handler: func(uc *UserConfig) { uc.Gender = "      "},
		},
		{
			Description: "the gender does not belong to the range of allowed", Handler: func(uc *UserConfig) { uc.Gender =  "Plant"},
		},
	}

	for _, test := range testsCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := config
			test.Handler(&dataCopy)

			// Act
			user, err := New(dataCopy)

			// Assert
			assert.Nil(t, user)
			require.Error(t, err)
			require.ErrorAs(t, err, &validationerrors.ValidationError{})
		})
	}
}
