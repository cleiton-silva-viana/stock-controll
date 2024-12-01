package unitary

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var Fake = faker.New()

type TestField[T any] struct {
	Description string
	Handler     func(*T)
}

type TestDependence[T any] struct {
	Description string
	Dependency  T
}

type Consistence struct {
	T            *testing.T
	Description  string
	Getter       func() interface{}
	Setter       func(value interface{}) error
	InvalidValue interface{}
}

func ConsistenceTest(config Consistence) {
	// Pre check
	require.NotNil(config.T, config.T)
	require.NotNil(config.T, config.Getter)
	require.NotNil(config.T, config.Setter)
	require.NotNil(config.T, config.InvalidValue)
	assert.NotEqual(config.T, config.InvalidValue, config.Getter())

	// Arrange
	initialValue := config.Getter()
	invalidValue := config.InvalidValue

	// Act
	err := config.Setter(invalidValue)

	// Assert
	require.Error(config.T, err)
	assert.Equal(config.T, initialValue, config.Getter())
}
