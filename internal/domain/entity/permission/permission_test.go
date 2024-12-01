package permission

import (
	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/test/unitary"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var data = Permission{
	uuid: "0192cfd8-9af2-79d7-9632-a8141ee5c7ad",
	name: "delete",
}

func TestNewNoError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Permission]{
		{
			Description: "permission name is capitalized",
			Handler:     func(p *Permission) { p.name = "UPDATE" },
		},
		{
			Description: "permission name with min length allowed",
			Handler:     func(p *Permission) { p.name = strings.Repeat("v", minNameLength) },
		},
		{
			Description: "permission name with max length allowed",
			Handler:     func(p *Permission) { p.name = strings.Repeat("a", maxNameLength) },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			instance, err := New(dataCopy.name)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, instance)
			assert.Equal(t, strings.ToLower(dataCopy.name), instance.Name())
		})
	}
}

func TestNewWithError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Permission]{
		{
			Description: "name is empty",
			Handler:     func(p *Permission) { p.name = "     " },
		},
		{
			Description: "name contains special chars",
			Handler:     func(p *Permission) { p.name = "DELET$" },
		},
		{
			Description: "name contains number",
			Handler:     func(p *Permission) { p.name = "add1ng" },
		},
		{
			Description: "name is less than min allowed",
			Handler:     func(p *Permission) { p.name = strings.Repeat("a", minNameLength-1) },
		},
		{
			Description: "name is greater than max allowed",
			Handler:     func(p *Permission) { p.name = strings.Repeat("b", maxNameLength+1) },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			instance, err := New(dataCopy.name)

			// Assert
			assert.Nil(t, instance)
			assert.Error(t, err)
			assert.ErrorAs(t, err, &validationerrors.ValidationError{})
		})
	}
}

func TestSetNameConsistence(t *testing.T) {
	// Arrange
	const invalidName = "#$$$$$"
	dataCopy := data

	// Act
	err := dataCopy.SetName(invalidName)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, dataCopy.Name(), data.Name())
	assert.Equal(t, dataCopy.UUID(), data.UUID())
}

type permissionTest struct {
	*Permission
	expectedEqual bool
}

var initialTestData = permissionTest{
	expectedEqual: false,
	Permission:    &data,
}

func TestEqual(t *testing.T) {
	testCases := []unitary.TestField[permissionTest]{
		{
			Description: "same permissions",
			Handler: func(t *permissionTest) {
				t.expectedEqual = true
			},
		},
		{
			Description: "permissions with different uuids",
			Handler: func(t *permissionTest) {
				t.name = "0192cfe9-6a7b-70a1-9066-f8858f33b51b"
				t.expectedEqual = false
			},
		},
		{
			Description: "permissions with different name",
			Handler: func(t *permissionTest) {
				t.name = "rollback"
				t.expectedEqual = false
			},
		},
		{
			Description: "different name and uuid",
			Handler: func(t *permissionTest) {
				t.uuid = "0192cfe9-6a7b-70a1-9066-f8858f33b51b"
				t.name = "rollback"
				t.expectedEqual = false
			},
		},
		{
			Description: "nil comparison",
			Handler: func(t *permissionTest) {
				t.Permission = nil
				t.expectedEqual = false
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := initialTestData
			test.Handler(&dataCopy)

			// Act
			result := data.Equals(dataCopy.Permission)

			// Assert
			assert.Equal(t, dataCopy, result)
		})
	}
}
