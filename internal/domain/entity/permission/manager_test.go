package permission

import (
	"strings"
	"testing"

	"stock-controll/internal/domain/valueobject/uuid"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require" 
)

func TestManagerNoError(t *testing.T) {
	// Act
	result := Manager()

	// Assert
	require.NotNil(t, result)
}

func TestAddPermissionNoError(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []unitary.TestField[Permission]{
		{
			Description: "name contains maximum caracters allowed",
			Handler: func(p *Permission) {
				p.name = strings.Repeat("a", minNameLength)
			},
		},
		{
			Description: "name contains mimimum caracters allowed",
			Handler: func(p *Permission) {
				p.name = strings.Repeat("b", maxNameLength)
			},
		},
	}

	for _, test := range testsCases {
		t.Run(test.Description, func(t *testing.T) {
			var manager = Manager()

			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			err := manager.AddPermission(&dataCopy)

			// Assert
			assert.Nil(t, err)
			require.NotEmpty(t, manager.permissions)
			assert.Contains(t, manager.permissions, dataCopy)
		})
	}
}

func Test_AddPermission_Error(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []unitary.TestField[Permission]{
		{
			Description: "name is empty",
			Handler: func(p *Permission) {
				p.name = "     "
			},
		},
		{
			Description: "name contains special characters",
			Handler: func(p *Permission) {
				p.name = "name@with#special$chars"
			},
		},
		{
			Description: "name is too short",
			Handler: func(p *Permission) {
				p.name = strings.Repeat("a", minNameLength-1)
			},
		},
		{
			Description: "name is too long",
			Handler: func(p *Permission) {
				p.name = strings.Repeat("c", maxNameLength+1)
			},
		},
	}

	for _, test := range testsCases {
		t.Run(test.Description, func(t *testing.T) {
			var manager = Manager()

			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			err := manager.AddPermission(&dataCopy)

			// Assert
			assert.NotNil(t, err)
			require.Empty(t, manager.permissions)
		})
	}
}

var permission1 = Permission{
	uuid: *uuid.New(),
	name: "create",
}

func TestPermissionByIDNoError(t *testing.T) {
	// Arrange
	var manager = Manager()

	manager.AddPermission(&permission1)

	// Act
	permission, err := manager.PermissionByUID(permission1.UUID())

	// Assert
	assert.Nil(t, err)
	require.NotNil(t, permission)
	assert.Equal(t, permission, permission1)
}

func TestPermissionByIDWithError(t *testing.T) {
	// Arrange
	var manager = Manager()

	// act
	permission, err := manager.PermissionByUID(permission1.UUID())

	// Assert
	assert.Nil(t, permission)
	require.NotNil(t, err)
}
