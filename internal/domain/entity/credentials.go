package entity

import (
	"time"
	
	"stock-controll/internal/domain/failure"
	vo "stock-controll/internal/domain/value_object"
)

type Credential struct {
	id            uint64
	user_ID       int
	password_hash []byte
	password_salt []byte
	reset_token   string
	created_at    time.Time
	updated_at    time.Time
}

func NewCredential( user_ID int, password string) (*Credential, []error) {
	var errorList []error

	if user_ID <= 0 {
		errorList = append(errorList, failure.FieldIsEmpty("user_ID"))
	}

	passwordParsed, err := vo.NewPassword(password)
	if err != nil {
		errorList = append(errorList, err)
	}

	if len(errorList) > 0 {
		return nil, errorList
	}

	return &Credential{
		user_ID:       user_ID,
		password_hash: passwordParsed.GetPasswordHash(),
		password_salt: passwordParsed.GetPasswordSalt(),
	}, nil
}
