package brand

import (
	"stock-controll/internal/domain/entity/common"
	"stock-controll/test/unitary"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var brandData = Brand{
	name:             "caola",
	description:      "this a company for your money",
	logo:             "...", // implementar
	status:           Active,
	manufacturerUUID: common.GenerateUUID(),
}

func Test_NewBrand_NoError(t *testing.T) {
	t.Parallel()

	testCases := []unitary.TestField[Brand]{
		{
			TestDescription: "brand name is equal than minimum alowed",
			Handler:         func(b Brand) { b.name = strings.Repeat("a", minNameLengthForBrand) },
		},
		{
			TestDescription: "brand name is equal than maximum allowed",
			Handler:         func(b Brand) { b.name = strings.Repeat("b", maxNameLengthForBrand) },
		},
		{
			TestDescription: "brand description is equal than minimum allowed",
			Handler:         func(b Brand) { b.description = strings.Repeat("a", minDescriptionLength) },
		},
		{
			TestDescription: "brand description is equal than maximum allowed",
			Handler:         func(b Brand) { b.description = strings.Repeat("a", maxDescriptionLength) },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			brand := brandData
			tt.Handler(brand)

			// Act
			brandInstance, err := NewBrand(brand.name, brand.description, brand.logo, brand.manufacturerUUID)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, brandInstance)
			assert.Equal(t, brand.name, brandInstance.GetName())
			assert.Equal(t, brand.description, brandInstance.GetDescription())
			assert.Equal(t, brand.logo, brandInstance.GetLogo())
			assert.Equal(t, brand.status, brandInstance.GetStatus())
			assert.Equal(t, brand.manufacturerUUID, brandInstance.GetManufacturerUUID())
			assert.NotEmpty(t, brandInstance.GetUUID())
		})
	}
}

func Test_NewBrand_WithError(t *testing.T) {
	t.Parallel()

	testCases := []unitary.TestField[Brand]{
		{
			TestDescription: "brand name length is short than minimum allowed",
			Handler:         func(b Brand) { b.name = strings.Repeat("a", minNameLengthForBrand-1) },
		},
		{
			TestDescription: "brand name length is greater than maximum allowed",
			Handler:         func(b Brand) { b.name = strings.Repeat("b", maxNameLengthForBrand+1) },
		},
		{
			TestDescription: "brand description length is short than minimum allowed",
			Handler:         func(b Brand) { b.description = strings.Repeat("a", minDescriptionLength-1) },
		},
		{
			TestDescription: "brand description length is greater than maximum allowed",
			Handler:         func(b Brand) { b.description = strings.Repeat("a", maxDescriptionLength+1) },
		},
		{
			TestDescription: "brand manufacturer uuid is invalid",
			Handler:         func(b Brand) { b.manufacturerUUID = "   " },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestDescription, func(t *testing.T) {
			brand := brandData
			tt.Handler(brand)

			// Act
			brandInstance, err := NewBrand(brand.name, brand.description, brand.logo, brand.manufacturerUUID)

			// Assert
			assert.Nil(t, brandInstance)
			assert.NotNil(t, err)
		})
	}
}
