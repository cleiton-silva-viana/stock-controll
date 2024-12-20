package role

import (
	"fmt"

	"stock-controll/internal/domain/entity/permission"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

type Role struct {
	uuid uuid.UUID
	name        string
	permissions map[string]permission.Permission
}

func New(name string) (*Role, error) {
	roleError := entity.Error("role").
		AddValidationError(validateName(name))

	if roleError.HasError() {
		return nil, roleError
	}

	return &Role{
		uuid: *uuid.New(),
		name: name,
		permissions: make(map[string]permission.Permission),
	}, nil
}

func (r *Role) UUID() string {
	return r.uuid.String()
}

func (r *Role) Name() string {
	return r.name
}

func (r *Role) Permissions() []permission.Permission {
	permissions := make([]permission.Permission, 0, len(r.permissions))
	for _, p := range r.permissions {
		permissions = append(permissions, p)
	}
	return permissions
}

const ErrPermissionAlreadyAssignedToRole = "ERR_PERMISSION_ALREADY_ASSIGNED_TO_ROLE"

func (r *Role) AddPermission(permission permission.Permission) error {
	permissionUUID := permission.UUID()
	exists := r.HasPermission(permissionUUID)
	if exists {
		return &field.FieldError{
			FieldName: "permissions",
			CodeError: fmt.Sprint(ErrPermissionAlreadyAssignedToRole, permissionUUID),
		}
	}
	r.permissions[permissionUUID] = permission
	return nil
}

const ErrPermissionNotAssociated = "ERR_PERMISSION_NOT_ASSOCIATED"

func (r *Role) RemovePermission(permissionUUID string) error {
	exists := r.HasPermission(permissionUUID)
	if !exists {
		return &field.FieldError{
			FieldName: "permissions",
			CodeError: ErrPermissionNotAssociated,
		}
	}
	delete(r.permissions, permissionUUID)
	return nil
}

func (r *Role) PermissionByID(uuid string) permission.Permission {
	return r.permissions[uuid]
}

func (r *Role) HasPermission(uuid string) bool {
	_, exists := r.permissions[uuid]
	return exists
}

const (
	minLength = 3
	maxLength = 24
)

func validateName(name string) error {
	return validate.New[string]("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minLength, maxLength),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
}