package factory

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
)

type IUserFactory interface {
	Create(userData dto.UserDTO) (*entity.User, error)
}

type UserFactory struct {}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func(u *UserFactory) Create(user dto.UserDTO) (*entity.User, error) {
	entity, err := entity.NewUser(user.Name, user.Gender, user.BirthDate)
    if err != nil {
        return nil, err
    }
    return entity, nil
}
