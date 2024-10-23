package validation

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

func Test_Validate_NoError(t *testing.T) {
	// Arrange
	const (
		fieldName  = "password"
		fieldValue = "123456"
	)

	// Act
	err := Validate(fieldName, fieldValue, IsBlank(ErrUnknown))

	// Assert
	assert.Nil(t, err)
}

// Refatorar !!!
// Verificar se a quantidade de erros esperados está sendo retornada
// Verificar se os erros estão sendo retornados
func Test_Validate_WithError(t *testing.T) {
	// Arrange
	const fieldValue = "123456"

	// Act
	err := Validate("field", fieldValue, CheckNumbers(Disallow, ErrUnknown))

	// Assert
	require.NotNil(t, err)
	require.Contains(t, err.CodeErrors, ErrUnknown)
}

func Test_IsBlank_NoError(t *testing.T) {
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
			err := Validate("name", tt.value, IsBlank(ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_IsBlank_WithError(t *testing.T) {
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
			err := Validate("field", tt.value, IsBlank(ErrUnknown))

			// Assert
			require.NotNil(t, err)
		})
	}
}

func Test_IsLengthEqualTo_NoError(t *testing.T) {
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
			err := Validate("field", tt.value, IsLengthEqualTo(tt.length, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_IsLengthEqualTo_WithError(t *testing.T) {
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
			err := Validate("field", tt.value, IsLengthEqualTo(tt.length, ErrUnknown))

			// Assert
			require.NotNil(t, err)
		})
	}
}

func Test_CheckSpecialChars_NoError(t *testing.T) {
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
			err := Validate("field", tt.value, CheckSpecialChars(tt.mode, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_CheckSpecialChars_WithError(t *testing.T) {
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
			err := Validate("field", tt.value, CheckSpecialChars(tt.mode, ErrUnknown))

			// Assert
			require.NotNil(t, err)
		})
	}
}

func Test_CheckNumbers_NoError(t *testing.T) {
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
			err := Validate("field", CheckNumbers(tt.mode, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_CheckNumbers_WithError(t *testing.T) {
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
			err := Validate("field", tt.value, CheckNumbers(tt.mode, ErrUnknown))

			// Assert
			assert.NotNil(t, err)
		})
	}
}

func Test_CheckLetters_NoError(t *testing.T) {
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
			err := Validate("field", tt.value, CheckLetters(tt.mode, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_CheckLetters_WithError(t *testing.T) {
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
			err := Validate("field", tt.value, CheckLetters(tt.mode, ErrUnknown))

			// Assert
			assert.NotNil(t, err)
		})
	}
}

func Test_CheckLowerCaseLetters_NoError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string not have lowercase letters, check mode set to 'disallow'",
			value:       "PINK",
			mode:        Disallow,
		},
		{
			description: "string have lowercase letters, check mode set to 'require'",
			value:       "pindaíbaSSSSS",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := Validate("field", tt.value, CheckLowerCaseLetters(tt.mode, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_CheckLowerCaseLetters_WithError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string with lower case letters, check mode set to 'disallow'",
			value:       "EMma",
			mode:        Disallow,
		},
		{
			description: "string not have lowercase letters, check mode set to 'require'",
			value:       "LAMMA",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := Validate("field", tt.value, CheckLowerCaseLetters(tt.mode, ErrUnknown))

			// Assert
			assert.NotNil(t, err)
		})
	}
}

func Test_CheckUpperCaseLetters_NoError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string not have uppercase letters, check mode set to 'disallow'",
			value:       "mindflow 2024 $",
			mode:        Disallow,
		},
		{
			description: "string have uppercase letters, check mode set to 'require'",
			value:       "killswitch ENGAGE 123 !@#",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := Validate("field", tt.value, CheckUpperCaseLetters(tt.mode, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_CheckUpperCaseLetters_WithError(t *testing.T) {
	// Arrange
	testsCases := []test{
		{
			description: "string have uppercase letters, check mode set to 'disallow'",
			value:       "RED HOT chil peper 1971",
			mode:        Disallow,
		},
		{
			description: "string not have uppercase letters, check mode set to 'require'",
			value:       "post malone #1",
			mode:        Require,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := Validate("field", tt.value, CheckUpperCaseLetters(tt.mode, ErrUnknown))

			// Assert
			assert.NotNil(t, err)
		})
	}
}

func Test_IsFormatValid_NoError(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`[a-z]`)
	const value = "valid: string contains only letters"

	// Act
	err := Validate("field", value, IsFormatValid(re, ErrUnknown))

	// Assert
	assert.Nil(t, err)
}

func Test_IsFormatValid_WithError(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`^[0-9]+$`)
	const value = "valid: string contains only letters"

	// Act
	err := Validate("field", value, IsFormatValid(re, ErrUnknown))

	// Assert
	assert.NotNil(t, err)
}

func Test_CheckWithRegex_NoError(t *testing.T) {
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
			err := Validate("field", tt.value, CheckWithRegex(re, tt.mode, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_CheckWithRegex_WithError(t *testing.T) {
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
			err := Validate("field", tt.value, CheckWithRegex(re, tt.mode, ErrUnknown))

			// Assert
			assert.NotNil(t, err)
		})
	}
}

func Test_IsValueInRange_NoError(t *testing.T) {
	// Arrange
	values := []string{"apple", "banana", "orange"}
	testsCases := []struct {
		test
		ignoreCase bool
	}{
		{
			test: test{
				description: "value is in the list, ingore case is true",
				value:       "APPLE",
			},
			ignoreCase: true,
		},
		{
			test: test{
				description: "valid: value is in the list, ignore case is false",
				value:       "banana",
			},
			ignoreCase: false,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := Validate("field", tt.value, IsValueInRange(values, tt.ignoreCase, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_IsValueInRange_WithError(t *testing.T) {
	// Arrange
	values := []string{"apple", "banana", "orange"}
	testsCases := []struct {
		test
		ignoreCase bool
	}{
		{
			test: test{
				description: "value is not in the list",
				value:       "sleeve",
			},
			ignoreCase: true,
		},
		{
			test: test{
				description: "value in the list, but ignore case is false",
				value:       "Apple",
			},
			ignoreCase: false,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {

			// Act
			err := Validate("field", tt.value, IsValueInRange(values, tt.ignoreCase, ErrUnknown))

			// Assert
			assert.NotNil(t, err)
		})
	}
}

func Test_IsFutureDate_NoError(t *testing.T) {
	// Arrange
	currentDate := time.Now()

	// Act
	err := Validate("current_date", currentDate, IsFutureDate(ErrUnknown))

	// Assert
	assert.Nil(t, err)
}

func Test_IsFutureDate_WithError(t *testing.T) {
	// Arrange
	currentDate := time.Now().AddDate(0, 0, +1)

	// Act
	err := Validate("current_date", currentDate, IsFutureDate(ErrUnknown))

	// Assert
	assert.NotNil(t, err)
}

func Test_IsBeforeThan_NoError(t *testing.T) {
	// Arrange
	minDate := time.Now().AddDate(-18, 0, 0)   // 2006
	birthDate := time.Now().AddDate(-25, 0, 0) // 1999

	// Act
	err := Validate("birth_date", birthDate, IsBeforeThan(minDate, ErrUnknown))

	// Assert
	assert.Nil(t, err)
}

func Test_IsBeforeThan_WithError(t *testing.T) {
	// Arrange
	minDate := time.Now().AddDate(-18, 0, 0)
	birthDate := time.Now().AddDate(-18, 0, +1)

	// Act
	err := Validate("current_date", birthDate, IsBeforeThan(minDate, ErrUnknown))

	// Assert
	assert.NotNil(t, err)
}

func Test_IIsAfterThan_NoError(t *testing.T) {
	// Arrange
	maxDate := time.Now().AddDate(-100, 0, 0)
	birthDate := time.Now().AddDate(-100, 0, 0)

	// Act
	err := Validate("current_date", birthDate, IsAfterThan(maxDate, ErrUnknown))

	// Assert
	assert.Nil(t, err)
}

func Test_IIsAfterThan_WithError(t *testing.T) {
	// Arrange
	maxDate := time.Now().AddDate(-100, 0, 0)
	birthDate := time.Now().AddDate(-100, 0, -1)

	// Act
	err := Validate("current_date", birthDate, IsAfterThan(maxDate, ErrUnknown))

	// Assert
	assert.NotNil(t, err)
}

type testString struct {
	testDescription string
	min             int
	max             int
	value           string
}

func Test_IsLengthInRange_NoError(t *testing.T) {
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
			err := Validate("quantity", tt.value, IsLengthInRange(tt.min, tt.max, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_IsLengthInRange_WithError(t *testing.T) {
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
			err := Validate("quantity", tt.value, IsLengthInRange(tt.min, tt.max, ErrUnknown))

			// Assert
			assert.NotNil(t, err)
		})
	}
}

type testNumber struct {
	testDescription string
	min             int
	max             int
	value           int
}

func Test_IsInRange_NoError(t *testing.T) {
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
			err := Validate("value", tt.value, IsInRange(tt.min, tt.max, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_IsInRange_WithError(t *testing.T) {
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
			err := Validate("value", tt.value, IsInRange(tt.min, tt.max, ErrUnknown))

			// Assert
			assert.NotNil(t, err)
		})
	}
}

type testGenerics struct {
	testDescription string
	value           any
	flag            any
}

func Test_IsEqualTo_NoError(t *testing.T) {
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
			err := Validate("quantity", tt.value, IsEqualTo(tt.flag, ErrUnknown))

			// Assert
			assert.Nil(t, err)
		})
	}
}

func Test_IsEqualTo_WithError(t *testing.T) {
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
			err := Validate[any]("quantity", tt.value, IsEqualTo(tt.flag, ErrUnknown))

			// Assert
			assert.NotNil(t, err)
		})
	}
}
