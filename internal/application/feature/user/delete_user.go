package userfeature

func (uf *UserFeature) Delete(userUID string) error {
	
	err := uf.uow.Begin()
	if err != nil {
		return err
	}

	err = uf.userRepository.Delete(userUID)
	if err != nil {
		return err
	}

	err = uf.credentialRepository.Delete(userUID)
	if err != nil {
		uf.uow.Rollback()
		return err
	}

	err = uf.contactRepository.Delete(userUID)
	if err != nil {
		uf.uow.Rollback()
		return err
	}

	uf.uow.Commit()
	return nil
}
