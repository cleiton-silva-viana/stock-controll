package factory

import (
	"stock-controll/internal/application/dto"
	entity "stock-controll/internal/domain/entity/user"
	"time"
)

type IUserFactory interface {
	CreateUser(dto dto.CreateUserRequestDTO) (entity.IUser, error)
}

type UserFactory struct{}

func (uf *UserFactory) createUser(firstName, lastName, gender, cpf string, birthDate time.Time) (entity.IUser, error) {
	return entity.NewUserBuilder().
		SetCPF(cpf).
		SetName(firstName, lastName).
		SetGender(gender).
		SetBirthDate(birthDate).
		Build()
}
