package create_contact

import (
	feature_errors "stock-controll/internal/common/features/errors"
	contact_DTO "stock-controll/internal/contact/DTO"
	contact_entity "stock-controll/internal/contact/entity"
	contact_repository "stock-controll/internal/contact/repository"
)

type ContactFeature struct {
	contactRepository contact_repository.ContactRepository
}

func CreateContact(repository contact_repository.ContactRepository) *ContactFeature  {
	return &ContactFeature{contactRepository: repository}
}

func(cf *ContactFeature)  Execute(contactDTO contact_DTO.CreateContactDTO) error {
	contact, errorsList := contact_entity.NewContact(contactDTO.Email, contactDTO.Phone)
	if len(errorsList) > 0 {
		return feature_errors.NewFeatureError("failure to create the contact", errorsList)
	}

	err := cf.contactRepository.Save(contact)
	if err != nil {
		return err
	} 

	return nil
}
