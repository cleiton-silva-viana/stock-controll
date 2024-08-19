package feature

import (
	"testing"
	
	"stock-controll/internal/application/dto"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateContact_UpdateWithSucess(t *testing.T) {
	// Arrange
	var email = "joao.batista@gmail.com"
	var phone = "(21) 983325432"

	var tests = []struct {
		name       string
		newContact dto.UpdateContactDTO
	}{
		{"Update only email contact", dto.UpdateContactDTO{ID: 1, Email: &email}},
		{"Update only phone contact", dto.UpdateContactDTO{ID: 1, Phone: &phone}},
		{"Update email and phone", dto.UpdateContactDTO{ID: 1, Email: &email, Phone: &phone}},
	}

	// Act
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			updateContactFeature := ContactFeature(contactRepositoryMock)
			contactRepositoryMock.On("UpdateEmail", email).Return(nil)
			contactRepositoryMock.On("UpdatePhone", phone).Return(nil)
			err := updateContactFeature.UpdateContact(test.newContact)

			// Assert
			assert.Nil(t, err)
			if test.newContact.Email != nil {
				contactRepositoryMock.AssertCalled(t, "UpdateEmail", email)
			}
			if test.newContact.Phone != nil {
				contactRepositoryMock.AssertCalled(t, "UpdatePhone", phone)
			}
		})
	}
}

/*
func Test_UpdateContact_WithInvalidParams(t *testing.T) {
	// Arrange
	contactRepositoryMock := new(mocks.ContactRepositoryMock)

	phoneValid := "(21) 983325432"
	phoneInvalid := "(21) 9833254321"
	emailValid := "joao.batista@gmail.com"
	emailInvalid := "joao.batista@gmail.com."

	var tests = []struct {
		name       string
		newContact dto.UpdateContactDTO
		expected_error error
	}{
		{name: "update with invalid phone and valid email",
			newContact: dto.UpdateContactDTO{
				ID:    1,
				Email: &emailValid,
				Phone: &phoneInvalid,
			},
		expected_error: errors.InvalidEmailFormat},
		{name: "update with invalid email and valid phone",
			newContact: dto.UpdateContactDTO{
				ID:    1,
				Email: &emailInvalid,
				Phone: &phoneValid,
			},
		expected_error: errors.PhoneWithSpecialCharacters},
		{name: "update with invalid email & phone",
			newContact: dto.UpdateContactDTO{
				ID:    1,
				Email: &emailInvalid,
				Phone: &phoneInvalid,
			},
		expected_error: errors.InvalidEmailFormat},}

	// Act
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			updateContactFeature := UpdateContact(contactRepositoryMock)
			contactRepositoryMock.On("UpateEmail").Return(nil)
			err := updateContactFeature.Execute(test.newContact)

			// Assert
			assert.Nil(t, err)
			assert.ErrorIs(t, err, test.expected_error)
		})
	}
}
*/
