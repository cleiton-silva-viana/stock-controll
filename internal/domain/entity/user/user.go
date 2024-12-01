package user

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"stock-controll/internal/domain/entity/permission"
	"stock-controll/internal/domain/entity/position"
	"stock-controll/internal/domain/entity/role"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
	validationError "stock-controll/internal/domain/services/error"
)

type IUser interface {
	UUID() string
	FirstName() string
	LastName() string
	FullName() string
	CPF() string
	BirthDate() time.Time
	Gender() string
	Role() role.Role
	Designation() string
}

type User struct {
	uuid      string
	fullName  FullName
	cpf       string
	gender    string
	birthDate time.Time
	role      role.Role
	position  position.Position
	active    bool
}

type FullName struct {
	firstName string
	lastName  string
}

func (fn *FullName) FirstName() string {
	return fn.firstName
}

func (fn *FullName) LastName() string {
	return fn.lastName
}

func (fn *FullName) FullName() string {
	return fmt.Sprintf("%s %s", fn.firstName, fn.lastName)
}

type UserConfig struct {
	FirstName string
	LastName  string
	CPF       string
	Gender    string
	BirthDate time.Time
}

// Tornar construtor privado, haja visto que tal estrutura não deve ser instanciada diretamente
func New(config UserConfig) (*User, error) {
	var userInstance = User{
		uuid:   uuid.New(),
		active: true,
		// TODO: definir função default
	}

	var userError = validationError.New("user").
		AddValidationError(userInstance.SetFirstName(config.FirstName)).
		AddValidationError(userInstance.SetLastName(config.LastName)).
		AddValidationError(userInstance.SetGender(config.Gender)).
		AddValidationError(userInstance.SetCPF(config.CPF)).
		AddValidationError(userInstance.SetBirthDate(config.BirthDate))

	if userError.HasError() {
		return nil, userError
	}

	return &userInstance, nil
}

func (u *User) UUID() string {
	return u.uuid
}

func (u *User) FirstName() string {
	return u.fullName.FirstName()
}

const (
	userNameMinLength = 3
	userNameMaxLength = 24
)

func (u *User) SetFirstName(name string) error {
	err := validate.New("first_name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(userNameMinLength, userNameMaxLength),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err == nil {
		u.fullName.firstName = strings.ToLower(name)
	}
	return err
}

func (u *User) LastName() string {
	return u.fullName.LastName()
}

func (u *User) SetLastName(name string) error {
	err := validate.New("last_name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(userNameMinLength, userNameMaxLength),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err == nil {
		u.fullName.lastName = strings.ToLower(name)
	}
	return err
}

func (u *User) FullName() string {
	return u.fullName.FullName()
}

func (u *User) CPF() string {
	return u.cpf
}

const (
	cpfLength               = 14
	ErrCPFWithInvalidFormat = "ERR_CPF_WITH_INVALID_FORMAR"
)

func (u *User) SetCPF(cpf string) error {
	re := `^\d{3}\.\d{3}\.\d{3}\-\d{2}$`
	err := validate.New[string]("cpf", cpf,
		validate.IsBlank(),
		validate.IsLengthEqualTo(cpfLength),
		validate.IsFormatValid(regexp.MustCompile(re), ErrCPFWithInvalidFormat),
	)
	if err == nil {
		u.cpf = cpf
	}
	return err
}

func (u *User) BirthDate() time.Time {
	return u.birthDate
}

var (
	userMinimumBirthDate   = time.Now().AddDate(-18, 0, 0)
	userMaximumBirthDate   = time.Now().AddDate(-100, 0, 0)
	ErrUserAgeExceedsLimit = "ERR_USER_AGE_EXCEEDS_LIMIT"
	ErrUserageBelowLimit   = "ERR_USER_AGE_BELOW_LIMIT"
)

func (u *User) SetBirthDate(date time.Time) error {
	err := validate.New("birth_date", date,
		validate.IsBeforeThan(userMinimumBirthDate, ErrUserageBelowLimit),
		validate.IsAfterThan(userMaximumBirthDate, ErrUserAgeExceedsLimit),
	)
	if err == nil {
		u.birthDate = date
	}
	return err
}

func (u *User) Gender() string {
	return fmt.Sprint(u.gender)
}

var (
	genders = []string{"female", "male"}
)

func (u *User) SetGender(gender string) error {
	err := validate.New("gender", gender,
		validate.IsBlank(),
		validate.IsInEnum(genders),
	)
	if err == nil {
		u.gender = strings.ToLower(gender)
	}
	return err
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

func (u *User) SetPermission(permission permission.Permission) error {
	return u.role.SetPermission(permission)
}
