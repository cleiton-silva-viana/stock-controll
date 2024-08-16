package contact_repository

import (
	contact_entity "stock-controll/internal/contact/entity"
	"stock-controll/internal/contact/value_object"
)

type ContactRepository interface {
	Save(contact *contact_entity.Contact) error
	Delete(contactID int) error
	UpdateEmail(email *value_object.Email) error
	UpdatePhone(phone *value_object.Phone) error
}
