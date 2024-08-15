package errors

import common_errors "stock-controll/internal/common/entity/errors"

var InvalidEmailFormat = &common_errors.Field{
	Field: "email",
	Description: common_errors.Description{
		Message: "email is on invalid format",
		Solution: "check if email contains the character '@' and the domain",
	},
}
