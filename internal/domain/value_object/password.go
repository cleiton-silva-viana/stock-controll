package value_object

import (
	"crypto/rand"

	"stock-controll/internal/domain/failure"
	"stock-controll/internal/domain/adapter"
)

type Password struct {
	password_salt   []byte
	password_hashed []byte
}

func NewPassword(password string) (*Password, error) {
	const MIN_LENGTH = 8
	const MAX_LENGTH = 24

	if len(password) < MIN_LENGTH {
		return nil, failure.PasswordIsShort
	}

	if len(password) > MAX_LENGTH {
		return nil, failure.PasswordIsLong
	}

	if !ContainsLowerCaseLetters(password) {
		return nil, failure.PasswordNotContainsLowerCases
	}

	if !ContainsUpperCaseLetters(password) {
		return nil, failure.PasswordNotContainsUpperCases
	}

	if !ContainsNumbers(password) {
		return nil, failure.PasswordNotContainsNumbers
	}

	if !ContainsSpecialChars(password) {
		return nil, failure.PasswordNotContainsSpecialChars
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

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
