package repository

import (
	"database/sql"
	"stock-controll/internal/domain/entity"
)

type CredentialRepository struct {
	db *sql.DB
}

func NewSQLCredential(DB *sql.DB) *CredentialRepository {
	return &CredentialRepository{db: DB}
}

func (cr *CredentialRepository) Save(credential entity.Credential) error {
	return nil
}

func (cr *CredentialRepository) Update(credential entity.Credential) error {
	return nil
}

func (cr *CredentialRepository) Delete(userID int) error {
	return nil
}

func (cr *CredentialRepository) GetByID(ID int) (*entity.Contact, error) {
	return nil, nil
}
