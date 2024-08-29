package persistence

import (
	"database/sql"
	"stock-controll/internal/domain/entity"
)

type ISQLContact interface {
	Save(contact entity.Contact) error
}

type SQLContact struct {
	db *sql.DB
}

func NewSQLContact(DB *sql.DB) *SQLContact {
	return &SQLContact{db: DB}
}

func (c *SQLContact) Save(contact entity.Contact) error {
	return nil
}
