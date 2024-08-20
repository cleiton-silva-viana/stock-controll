package failure

var CPNJWithInvalidFormat = &Field{
	Field: "CNPJ",
	Description: Description{
		Message:  "the CNPJ is invalid format",
		Solution: "check if cpnj have the format XX.XXX.XXX/0001-XX",
	},
}

var CPNJWithInvalidCheckerDigits = &Field{
	Field: "CNPJ",
	Description: Description{
		Message:  "the CNPJ is not valid",
		Solution: "check if the cnpj is correct filled",
	},
}

var CPNJWithInvalidCharacters = &Field{
	Field: "CNPJ",
	Description: Description{
		Message:  "the CNPJ contains invalid characters",
		Solution: "check if cpnj have the format XX.XXX.XXX/0001-XX",
	},
}

