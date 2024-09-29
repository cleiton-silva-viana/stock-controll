package factory

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
)

type ICredentialFactory interface {
	Create(dto.CreateCredentialDTO) (*entity.Credential, error)
}

type CredentialFactory struct{}

func NewCredentialFactory() *CredentialFactory {
	return &CredentialFactory{}
}

func (c *CredentialFactory) Create(credential dto.CreateCredentialDTO) (*entity.Credential, error) {
	entity, err := entity.NewCredential(credential.UID, credential.Password)
    if err != nil {
        return nil, err
    }
    return entity, nil
}
