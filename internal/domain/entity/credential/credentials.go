package credential

import (
	"crypto/rand"
	"regexp"
	"time"

	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/presentation/adapter"
)

const (
	ErrPaswordMissingLetters           = "ERR_PASSWORD_MISSING_LETTERS"
	ErrPaswordMissingLowercase         = "ERR_PASSWORD_MISSING_LOWERCASE"
	ErrPaswordMissingUppercase         = "ERR_PASSWORD_MISSING_UPPERCASE"
	ErrPaswordMissingNumbers           = "ERR_PASSWORD_MISSING_NUMBERS"
	ErrPaswordMissingSpecialCharacters = "ERR_PASSWORD_MISSING_SPECIAL_CHARACTERS"
	ErrPaswordContainsWhitespace       = "ERR_PASSWORD_CONTAINS_WHITESPACE"
	ErrPaswordEqualsUsername           = "ERR_PASSWORD_EQUALS_USERNAME"
	ErrPaswordEqualsPrevious           = "ERR_PASSWORD_EQUALS_PREVIOUS"
	ErrSaltGenerationFailed            = "ERR_SALT_GENERATION_FAILED"
)

type ICredential interface {
	GetUUID() string
	GetPasswordHash() []byte
	GetPasswordSalt() []byte
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type Credential struct {
	uuid         string
	passwordHash []byte
	passwordSalt []byte
	resetToken   string
	createdAt    time.Time
	updatedAt    time.Time
}

func NewCredential(userUUID, password string) (ICredential, error) {
	var credentialError = validationerrors.New("credential")
	var c = Credential{}

	credentialError.
		AddValidationError(c.SetPassword(password)).
		AddValidationError(uuid.IsValid("user_uuid", userUUID))

	if credentialError.HasError() {
		return nil, credentialError
	}

	c.createdAt = time.Now()
	return &c, nil
}

const (
	passwordMinLength  = 8
	passwordMaxLength  = 24
	passwordSaltLength = 24
)

func (c *Credential) GetUUID() string {
	return c.uuid
}

func (c *Credential) SetPassword(password string) error {
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

func (c *Credential) validatePassword(password string) error {
	return validate.New("password", password,
		validate.IsBlank(),
		validate.IsLengthInRange(passwordMinLength, passwordMaxLength),
		validate.CheckNumbers(validate.Require),
		validate.CheckLetters(validate.Require),
		validate.CheckWithRegex(regexp.MustCompile(`[a-z]`), validate.Require, ErrPaswordMissingLowercase),
		validate.CheckWithRegex(regexp.MustCompile(`[a-z]`), validate.Require, ErrPaswordMissingUppercase),
		// aDICIONAR VALIDAÇÃO PARA CHECAR S ESENHA TEM CARACTERES ESPECIAIS
		// pASSAR ERRO PERSONALIZADO AO CASO
		validate.CheckSpecialChars(validate.Require))
}

func (c *Credential) saltForPassword(length uint) ([]byte, error) {
	var salt = make([]byte, length)
	salt, err := generateSalt(int(length))
	if err != nil {
		return nil, &validate.FieldError{
			FieldName: "password",
			CodeError: ErrSaltGenerationFailed,
		}
	}
	return salt, nil
}

func (c *Credential) hashPassword(password string, salt []byte) []byte {
	return []byte(adapter.NewHasher().GenerateHash([]byte(password), salt))
}

func (c *Credential) GetPasswordHash() []byte {
	return c.passwordHash
}

func (c *Credential) GetPasswordSalt() []byte {
	return c.passwordSalt
}

func (c *Credential) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Credential) GetUpdatedAt() time.Time {
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
