package permission

// TODO: Usar entityError

import (
	"fmt"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
	"sync"
)

type Permission struct {
	id   int
	name string
}

func NewPermission(id int, name string) (*Permission, validationError.IValidationError) {
	var permissionError = validationError.NewValidationError("permission")
	var permissionInstance = Permission{}

	permissionError.
		AddValidationError(permissionInstance.SetID(id)).
		AddValidationError(permissionInstance.SetName(name))

	if permissionError.HasError() {
		return nil, permissionError
	}
	return &permissionInstance, nil
}

func (p *Permission) String() string {
	return fmt.Sprintf("Permission(ID: %d, Name: %s)", p.id, p.name)
}

func (p *Permission) GetID() int {
	return p.id
}

const (
	minIDNumberForPermission = 1
	maxIDNumberForPermission = 1000
)

func (p *Permission) SetID(id int) *validation.FieldError {
	err := validation.Validate[int]("id", id,
		validation.IsInRange(
			minIDNumberForPermission,
			maxIDNumberForPermission,
			validation.ErrUnknown),
	)
	if err == nil {
		p.id = id
	}
	return err
}

func (p *Permission) GetName() string {
	return p.name
}

const (
	minNameLengthForPermission = 6
	maxNameLengthForPermission = 24
)

func (p *Permission) SetName(name string) *validation.FieldError {
	err := validation.Validate[string]("name", name,
		validation.IsBlank(
			validation.ErrUnknown),
		validation.IsLengthInRange(
			minNameLengthForPermission,
			maxNameLengthForPermission,
			validation.ErrUnknown),
		validation.CheckNumbers(
			validation.Disallow,
			validation.ErrUnknown),
		validation.CheckSpecialChars(
			validation.Disallow,
			validation.ErrUnknown),
	)
	if err == nil {
		p.name = name
	}
	return err
}

func (p *Permission) Equals(other *Permission) bool {
	return p.id == other.id && p.name == other.name
}

// mover

var instance *permissionManager
var once sync.Once

type permissionManager struct {
	permissions map[int]*Permission
}

func newPermissionManager() *permissionManager {
	return &permissionManager{
		permissions: make(map[int]*Permission),
	}
}

func GetPermissionManager() *permissionManager {
	once.Do(func() {
		instance = newPermissionManager()
	})
	return instance
}

// TODO: Refatorar erro
func (manager *permissionManager) AddPermission(ID int, name string) validationError.IValidationError {
	var err = validationError.NewValidationError("permission")
	_, exists := manager.permissions[ID]
	if exists {
		/* validation.ErrPermissionAlreadyRegistered */
		err.AddValidationError(&validation.FieldError{
			FieldName:  "id",
			CodeErrors: []string{string(validation.ErrUnknown)},
		})
	}

	permission, err := NewPermission(ID, name)
	if err != nil {
		return err
	}

	manager.permissions[ID] = permission
	return nil
}

func (manager *permissionManager) GetPermissionByID(ID int) (*Permission, validationError.IValidationError) {
	permission, exists := manager.permissions[ID]
	if !exists {
		var err = validationError.NewValidationError("permission")
		err.AddValidationError(&validation.FieldError{
			// not found
			FieldName:  "id",
			CodeErrors: []string{string(validation.ErrUnknown)},
		})
		return nil, err
	}
	return permission, nil
}

// Método para recuperar permissões oriundas de um banco de dados
// Mudar de pacote
func (manager *permissionManager) loadPermissions() error {
	return nil
}
