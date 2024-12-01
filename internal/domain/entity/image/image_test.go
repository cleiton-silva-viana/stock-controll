package img

import (
	"testing"

	imagemock "stock-controll/test/mock/entity/image"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNoError(t *testing.T) {
	// Arrange
	file, err := imagemock.FileJPEG()
	require.NoError(t, err)

	const width = uint(100)
	const heigth = uint(100)

	validator := &imagemock.Validator{}
	validator.On("Dimension", width, heigth).Return(nil)
	validator.On("Resolution", width, heigth).Return(nil)

	// Act
	image, err := new(file, validator)

	// Arrange
	require.NoError(t, err)
	require.NotNil(t, image)
	assert.Equal(t, "jpeg", image.Extension())
}

func TestNewWhitError(t *testing.T) {
	t.Run("The file is nil", func(t *testing.T) {
		// Arrange
		validator := &imagemock.Validator{}

		// Act
		image, err := new(nil, validator)

		// Assert
		assert.Nil(t, image)
		require.Error(t, err)
	})

	t.Run("The file extension is invalid", func(t *testing.T) {
		// Arrange
		file, err := imagemock.FileGIF()
		require.NoError(t, err)

		validator := &imagemock.Validator{}

		// Act
		image, err := new(file, validator)

		// Arrange
		assert.Nil(t, image)
		require.Error(t, err)
	})
}
