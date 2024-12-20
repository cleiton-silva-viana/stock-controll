package permission

import (
	"sync"

	"stock-controll/internal/domain/services/error/field"
)

var (
	instance *manager
	once     sync.Once
)

type manager struct {
	permissions map[string]*Permission
}

func Manager() *manager {
	once.Do(func() {
		instance = newManager()
	})
	return instance
}

func newManager() *manager {
	return &manager{
		permissions: make(map[string]*Permission),
	}
}

func (m *manager) HasPermissionByUUID(permissionUUID string) (*Permission, bool) {
	permission, exists := m.permissions[permissionUUID]
	return permission, exists
}

const ErrPermissionAlreadyRegistered = "ERR_PERMISSION_ALREADY_REGISTERED"

func (m *manager) AddPermission(permission *Permission) error {
	permissionUUID := permission.UUID()
	_, exists := m.HasPermissionByUUID(permissionUUID)
	if exists {
		return &field.FieldError{
			FieldName: "uuid",
			CodeError: ErrPermissionAlreadyRegistered,
		}
	}
	m.permissions[permissionUUID] = permission
	return nil
}

const ErrPermissionNotFound = "ERR_PERMISSION_NOT_FOUND"

func (m *manager) PermissionByUID(permissionUUID string) (*Permission, error) {
	permissionInstance, exists := m.HasPermissionByUUID(permissionUUID)
	if !exists {
		return nil, &field.FieldError{
			// not found
			FieldName: "permission_uuid",
			CodeError: ErrPermissionNotFound,
		}
	}
	return permissionInstance, nil
}

/*
O método load permissions deve ser responsável por carregar as permissões oriundas de um banco de dados...
*/
func (m *manager) loadPermissions() error {
	return nil
}
