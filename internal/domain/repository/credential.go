package repository

import "stock-controll/internal/domain/entity"

type ICredentialRepository interface {
	Save(credential entity.Credential) error
}
