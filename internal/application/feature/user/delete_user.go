package userfeature

import "stock-controll/internal/infrastructure/repository"

type UserCommand struct {
	repository repository.UserRepository
}

func NewUserRepository(repo repository.UserRepository) *UserCommand {
	return &UserCommand{
		repository: repo,
	}
}

func (u *UserCommand) Delete(userID int) error {
	err := u.repository.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}
