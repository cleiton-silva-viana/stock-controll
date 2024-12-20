package user

import (
	"fmt"
	"time"

	"stock-controll/internal/domain/entity/permission"
	"stock-controll/internal/domain/entity/position"
	"stock-controll/internal/domain/entity/role"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/cpf"
	"stock-controll/internal/domain/valueobject/name"
	"stock-controll/internal/domain/valueobject/uuid"
)

type IUser interface {
	FullName() string
	CPF() string
	BirthDate() time.Time
	Gender() string
	Role() role.Role
	Designation() string
}

type User struct {
	uuid      uuid.UUID
	name      name.Name
	cpf       cpf.CPF
	gender    string
	birthDate time.Time
	role      role.Role
	position  position.Position
	active    bool
}

type UserConfig struct {
	FirstName string
	LastName  string
	CPF       string
	Gender    string
	BirthDate time.Time
}

func New(config UserConfig) (*User, error) {

	var firstNameErr, lastNameErr error

	nameVO, nameErrors := name.New(config.FirstName, config.LastName)

	switch len(nameErrors) {
	case 1:
		firstNameErr = nameErrors[0]
	case 2:
		lastNameErr = nameErrors[1]
	}

	cpfVO, cpfErr := cpf.New(config.CPF)

	genderErr := validateGender(config.Gender)
	birthDateErr := validateBirthDate(config.BirthDate)

	userErrors := entity.Error("user").
		AddValidationError(firstNameErr).
		AddValidationError(lastNameErr).
		AddValidationError(cpfErr).
		AddValidationError(genderErr).
		AddValidationError(birthDateErr)

	if userErrors.HasError() {
		return nil, userErrors
	}

	return &User{
		uuid:   *uuid.New(),
		active: true,
		name:   *nameVO,
		cpf:    *cpfVO,
		// TODO: definir função default
	}, nil
}

func (u *User) UUID() string {
	return u.uuid.String()
}

func (u *User) FullName() string {
	return fmt.Sprint("%s %s", u.name.FirstName(), u.name.LastName())
}

func (u *User) BirthDate() time.Time {
	return u.birthDate
}

func (u *User) CPF() string {
	return u.cpf.CPF()
}

func (u *User) Gender() string {
	return fmt.Sprint(u.gender)
}

func (u *User) Role() role.Role {
	return u.role
}

// TODO: adicionar verificações
func (u *User) SetRole(role role.Role) error {
	u.role = role
	return nil
}

func (u *User) Position() string {
	return u.position.Name()
}

func (u *User) Permissions() []permission.Permission {
	return u.role.Permissions()
}

/*
Pensar a respeito
func (u *User) SetPermission(permission permission.Permission) error {
	return u.role.SetPermission(permission)
}
*/

var (
	minBirthDate           = time.Now().AddDate(-18, 0, 0)
	maxBirthDate           = time.Now().AddDate(-100, 0, 0)
	ErrUserAgeExceedsLimit = "ERR_USER_AGE_EXCEEDS_LIMIT"
	ErrUserageBelowLimit   = "ERR_USER_AGE_BELOW_LIMIT"
)

func validateBirthDate(date time.Time) error {
	return validate.New("birth_date", date,
		validate.IsBeforeThan(minBirthDate, ErrUserageBelowLimit),
		validate.IsAfterThan(maxBirthDate, ErrUserAgeExceedsLimit),
	)
}

var (
	genders = []string{"female", "male"}
)

func validateGender(gender string) error {
	return validate.New("gender", gender,
		validate.IsBlank(),
		validate.IsInEnum(genders),
	)
}
