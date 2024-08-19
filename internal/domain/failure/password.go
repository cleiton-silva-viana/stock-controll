package failure

func NewPasswordError(error_message, tip string) error {
	return NewFieldError("password", error_message, tip)
}

var PasswordIsShort = NewPasswordError("the password cannot contain less than 8 characters",
	"create a new password with more chars")

var PasswordIsLong = NewPasswordError("the password cannot contain more than 24 characters",
	"create a new password with more chars")

var PasswordNotContainsLowerCases = NewPasswordError("the password must have contains lower case letters",
	"add lower cases letters in your password")

var PasswordNotContainsUpperCases = NewPasswordError("the password must have contains upper case letters",
	"add upper cases letters in your password")

var PasswordNotContainsNumbers = NewPasswordError("the password must have contains numbers",
	"add numbers in your password")

var PasswordNotContainsSpecialChars = NewPasswordError("the password must have contains special characters",
"add special characters (e.g: $, #, &) in your password")

var GenerateSaltError = NewPasswordError("the salt is error in the generate", "...")
