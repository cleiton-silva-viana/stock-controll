package validate

// Adotar uso opcional de códigos de erros nas funções validadoras
// Criar um wrapper que embrulha todas as funções validadoras

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"stock-controll/internal/domain/services/error/field"
)

const (
	ErrUnsupportedType           = "ERR_UNSUPPORTED_TYPE"
	ErrResourceAlreadyRegistered = "ERR_ALREADY_REGISTERED"
	ErrResourceNotFound          = "ERR_FIELD_NOT_FOUND"
)

var (
	ValidatorPool = sync.Pool{
		New: func() interface{} {
			return &Validator[any]{}
		},
	}
)

func LoadValidatorPoll[T any](value T) *Validator[T] {
	var v, ok = ValidatorPool.Get().(*Validator[T])
	if !ok {
		v = &Validator[T]{}
	}
	defer ValidatorPool.Put(v)
	v.value = value
	v.errorCode = field.NoError
	return v
}

type Validator[T any] struct {
	value     T
	errorCode string
}

type Options[T any] func(*Validator[T])

func New[T any](fieldName string, fieldValue T, Validators ...Options[T]) error {
	Validator := LoadValidatorPoll(fieldValue)

	err := &field.FieldError{
		FieldName: fieldName,
		CodeError: field.NoError,
	}

	for _, validate := range Validators {
		validate(Validator)
		if Validator.errorCode != field.NoError {
			err.AddErrorCode(Validator.errorCode)
			break
		}
	}

	if err.HasError() {
		return err
	}
	return nil
}

type checkMode int

const (
	Require checkMode = iota
	Disallow
)

func setErrorCodeBasedOnMatch(v *Validator[string], match bool, mode checkMode, codeForRequire, codeForDisallow string) {
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

const (
	ErrFieldMustContainSpecialCharacters   = "ERR_FIELD_CONTAINS_SPECIAL_CHARACTERS"
	ErrFieldCannotContainSpecialCharacters = "ERR_FIELD_CANNOT_CONTAINS_SPECIAL_CHARACTERS"
)

func CheckSpecialChars(mode checkMode) func(v *Validator[string]) {
	caseRequire := ErrFieldMustContainSpecialCharacters
	caseDisallow := ErrFieldCannotContainSpecialCharacters
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]`)
	return func(v *Validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(v, match, mode, caseRequire, caseDisallow)
	}
}

const (
	ErrFieldMustContainNumbers   = "ERR_FIELD_CONTAINS_NUMBERS"
	ErrFieldCannotContainNumbers = "ERR_FIELD_CONTAINS_NUMBERS"
)

func CheckNumbers(mode checkMode) func(v *Validator[string]) {
	caseRequire := ErrFieldMustContainNumbers
	caseDisallow := ErrFieldCannotContainNumbers
	re := regexp.MustCompile(`[0-9]`)
	return func(v *Validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(v, match, mode, caseRequire, caseDisallow)
	}
}

const (
	ErrFieldMustContainLetters    = "ERR_FIELD_CONTAINS_LETTERS"
	ErrFieldCannotContainsLetters = "ERR_FIELD_CONTAINS_LETTERS"
)

func CheckLetters(mode checkMode) func(v *Validator[string]) {
	caseRequire := ErrFieldMustContainLetters
	caseDisallow := ErrFieldCannotContainsLetters
	re := regexp.MustCompile(`[a-zA-z]`)
	return func(v *Validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(v, match, mode, caseRequire, caseDisallow)
	}
}

const ErrFieldCannotBeEmpty = "ERR_FIELD_CANNOT_BE_EMPTY"

func IsBlank() func(v *Validator[string]) {
	return func(v *Validator[string]) {
		if strings.Trim(v.value, " ") == "" {
			v.errorCode = ErrFieldCannotBeEmpty
		}
	}
}

const ErrFieldInvalidLength = "ERR_FIELD_INVALID_LENGTH"

func IsLengthEqualTo(length int) func(v *Validator[string]) {
	return func(v *Validator[string]) {
		if len(v.value) != length {
			v.errorCode = ErrFieldInvalidLength
		}
	}
}

const ErrFieldLengthOutOfRange = "ERR_FIELD_LENGTH_OUT_OF_RANGE"

func IsLengthInRange(minLength, maxLength int) func(v *Validator[string]) {
	return func(v *Validator[string]) {
		if len(v.value) < minLength || len(v.value) > maxLength {
			v.errorCode = ErrFieldLengthOutOfRange
		}
	}
}

// Regexp

func IsFormatValid(re *regexp.Regexp, code string) func(v *Validator[string]) {
	return func(v *Validator[string]) {
		match := re.MatchString(v.value)
		if !match {
			v.errorCode = code
		}
	}
}

func CheckWithRegex(re *regexp.Regexp, mode checkMode, code string) func(v *Validator[string]) {
	return func(v *Validator[string]) {
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

func IsInEnum(object []string) func(v *Validator[string]) {
	return func(v *Validator[string]) {
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

func IsBeforeThan(minimunDate time.Time) func(v *Validator[time.Time]) {
	return func(v *Validator[time.Time]) {
		if minimunDate.Unix() < v.value.Unix() {
			v.errorCode = fmt.Sprint(ErrTimeIsBeforeTheLimit, minimunDate)
		}
	}
}

const ErrTimeIsAfterTheLimit = "ERR_TIME_IS_AFTER_THE_LIMIT"

func IsAfterThan(maximumDate time.Time) func(v *Validator[time.Time]) {
	return func(v *Validator[time.Time]) {
		if maximumDate.Unix() > v.value.Unix() {
			v.errorCode = fmt.Sprint(ErrTimeIsAfterTheLimit, maximumDate)
		}
	}
}

// Int

const ErrNumberOutOfRange = "ERR_NUMBER_OUT_OF_RANGE"

func IsInRange(min, max int) func(v *Validator[int]) {
	return func(v *Validator[int]) {
		if v.value < min || v.value > max {
			v.errorCode = ErrNumberOutOfRange
		}
	}
}

// float64

func IsInRangeFloat64(min, max float64) func(v *Validator[float64]) {
	return func(v *Validator[float64]) {
		if v.value < min || v.value > max {
			v.errorCode = ErrNumberOutOfRange
		}
	}
}

// Generics

func IsEqualTo(value any, code string) func(v *Validator[any]) {
	return func(v *Validator[any]) {
		if v.value != value {
			v.errorCode = code
		}
	}
}

const ErrNilObject = "ERR_NIL_OBJECT"

func IsNil() func(v *Validator[any]) {
	return func(v *Validator[any]) {
		if v.value == nil {
			v.errorCode = ErrNilObject
		}
	}
}
