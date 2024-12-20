package productimg

import (
	"image"
	"mime/multipart"

	"stock-controll/internal/domain/entity/image/img"
	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/valueobject/uuid"
)

/*

Deve ser usado algum mecanismos que ofusque, borre ou torne transparente o fundo da imagem

exemplo
coca-cola: {
	"654896964-547987-49867-9879874.698-87"
	metadatas: {
		title: "coca-cola 2L s/a"
		description: "coca-cola sem açúcar de 2L"
		keyWords: ["coca-cola", "sem açúcar", "2 litros", "refrigerante"]
	}
	imageResolutions: {
		"400x400": "localhost:8000/image/product/123496514gre79"
		"800x800": "localhost:8000/image/product/1234654r9eger9"
		"1200x1200": "localhost:8000/image/product/65v4r5gre6545"
	}
}

*/

type Product struct {
	productUUID       string
	metadatas         Metadatas
	imagesResolutions map[string]image.Image
}

func createSizeConfiguration(name string, width, heigth uint) img.ResizeConfig {
	return img.ResizeConfig{
		Name: name,
		Size: img.Size{
			Width:  width,
			Height: heigth,
		},
		Mode: img.KeepAspectRatioByWidth,
		ConverterConfig: img.ConverterConfig{
			TargetFormat:    "webp",
			Compress:        true,
			Lossless:        true,
			CompressQuality: 80,
		},
	}
}

var productSmallSizeConfig = createSizeConfiguration("small", 150, 150)
var productMediumSizeConfig = createSizeConfiguration("medium", 400, 400)
var productBigSizeConfig = createSizeConfiguration("big", 800, 800)

var sizes = []img.ResizeConfig{
	productSmallSizeConfig,
	productMediumSizeConfig,
	productBigSizeConfig,
}

type ConfigsProductImage struct {
	ProductUUID string
	File        multipart.File
	Metadatas   struct {
		Title       string
		Description string
		Keywords    []string
	}
	Resizer   img.IResizer
	Converter img.IConverter
}

func New(config ConfigsProductImage) (*Product, error) {

	entityErrs := entity.Error("product image")
	entityErrs.AddValidationError(uuid.IsValid("productUUID", config.ProductUUID))

	metadatas, errs := NewMetadatas(config.Metadatas.Title, config.Metadatas.Description, config.Metadatas.Keywords)
	if errs != nil {
		for _, e := range errs {
			entityErrs.AddValidationError(e)
		}
	}

	image, err := img.New(config.File, &ProductValidator{})
	if err != nil {
		entityErrs.AddValidationError(err)
		return nil, entityErrs
	}

	resizedImages, err := img.Resize(*image, config.Resizer, config.Converter, sizes)
	if err != nil {
		entityErrs.AddValidationError(err)
		return nil, entityErrs
	}

	return &Product{
		productUUID:       config.ProductUUID,
		metadatas:         *metadatas,
		imagesResolutions: resizedImages,
	}, nil
}
