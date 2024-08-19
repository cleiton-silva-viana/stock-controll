package failure

var NameIsShort = NewFieldError("name",
	"the name cannot contain less than 3 letters",
	"Check if name is correct")

var NameIsLong = NewFieldError("name",
	"the name cannot contain more than 20 chars",
	"Check if name is correct")

var NameWithInvalidChars = NewFieldError("name",
	"the name cannot contains number or special characters",
	"Check if name is correct")

var NameIsEmpty = NewFieldError("name",
	"the name is empty",
	"Check if name is correct")
