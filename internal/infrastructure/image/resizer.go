package infrastructure

import (
	"image"

	"github.com/nfnt/resize"
)

type ResizeService struct{}

// Devo entender o tipo de interpolação de imagem ... 
func (rs *ResizeService) Resize(img image.Image, width, height uint) (image.Image, error) {
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)
	return resizedImg, nil
}
