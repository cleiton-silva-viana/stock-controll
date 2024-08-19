package repository

import (
	"stock-controll/internal/domain/entity"
	"stock-controll/internal/domain/value_object"
)

type ContactRepository interface {
	Save(contact *entity.Contact) error
	Delete(contactID int) error
	UpdateEmail(email *value_object.Email) error
	UpdatePhone(phone *value_object.Phone) error
}
