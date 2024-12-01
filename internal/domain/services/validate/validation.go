package validate

// Adotar uso opcional de códigos de erros nas funções validadoras
// Criar um wrapper que embrulha todas as funções validadoras

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
)

type ErrorCode string

const (
	NoError               = "NO_ERROR"
	ErrFieldCannotBeEmpty = "ERR_FIELD_CANNOT_BE_EMPTY"
	ErrFieldInvalidLength = "ERR_FIELD_INVALID_LENGTH"

	ErrFieldMustContainSpecialCharacters   = "ERR_FIELD_CONTAINS_SPECIAL_CHARACTERS"
	ErrFieldCannotContainSpecialCharacters = "ERR_FIELD_CANNOT_CONTAINS_SPECIAL_CHARACTERS"

	ErrFieldMustContainNumbers   = "ERR_FIELD_CONTAINS_NUMBERS"
	ErrFieldCannotContainNumbers = "ERR_FIELD_CONTAINS_NUMBERS"

	ErrFieldMustContainLetters    = "ERR_FIELD_CONTAINS_LETTERS"
	ErrFieldCannotContainsLetters = "ERR_FIELD_CONTAINS_LETTERS"

	ErrFieldLengthOutOfRange     = "ERR_FIELD_LENGTH_OUT_OF_RANGE"
	ErrUnsupportedType           = "ERR_UNSUPPORTED_TYPE"
	ErrResourceAlreadyRegistered = "ERR_ALREADY_REGISTERED"
	ErrResourceNotFound          = "ERR_FIELD_NOT_FOUND"

	ErrNilObject = "ERR_NIL_OBJECT"
	/*
		"ERR_NIL_OBJECT": {
		    "Message": "The object is null and cannot be modified.",
		    "Solution": "Please ensure the object is initialized before attempting to modify it."
		}
	*/

	// Mudar de pacote >>>
	ErrInvalidUUIDFormat = "ERR_INVALID_UUID_FORMAT"
)

var (
	validatorPool = sync.Pool{
		New: func() interface{} {
			return &validator[any]{}
		},
	}
)

func LoadValidatorPoll[T any](value T) *validator[T] {
	var v, ok = validatorPool.Get().(*validator[T])
	if !ok {
		v = &validator[T]{}
	}
	defer validatorPool.Put(v)
	v.value = value
	v.errorCode = NoError
	return v
}

type FieldError struct {
	FieldName string
	CodeError string
}

func NewFieldError(fieldName string) *FieldError {

	// Adicionar validações ...

	return &FieldError{
		FieldName: fieldName,
		CodeError: NoError,
	}
}

func (fe *FieldError) Error() string {
	return fe.CodeError
}

func (fe *FieldError) AddErrorCode(code string) *FieldError {
	fe.CodeError = code
	return fe
}

func (fe *FieldError) HasError() bool {
	return fe.CodeError != NoError
}

type validator[T any] struct {
	value     T
	errorCode string
}

type ValidateOptions[T any] func(*validator[T])

func New[T any](fieldName string, fieldValue T, validators ...ValidateOptions[T]) *FieldError {
	validator := LoadValidatorPoll(fieldValue)

	var errs = FieldError{
		FieldName: fieldName,
	}

	for _, validate := range validators {
		validate(validator)
		errs.AddErrorCode(validator.errorCode)
		validator.errorCode = NoError
	}

	if errs.HasError() {
		return &errs
	}
	return nil
}

type checkMode int

const (
	Require checkMode = iota
	Disallow
)

func setErrorCodeBasedOnMatch(v *validator[string], match bool, mode checkMode, codeForRequire, codeForDisallow string) {
	switch mode {
	case Require:
		if !match {
			v.errorCode = codeForRequire
		}
	case Disallow:
		if match {
			v.errorCode = codeForDisallow
		}
	}
}

func CheckSpecialChars(mode checkMode) func(v *validator[string]) {
	caseRequire := ErrFieldMustContainSpecialCharacters
	caseDisallow := ErrFieldCannotContainSpecialCharacters
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]`)
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(v, match, mode, caseRequire, caseDisallow)
	}
}

func CheckNumbers(mode checkMode) func(v *validator[string]) {
	caseRequire := ErrFieldMustContainNumbers
	caseDisallow := ErrFieldCannotContainNumbers
	re := regexp.MustCompile(`[0-9]`)
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(v, match, mode, caseRequire, caseDisallow)
	}
}

func CheckLetters(mode checkMode) func(v *validator[string]) {
	caseRequire := ErrFieldMustContainLetters
	caseDisallow := ErrFieldCannotContainsLetters
	re := regexp.MustCompile(`[a-zA-z]`)
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(v, match, mode, caseRequire, caseDisallow)
	}
}

func IsBlank() func(v *validator[string]) {
	return func(v *validator[string]) {
		if strings.Trim(v.value, " ") == "" {
			v.errorCode = ErrFieldCannotBeEmpty
		}
	}
}

func IsLengthEqualTo(length int) func(v *validator[string]) {
	return func(v *validator[string]) {
		if len(v.value) != length {
			v.errorCode = ErrFieldInvalidLength
		}
	}
}

func IsLengthInRange(minLength, maxLength int) func(v *validator[string]) {
	return func(v *validator[string]) {
		if len(v.value) < minLength || len(v.value) > maxLength {
			v.errorCode = ErrFieldLengthOutOfRange
		}
	}
}

/*
func IsValueInRange(values []string, ignoreCase bool, code string) func(v *validator[string]) {
	return func(v *validator[string]) {
		for _, value := range values {
			if (ignoreCase && strings.EqualFold(value, v.value)) || (!ignoreCase && value == v.value) {
				return
			}
		}
		v.errorCode = code
	}
}
*/

// Regexp

func IsFormatValid(re *regexp.Regexp, code string) func(v *validator[string]) {
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		if !match {
			v.errorCode = code
		}
	}
}

func CheckWithRegex(re *regexp.Regexp, mode checkMode, code string) func(v *validator[string]) {
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		switch mode {
		case Require:
			if !match {
				v.errorCode = code
			}
		case Disallow:
			if match {
				v.errorCode = code
			}
		}
	}
}

const ErrValueNotInTheEnum = "ERR_VALUE_NOT_IN_ENUM"

func IsInEnum(object []string) func(v *validator[string]) {
	return func(v *validator[string]) {
		for _, obj := range object {
			if v.value == obj {
				return
			}
		}
		v.errorCode = ErrValueNotInTheEnum
	}
}

// Time
const ErrTimeIsBeforeTheLimit = "ERR_TIME_IS_BEFORE_THE_LIMIT"

func IsBeforeThan(minimunDate time.Time, code string) func(v *validator[time.Time]) {
	return func(v *validator[time.Time]) {
		if minimunDate.Unix() < v.value.Unix() {
			v.errorCode = fmt.Sprint(ErrTimeIsBeforeTheLimit, minimunDate)
		}
	}
}

const ErrTimeIsAfterTheLimit = "ERR_TIME_IS_AFTER_THE_LIMIT"

func IsAfterThan(maximumDate time.Time, code string) func(v *validator[time.Time]) {
	return func(v *validator[time.Time]) {
		if maximumDate.Unix() > v.value.Unix() {
			v.errorCode = fmt.Sprint(ErrTimeIsAfterTheLimit, maximumDate)
		}
	}
}

// Int

const ErrNumberOutOfRange = "ERR_NUMBER_OUT_OF_RANGE"

func IsInRange(min, max int) func(v *validator[int]) {
	return func(v *validator[int]) {
		if v.value < min || v.value > max {
			v.errorCode = ErrNumberOutOfRange
		}
	}
}

// float64

func IsInRangeFloat64(min, max float64) func(v *validator[float64]) {
	return func(v *validator[float64]) {
		if v.value < min || v.value > max {
			v.errorCode = ErrNumberOutOfRange
		}
	}
}

// Generics

func IsEqualTo(value any, code string) func(v *validator[any]) {
	return func(v *validator[any]) {
		if v.value != value {
			v.errorCode = code
		}
	}
}

func IsNil(object any) func(v *validator[any]) {
	return func(v *validator[any]) {
		if v.value == nil {
			v.errorCode = ErrNilObject
		}
	}
}
