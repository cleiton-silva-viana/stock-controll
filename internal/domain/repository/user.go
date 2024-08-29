package repository

import "stock-controll/internal/domain/entity"

type IUserRepository interface {
	Save(user entity.User) (int, error)
}
