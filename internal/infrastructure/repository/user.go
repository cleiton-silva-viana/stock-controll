package repository

import (
	"database/sql"
	"stock-controll/internal/domain/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{db: DB}
}

func (ur *UserRepository) Save(user entity.User) error {
	return nil
}

func (ur *UserRepository) Update(user entity.User) error {
	return nil
}

func (ur *UserRepository) Delete(userID int) error {
	return nil
}

func (ur *UserRepository) GetByID(userID int) error {
	return nil
}