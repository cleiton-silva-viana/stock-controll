package password

import (
	"crypto/rand"
	"regexp"

	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/presentation/adapter"
)

type Password struct {
	salt []byte
	hash []byte
}

func New(pass string) (*Password, error) {
	err := validatePassword(pass)
	if err != nil {
		return nil, err
	}

	salt, err := saltForPassword(passwordSaltLength)
	if err != nil {
		return nil, err
	}

	hash := hashPassword(pass, salt)

	return &Password{
		salt: []byte(salt),
		hash: hash,
	}, nil
}
func (p *Password) Salt() []byte {
	return p.salt
}

func (p *Password) Hash() []byte {
	return p.hash
}

const (
	passwordMinLength                   = 8
	passwordMaxLength                   = 24
	passwordSaltLength                  = 24
	ErrPasswordMissingLetters           = "ERR_PASSWORD_MISSING_LETTERS"
	ErrPasswordMissingLowercase         = "ERR_PASSWORD_MISSING_LOWERCASE"
	ErrPasswordMissingUppercase         = "ERR_PASSWORD_MISSING_UPPERCASE"
	ErrPasswordMissingNumbers           = "ERR_PASSWORD_MISSING_NUMBERS"
	ErrPasswordMissingSpecialCharacters = "ERR_PASSWORD_MISSING_SPECIAL_CHARACTERS"
	ErrPasswordContainsWhitespace       = "ERR_PASSWORD_CONTAINS_WHITESPACE"
	ErrSaltGenerationFailed             = "ERR_SALT_GENERATION_FAILED"
)

func validatePassword(password string) error {
	return validate.New(
		"password", password,
		validate.IsBlank(),
		validate.IsLengthInRange(passwordMinLength, passwordMaxLength),
		validate.CheckNumbers(validate.Require),
		validate.CheckLetters(validate.Require),
		validate.CheckWithRegex(regexp.MustCompile(`[a-z]`), validate.Require, ErrPasswordMissingLowercase),
		validate.CheckWithRegex(regexp.MustCompile(`[A-Z]`), validate.Require, ErrPasswordMissingUppercase),
		validate.CheckSpecialChars(validate.Require))
}

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func saltForPassword(length uint) ([]byte, error) {
	var salt = make([]byte, length)
	salt, err := generateSalt(int(length))
	if err != nil {
		return nil, &field.FieldError{
			FieldName: "password",
			CodeError: ErrSaltGenerationFailed,
		}
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) []byte {
	return []byte(adapter.NewHasher().GenerateHash([]byte(password), salt))
}
