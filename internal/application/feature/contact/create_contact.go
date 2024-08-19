package feature

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/application/feature"
	"stock-controll/internal/domain/entity"
	"stock-controll/internal/domain/repository"
)

type contactFeature struct {
	contactRepository repository.ContactRepository
}

func ContactFeature(repository repository.ContactRepository) *contactFeature {
	return &contactFeature{contactRepository: repository}
}

func (cf *contactFeature) CreateContact(contactDTO dto.CreateContactDTO) error {
	contact, errorsList := entity.NewContact(contactDTO.Email, contactDTO.Phone)
	if len(errorsList) > 0 {
		return feature.NewFeatureError("failure to create the contact", errorsList)
	}

	err := cf.contactRepository.Save(contact)
	if err != nil {
		return err
	}

	return nil
}
