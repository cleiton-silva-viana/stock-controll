package factory

import (
	contactEntity "stock-controll/internal/domain/entity/contact"
)

type ContactFactory struct{}

func NewcontactFactory() *ContactFactory {
	return &ContactFactory{}
}

func (c *ContactFactory) Create(uuid, email, phone string) (contactEntity.IContact, error) {
	return contactEntity.NewContact(uuid, email, phone)
}
