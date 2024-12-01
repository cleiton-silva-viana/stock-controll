package infrastructure

import (
	"image"
	img "stock-controll/internal/domain/entity/image"

	"github.com/chai2010/webp"
)

// convert(img image.Image, config ConverterConfig)

type ConverterWEBP struct{}

func (cw *ConverterWEBP) Convert(img image.Image, config img.ConverterConfig) (image.Image, error) {
	// var webpBuff bytes.Buffer

	webp.EncodeRGB(img, float32(config.CompressQuality))

}
