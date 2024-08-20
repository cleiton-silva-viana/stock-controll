package entity

import "stock-controll/internal/domain/value_object"

type Contact struct {
	Id    int
	Phone value_object.Phone
	Email value_object.Email
}

func NewContact(email, phone string) (*Contact, []error) {
	var errorsList []error

	userEmail, err := value_object.NewEmail(email)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	userPhone, err := value_object.NewPhone(phone)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	if len(errorsList) > 0 {
		return nil, errorsList
	}

	return &Contact{
		Id:    0,
		Email: *userEmail,
		Phone: *userPhone,
	}, nil
}
