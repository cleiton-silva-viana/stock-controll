package userfeature

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
)

func (u *UserCommand) Update(dto dto.UserDTO) error {
	user, err := entity.NewUser(dto.Name, dto.Gender, dto.BirthDate)
	if err != nil {
		return err
	}
	return u.repository.Update(*user)
}
