package role

// TODO: Refatorar - melhora ros testes, realizar testes para função não testadas

import (
	permissionEntity "stock-controll/internal/domain/entity/permission"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type roleStructForTest struct {
	testDescription string
	id              int
	name            string
}

func Test_Role_NewRole_NoError(t *testing.T) {
	testsCases := []roleStructForTest{
		{
			testDescription: "name with minimum allowed characters",
			id:              minIDNumberForRole,
			name:            strings.Repeat("a", minNameLengthForRole),
		},
		{
			testDescription: "name with maximum allowed characters",
			id:              maxIDNumberForRole,
			name:            strings.Repeat("b", maxNameLengthForRole),
		},
		{
			testDescription: "id with minimum allowed number",
			id:              minIDNumberForRole,
			name:            "valid role name",
		},
		{
			testDescription: "id with maximum allowed number",
			id:              maxIDNumberForRole,
			name:            "another valid role name",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			role, err := NewRole(tt.id, tt.name)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, role)
			assert.Equal(t, role.GetID(), tt.id)
			assert.Equal(t, role.GetName(), tt.name)
			assert.NotNil(t, role.GetPermissions())
			assert.Empty(t, role.GetPermissions())
		})
	}
}

func Test_Role_NewRole_WithError(t *testing.T) {
	// Arrange
	testsCases := []struct {
		testDescription string
		roleID          int
		roleName        string
	}{
		{
			testDescription: "role name is empty",
			roleID:          10,
			roleName:        "",
		},
		{
			testDescription: "role name contain special characters",
			roleID:          10,
			roleName:        "Sell Pr@ducts",
		},
		{
			testDescription: "role name is too shoort",
			roleID:          9,
			roleName:        strings.Repeat("a", minNameLengthForRole-1),
		},
		{
			testDescription: "role name is too long",
			roleID:          8,
			roleName:        strings.Repeat("b", maxNameLengthForRole+1),
		},
		{
			testDescription: "role id is less than allowed",
			roleID:          minIDNumberForRole - 1,
			roleName:        "sell product",
		},
		{
			testDescription: "role id is greater than allowed",
			roleID:          maxIDNumberForRole + 1,
			roleName:        "sell product",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			role, err := NewRole(tt.roleID, tt.roleName)

			// Assert
			assert.Nil(t, role)
			require.NotNil(t, err)
		})
	}
}

var role, _ = NewRole(30, "client")
var manager = permissionEntity.GetPermissionManager()

func Test_Role_AddPermission_NoError(t *testing.T) {
	// Arrange
	manager.AddPermission(1, "view product")
	permission, _ := manager.GetPermissionByID(1)

	// Act
	role.SetPermissions(*permission)

	// Assert
	assert.True(t, role.HasPermission(permission.GetID()))
	assert.Equal(t, role.GetPermissionByID(permission.GetID()), *permission)
}

func Test_Role_SetPermissions_WithError(t *testing.T) {
	// Arrange
	manager.AddPermission(10, "read product")
	permission, _ := manager.GetPermissionByID(10)

	// Act
	err1 := role.SetPermissions(*permission)
	err2 := role.SetPermissions(*permission)

	// Assert
	assert.Nil(t, err1)
	require.NotNil(t, err2)
	assert.True(t, role.HasPermission(permission.GetID()))
}

func Test_role_RemovePermission(t *testing.T) {
	// Arrange
	manager.AddPermission(10, "read product")
	permission, _ := manager.GetPermissionByID(10)
	role.SetPermissions(*permission)

	// Act
	role.RemovePermission(permission.GetID())

	// Assert
	assert.NotContains(t, role.GetPermissions(), permission)
	assert.False(t, role.HasPermission(permission.GetID()))
}
