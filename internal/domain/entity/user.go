package entity

import (
	"fmt"
	"stock-controll/internal/domain/failure"
	vo "stock-controll/internal/domain/value_object"
)

type User struct {
	id         int
	// firstName vo.Name
	// lastName vo.Name
	Name       vo.Name
	Gender     vo.Gender
	BirthDate  vo.Date
	Occupation string
}

func NewUser(name, gender, birthDate string) (*User, []error) {
	var errorsList []error

	nameParsed, err := vo.NameDefaultValidation("name", name)
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
		return nil, errorsList
	}

	return &User{
		Name:      *nameParsed,
		Gender:    *genderParsed,
		BirthDate: *dateParsed,
	}, nil
}

func (u *User) GetName() string {
	nameParsed := fmt.Sprint(u.Name)
	return nameParsed
}

func (u *User) GetGender() string { 
	genderParsed := fmt.Sprint(u.Gender)
	return genderParsed
}