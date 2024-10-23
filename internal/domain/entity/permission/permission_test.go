package permission

// TODO: Checar o tipo de erro retornado

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetPermissionManager_NoError(t *testing.T) {
	// Act
	result := GetPermissionManager()

	// Assert
	require.NotNil(t, result)
}

type permissionTest struct {
	testDescription string
	id              int
	name            string
}

func Test_AddPermission_NoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []permissionTest{
		{
			testDescription: "name contains maximum caracters allowed",
			id:              37,
			name:            strings.Repeat("a", minNameLengthForPermission),
		},
		{
			testDescription: "name contains mimimum caracters allowed",
			id:              38,
			name:            strings.Repeat("b", maxNameLengthForPermission),
		},
		{
			testDescription: "name contains minimum number for id allowed",
			id:              minIDNumberForPermission,
			name:            "sell product",
		},
		{
			testDescription: "name contains maximum number fo id allowed",
			id:              maxIDNumberForPermission,
			name:            "buy product",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {
			var manager = GetPermissionManager()

			// Act
			err := manager.AddPermission(tt.id, tt.name)

			// Assert
			assert.Nil(t, err)
			require.NotEmpty(t, manager.permissions[tt.id])
			assert.Equal(t, manager.permissions[tt.id].name, tt.name)
		})
	}
}

func Test_AddPermission_Error(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []permissionTest{

		{
			testDescription: "name is empty",
			id:              5,
			name:            "", 
		},
		{
			testDescription: "name contains special characters",
			id:              6,
			name:            "name@with#special$chars",
		},
		{
			testDescription: "name is too short",
			id:              3,
			name:            strings.Repeat("a", minNameLengthForPermission - 1),
		},
		{
			testDescription: "name is too long",
			id:              4,
			name:            strings.Repeat("c", maxNameLengthForPermission + 1),
		},
		{
			testDescription: "id is negative",
			id:              minIDNumberForPermission - 1,
			name:            "valid name",
		},
		{
			testDescription: "id is above the allowed limit",
			id:              maxIDNumberForPermission + 1,
			name:            "another valid name",
		},
	}
	
	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {
			var manager = GetPermissionManager()

			// Act
			err := manager.AddPermission(tt.id, tt.name)

			// Assert
			assert.NotNil(t, err)
			require.Empty(t, manager.permissions[tt.id])
		})
	}
}

var registerProduct, _ = NewPermission(1, "register product")

func Test_GetPermissionByID_NoError(t *testing.T) {
	// Arrange
	var manager = GetPermissionManager()
	manager.AddPermission(registerProduct.id, registerProduct.name)

	// Act
	permission, err := manager.GetPermissionByID(registerProduct.id)

	// Assert
	assert.Nil(t, err)
	require.NotNil(t, permission)
	assert.Equal(t, registerProduct.name, permission.name)
}

func Test_GetPermissionByID_WithError(t *testing.T) {
	// Arrange
	var manager = GetPermissionManager()

	// act
	permission, err := manager.GetPermissionByID(111)

	// Assert
	assert.Nil(t, permission)
	require.NotNil(t, err)
}
