package role

import (
	"fmt"
	"stock-controll/internal/domain/entity/permission"
	validationError "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"
)

type Role struct {
	uuid        string
	name        string
	permissions map[string]permission.Permission
}

func New(name string) (*Role, error) {
	var roleInstance = Role{
		uuid:        uuid.New(),
		name:        name,
		permissions: make(map[string]permission.Permission),
	}

	var roleError = validationError.New("role").
		AddValidationError(roleInstance.SetName(name))

	if roleError.HasError() {
		return nil, roleError
	}

	roleInstance.uuid = uuid.New()
	return &roleInstance, nil
}

func (r *Role) UUID() string {
	return r.uuid
}

func (r *Role) Name() string {
	return r.name
}

const (
	minNameLengthForRole = 3
	maxNameLengthForRole = 24
)

func (r *Role) SetName(name string) error {
	err := validate.New[string]("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLengthForRole, maxNameLengthForRole),
		validate.CheckNumbers(validate.Disallow),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err == nil {
		r.name = name
	}
	return err
}

func (r *Role) Permissions() []permission.Permission {
	permissions := make([]permission.Permission, 0, len(r.permissions))
	for _, p := range r.permissions {
		permissions = append(permissions, p)
	}
	return permissions
}

const ErrPermissionAlreadyAssignedToRole = "ERR_PERMISSION_ALREADY_ASSIGNED_TO_ROLE"

func (r *Role) SetPermission(permission permission.Permission) error {
	permissionUUID := permission.UUID()
	exists := r.HasPermission(permissionUUID)
	if exists {
		return &validate.FieldError{
			FieldName: "permissions",
			CodeError: fmt.Sprint(ErrPermissionAlreadyAssignedToRole, permissionUUID),
		}
	}
	r.permissions[permissionUUID] = permission
	return nil
}

func (r *Role) PermissionByID(uuid string) permission.Permission {
	return r.permissions[uuid]
}

const ErrPermissionNotAssociated = "ERR_PERMISSION_NOT_ASSOCIATED"

func (r *Role) RemovePermission(permissionUUID string) error {
	exists := r.HasPermission(permissionUUID)
	if !exists {
		return &validate.FieldError{
			FieldName: "permissions",
			CodeError: ErrPermissionNotAssociated,
		}
	}
	delete(r.permissions, permissionUUID)
	return nil
}

func (r *Role) HasPermission(uuid string) bool {
	_, exists := r.permissions[uuid]
	return exists
}
