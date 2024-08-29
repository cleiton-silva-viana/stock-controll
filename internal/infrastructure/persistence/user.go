package persistence

import (
	"database/sql"
	"stock-controll/internal/domain/entity"
)

type ISQLUser interface {
	Save(user entity.User) (int, error)
}

type SQLUser struct {
	db *sql.DB
}

func NewSQLUser(DB *sql.DB) *SQLUser {
	return &SQLUser{db: DB}
}

func (s *SQLUser) Save(user entity.User) (int, error) {
	return -1, nil
}
