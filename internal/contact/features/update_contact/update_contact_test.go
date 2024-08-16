package update_contact

import (
	contact_DTO "stock-controll/internal/contact/DTO"
	"stock-controll/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateContact_UpdateWithSucess(t *testing.T) {
	// Arrange
	contactRepositoryMock := new(mocks.ContactRepositoryMock)

	var email = "joao.batista@gmail.com"
	var phone = "(21) 983325432"

	var tests = []struct {
		name       string
		newContact contact_DTO.UpdateContactDTO
	}{
		{"Update only email contact", contact_DTO.UpdateContactDTO{ID: 1, Email: &email}},
		{"Update only phone contact", contact_DTO.UpdateContactDTO{ID: 1, Phone: &phone}},
		{"Update email and phone", contact_DTO.UpdateContactDTO{ID: 1, Email: &email, Phone: &phone}},
	}

	// Act
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			updateContactFeature := UpdateContact(contactRepositoryMock)
			contactRepositoryMock.On("UpdateEmail", email).Return(nil)
			contactRepositoryMock.On("UpdatePhone", phone).Return(nil)
			err := updateContactFeature.Execute(test.newContact)

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
		newContact contact_DTO.UpdateContactDTO
		expected_error error
	}{
		{name: "update with invalid phone and valid email",
			newContact: contact_DTO.UpdateContactDTO{
				ID:    1,
				Email: &emailValid,
				Phone: &phoneInvalid,
			},
		expected_error: errors.InvalidEmailFormat},
		{name: "update with invalid email and valid phone",
			newContact: contact_DTO.UpdateContactDTO{
				ID:    1,
				Email: &emailInvalid,
				Phone: &phoneValid,
			},
		expected_error: errors.PhoneWithSpecialCharacters},
		{name: "update with invalid email & phone",
			newContact: contact_DTO.UpdateContactDTO{
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
