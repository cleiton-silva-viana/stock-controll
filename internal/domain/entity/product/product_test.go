package product

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type test struct {
	testDescription string
	productDTO
}

type productDTO struct {
	name             string
	description      string
	barcode          string
	brandUUID        string
	categoryUUID     string
	manufacturerUUID string
}

func Test_NewProcut_Sucess(t *testing.T) {
	t.Parallel()

	// Arrange
	testsCases := []test{
		{
			testDescription: "Product with minimum name length",
			productDTO: productDTO{
				name:             strings.Repeat("a", productNameMinLength),
				description:      "Descrição do produto",
				barcode:          "1234567890123",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "Product with maximum name length",
			productDTO: productDTO{
				name:             strings.Repeat("b", productNameMaxLength),
				description:      "Descrição do produto",
				barcode:          "1234567890123",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "Product with min description length",
			productDTO: productDTO{
				name:             "Product Test",
				description:      strings.Repeat("s", productDescriptionMinLength),
				barcode:          "1234567890123",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "Product with maximum description length",
			productDTO: productDTO{
				name:             "Product test",
				description:      strings.Repeat("z", productDescriptionMaxLength),
				barcode:          "1234567890123",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			product, err := NewProductBuilder().
				SetName(tt.name).
				SetDescription(tt.description).
				SetBarcode(tt.barcode).
				SetBrandUUID(tt.brandUUID).
				SetCategoryUUID(tt.categoryUUID).
				SetManufacturerUUID(tt.manufacturerUUID).
				Build()

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, product)
			assert.Equal(t, strings.ToLower(tt.name), strings.ToLower(product.GetName()))
			assert.Equal(t, strings.ToLower(tt.description), strings.ToLower(product.GetDescription()))
			assert.Equal(t, tt.brandUUID, product.GetBrandUUID())
			assert.Equal(t, tt.categoryUUID, product.GetCategoryUUID())
			assert.Equal(t, tt.manufacturerUUID, product.GetManufacturerUUID())
		})
	}
}

func Test_NewProduct_Error(t *testing.T) {
	t.Parallel()

	testsCases := []test{
		{
			testDescription: "empty product name",
			productDTO: productDTO{
				name:             "            ",
				description:      strings.Repeat("abc", 10),
				barcode:          "1946753214860",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "product name is short than minimum allowed",
			productDTO: productDTO{
				name:             strings.Repeat("a", productNameMinLength-1),
				description:      "Savor the sweetness of our Juicy Navel Oranges, handpicked for peak freshness",
				barcode:          "6548964631668",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "product name is greater than maximum allowed",
			productDTO: productDTO{
				name:             strings.Repeat("a", productNameMaxLength+1),
				description:      "Savor the sweetness of our Juicy Navel Oranges, handpicked for peak freshness",
				barcode:          "8949461894984",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "product description is short than minimum allowed",
			productDTO: productDTO{
				name:             "banana",
				description:      strings.Repeat("d", productDescriptionMinLength - 1),
				barcode:          "1234567890123",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "product description is greater than maximum allowed",
			productDTO: productDTO{
				name:             "banana",
				description:      strings.Repeat("A", productDescriptionMaxLength + 1),
				barcode:          "1234567890123",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "barcode contain letter",
			productDTO: productDTO{
				name:             "orange",
				description:      "Savor the sweetness of our Juicy Navel Oranges, handpicked for peak freshness. These easy-to-peel oranges are rich in vitamin C and perfect for snacking or juicing. Enjoy their vibrant flavor and refreshing juiciness in salads or on their own. A delicious, healthy treat for everyone!",
				barcode:          "1234565789a",
				brandUUID:        "01928cee-b413-72f3-ad15-a3a297f0a114",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "brand uuid contain invalid format",
			productDTO: productDTO{
				name:             "orange",
				description:      "Savor the sweetness of our Juicy Navel Oranges, handpicked for peak freshness. These easy-to-peel oranges are rich in vitamin C and perfect for snacking or juicing. Enjoy their vibrant flavor and refreshing juiciness in salads or on their own. A delicious, healthy treat for everyone!",
				barcode:          "1234567890123",
				brandUUID:        "0192589b-33df-7408/b68e-e95b7b1cf94e",
				manufacturerUUID: "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "manufacturer UUID contain invalid format",
			productDTO: productDTO{
				name:             "orange",
				description:      "Savor the sweetness of our Juicy Navel Oranges, handpicked for peak freshness. These easy-to-peel oranges are rich in vitamin C and perfect for snacking or juicing. Enjoy their vibrant flavor and refreshing juiciness in salads or on their own. A delicious, healthy treat for everyone!",
				barcode:          "1234567890123",
				brandUUID:        "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				manufacturerUUID: "092589b-a7c7-79d3-842d-9318f5f45961",
				categoryUUID:     "01928d22-18d1-701a-bad2-8a49848d97f4",
			},
		},
		{
			testDescription: "category UUID contain invalid format",
			productDTO: productDTO{
				name:             "orange",
				description:      "Savor the sweetness of our Juicy Navel Oranges, handpicked for peak freshness. These easy-to-peel oranges are rich in vitamin C and perfect for snacking or juicing. Enjoy their vibrant flavor and refreshing juiciness in salads or on their own. A delicious, healthy treat for everyone!",
				barcode:          "1234567890123",
				brandUUID:        "01928d22-004d-7d8e-92ac-9f4be5c3cddf",
				manufacturerUUID: "01928cee-b413-72f3-ad15-a3a297f0a114",
				categoryUUID:     "0192589c-174b-77ce-8395",
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.testDescription, func(t *testing.T) {

			// Act
			product, err := NewProductBuilder().
				SetName(tt.productDTO.name).
				SetDescription(tt.productDTO.description).
				SetBarcode(tt.productDTO.barcode).
				SetBrandUUID(tt.productDTO.brandUUID).
				SetCategoryUUID(tt.productDTO.categoryUUID).
				SetManufacturerUUID(tt.productDTO.manufacturerUUID).
				Build()

			// Assert
			assert.Nil(t, product)
			assert.NotNil(t, err)
		})
	}
}
