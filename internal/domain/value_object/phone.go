package value_object

import (
	"fmt"
	"regexp"
	"stock-controll/internal/domain/failure"
	"strconv"
	"strings"
)

type Phone struct {
	areaCode int
	number   int
	device   string
}

func (p *Phone) GetPhone() string {
	return fmt.Sprintf("(%d) %d", p.areaCode, p.number)
}

func NewPhone(phone string) (*Phone, error) {
	if ContainsLetters(phone) {
		return nil, failure.PhoneWithLetters
	}

	areaCode := getAreaCode(phone)
	if areaCode == 0 {
		return nil, failure.PhoneWithoutAreaCode
	}

	phoneParsed := removeSpecialCharsOfPhone(phone)
	phoneParsed = strings.ReplaceAll(phoneParsed, " ", "")
	if ContainsSpecialChars(phoneParsed) {
		return nil, failure.PhoneWithSpecialCharacters
	}
	
	device := getDevice(phone)
	if device == "unknow" {
		return nil, failure.PhoneWithInvalidLength
	}

	number := getNumber(phone)

	return &Phone{
		areaCode: areaCode,
		number:   number,
		device:   device,
	}, nil
}

func removeSpecialCharsOfPhone(phone string) string {
	replacer := strings.NewReplacer(
		"(", "",
		")", "",
		"-", "",
	)
	return replacer.Replace(phone)
}

func getPhoneLength(phone string) int {
	phoneParsed := removeSpecialCharsOfPhone(phone)
	phoneParsed = strings.ReplaceAll(phoneParsed, " ", "")
	return len(phoneParsed)
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
	if phoneLength == 11 {
		return "cellphone"
	}
	return "unknow"
}

func getNumber(phone string) int {
	phoneParts := strings.Split(phone, " ")
	phoneNumber := removeSpecialCharsOfPhone(phoneParts[1])
	number, _ := strconv.Atoi(phoneNumber)
	return number
}
