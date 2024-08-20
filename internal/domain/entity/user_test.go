package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type user_data struct {
	name      string
	gender    string
	birthDate string
	password  string
}

func Test_NewUser_MustReturnAUser(t *testing.T) {
	// Arrange
	tests := []struct {
		descritipn string
		user_data
	}{
		{descritipn: "create a user of male gender", user_data: user_data{
			name: "Johan", gender: "male", birthDate: "1999-05-05", password: "#$@ASDbcvx6596#",
		}},
		{descritipn: "create a user of female gender", user_data: user_data{
			name: "Ana Maria", gender: "female", birthDate: "2004-08-01", password: "AnaM4r14@#$2004",
		}},
	}

	// act
	for _, test := range tests {
		t.Run(test.descritipn, func(t *testing.T) {

			result, err := NewUser(test.name, test.gender, test.birthDate, test.password)

			// Assert
			assert.Nil(t, err)
			assert.IsType(t, User{}, *result)
		})
	}
}

func Test_NewUser_MustReturnAError(t *testing.T) {
	// Arrange
	tests := []struct {
		descritipn string
		user_data
	}{
		{descritipn: "create a user of invalid gender", user_data: user_data{
			name: "Abrahm Lincon", gender: "cat", birthDate: "1999-05-05", password: "V3e5$#vwEFw6356",
		}},
		{descritipn: "create a user of invalid password", user_data: user_data{
			name: "Donald Trump", gender: "male", birthDate: "1946-06-14", password: "DonaldTrump",
		}},
		{descritipn: "create a user of invalid birthDate", user_data: user_data{
			name: "Barack Obama", gender: "male", birthDate: "2009-06-14", password: "FEW5#$fwe$#656fweERT",
		}},
		{descritipn: "create a user of invalid name", user_data: user_data{
			name: "Joe Biden 99", gender: "male", birthDate: "1942-11-20", password: "@rg4465$#%regf.REGRE",
		}},
	}

	// Act
	for _, test := range tests {
		t.Run(test.descritipn, func(t *testing.T) {
			result, err := NewUser(test.name, test.gender, test.birthDate, test.password)

			// Assert
			assert.Nil(t, result)
			assert.NotEmpty(t, err)
		})
	}
}
