package factory

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
)

type ICredentialFactory interface {
	Create(dto.CreateCredentialDTO) (*entity.Credential, []error)
}

type CredentialFactory struct{}

func NewCredentialFactory() *CredentialFactory {
	return &CredentialFactory{}
}

func (c *CredentialFactory) Create(credential dto.CreateCredentialDTO) (*entity.Credential, []error) {
	entity, errors := entity.NewCredential(credential.ID, credential.Password)
    if len(errors) > 0 {
        return nil, errors
    }
    return entity, nil
}
