package mocks

import (
	"stock-controll/internal/domain/entity"
	"stock-controll/internal/domain/value_object"

	"github.com/stretchr/testify/mock"
)

type ContactRepositoryMock struct {
	mock.Mock
}

func (c *ContactRepositoryMock) Save(contact *entity.Contact) error {
	args := c.MethodCalled("Save")
	return args.Error(0)
}

func (c *ContactRepositoryMock) Delete(contactID int) error {
	args := c.Called(contactID)
	return args.Error(0)
}

func (c *ContactRepositoryMock) UpdateEmail(email *value_object.Email) error {
	emailString := email.GetEmail()
	args := c.Called(emailString)
	return args.Error(0)
}

func (c *ContactRepositoryMock) UpdatePhone(phone *value_object.Phone) error {
	phoneString := phone.GetPhone()
	args := c.Called(phoneString)
	return args.Error(0)
}
