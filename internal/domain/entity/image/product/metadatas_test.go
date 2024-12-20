package productimg

import (
	"strings"
	"testing"
	
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var metadadas = Metadatas{
	title:       "A valid title",
	description: "A valid description for this test",
}

var keywords = []string{"food", "diet", "green", "low calorie"}

func TestNewMetadatasNoError(t *testing.T) {
	testCases := []unitary.TestField[Metadatas]{
		{
			Description: "metadatas with min title length allowed",
			Handler:     func(m *Metadatas) { m.title = strings.Repeat("a", minMetadatasTitleLength) },
		},
		{
			Description: "metadatas with max title length allowed",
			Handler:     func(m *Metadatas) { m.title = strings.Repeat("b", maxMetadatasTitleLength) },
		},
		{
			Description: "metadatas with min description length allowed",
			Handler:     func(m *Metadatas) { m.description = strings.Repeat("c", minMetadatasDescriptionLength) },
		},
		{
			Description: "metadatas with max description length allowed",
			Handler:     func(m *Metadatas) { m.description = strings.Repeat("c", maxMetadatasDescriptionLength) },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			metadatasCopy := metadadas
			tt.Handler(&metadatasCopy)

			// Act
			result, err := NewMetadatas(metadatasCopy.title, metadatasCopy.description, keywords)

			// Assert
			assert.Nil(t, err)
			require.NotNil(t, result)
			assert.Equal(t, result.title, metadatasCopy.Title())
			assert.Equal(t, result.description, metadatasCopy.Description())
			assert.Len(t, result.keywords, len(keywords))
		})
	}
}

func TestNewMetadatasWithError(t *testing.T) {
	testCases := []unitary.TestField[Metadatas]{
		{
			Description: "title length less than min allowed",
			Handler:     func(m *Metadatas) { m.title = strings.Repeat("a", minMetadatasTitleLength-1) },
		},
		{
			Description: "title length greater than max allowed",
			Handler:     func(m *Metadatas) { m.title = strings.Repeat("a", maxMetadatasTitleLength+1) },
		},
		{
			Description: "title width special characters",
			Handler:     func(m *Metadatas) { m.title = "width $pecial Char$" },
		},
		{
			Description: "description length is less than allowed",
			Handler:     func(m *Metadatas) { m.description = strings.Repeat("a", minMetadatasDescriptionLength-1) },
		},
		{
			Description: "description length  is greater than allowed",
			Handler:     func(m *Metadatas) { m.description = strings.Repeat("a", maxMetadatasDescriptionLength+1) },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			metadatasCopy := metadadas
			tt.Handler(&metadatasCopy)

			// Act
			instance, err := NewMetadatas(metadatasCopy.title, metadatasCopy.description, keywords)

			// Assert
			assert.Nil(t, instance)
			require.NotNil(t, err)
			assert.Len(t, err, 1)
		})
	}
}

// Description: "min keywords quantity allowed"
// Description: "max keywords quantity allowed"
// Description: "keyword with compoust name"
// Description: "keyword with number"
// Description: "keywords quantity is less than min allowed",
// Description: "keywords quantity is greater than max allowed",
// Description: "keyword with special characters",
// Description: "empty keyword",
