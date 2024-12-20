package credential

import (
	"time"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/valueobject/password"
	"stock-controll/internal/domain/valueobject/uuid"
)

const (
	ErrPaswordEqualsUsername = "ERR_PASSWORD_EQUALS_USERNAME"
	ErrPaswordEqualsPrevious = "ERR_PASSWORD_EQUALS_PREVIOUS"
)

type ICredential interface {
	UUID() string
	PasswordHash() []byte
	PasswordSalt() []byte
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

type Credential struct {
	uuid uuid.UUID
	pass password.Password
	createdAt time.Time
	updatedAt time.Time
}

func New(userUUID, pass string) (*Credential, error) {
	var credentialError = entity.Error("credential")

	parsedUUID, uuidErr := uuid.Parse("user_uuid", userUUID)
	passwordVO, passErr := password.New(pass)

	credentialError.
		AddValidationError(uuidErr).
		AddValidationError(passErr)

	if credentialError.HasError() {
		return nil, credentialError
	}

	credential := &Credential{
		uuid:      *parsedUUID,
		pass:  *passwordVO,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	return credential, nil
}

func (c *Credential) UUID() string {
	return c.uuid.String()
}

func (c *Credential) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Credential) UpdatedAt() time.Time {
	return c.updatedAt
}

// Lógica para verificar se a senha não é igual a senha anterior
func (c *Credential) UpdatePassword(pass string) error {
	newPass, err := password.New(pass)
	if err == nil {
		c.pass = *newPass
	}
	return err
}
