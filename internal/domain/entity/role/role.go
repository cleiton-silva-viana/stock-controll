package role

import (
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/entity/permission"
	"stock-controll/internal/domain/validation"
)

type Role struct {
	id          int
	name        string
	permissions map[int]permission.Permission
}

func NewRole(id int, name string, permissions ...permission.Permission) (*Role, validationError.IValidationError) {
	var roleError = validationError.NewValidationError("role")
	var roleInstance = Role{}
	roleInstance.permissions = make(map[int]permission.Permission)

	roleError.
		AddValidationError(roleInstance.SetID(id)).
		AddValidationError(roleInstance.SetName(name)).
		AddValidationError(roleInstance.SetPermissions(permissions...))

	if roleError.HasError() {
		return nil, roleError
	}
	return &roleInstance, nil
}

func (r *Role) GetID() int {
	return r.id
}

const (
	minIDNumberForRole = 1
	maxIDNumberForRole = 100
)

func (r *Role) SetID(id int) *validation.FieldError {
	err := validation.Validate[int]("id", id,
		validation.IsInRange(minIDNumberForRole, maxIDNumberForRole, validation.ErrUnknown))
	if err == nil {
		r.id = id
	}
	return err
}

func (r *Role) GetName() string {
	return r.name
}

const (
	minNameLengthForRole = 3
	maxNameLengthForRole = 24
)

func (r *Role) SetName(name string) *validation.FieldError {
	err := validation.Validate[string]("name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minNameLengthForRole, maxNameLengthForRole, validation.ErrUnknown),
		validation.CheckNumbers(validation.Disallow, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
	if err == nil {
		r.name = name
	}
	return err
}

func (r *Role) GetPermissions() []permission.Permission {
	permissions := make([]permission.Permission, 0, len(r.permissions))
	for _, p := range r.permissions {
		permissions = append(permissions, p)
	}
	return permissions
}

// ao invés de retornar um error, podemos retornar um error code ...
func (r *Role) SetPermissions(permissions ...permission.Permission) *validation.FieldError {
	for _, permission := range permissions {
		if r.HasPermission(permission.GetID()) {
			err := validation.FieldError{
				FieldName: "permissions",
				CodeErrors: []string{string(validation.ErrUnknown)},
			}
			return &err
		}
		r.permissions[permission.GetID()] = permission
	}
	return nil
}

func (r *Role) GetPermissionByID(id int) permission.Permission {
	return r.permissions[id]
}

func (r *Role) RemovePermission(PermissionID int) {
	delete(r.permissions, PermissionID)
}

func (r *Role) HasPermission(id int) bool {
	_, exits := r.permissions[id]
	return exits
}
