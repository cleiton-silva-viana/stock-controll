package persistence

import (
	"database/sql"
	"stock-controll/internal/domain/entity"
)

type ISQLCredential interface {
	Save(credential entity.Credential) error
}

type SQLCredential struct {
	db *sql.DB
}

func NewSQLCredential(DB *sql.DB) *SQLCredential {
	return &SQLCredential{db: DB}
}

func (c *SQLCredential) Save(credential entity.Credential) error {
	//
	return nil
}
