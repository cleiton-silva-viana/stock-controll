package feature

func (cf *contactFeature) DeleteContact(contactID int) error {
	err := cf.contactRepository.Delete(contactID)
	if err != nil {
		return err
	}
	return nil
}
