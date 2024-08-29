package valueobject

import (
	"crypto/rand"

	"stock-controll/internal/domain/failure"
	"stock-controll/internal/domain/validation"
	"stock-controll/internal/presentation/adapter"
)

type Password struct {
	password_salt   []byte
	password_hashed []byte
}

func NewPassword(password string) (*Password, error) {
	const minLength = 8
	const maxLength = 24

	if len(password) < minLength {
		return nil, failure.PasswordIsShort(minLength)
	}

	if len(password) > maxLength {
		return nil, failure.PasswordIsLong(maxLength)
	}

	if !validation.ContainsLowerCaseLetters(password) {
		return nil, failure.PasswordNotContainsLowerCases
	}

	if !validation.ContainsUpperCaseLetters(password) {
		return nil, failure.PasswordNotContainsUpperCases
	}

	if !validation.ContainsNumbers(password) {
		return nil, failure.FieldWithoutNumber("password")
	}

	if !validation.ContainsSpecialChars(password) {
		return nil, failure.FieldWithoutSpecialChars("password")
	}

	salt, err := generateSalt(24)
	if err != nil {
		return nil, failure.GenerateSaltError
	}

	passwordHarshed := adapter.NewHasher().GenerateHash([]byte(password), salt)

	return &Password{
		password_salt:   salt,
		password_hashed: []byte(passwordHarshed),
	}, nil
}

func (p *Password) GetPasswordHash() []byte {
	return p.password_hashed
}

func (p *Password) GetPasswordSalt() []byte {
	return p.password_salt
}

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
