package user

import (
	"strings"
	"testing"
	"time"

	"stock-controll/internal/domain/services/error/entity"
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
			Description: "Compound last name",
			Handler:         func(uc *UserConfig) { uc.LastName = "Maria Clara" },
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
			assert.Contains(t, user.FullName(), config.FirstName)
			assert.Contains(t, user.FullName(), config.LastName)
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
			Description: "invalid first name",
			Handler:         func(uc *UserConfig) { uc.FirstName = "            " },
		},
		{
			Description: "invalid last name ",
			Handler:         func(uc *UserConfig) { uc.LastName = "Maria @#" },
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
			require.ErrorAs(t, err, &entity.EntityError{})
		})
	}
}
