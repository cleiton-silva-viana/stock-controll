package update_contact

import (
	"errors"
	contact_DTO "stock-controll/internal/contact/DTO"
	contact_repository "stock-controll/internal/contact/repository"
	"stock-controll/internal/contact/value_object"
)

type ContactFeature struct {
	contactRepository contact_repository.ContactRepository
}

func UpdateContact(repository contact_repository.ContactRepository) *ContactFeature {
	return &ContactFeature{contactRepository: repository}
}

func (cf *ContactFeature) Execute(newContact contact_DTO.UpdateContactDTO) error {
	// Verificar as permissões do usuário

	if newContact.Email == nil && newContact.Phone == nil {
		return errors.New("all fields are empyt")
	}

	if newContact.Email != nil {
		email, err := value_object.NewEmail(*newContact.Email)
		if err != nil {
			return err
		}

		err = cf.contactRepository.UpdateEmail(email)
		if err != nil {
			return err
		}
	}
	
	if newContact.Phone != nil {
		phone, err := value_object.NewPhone(*newContact.Phone)
		if err != nil {
			return err
		}

		err = cf.contactRepository.UpdatePhone(phone)
		if err != nil {
			return err
		}
	}

	return nil
}
