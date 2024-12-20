package tag

import (
	"strings"
	"testing"
	"time"

	"stock-controll/internal/domain/valueobject/uuid"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var product1 = uuid.New().String()
var product2 = uuid.New().String()
var product3 = uuid.New().String()
var product4 = uuid.New().String()

var data = Tag{
	name:        "food",
	description: "for foods products",
	tagType:     Category,
	createdAt:   time.Now(),
	updatedAt:   time.Now(),
	status:      "active",
	associatedProduct: map[string]struct{}{
		product1: struct{}{},
		product2: struct{}{},
		product3: struct{}{},
	},
}

func TestNewNoError(t *testing.T) {
	t.Parallel()

	// Arrange
	tests := []unitary.TestField[Tag]{
		{
			Description: "name with minimun length allowed",
			Handler:     func(t *Tag) { t.name = strings.Repeat("a", minNameLength) },
		},
		{
			Description: "name with maximum length allowed",
			Handler:     func(t *Tag) { t.name = strings.Repeat("b", maxNameLength) },
		},
		{
			Description: "compoust name for tag",
			Handler:     func(t *Tag) { t.name = "new heroes" },
		},
		{
			Description: "tag name with number",
			Handler:     func(t *Tag) { t.name = "S1mpson" },
		},
		{
			Description: "tag description with min length allowed",
			Handler:     func(t *Tag) { t.description = strings.Repeat("a", minDescriptionLength) },
		},
		{
			Description: "tag description with max length allowed",
			Handler:     func(t *Tag) { t.description = strings.Repeat("b", maxDescriptionLength) },
		},
		{
			Description: "tag description with number",
			Handler:     func(t *Tag) { t.description = "the homer simpson are cool 1111" },
		},
		{
			Description: "tag description with special characters",
			Handler:     func(t *Tag) { t.description = "tag description with special characters ;)" },
		},
		// TODO: testar tag type
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			dataCopy := data
			tt.Handler(&dataCopy)

			// Act
			tag, err := New(dataCopy.name, dataCopy.description, string(dataCopy.tagType))

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, tag)
			assert.Equal(t, dataCopy.name, tag.Name())
			assert.Equal(t, dataCopy.description, tag.Description())
			assert.Equal(t, dataCopy.tagType, tag.TagType())
			assert.LessOrEqual(t, time.Now(), tag.CreatedAt())
			assert.LessOrEqual(t, time.Now(), tag.UpdatedAt())
			require.NotNil(t, tag.AssociatedProductsUUIDs())
			assert.Len(t, tag.AssociatedProductsUUIDs(), 0)
			// como validar o status?
		})
	}
}

func TestNewWithError(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := []unitary.TestField[Tag]{
		{
			Description: "tag name is empty",
			Handler:     func(t *Tag) { t.name = "      " },
		},
		{
			Description: "tag name is less than minimum allowed",
			Handler:     func(t *Tag) { t.name = strings.Repeat("b", minNameLength-1) },
		},
		{
			Description: "tag name is greater than maximum allowed",
			Handler:     func(t *Tag) { t.name = strings.Repeat("b", maxNameLength+1) },
		},
		{
			Description: "tag name is empty",
			Handler:     func(t *Tag) { t.name = "" },
		},
		{
			Description: "tag name contains special chars",
			Handler:     func(t *Tag) { t.name = "new heroes #%," },
		},
		{
			Description: "tag description with min length less than allowed",
			Handler:     func(t *Tag) { t.description = strings.Repeat("a", minDescriptionLength-1) },
		},
		{
			Description: "tag description with max length greater than allowed",
			Handler:     func(t *Tag) { t.description = strings.Repeat("b", maxDescriptionLength+1) },
		},
		{
			Description: "tag type is invalid",
			Handler:     func(t *Tag) { t.tagType = "unkown" },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Description, func(t *testing.T) {
			dataCopy := data
			tt.Handler(&dataCopy)

			// Act
			tag, err := New(dataCopy.name, dataCopy.description, string(dataCopy.tagType))

			// Assert
			assert.Nil(t, tag)
			require.Error(t, err)
		})
	}
}

func TestAssociateProductsNoError(t *testing.T) {
	// Arrange
	dataCopy := data

	// Act
	err := dataCopy.AssociateProduct(product4)

	// Assert
	assert.NoError(t, err)
	require.Equal(t, data.associatedProduct, 4)
	assert.True(t, data.IsAssociated(product4))
}

type uuidtest struct {
	uuid string
}

func TestAssociateProductsWithError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[uuidtest]{
		{
			Description: "empty uuid",
			Handler:     func(t *uuidtest) { t.uuid = "                    " },
		},
		{
			Description: "product uuid equal than tag uuid",
			Handler:     func(t *uuidtest) { t.uuid = data.uuid.String() },
		},
		{
			Description: "product uuid already registered in tag",
			Handler:     func(t *uuidtest) { t.uuid = product1 },
		},
		{
			Description: "invalid uuid format",
			Handler:     func(t *uuidtest) { t.uuid = "0192cb012-7aad-75fa-92cf-bd9e900b66dd" },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Description, func(t *testing.T) {
			uuidForTest := uuidtest{}
			dataCopy := data
			tt.Handler(&uuidForTest)

			// Act
			err := dataCopy.AssociateProduct(uuidForTest.uuid)

			// Assert
			require.Error(t, err)
		})
	}
}

func TestDisassociateProductNoError(t *testing.T) {
	// Arrange
	dataCopy := data

	// Act
	err := dataCopy.DisassociateProduct(product3)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, dataCopy.associatedProduct, 2)
	assert.NotContains(t, dataCopy.associatedProduct, product3)
}

func TestDisassociateProductWithError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[uuidtest]{
		{
			Description: "empyt uuid",
			Handler:     func(u *uuidtest) { u.uuid = "              " },
		},
		{
			Description: "uuid not associated",
			Handler:     func(u *uuidtest) { u.uuid = product4 },
		},
		{
			Description: "uuid equal tha uuid of the tag",
			Handler:     func(u *uuidtest) { u.uuid = data.UUID() },
		},
		{
			Description: "uuid with invalid format",
			Handler:     func(u *uuidtest) { u.uuid = "321-5dwed650-5EWF54f" },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Description, func(t *testing.T) {
			dataCopy := data
			uuidFortest := uuidtest{}
			tt.Handler(&uuidFortest)

			// Act
			err := dataCopy.DisassociateProduct(uuidFortest.uuid)

			// Assert
			assert.Error(t, err)
			assert.Equal(t, dataCopy, 3)
		})
	}
}
