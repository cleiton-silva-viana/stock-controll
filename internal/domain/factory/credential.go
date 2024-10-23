package factory

import (
	credentialEntity "stock-controll/internal/domain/entity/credential"
)

type CredentialFactory struct{}

func NewCredentialFactory() *CredentialFactory {
	return &CredentialFactory{}
}

func (c *CredentialFactory) Create(uuid, password string) (credentialEntity.ICredential, error) {
	return credentialEntity.NewCredential(uuid, password)
}
