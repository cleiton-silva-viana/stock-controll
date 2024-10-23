package userfeature

import (
	"stock-controll/internal/application/dto"
)

func (uf *UserFeature) Update(dto dto.UserDTO) error {
	user, err := uf.userFactory.Create(dto)
	if err != nil {
		return err
	}

	err = uf.userRepository.Update(*user)
	return err
}
