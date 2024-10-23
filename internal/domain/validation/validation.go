package validation

// Adotar uso opcional de códigos de erros nas funções validadoras
// Criar um wrapper que embrulha todas as funções validadoras

import (
	"regexp"
	"strings"
	"sync"
	"time"
)

type errorCode string

const (
	NoError   errorCode = "NO_ERROR"
	ErrUnknown errorCode = "UNKNOW_ERROR"
)

var (
	validatorPool = sync.Pool{
		New: func() interface{} {
			return &validator[any]{}
		},
	}
)

func loadValidatorPoll[T any](value T) *validator[T] {
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
	FieldName  string
	CodeErrors []string
}

func NewFieldError(fieldName string) *FieldError {

	// Adicionar validações ...

	return &FieldError{
		FieldName:  fieldName,
		CodeErrors: make([]string, 5),
	}
}

func (v *FieldError) AddErrorCode(code string) *FieldError {
	if code != string(NoError) {
		v.CodeErrors = append(v.CodeErrors, code)
	}
	return v
}

func (v *FieldError) HasError() bool {
	return len(v.CodeErrors) > 0
}

type validator[T any] struct {
	value     T
	errorCode errorCode
}

type ValidateOptions[T any] func(*validator[T])

func Validate[T any](fieldName string, fieldValue T, validators ...ValidateOptions[T]) *FieldError {
	validator := loadValidatorPoll(fieldValue)

	var errs = FieldError{
		FieldName:  fieldName,
	}

	for _, validate := range validators {
		validate(validator)
		errs.AddErrorCode(string(validator.errorCode))
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

func setErrorCodeBasedOnMatch(match bool, mode checkMode, code errorCode, v *validator[string]) {
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

func CheckSpecialChars(mode checkMode, code errorCode) func(v *validator[string]) {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]`)
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(match, mode, code, v)
	}
}

func CheckNumbers(mode checkMode, code errorCode) func(v *validator[string]) {
	re := regexp.MustCompile(`[0-9]`)
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(match, mode, code, v)
	}
}

func CheckLetters(mode checkMode, code errorCode) func(v *validator[string]) {
	re := regexp.MustCompile(`[a-zA-z]`)
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(match, mode, code, v)
	}
}

func CheckLowerCaseLetters(mode checkMode, code errorCode) func(v *validator[string]) {
	re := regexp.MustCompile(`[a-z]`)
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(match, mode, code, v)
	}
}

func CheckUpperCaseLetters(mode checkMode, code errorCode) func(v *validator[string]) {
	re := regexp.MustCompile(`[A-Z]`)
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		setErrorCodeBasedOnMatch(match, mode, code, v)
	}
}

func IsBlank(code errorCode) func(v *validator[string]) {
	return func(v *validator[string]) {
		if strings.Trim(v.value, " ") == "" {
			v.errorCode = code
		}
	}
}

func IsLengthEqualTo(length int, code errorCode) func(v *validator[string]) {
	return func(v *validator[string]) {
		if len(v.value) != length {
			v.errorCode = code
		}
	}
}

func IsLengthInRange(minLength, maxLength int, code errorCode) func(v *validator[string]) {
	return func(v *validator[string]) {
		if len(v.value) < minLength || len(v.value) > maxLength {
			v.errorCode = code
		}
	}
}

func IsValueInRange(values []string, ignoreCase bool, code errorCode) func(v *validator[string]) {
	return func(v *validator[string]) {
		for _, value := range values {
			if (ignoreCase && strings.EqualFold(value, v.value)) || (!ignoreCase && value == v.value) {
				return
			}
		}
		v.errorCode = code
	}
}

// Regexp

func IsFormatValid(re *regexp.Regexp, code errorCode) func(v *validator[string]) {
	return func(v *validator[string]) {
		match := re.MatchString(v.value)
		if !match {
			v.errorCode = code
		}
	}
}

func CheckWithRegex(re *regexp.Regexp, mode checkMode, code errorCode) func(v *validator[string]) {
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

// Time

func IsFutureDate(code errorCode) func(v *validator[time.Time]) {
	return func(v *validator[time.Time]) {
		currentDate := time.Now()
		if v.value.After(currentDate) {
			v.errorCode = code
		}
	}
}

func IsBeforeThan(minimunDate time.Time, code errorCode) func(v *validator[time.Time]) {
	return func(v *validator[time.Time]) {
		if minimunDate.Unix() < v.value.Unix() {
			v.errorCode = code
		}
	}
}

func IsAfterThan(maximumDate time.Time, code errorCode) func(v *validator[time.Time]) {
	return func(v *validator[time.Time]) {
		if maximumDate.Unix() > v.value.Unix() {
			v.errorCode = code
		}
	}
}

// Int

func IsInRange(min, max int, code errorCode) func(v *validator[int]) {
	return func(v *validator[int]) {
		if v.value < min || v.value > max {
			v.errorCode = code
		}
	}
}

// Generics

func IsEqualTo(value any, code errorCode) func(v *validator[any]) {
	return func(v *validator[any]) {
		if v.value != value {
			v.errorCode = code
		}
	}
}

func IsNil(object any, code errorCode) func(v *validator[any]) {
	return func(v *validator[any]) {
		if v.value == nil {
			v.errorCode = code
		}
	}
}

func IsInEnum(object []any, code errorCode) func(v *validator[any]) {
	return func(v *validator[any]) {
		for _, obj := range object {
			if v.value == obj {
				return
			}
		}
		v.errorCode = code
	}
}

