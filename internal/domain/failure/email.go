package failure


var EmailWithInvalidFormat = &Field{
	Field: "email",
	Description: Description{
		Message:  "email is on invalid format",
		Solution: "check if email contains the character '@' and the domain",
	},
}
