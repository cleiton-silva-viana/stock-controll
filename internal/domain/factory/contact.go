package factory

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
)

type IContactFactory interface {
	Create(dto.CreateContactDTO) (*entity.Contact, error)
}

type ContactFactory struct{}

func NewcontactFactory() *ContactFactory {
	return &ContactFactory{}
}

func (c *ContactFactory) Create(contact dto.CreateContactDTO) (*entity.Contact, error) {
	entity, err := entity.NewContact(contact.UID, contact.Email, contact.Phone)
    if err != nil {
		return nil, err
	}
    return entity, nil
}