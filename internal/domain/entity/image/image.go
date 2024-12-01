package img

/*
O que está abaixo deve ser considerado no use case???
Validação de URLs: Verifique se os URLs das imagens estão corretos e acessíveis. URLs inválidos podem resultar em imagens quebradas na promoção.
Versionamento: Considere implementar um sistema de versionamento para as imagens, especialmente se as promoções forem atualizadas com frequência. Isso ajuda a manter um histórico das alterações.
Data de Upload: O formato da data deve ser consistente e preferencialmente em um padrão ISO (YYYY-MM-DD) para facilitar a comparação e a ordenação.
*/

import (
	"bytes"
	"image"
	"io"
	"mime/multipart"

	validationerrors "stock-controll/internal/domain/services/error"
	"stock-controll/internal/domain/services/validate"
)

type Size struct {
	width  uint
	height uint
}

func (i *Size) Width() uint {
	return i.width
}

func (i *Size) Heigth() uint {
	return i.height
}

type Image struct {
	extension string
	Size
	image.Image
}

func (i *Image) Extension() string {
	return i.extension
}

type IValidator interface {
	Dimension(width, height uint) error
	Resolution(width, height uint) error
}

type Config struct {
	Image       multipart.File
	Title       string
	Description string
}

const ErrFileReadFailed = "ERR_FILE_READ_FAILED"

func new(file multipart.File, validator IValidator) (*Image, error) {
	imageErrors := validationerrors.New("image").
		AddValidationError(isNil(file))
	if imageErrors.HasError() {
		return nil, imageErrors
	}

	defer file.Close()

	buff, err := io.ReadAll(file)
	if err != nil {
		return nil, &validate.FieldError{
			FieldName: "image",
			CodeError: ErrFileReadFailed,
		}
	}

	img, format, err := image.Decode(bytes.NewReader(buff))
	if err != nil {
		return nil, err
	}

	imageErrors.AddValidationError(validateImageType(format))
	if imageErrors.HasError() {
		return nil, imageErrors
	}

	width := uint(img.Bounds().Dx())
	height := uint(img.Bounds().Dy())

	imageErrors.
		AddValidationError(validator.Dimension(width, height)).
		AddValidationError(validator.Resolution(width, height))
	if imageErrors.HasError() {
		return nil, imageErrors
	}

	return &Image{
		extension: format,
		Size: Size{
			width:  width,
			height: height,
		},
		Image: img,
	}, nil
}

const ErrImageRequired = "ERR_IMAGE_REQUIRED"

func isNil(file any) error {
	if file == nil {
		return &validate.FieldError{
			FieldName: "image",
			CodeError: ErrImageRequired,
		}
	}
	return nil
}

var ValidImageTypes = map[string]struct{}{
	"png":  {},
	"jpg":  {},
	"jpeg": {},
	"webp": {},
}

const ErrInvalidFileExtension = "ERR_INVALID_FILE_EXTENSION"

func validateImageType(fileType string) error {
	_, ok := ValidImageTypes[fileType]
	if !ok {
		return &validate.FieldError{
			FieldName: "image",
			CodeError: ErrInvalidFileExtension,
		}
	}
	return nil
}

type ResizeMode int

const (
	KeepAspectRatioByWidth ResizeMode = iota
	KeepAspectRatioByHeight
	KeepExactDimensions
)

type ResizeConfig struct {
	Name    string
	Mode    ResizeMode
	NewSize Size
	ConverterConfig
}

type IResizer interface {
	Resize(img image.Image, size Size) (image.Image, error)
}

func resize(img Image, resizer IResizer, converter IConverter, configs []ResizeConfig) (map[string]image.Image, error) {
	var resizedImages = make(map[string]image.Image, len(configs))

	for _, config := range configs {
		newSize := calculateDimensions(img, config.Mode, config.NewSize)
		resizedImage, err := resizer.Resize(img, newSize)
		if err != nil {
			return nil, err
		}

		convertedImage, err := converter.Convert(resizedImage, config.ConverterConfig)
		if err != nil {
			return nil, err
		}

		resizedImages[config.Name] = convertedImage
	}

	return resizedImages, nil
}

type DimensionCalculator func(aspectRatio, width, heigth uint) Size

var dimensionCalculator = map[ResizeMode]DimensionCalculator{
	KeepAspectRatioByWidth: func(aspectRatio, width, height uint) Size {
		return Size{width: width, height: width / aspectRatio}
	},
	KeepAspectRatioByHeight: func(aspectRatio, width, height uint) Size {
		return Size{width: height, height: height * aspectRatio}
	},
	KeepExactDimensions: func(aspectRatio, width, height uint) Size {
		return Size{width: width, height: height}
	},
}

func calculateDimensions(img Image, mode ResizeMode, size Size) Size {
	aspectRatio := uint(img.width / img.height)
	return dimensionCalculator[mode](aspectRatio, size.width, size.height)
}

type IConverter interface {
	Convert(img image.Image, config ConverterConfig) (image.Image, error)
}

type ConverterConfig struct {
	TargetFormat    string
	Compress        bool
	Lossless        bool
	CompressQuality uint
}
