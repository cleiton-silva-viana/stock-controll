package entity

import (
	"fmt"
	"net/http"
	"stock-controll/internal/domain/failure"
	vo "stock-controll/internal/domain/value_object"
)


type User struct {
	uid         int
	firstName  vo.Name
	lastName   vo.Name
	Gender     vo.Gender
	BirthDate  vo.Date
	Occupation string
}

func NewUser(fistName, gender, birthDate string) (*User, *failure.Fields) {
	var errorsList []error

	// Criar método para criação de UID 
	// *** UID v7 ***
	UID := 0 


	fistNameParsed, err := vo.NameDefaultValidation("first name", fistName)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	genderParsed, err := vo.NewGender(gender)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	dateParsed, err := vo.NewDate(birthDate)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	if dateParsed != nil {
		isOlder := dateParsed.IsOlderThan(18)
		if !isOlder {
			errorsList = append(errorsList, failure.InsufficientAge)
		}
	}

	if len(errorsList) > 0 {
		return nil, &failure.Fields{
			Status:  http.StatusBadRequest,
			ErrList: errorsList,
		}
	}

	return &User{
		uid: UID,
		firstName: *fistNameParsed,
		Gender:    *genderParsed,
		BirthDate: *dateParsed,
	}, nil
}

func (u *User) GetUID() int {
	return u.uid
}

func (u *User) GetName() string {
	nameParsed := fmt.Sprint(u.firstName, u.lastName)
	return nameParsed
}

func (u *User) GetGender() string {
	genderParsed := fmt.Sprint(u.Gender)
	return genderParsed
}
