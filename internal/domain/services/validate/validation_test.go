package validate

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type test struct {
	description string
	value       string
	length      int
	mode        checkMode
}

func TestValidateNoError(t *testing.T) {
	// Arrange
	const (
		fieldName  = "password"
		fieldValue = "123456"
	)

	// Act
	err := New(fieldName, fieldValue)

	// Assert
	assert.NoError(t, err)
}

// Refatorar !!!
// Verificar se a quantidade de erros esperados está sendo retornada
// Verificar se os erros estão sendo retornados
func TestValidateWithError(t *testing.T) {
	// Arrange
	const fieldValue = "123456"

	// Act
	err := New("field", fieldValue, CheckNumbers(Disallow))

	// Assert
	require.Error(t, err)
}

func TestIsBlankNoError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "the field is filled with digit",
			value:       " 1 ",
		},
		{
			description: "the field is filled with letter",
			value:       " a ",
		},
		{
			description: "the field is filled with special character",
			value:       " & ",
		},
		{
			description: "the field is filled",
			value:       " Roma@123 ",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("name", tt.value)

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestIsBlankWithError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "Must have error, the field is empty",
			value:       "",
		},
		{
			description: "Must have error, the field is filled with empty characters",
			value:       "      ",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value)

			// Assert
			require.NotNil(t, err)
		})
	}
}

func TestIsLengthEqualToNoError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "field filled with spaces characters",
			value:       strings.Repeat(" ", 10),
			length:      10,
		},
		{
			description: "field filled with digits",
			value:       strings.Repeat("1", 10),
			length:      10,
		},
		{
			description: "field filled with letters",
			value:       strings.Repeat("a", 10),
			length:      10,
		},
		{
			description: "field filled with special characters",
			value:       strings.Repeat("@", 10),
			length:      10,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value, IsLengthEqualTo(tt.length))

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestIsLengthEqualToWithError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "field filled with 1 character shorter than the expected length",
			value:       strings.Repeat(" ", 9),
			length:      10,
		},
		{
			description: "field filled with 1 character more than the expected length",
			value:       strings.Repeat(" ", 11),
			length:      10,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value, IsLengthEqualTo(tt.length))

			// Assert
			require.NotNil(t, err)
		})
	}
}

func TestCheckSpecialCharsNoError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string with letters, spaces and number, check mode set to 'disallow'",
			value:       "orange 12",
			mode:        Disallow,
		},
		{
			description: "string with special chars - check mode set to 'required'",
			value:       "H@la",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value, CheckSpecialChars(tt.mode))

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestCheckSpecialCharsWithError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string with special characters only",
			value:       "orange 13",
			mode:        Require,
		},
		{
			description: "string have special characters",
			value:       "$$$apple$$$",
			mode:        Disallow,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value, CheckSpecialChars(tt.mode))

			// Assert
			require.NotNil(t, err)
		})
	}
}

func TestCheckNumbersNoError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string not contains digits, check mode is setted with 'disallow'",
			value:       "World War II",
			mode:        Disallow,
		},
		{
			description: "string contains digits, check mod is setted with 'required'",
			value:       "World War 2",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", CheckNumbers(tt.mode))

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestCheckNumbersWithError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string with number, check mode set to 'disallow'",
			value:       "1º position",
			mode:        Disallow,
		},
		{
			description: "string not contain number, check mode set to 'required'",
			value:       "first position",
			mode:        Require,
		},
		{
			description: "string filled with spaces characters, check mode set to 'required'",
			value:       "         ",
			mode:        Require,
		},
		{
			description: "string filled with special characters, check mode set to 'required'",
			value:       "#$@¨%$",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value, CheckNumbers(tt.mode))

			// Assert
			assert.Error(t, err)
		})
	}
}

func TestCheckLettersNoError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string not have letters, check mode set to 'disallow'",
			value:       "1945 @#$",
			mode:        Disallow,
		},
		{
			description: "string have letters, check mode set to 'required'",
			value:       "pink",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value, CheckLetters(tt.mode))

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestCheckLettersWithError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string have letters, check mode set to 'disallow'",
			value:       "Spanish",
			mode:        Disallow,
		},
		{
			description: "string not have letters, check mode set to 'required'",
			value:       "51",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value, CheckLetters(tt.mode))

			// Assert
			assert.Error(t, err)
		})
	}
}


func TestIsFormatValidNoError(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`[a-z]`)
	const value = "valid: string contains only letters"

	// Act
	err := New("field", value, IsFormatValid(re, ErrFieldCannotBeEmpty))

	// Assert
	assert.NoError(t, err)
}

func TestIsFormatValidWithError(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`^[0-9]+$`)
	const value = "valid: string contains only letters"

	// Act
	err := New("field", value, IsFormatValid(re, ErrFieldCannotBeEmpty))

	// Assert
	assert.Error(t, err)
}

func TestCheckWithRegexNoError(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`^[0-9]{3,5}$`)
	testsCases := []test{
		{
			description: "Is valid",
			value:       "abcdef1s2",
			mode:        Disallow,
		},
		{
			description: "Is valid",
			value:       "123",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value, CheckWithRegex(re, tt.mode, ErrFieldCannotBeEmpty))

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestCheckWithRegexWithError(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`^[0-9]{3,5}$`)
	testsCases := []test{
		{
			description: "",
			value:       "123",
			mode:        Disallow,
		},
		{
			description: "",
			value:       "abcded",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := New("field", tt.value, CheckWithRegex(re, tt.mode, ErrFieldCannotBeEmpty))

			// Assert
			assert.Error(t, err)
		})
	}
}

func TestIsBeforeThanNoError(t *testing.T) {
	// Arrange
	minDate := time.Now().AddDate(-18, 0, 0)   // 2006
	birthDate := time.Now().AddDate(-25, 0, 0) // 1999

	// Act
	err := New("birth_date", birthDate, IsBeforeThan(minDate, ErrFieldCannotBeEmpty))

	// Assert
	assert.NoError(t, err)
}

func TestIsBeforeThanWithError(t *testing.T) {
	// Arrange
	minDate := time.Now().AddDate(-18, 0, 0)
	birthDate := time.Now().AddDate(-18, 0, +1)

	// Act
	err := New("current_date", birthDate, IsBeforeThan(minDate, ErrFieldCannotBeEmpty))

	// Assert
	assert.Error(t, err)
}

func TestIsAfterThanNoError(t *testing.T) {
	// Arrange
	maxDate := time.Now().AddDate(-100, 0, 0)
	birthDate := time.Now().AddDate(-100, 0, 0)

	// Act
	err := New("current_date", birthDate, IsAfterThan(maxDate, ErrFieldCannotBeEmpty))

	// Assert
	assert.NoError(t, err)
}

func TestIsAfterThanWithError(t *testing.T) {
	// Arrange
	maxDate := time.Now().AddDate(-100, 0, 0)
	birthDate := time.Now().AddDate(-100, 0, -1)

	// Act
	err := New("current_date", birthDate, IsAfterThan(maxDate, ErrFieldCannotBeEmpty))

	// Assert
	assert.Error(t, err)
}

type testString struct {
	testDescription string
	min             int
	max             int
	value           string
}

func TestIsLengthInRangeNoError(t *testing.T) {
	// Arrange
	testsCases := []testString{
		{
			testDescription: "value is equal to minimum",
			min:             10,
			max:             20,
			value:           strings.Repeat("b", 10),
		},
		{
			testDescription: "value is equal than maximum",
			min:             0,
			max:             10,
			value:           strings.Repeat("a", 10),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			err := New("quantity", tt.value, IsLengthInRange(tt.min, tt.max))

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestIsLengthInRangeWithError(t *testing.T) {
	// Arrange
	testsCases := []testString{
		{
			testDescription: "value is less to than minimum allowed",
			min:             10,
			max:             20,
			value:           strings.Repeat("b", 9),
		},
		{
			testDescription: "value is greater than maximum allowed",
			min:             0,
			max:             10,
			value:           strings.Repeat("a", 11),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			err := New("quantity", tt.value, IsLengthInRange(tt.min, tt.max))

			// Assert
			assert.Error(t, err)
		})
	}
}

type testNumber struct {
	testDescription string
	min             int
	max             int
	value           int
}

func TestIsInRangeNoError(t *testing.T) {
	// Arrange
	testsCases := []testNumber{
		{
			testDescription: "the value is equal to minimum  allowed",
			min:             -1,
			max:             1,
			value:           -1,
		},
		{
			testDescription: "the value is equal to maximum allowed",
			min:             -2,
			max:             2,
			value:           2,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			err := New("value", tt.value, IsInRange(tt.min, tt.max))

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestIsInRangeWithError(t *testing.T) {
	// Arrange
	testsCases := []testNumber{
		{
			testDescription: "the value is short to minimum allowed",
			min:             -1,
			max:             1,
			value:           -2,
		},
		{
			testDescription: "the value is greater to maximum allowed",
			min:             -2,
			max:             2,
			value:           3,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			err := New("value", tt.value, IsInRange(tt.min, tt.max))

			// Assert
			assert.Error(t, err)
		})
	}
}

type testGenerics struct {
	testDescription string
	value           any
	flag            any
}

func TestIsEqualToNoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []testGenerics{
		{
			testDescription: "compare string",
			value:           "abc",
			flag:            "abc",
		},
		{
			testDescription: "compare number",
			value:           12,
			flag:            12,
		},
		{
			testDescription: "compare time",
			value:           time.Now().UTC(),
			flag:            time.Now().UTC(),
		},
		{
			testDescription: "compare booleans",
			value:           true,
			flag:            true,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			err := New("quantity", tt.value, IsEqualTo(tt.flag, ErrFieldCannotBeEmpty))

			// Assert
			assert.NoError(t, err)
		})
	}
}

func TestIsEqualToWithError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []testGenerics{
		{
			testDescription: "strings not equals",
			value:           "123",
			flag:            "456",
		},
		{
			testDescription: "number not equals",
			flag:            12,
			value:           13,
		},
		{
			testDescription: "time not equals",
			flag:            time.Now().AddDate(-2, 0, 0),
			value:           time.Now(),
		},
		{
			testDescription: "boolean not equals",
			flag:            true,
			value:           false,
		},
		{
			testDescription: "types not are equals",
			flag:            0,
			value:           "string",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			err := New[any]("quantity", tt.value, IsEqualTo(tt.flag, ErrFieldCannotBeEmpty))

			// Assert
			assert.Error(t, err)
		})
	}
}
