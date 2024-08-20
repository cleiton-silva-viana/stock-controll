package failure

var DateWithInvalidFormat = &Field{
	Field: "date",
	Description: Description{
		Message:  "date is on invalid format",
		Solution: "check if date have the format yyyy-mm-dd",
	},
}

var InsufficientAge = &Field{
	Field: "date",
	Description: Description{
		Message: "just for adults",
		Solution: "grow up first",
	},
}