package brand

import (
	"strings"
	"testing"

	"stock-controll/internal/domain/valueobject/uuid"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type brandConfig struct {
	name             string
	description      string
	manufacturerUUID string
	status           bool
}

func Setup() *brandConfig {
	return &brandConfig{
		name:             "caola",
		description:      "this a company for your money",
		status:           Active,
		manufacturerUUID: uuid.New().String(),
	}
}

func TestNewNoError(t *testing.T) {
	t.Parallel()

	testCases := []unitary.TestField[brandConfig]{
		{
			Description: "brand name is equal than minimum alowed",
			Handler:     func(bc *brandConfig) { bc.name = strings.Repeat("a", minNameLengthForBrand) },
		},
		{
			Description: "brand name is equal than maximum allowed",
			Handler:     func(bc *brandConfig) { bc.name = strings.Repeat("b", maxNameLengthForBrand) },
		},
		{
			Description: "brand description is equal than minimum allowed",
			Handler:     func(bc *brandConfig) { bc.description = strings.Repeat("a", minDescriptionLength) },
		},
		{
			Description: "brand description is equal than maximum allowed",
			Handler:     func(bc *brandConfig) { bc.description = strings.Repeat("a", maxDescriptionLength) },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			b := Setup()
			test.Handler(b)

			// Act
			result, err := New(b.name, b.description, b.manufacturerUUID)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, result)
			assert.Equal(t, b.name, result.Name())
			assert.Equal(t, b.description, result.Description())
			assert.Equal(t, b.status, result.Status())
			assert.Equal(t, b.manufacturerUUID, result.ManufacturerUUID())
			assert.NotEmpty(t, result.UUID())
		})
	}
}

func TestNewWithError(t *testing.T) {
	t.Parallel()

	testCases := []unitary.TestField[brandConfig]{
		{
			Description: "brand name length is short than minimum allowed",
			Handler:     func(bc *brandConfig) { bc.name = strings.Repeat("a", minNameLengthForBrand-1) },
		},
		{
			Description: "brand name length is greater than maximum allowed",
			Handler:     func(bc *brandConfig) { bc.name = strings.Repeat("b", maxNameLengthForBrand+1) },
		},
		{
			Description: "brand description length is short than minimum allowed",
			Handler:     func(bc *brandConfig) { bc.description = strings.Repeat("a", minDescriptionLength-1) },
		},
		{
			Description: "brand description length is greater than maximum allowed",
			Handler:     func(bc *brandConfig) { bc.description = strings.Repeat("a", maxDescriptionLength+1) },
		},
		{
			Description: "brand manufacturer uuid is invalid",
			Handler:     func(bc *brandConfig) { bc.manufacturerUUID = "   " },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			b := Setup()
			test.Handler(b)

			// Act
			result, err := New(b.name, b.description, b.manufacturerUUID)

			// Assert
			assert.Nil(t, result)
			assert.NotNil(t, err)
		})
	}
}
