package factory

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
)

type IUserFactory interface {
	Create(userData dto.CreateUserDTO) (*entity.User, []error)
}

type UserFactory struct {}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func(u *UserFactory) Create(user dto.CreateUserDTO) (*entity.User, []error) {
	entity, errors := entity.NewUser(user.Name, user.Gender, user.BirthDate)
    if len(errors) > 0 {
        return nil, errors
    }
    return entity, nil
}
