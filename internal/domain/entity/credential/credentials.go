package credential

import (
	"crypto/rand"
	"time"

	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
	"stock-controll/internal/presentation/adapter"
)

type ICredential interface {
	GetUUID() string
	GetPasswordHash() []byte
	GetPasswordSalt() []byte
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type credential struct {
	uuid string
	passwordHash []byte
	passwordSalt []byte
	resetToken   string
	createdAt    time.Time
	updatedAt    time.Time
}

func NewCredential(uuid, password string) (ICredential, validationError.IValidationError) {
	var credentialError = validationError.NewValidationError("credential")
	var c = credential{}

	credentialError.AddValidationError(c.SetPassword(password))

	isValid := common.IsValidUUUID(uuid)
	if !isValid {
		credentialError.AddValidationError(&validation.FieldError{
			FieldName: "uuid",
			CodeErrors: []string{string(validation.ErrUnknown)},
		})
	}

	if credentialError.HasError() {
		return nil, credentialError
	}

	c.createdAt = time.Now()

	return &c, nil
}

const (
	passwordMinLength = 8
	passwordMaxLength = 24
	passwordSaltLength = 24
)

func (c *credential) GetUUID() string {
	return c.uuid
}

func (c *credential) SetPassword(password string) *validation.FieldError {
	err := c.validatePassword(password)
	if err != nil {
		return err
	}

	salt, err := c.saltForPassword(passwordSaltLength)
	if err != nil {
		return err
	}

	c.passwordSalt = []byte(salt)
	c.passwordHash = c.hashPassword(password, salt)
	c.updatedAt = time.Now()
	return nil
}

func (c *credential) validatePassword(password string) *validation.FieldError {
	return validation.Validate("password", password,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(passwordMinLength, passwordMaxLength, validation.ErrUnknown),
		validation.CheckNumbers(validation.Require, validation.ErrUnknown),
		validation.CheckLetters(validation.Require, validation.ErrUnknown),
		validation.CheckLowerCaseLetters(validation.Require, validation.ErrUnknown),
		validation.CheckUpperCaseLetters(validation.Require, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Require, validation.ErrUnknown))
}

func (c *credential) saltForPassword(length uint) ([]byte, *validation.FieldError){
	var salt = make([]byte, length) 
	salt, err := generateSalt(int(length))
	if err != nil {
		return nil, &validation.FieldError{
			FieldName:  "password",
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}
	return salt, nil
}

func (c credential) hashPassword(password string, salt []byte) []byte {
	return []byte(adapter.NewHasher().GenerateHash([]byte(password), salt))
}

func (c *credential) GetPasswordHash() []byte {
	return c.passwordHash
}

func (c *credential) GetPasswordSalt() []byte {
	return c.passwordSalt
}

func (c *credential) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *credential) GetUpdatedAt() time.Time {
	return c.updatedAt
}


// TODO: mover para método da senha
func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
