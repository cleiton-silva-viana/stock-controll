package repository

import (
	"stock-controll/internal/domain/entity"
	vo "stock-controll/internal/domain/value_object"
)

type IContactRepository interface {
	Save(contact entity.Contact) error
	Delete(ID int) error
	UpdateEmail(email vo.Email) error
	UpdatePhone(phone vo.Phone) error
}
