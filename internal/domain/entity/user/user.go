package user

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/entity/permission"
	"stock-controll/internal/domain/entity/role"
	"stock-controll/internal/domain/validation"
)

type IUser interface {
	GetUUID() string
	GetFirstName() string
	GetLastName() string
	GetFullName() string
	GetCPF() string
	GetBirthDate() time.Time
	GetGender() string
	GetRole() role.Role
	GetDesignation() string
}

// transformar em VO
type designation string

const (
	Admin         designation = "admin"
	Developer     designation = "developer"
	Client        designation = "client"
	Manager       designation = "manager"
	Seller        designation = "seller"
	Buyer         designation = "buyer"
	Conference    designation = "conference"
	HumanResource designation = "rh"
)

type user struct {
	uuid        string
	fullName    fullName
	cpf         string
	gender      string
	birthDate   time.Time
	role        role.Role
	designation designation
	active      bool
}

type fullName struct {
	firstName string
	lastName  string
}

func (fn *fullName) GetFirstName() string {
	return fn.firstName
}

func (fn *fullName) GetLastName() string {
	return fn.lastName
}

func (fn *fullName) GetFullName() string {
	return fmt.Sprintf("%s %s", fn.firstName, fn.lastName)
}

// Tornar construtor privado, haja visto que tal estrutura não deve ser instanciada diretamente
func NewUser(firstName, lastName, cpf, gender string, birthDate time.Time) (*user, validationError.IValidationError) {
	var userError = validationError.NewValidationError("user")
	var userInstance = user{
		uuid: common.GenerateUUID(),
		active: true,
		// TODO: definir função default
	}

	userError.
		AddValidationError(userInstance.SetFirstName(firstName)).
		AddValidationError(userInstance.SetLastName(lastName)).
		AddValidationError(userInstance.SetGender(gender)).
		AddValidationError(userInstance.SetCPF(cpf)).
		AddValidationError(userInstance.SetBirthDate(birthDate))

	if userError.HasError() {
		return nil, userError
	}

	return &userInstance, nil
}

func (u *user) GetUUID() string {
	return u.uuid
}

func (u *user) GetFirstName() string {
	return u.fullName.GetFirstName()
}

const (
	userNameMinLength = 3
	userNameMaxLength = 24
)

func (u *user) SetFirstName(name string) *validation.FieldError {
	err := validation.Validate("first_name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(userNameMinLength, userNameMaxLength, validation.ErrUnknown),
		validation.CheckNumbers(validation.Disallow, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
	if err == nil {
		u.fullName.firstName = strings.ToLower(name)
	}
	return err
}

func (u *user) GetLastName() string {
	return u.fullName.GetLastName()
}

func (u *user) SetLastName(name string) *validation.FieldError {
	err := validation.Validate("last_name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(userNameMinLength, userNameMaxLength, validation.ErrUnknown),
		validation.CheckNumbers(validation.Disallow, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
	if err == nil {
		u.fullName.lastName = strings.ToLower(name)
	}
	return err
}

func (u *user) GetFullName() string {
	return u.fullName.GetFullName()
}

func (u *user) GetCPF() string {
	return u.cpf
}

func (u *user) SetCPF(cpf string) *validation.FieldError {
	re := `^\d{3}\.\d{3}\.\d{3}\-\d{2}$`
	err := validation.Validate[string]("cpf", cpf,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthEqualTo(14, validation.ErrUnknown),
		validation.IsFormatValid(regexp.MustCompile(re), validation.ErrUnknown),
	)
	if err == nil {
		u.cpf = cpf
	}
	return err
}

func (u *user) GetBirthDate() time.Time {
	return u.birthDate
}

var (
	userMinimumBirthDate = time.Now().AddDate(-18, 0, 0)
	userMaximumBirthDate = time.Now().AddDate(-100, 0, 0)
)

func (u *user) SetBirthDate(date time.Time) *validation.FieldError {
	err := validation.Validate("birth_date", date,
		validation.IsFutureDate(validation.ErrUnknown),
		validation.IsBeforeThan(userMinimumBirthDate, validation.ErrUnknown),
		validation.IsAfterThan(userMaximumBirthDate, validation.ErrUnknown),
	)
	if err == nil {
		u.birthDate = date
	}
	return err
}

func (u *user) GetGender() string {
	return fmt.Sprint(u.gender)
}

var (
	genders = []string{"female", "male"}
)

func (u *user) SetGender(gender string) *validation.FieldError {
	err := validation.Validate("gender", gender,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsValueInRange(genders, true, validation.ErrUnknown),
	)
	if err == nil {
		u.gender = strings.ToLower(gender)
	}
	return err
}

func (u *user) GetRole() role.Role {
	return u.role
}

// TODO: adicionar verificações
func (u *user) SetRole(newRole role.Role) error {
	u.role = newRole
	return nil
}

func (u *user) GetDesignation() string {
	return string(u.designation)
}

// TODO: Adicionar verificações
func (u *user) SetDesignation(designation string) error {
	u.designation = u.designation
	return nil
}

func (u *user) GetPermissions() []permission.Permission {
	return u.role.GetPermissions()
}

func (u *user) SetPermission(newPermission permission.Permission) *validation.FieldError {
	return u.role.SetPermissions(newPermission)
}

type UserBuilder struct {
	firstName string
	lastName  string
	cpf       string
	gender    string
	birthDate time.Time
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (ub *UserBuilder) SetName(firstName, lastName string) *UserBuilder {
	ub.firstName = firstName
	ub.lastName = lastName
	return ub
}

func (ub *UserBuilder) SetCPF(cpf string) *UserBuilder {
	ub.cpf = cpf
	return ub
}

func (ub *UserBuilder) SetGender(gender string) *UserBuilder {
	ub.gender = gender
	return ub
}

func (ub *UserBuilder) SetBirthDate(birthDate time.Time) *UserBuilder {
	ub.birthDate = birthDate
	return ub
}

func (ub *UserBuilder) Build() (IUser, validationError.IValidationError) {
	return NewUser(ub.firstName, ub.lastName, ub.cpf, ub.gender, ub.birthDate)
}
