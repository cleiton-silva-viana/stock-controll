package value_object

import (
	"regexp"
	"stock-controll/internal/contact/errors"
	"strconv"
	"strings"
)

type Phone struct {
	areaCode int
	number   int
	device   string
}

func NewPhone(phone string) (*Phone, error) {
	err := checkPhoneLength(phone)
	if err != nil {
		return nil, err
	}

	err = checkIfPhoneContainsLetters(phone)
	if err != nil {
		return nil, err
	}

	err = checkIfPhoneContainsInvalidCharacters(phone)
	if err != nil {
		return nil, err
	}

	err = checkIfPhoneContainsAreaCode(phone)
	if err != nil {
		return nil, err
	}

	device := getDevice(phone)
	areaCode := getAreaCode(phone)
	number := getNumber(phone)

	return &Phone{
		areaCode: areaCode,
		number:   number,
		device:   device,
	}, nil
}

func checkIfPhoneContainsLetters(phone string) error {
	re := regexp.MustCompile(`[a-z]`)
	if re.MatchString(phone) {
		return errors.PhoneWithLetters
	}
	return nil
}

func removeSpecialCharsOfPhone(phone string) string {
	replacer := strings.NewReplacer(
		"(", "",
		")", "",
		"-", "",
	)
	return replacer.Replace(phone)
}

func checkIfPhoneContainsInvalidCharacters(phone string) error {
	phoneParsed := removeSpecialCharsOfPhone(phone)
	re := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)

	if re.MatchString(phoneParsed) {
		return errors.PhoneWithSpecialCharacters
	}
	return nil
}

func checkIfPhoneContainsAreaCode(phone string) error {
	if getAreaCode(phone) == 0 {
		return errors.PhoneWithoutAreaCode
	}
	return nil
}

func getPhoneLength(phone string) int {
	phoneParsed := removeSpecialCharsOfPhone(phone)
	phoneParsed = strings.ReplaceAll(phoneParsed, " ", "")
	return len(phoneParsed)
}

func checkPhoneLength(phone string) error {
	phoneLength := getPhoneLength(phone)
	if phoneLength < 10 || phoneLength > 11 {
		return errors.PhoneWithInvalidLength
	}
	return nil
}

func getAreaCode(phone string) int {
	hasAreaCode := strings.Contains(phone, "(")
	if hasAreaCode {
		re := regexp.MustCompile(`(\d+)`)
		result := re.FindString(phone)
		areaCode, _ := strconv.Atoi(result)
		return areaCode
	}
	return 0
}

func getDevice(phone string) string {
	phoneLength := getPhoneLength(phone)
	if phoneLength == 10 {
		return "landline"
	}
	return "cellphone"
}

func getNumber(phone string) int {
	phoneParts := strings.Split(phone, " ")
	phoneNumber := removeSpecialCharsOfPhone(phoneParts[1])
	number, _ := strconv.Atoi(phoneNumber)
	return number
}
