package factory

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
)

type IContactFactory interface {
	Create(dto.CreateContactDTO) (*entity.Contact, []error)
}

type ContactFactory struct{}

func NewcontactFactory() *ContactFactory {
	return &ContactFactory{}
}

func (c *ContactFactory) Create(contact dto.CreateContactDTO) (*entity.Contact, []error) {
	entity, errors := entity.NewContact(contact.ID, contact.Email, contact.Phone)
    if len(errors) > 0 {
        return nil, errors
    }
    return entity, nil
}