package failure

import "fmt"

func PasswordIsShort(min int) error {
	return &Field{
		Field: "password",
		Description: Description{
			Message:  fmt.Sprintf("the password cannot contain less than %d characters", min),
			Solution: "create a new password with more characters",
		},
	}
}

func PasswordIsLong(max int) error {
	return &Field{
		Field: "password",
		Description: Description{
			Message:  fmt.Sprintf("the password cannot contain more than %d characters", max),
			Solution: "create a new password with less characters",
		},
	}
}

var PasswordNotContainsLowerCases = &Field{
	Field: "password",
	Description: Description{
		Message:  "the password must have contains lower case letters",
		Solution: "add lower cases letters in your password",
	},
}

var PasswordNotContainsUpperCases = &Field{
	Field: "password",
	Description: Description{
		Message:  "the password must have contains upper case letters",
		Solution: "add upper cases letters in your password",
	},
}

var GenerateSaltError = &Field{
	Field: "password",
	Description: Description{
		Message:  "the salt is error in the generate",
		Solution: "",
	},
}
