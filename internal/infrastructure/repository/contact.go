package repository

import (
	"database/sql"
	"stock-controll/internal/domain/entity"
)


type ContactRepository struct {
	db *sql.DB
}

func NewSQLContact(DB *sql.DB) *ContactRepository {
	return &ContactRepository{db: DB}
}

func (cr *ContactRepository) Save(contact entity.Contact) error {
	return nil
}

func (cr *ContactRepository) Update(contact entity.Contact) error {
	return nil
}

func (cr *ContactRepository) Delete(userUID int) error {
	return nil
}

func (cr *ContactRepository) GetByID(ID int) (*entity.Contact, error) {
	return nil, nil
}
