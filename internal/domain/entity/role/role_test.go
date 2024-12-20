package role

import (
	"fmt"
	"strings"
	"testing"

	"stock-controll/internal/domain/entity/permission"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/valueobject/uuid"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var permission1, _ = permission.New("create")
var permission2, _ = permission.New("update")
var permission3, _ = permission.New("delete")
var permission4, _ = permission.New("view")

func Setup() *Role {
	return &Role{
		uuid: *uuid.New(),
		name: "admin",
		permissions: map[string]permission.Permission{
			permission1.UUID(): *permission1,
			permission2.UUID(): *permission2,
			permission3.UUID(): *permission3,
		},
	}
}

func TestNewNoError(t *testing.T) {
	tests := []unitary.TestField[Role]{
		{
			Description: "name with minimum allowed characters",
			Handler:     func(r *Role) { r.name = strings.Repeat("a", minLength) },
		},
		{
			Description: "name with maximum allowed characters",
			Handler:     func(r *Role) { r.name = strings.Repeat("b", maxLength) },
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("Test function New role - no error expected - %s", tt.Description), func(t *testing.T) {
			r := Setup()
			tt.Handler(r)

			// Act
			role, err := New(r.name)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, role)
			assert.Equal(t, role.Name(), r.name)
			assert.NotNil(t, role.Permissions())
			assert.Empty(t, role.Permissions())
		})
	}
}

func TestNewWithError(t *testing.T) {
	// Arrange
	tests := []unitary.TestField[Role]{
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
			Handler:     func(r *Role) { r.name = strings.Repeat("a", minLength-1) },
		},
		{
			Description: "role name is too long",
			Handler:     func(r *Role) { r.name = strings.Repeat("b", maxLength+1) },
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("Test function New role - error expected because the %s", tt.Description), func(t *testing.T) {
			// Arrange
			r := Setup()
			tt.Handler(r)

			// Act
			role, err := New(r.name)
			// Assert
			assert.Nil(t, role)
			require.Error(t, err)
			assert.ErrorAs(t, err, &entity.EntityError{})
		})
	}
}

func TestSetPermissionNoError(t *testing.T) {
	// Arrange
	r := Setup()
	rr, _ := New(r.name)

	// Act
	err := rr.AddPermission(*permission4)

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, rr.Permissions(), permission4)
	assert.Len(t, rr.Permissions(), 4)
}

func TestSetPermissionsWithError(t *testing.T) {
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
		// Arrange
		t.Run(test.Description, func(t *testing.T) {
			r := Setup()
			rr, _ := New(r.name)
			// Act
			err := rr.AddPermission(test.Dependency)

			// Assert
			assert.Error(t, err)
			assert.Len(t, rr.Permissions(), 3)
		})
	}
}

func TestRemovePermissionNoError(t *testing.T) {
	// Arrange
	permissionUUID := permission1.UUID()
	rr := Setup()

	// Act
	err := rr.RemovePermission(permissionUUID)

	// Assert
	assert.NoError(t, err)
	assert.NotContains(t, rr.Permissions(), permission1)
	assert.False(t, rr.HasPermission(permissionUUID))
	assert.Len(t, rr.Permissions(), 2)
}

func TestRemovePermissionWithError(t *testing.T) {
	testCases := []unitary.TestDependence[permission.Permission]{
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
			rr := Setup()

			// Act
			err := rr.RemovePermission(permissionUUID)

			// Assert
			require.Error(t, err)
			assert.ErrorIs(t, err, &field.FieldError{})
			assert.Len(t, rr.Permissions(), 3)
			assert.False(t, rr.HasPermission(permissionUUID))
		})
	}
}
