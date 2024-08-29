package entity

import (
	vo "stock-controll/internal/domain/value_object"
)

type Contact struct {
	id    int
	phone vo.Phone
	email vo.Email
}

func NewContact(ID int, email, phone string) (*Contact, []error) {
	var errorsList []error

	// Validar o ID

	userEmail, err := vo.NewEmail(email)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	userPhone, err := vo.NewPhone(phone)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	if len(errorsList) > 0 {
		return nil, errorsList
	}

	return &Contact{
		id: ID,
		email: *userEmail,
		phone: *userPhone,
	}, nil
}
