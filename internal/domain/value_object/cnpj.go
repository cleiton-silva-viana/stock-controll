package valueobject

import (
	"regexp"
	"strconv"
	"strings"
	
	"stock-controll/internal/domain/failure"
)

type CNPJ struct {
	cnpj string
}

func NewCNPJ(cnpj string) (*CNPJ, error) {
	isValidFormat := validateFormat(cnpj)
	if !isValidFormat {
		return nil, failure.FieldWithInvalidFormat("cnpj", "XX.XXX.XXX/0001-XX")
	}
	
	var CNPJ_digits = getDigits(cnpj)
	firstCheckerDigit := strconv.Itoa(validateDigits(CNPJ_digits, true))
	secondCheckDigit := strconv.Itoa(validateDigits(CNPJ_digits, false))
	checkerDigits := strconv.Itoa(extractValidatorDigitsOnCnpj(cnpj))

	if checkerDigits != firstCheckerDigit+secondCheckDigit {
		return nil, failure.CPNJWithInvalidCheckerDigits
	}

	return &CNPJ{
		cnpj: cnpj,
	}, nil
}

func validateFormat(cnpj string) bool {
	layout := `^(\d{2})(?:\.(\d{3}))(?:\.(\d{3}))(?:\/(\d{4}))(?:\-(\d{2})$)`
	re := regexp.MustCompile(layout)
	match := re.Match([]byte(cnpj))
	return match
}

func getDigits(cnpj string) string {
	re := regexp.MustCompile(`\d+`)
	matchs := re.FindAllString(cnpj, -1)
	digits := strings.Join(matchs, "")
	return digits
}

func extractValidatorDigitsOnCnpj(cnpj string) int {
	re := regexp.MustCompile(`\d{2}$`)
	match := re.FindString(cnpj)
	CheckerDigits, err := strconv.Atoi(match)
	if err != nil {
		return 0
	}
	return CheckerDigits
}

func validateDigits(cnpj string, first bool) int {
	var digits int
	var multiplicator int
	var flag int

	if first {
		digits = 12
		multiplicator = 5
		flag = 3
	} else {
		digits = 13
		multiplicator = 6
		flag = 4
	}

	CNPJ_string := getDigits(cnpj)
	var slice []int
	for _, char := range CNPJ_string {
		digit := int(char - '0')
		slice = append(slice, digit)
	}

	var somatorie int
	for i := 0; i < digits; i++ {
		digit := slice[i]
		result := digit * multiplicator
		somatorie += result

		if i == flag {
			multiplicator = 10
		}
		multiplicator -= 1
	}

	rest := somatorie % 11
	var DigitChecker int
	if (rest) < 2 {
		DigitChecker = 0
	} else {
		DigitChecker = 11 - rest
	}
	return DigitChecker
}
