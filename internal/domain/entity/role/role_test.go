package role

import (
	"strings"
	"testing"

	"stock-controll/internal/domain/entity/permission"
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/uuid"
	"stock-controll/internal/domain/services/validate"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var permission1, _ = permission.New("create")
var permission2, _ = permission.New("update")
var permission3, _ = permission.New("delete")
var permission4, _ = permission.New("view")

var data = Role{
	uuid: uuid.New(),
	name: "admin",
	permissions: map[string]permission.Permission{
		permission1.UUID(): *permission1,
		permission2.UUID(): *permission2,
		permission3.UUID(): *permission3,
	},
}

func TestNewNoError(t *testing.T) {
	testsCases := []unitary.TestField[Role]{
		{
			Description: "name with minimum allowed characters",
			Handler:     func(r *Role) { r.name = strings.Repeat("a", minNameLengthForRole) },
		},
		{
			Description: "name with maximum allowed characters",
			Handler:     func(r *Role) { r.name = strings.Repeat("b", maxNameLengthForRole) },
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.Description, func(t *testing.T) {
			dataCopy := data
			tt.Handler(&dataCopy)

			// Act
			role, err := New(dataCopy.name)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, role)
			assert.NoError(t, uuid.IsValid("", role.UUID()))
			assert.Equal(t, role.Name(), dataCopy.name)
			assert.NotNil(t, role.Permissions())
			assert.Empty(t, role.Permissions())
		})
	}
}

func TestNewWithError(t *testing.T) {
	// Arrange
	testsCases := []unitary.TestField[Role]{
		{
			Description: "role name is empty",
			Handler:     func(r *Role) { r.name = "" },
		},
		{
			Description: "role name contain special characters",
			Handler:     func(r *Role) { r.name = "Sell Pr@ducts" },
		},
		{
			Description: "role name is too shoort",
			Handler:     func(r *Role) { r.name = strings.Repeat("a", minNameLengthForRole-1) },
		},
		{
			Description: "role name is too long",
			Handler:     func(r *Role) { r.name = strings.Repeat("b", maxNameLengthForRole+1) },
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.Description, func(t *testing.T) {
			dataCopy := data
			tt.Handler(&dataCopy)

			// Act
			role, err := New(dataCopy.name)

			// Assert
			assert.Nil(t, role)
			require.Error(t, err)
			assert.ErrorAs(t, err, &validationerrors.ValidationError{})
		})
	}
}

func TestSetPermissionNoError(t *testing.T) {
	// Arrange
	role := data

	// Act
	err := role.SetPermission(*permission4)

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, role.Permissions(), permission4)
	assert.Len(t, role.Permissions(), 4)
}

func TestSetPermissionsWithError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestDependence[permission.Permission]{
		{
			Description: "permissions is empty filled",
			Dependency:  permission.Permission{},
		},
		{
			Description: "add already registered permission",
			Dependency:  *permission1,
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			role := data
			// Act
			err := role.SetPermission(test.Dependency)

			// Assert
			assert.Error(t, err)
			assert.Len(t, role.Permissions(), 3)
		})
	}
}

func TestRemovePermissionNoError(t *testing.T) {
	// Arrange
	permissionUUID := permission1.UUID()
	role := data

	// Act
	err := role.RemovePermission(permissionUUID)

	// Assert
	assert.NoError(t, err)
	assert.NotContains(t, role.Permissions(), permission1)
	assert.False(t, role.HasPermission(permissionUUID))
	assert.Len(t, role.Permissions(), 2)
}

func TestRemovePermissionWithError(t *testing.T) {
	testCases := []unitary.TestDependence[permission.Permission]{
		// Arrange
		{
			Description: "the permission not available in role",
			Dependency:  *permission4,
		},
		{
			Description: "the permission is empty filled",
			Dependency:  permission.Permission{},
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			// Arrange
			permissionUUID := test.Dependency.UUID()
			role := data

			// Act
			err := data.RemovePermission(permissionUUID)

			// Assert
			require.Error(t, err)
			assert.ErrorIs(t, err, &validate.FieldError{})
			assert.Len(t, role.Permissions(), 3)
			assert.False(t, role.HasPermission(permissionUUID))
		})
	}
}
