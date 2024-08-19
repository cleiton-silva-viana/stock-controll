package feature

import (
	"errors"

	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/value_object"
)


func (cf *contactFeature) UpdateContact(newContact dto.UpdateContactDTO) error {
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
