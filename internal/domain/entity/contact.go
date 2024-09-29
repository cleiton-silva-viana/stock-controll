package entity

import (
	"net/http"
	"stock-controll/internal/domain/failure"
	vo "stock-controll/internal/domain/value_object"
)

type Contact struct {
	uid    int
	phone vo.Phone
	email vo.Email
}

func NewContact(uid int, email, phone string) (*Contact, *failure.Fields) {
	var errorsList []error

	userEmail, err := vo.NewEmail(email)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	userPhone, err := vo.NewPhone(phone)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	if len(errorsList) > 0 {
		return nil, &failure.Fields{
			Status: http.StatusBadRequest,
			ErrList: errorsList,
		}
	}

	return &Contact{
		uid:    uid,
		email: *userEmail,
		phone: *userPhone,
	}, nil
}
