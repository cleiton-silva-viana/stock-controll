package entity

import (
	"stock-controll/internal/domain/value_object"
)

type User struct {
	ID         int
	Name       value_object.Name
	Gender     value_object.Gender
	BirthDate  value_object.Date
	Occupation string
	Password   value_object.Password
}

func NewUser(name, gender, birthDate, password string) (*User, []error) {
	var errorsList []error

	nameVO, err := value_object.NewName(name, "name", 20)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	genderVO, err := value_object.NewGender(gender)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	dateVO, err := value_object.NewDate(birthDate)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	_, err = dateVO.IsOlderThan(18, true)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	passwordVO, err := value_object.NewPassword(password)
	if err != nil {
		errorsList = append(errorsList, err)
	}

	if len(errorsList) > 0 {
		return nil, errorsList
	}

	return &User{
		Name:      *nameVO,
		Gender:    *genderVO,
		BirthDate: *dateVO,
		Password:  *passwordVO,
	}, nil
}
