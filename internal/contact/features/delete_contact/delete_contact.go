package delete_contact

import contact_repository "stock-controll/internal/contact/repository"

type ContactFeature struct {
	contactRepository contact_repository.ContactRepository
}

func DeleteContact(repository contact_repository.ContactRepository) *ContactFeature {
	return &ContactFeature{
		contactRepository: repository,
	}
}

func (cf *ContactFeature) Execute(contactID int) error {
	err := cf.contactRepository.Delete(contactID)
	if err != nil {
		return err
	}
	return nil
}