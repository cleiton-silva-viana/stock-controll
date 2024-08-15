package errors

import common_errors "stock-controll/internal/common/entity/errors"

var PhoneWithSpecialCharacters = NewPhoneError("the phone number can't contains special characters")
var PhoneWithLetters = NewPhoneError("the phone number cannot contains letters")
var PhoneWithInvalidLength = NewPhoneError("the phone number cannot less 10 digits and more 11 digits")
var PhoneWithoutAreaCode = NewPhoneError("The phone number must have contains the area code")

func NewPhoneError(message string) error {
	return &common_errors.Field{
		Field: "phone",
		Description: common_errors.Description{
			Message:  message,
			Solution: "the telephone number must be as follows: (00) 00000-0000",
		},
	}
}