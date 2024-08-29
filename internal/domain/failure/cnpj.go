package failure

var CPNJWithInvalidCheckerDigits = &Field{
	Field: "CNPJ",
	Description: Description{
		Message:  "the CNPJ is not valid",
		Solution: "check if the cnpj is correct filled",
	},
}
