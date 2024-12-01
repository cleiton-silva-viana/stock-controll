package brand

import (
	"strings"
	"testing"

	"stock-controll/internal/domain/services/uuid"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var data = Brand{
	name:             "caola",
	description:      "this a company for your money",
	logo:             "...", // implementar
	status:           Active,
	manufacturerUUID: uuid.New(),
}

func TestNewNoError(t *testing.T) {
	t.Parallel()

	testCases := []unitary.TestField[Brand]{
		{
			Description: "brand name is equal than minimum alowed",
			Handler:     func(b *Brand) { b.name = strings.Repeat("a", minNameLengthForBrand) },
		},
		{
			Description: "brand name is equal than maximum allowed",
			Handler:     func(b *Brand) { b.name = strings.Repeat("b", maxNameLengthForBrand) },
		},
		{
			Description: "brand description is equal than minimum allowed",
			Handler:     func(b *Brand) { b.description = strings.Repeat("a", minDescriptionLength) },
		},
		{
			Description: "brand description is equal than maximum allowed",
			Handler:     func(b *Brand) { b.description = strings.Repeat("a", maxDescriptionLength) },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			copy := data
			test.Handler(&copy)

			// Act
			brandInstance, err := New(copy.name, copy.description, copy.logo, copy.manufacturerUUID)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, brandInstance)
			assert.Equal(t, copy.name, brandInstance.Name())
			assert.Equal(t, copy.description, brandInstance.Description())
			assert.Equal(t, copy.logo, brandInstance.Logo())
			assert.Equal(t, copy.status, brandInstance.Status())
			assert.Equal(t, copy.manufacturerUUID, brandInstance.ManufacturerUUID())
			assert.NotEmpty(t, brandInstance.UUID())
		})
	}
}

func TestNewWithError(t *testing.T) {
	t.Parallel()

	testCases := []unitary.TestField[Brand]{
		{
			Description: "brand name length is short than minimum allowed",
			Handler:     func(b *Brand) { b.name = strings.Repeat("a", minNameLengthForBrand-1) },
		},
		{
			Description: "brand name length is greater than maximum allowed",
			Handler:     func(b *Brand) { b.name = strings.Repeat("b", maxNameLengthForBrand+1) },
		},
		{
			Description: "brand description length is short than minimum allowed",
			Handler:     func(b *Brand) { b.description = strings.Repeat("a", minDescriptionLength-1) },
		},
		{
			Description: "brand description length is greater than maximum allowed",
			Handler:     func(b *Brand) { b.description = strings.Repeat("a", maxDescriptionLength+1) },
		},
		{
			Description: "brand manufacturer uuid is invalid",
			Handler:     func(b *Brand) { b.manufacturerUUID = "   " },
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			copy := data
			test.Handler(&copy)

			// Act
			brandInstance, err := New(copy.name, copy.description, copy.logo, copy.manufacturerUUID)

			// Assert
			assert.Nil(t, brandInstance)
			assert.NotNil(t, err)
		})
	}
}

func TestStateConsistencyAfterInvalidSet(t *testing.T) {
	testCases := []unitary.Consistence{
		{
			T:            t,
			Description:  "test consistence of brand name",
			Getter:       func() interface{} { return data.Name },
			Setter:       func(value interface{}) error { return data.SetName(value.(string)) },
			InvalidValue: "$Roll$Royce",
		},
		{
			T:            t,
			Description:  "test consistence of brand description",
			Getter:       func() interface{} { return data.Description },
			Setter:       func(value interface{}) error { return data.SetDescription(value.(string)) },
			InvalidValue: strings.Repeat("b", minDescriptionLength-1),
		},
		{
			T:            t,
			Description:  "test consistence of manufacturer uuid",
			Getter:       func() interface{} { return data.ManufacturerUUID },
			Setter:       func(value interface{}) error { return data.SetManufacturerUUID(value.(string)) },
			InvalidValue: "6654.9498;79/7949874.98789",
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {

			// Act & Arrange
			unitary.ConsistenceTest(test)
		})
	}
}
