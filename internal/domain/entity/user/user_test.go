package user

import (
	"strings"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var fake = faker.New()

type userDatas struct {
	firstName string
	lastName  string
	cpf       string
	gender    string
	birthDate time.Time
}

type test struct {
	testDescription string
	userDatas
}

func Test_NewUser_NoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []test{
		{
			testDescription: "user valid",
			userDatas: userDatas{
				firstName: fake.Person().FirstName(),
				lastName:  fake.Person().LastName(),
				cpf:       "186.887.817-89",
				gender:    fake.Person().Gender(),
				birthDate: time.Now().AddDate(-55, 0, 0),
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			user, err := NewUserBuilder().
				SetCPF(tt.cpf).
				SetName(tt.firstName, tt.lastName).
				SetGender(tt.gender).
				SetBirthDate(tt.birthDate).
				Build()

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, user)
			require.Implements(t, (*IUser)(nil), user)
			assert.NotEmpty(t, user.GetUUID())
			assert.Equal(t, strings.ToLower(tt.firstName), strings.ToLower(user.GetFirstName()))
			assert.Equal(t, strings.ToLower(tt.lastName), strings.ToLower(user.GetLastName()))
			assert.Equal(t, strings.ToLower(tt.gender), strings.ToLower(user.GetGender()))
			assert.Equal(t, tt.birthDate, user.GetBirthDate())
		})
	}
}

func Test_NewUser_WithError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []test{
		{
			testDescription: "first name empty",
			userDatas: userDatas{
				firstName: "    ",
				lastName:  fake.Person().LastName(),
				cpf:       "186.817.817-89",
				gender:    fake.Person().Gender(),
				birthDate: time.Now().AddDate(-30, 0, 0),
			},
		},
		{
			testDescription: "invalid gender",
			userDatas: userDatas{
				firstName: fake.Person().FirstName(),
				lastName:  fake.Person().LastName(),
				cpf:       "186.817.817-89",
				gender:    "unknown",
				birthDate: time.Now().AddDate(-30, 0, 0),
			},
		},
		{
			testDescription: "invalid CPF",
			userDatas: userDatas{
				firstName: fake.Person().FirstName(),
				lastName:  fake.Person().LastName(),
				cpf:       "186.817.817.89",
				gender:    "unknown",
				birthDate: time.Now().AddDate(-30, 0, 0),
			},
		},

		{
			testDescription: "first name & last name with invalid characters ",
			userDatas: userDatas{
				firstName: "John123",
				lastName:  "kratos@",
				cpf:       "186.817.817-89",
				gender:    fake.Person().Gender(),
				birthDate: time.Now().AddDate(-30, 0, 0),
			},
		},
		{
			testDescription: "gender unknow and age is lower than 18 years old",
			userDatas: userDatas{
				firstName: fake.Person().FirstName(),
				lastName:  fake.Person().LastName(),
				cpf:       "186.817.817-89",
				gender:    "zombie",
				birthDate: time.Now().AddDate(-17, 0, 0),
			},
		},
		{
			testDescription: "all fields are invalid",
			userDatas: userDatas{
				firstName: "123###",
				lastName:  "",
				cpf:       "186.817.817-89",
				gender:    "zombie",
				birthDate: time.Now().AddDate(-10, 0, 0),
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			user, err := NewUserBuilder().
				SetCPF(tt.cpf).
				SetName(tt.firstName, tt.lastName).
				SetGender(tt.gender).
				SetBirthDate(tt.birthDate).
				Build()

			// Assert
			assert.Nil(t, user)
			require.NotNil(t, err)
			// require.ErrorAs(t, err, &entityError)
		})
	}
}

var person = userDatas{
	firstName: fake.Person().FirstName(),
	lastName:  fake.Person().LastName(),
	gender:    fake.Person().Gender(),
	birthDate: time.Now().AddDate(-25, 0, 0),
	cpf:       "177.777.897-98",
}

type fieldTest struct {
	testDescription string
	field           string
}

func Test_NewUser_BirthDate_NoError(t *testing.T) {
	var testsCases = []struct {
		testDescription string
		field           time.Time
	}{
		{
			testDescription: "user with minimum age",
			field:           time.Now().AddDate(-18, 0, 0),
		},
		{
			testDescription: "user with maximum age",
			field:           time.Now().AddDate(-100, 0, 0),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			user, err := NewUser(person.firstName, person.lastName, person.cpf, person.gender, tt.field)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, user)
			assert.Equal(t, tt.field, user.GetBirthDate())
		})
	}
}

func Test_NewUser_BirthDate_WithError(t *testing.T) {
	var testsCases = []struct {
		testDescription string
		field           time.Time
	}{
		{
			testDescription: "age under minimum",
			field:           time.Now().AddDate(-17, 0, 0),
		},
		{
			testDescription: "age over maximum",
			field:           time.Now().AddDate(-100, 0, -1),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {
			person.birthDate = tt.field

			// Act
			user, err := NewUser(person.firstName, person.lastName, person.cpf, person.gender, person.birthDate)

			// Assert
			assert.Nil(t, user)
			assert.NotNil(t, err)
		})
	}
}

func Test_NewUser_Name_NoError(t *testing.T) {
	// Arrange
	var testsCases = []fieldTest{
		{
			testDescription: "Compound first name",
			field:           "Maria Clara",
		},
		{
			testDescription: "name with minimum length",
			field:           strings.Repeat("a", userNameMinLength),
		},
		{
			testDescription: "name with maximum length",
			field:           strings.Repeat("a", userNameMaxLength),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {
			person.firstName = tt.field

			// Act
			user, err := NewUser(person.firstName, person.lastName, person.cpf, person.gender, person.birthDate)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, user)
			assert.Contains(t, user.GetFullName(), tt.field)
		})
	}
}

func Test_NewUser_Name_WithError(t *testing.T) {
	// Arrange
	var testsCases = []fieldTest{
		{
			testDescription: "name with special characters",
			field:           "Maria @#",
		},
		{
			testDescription: "name with number",
			field:           "Otavio 123",
		},
		{
			testDescription: "name is filled with spaces characters",
			field:           "           ",
		},
		{
			testDescription: "name length greater than minimum",
			field:           strings.Repeat("d", userNameMinLength-1),
		},
		{
			testDescription: "name length greater than maximum",
			field:           strings.Repeat("c", userNameMaxLength+1),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {
			person.firstName = tt.field

			// Act
			user, err := NewUser(person.firstName, person.lastName, person.cpf, person.gender, person.birthDate)

			// Assert
			assert.Nil(t, user)
			assert.NotNil(t, err)
		})
	}
}

func Test_NewUser_CPF_NoError(t *testing.T) {
	// Arrange
	person.cpf = "435.116.860-91"

	// Act
	user, err := NewUser(person.firstName, person.lastName, person.cpf, person.gender, person.birthDate)

	// Assert
	assert.Nil(t, err)
	require.NotNil(t, user)
	assert.Equal(t, user.GetCPF(), person.cpf)
}

func Test_NewUser_CPF_WithError(t *testing.T) {
	// Arrange
	testsCases := []fieldTest{
		{
			testDescription: "cpf empty",
			field:           "                ",
		},
		{
			testDescription: "cpf is short",
			field:           "177.868-97",
		},
		{
			testDescription: "cpf is long",
			field:           "177.868.886-978",
		},
		{
			testDescription: "cpf with invalid format",
			field:           "167.868.886.97",
		},
		{
			testDescription: "cpf with invalid checker digits",
			field:           "123.456.789-01",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {
			person.cpf = tt.field

			// Act
			user, err := NewUser(person.firstName, person.lastName, person.cpf, person.gender, person.birthDate)

			// Assert
			assert.Nil(t, user)
			assert.NotNil(t, err)
		})
	}
}

func Test_NewUser_Gender_NoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []fieldTest{
		{
			testDescription: "create a male gender",
			field:           "Male",
		},
		{
			testDescription: "create a female gender",
			field:           "Female",
		},
		{
			testDescription: "create gender with uppercases letters",
			field:           "FEMALE",
		},
		{
			testDescription: "create gender with lowercases letters",
			field:           "male",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {
			person.gender = tt.field

			// Act
			user, err := NewUser(person.firstName, person.lastName, person.cpf, person.gender, person.birthDate)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, user)
			assert.Equal(t, strings.ToLower(tt.field), user.GetGender())
		})
	}
}

func Test_NewUser_Gender_WithError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []fieldTest{
		{

			testDescription: "the gender is empty",
			field:           "",
		},

		{
			testDescription: "the gender field is filled with is empty characters",
			field:           "      ",
		},
		{
			testDescription: "the gender does not belong to the range of allowed values",
			field:           "Plant",
		},
	}
	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {
			person.gender = tt.field

			// Act
			user, err := NewUser(person.firstName, person.lastName, person.cpf, person.gender, person.birthDate)

			// Arrange
			assert.Nil(t, user)
			assert.NotNil(t, err)
		})
	}
}
