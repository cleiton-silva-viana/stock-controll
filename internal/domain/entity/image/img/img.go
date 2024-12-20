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

	"stock-controll/internal/domain/services/error/field"
)

type IValidator interface {
	Dimension(width, height uint) error
	Resolution(width, height uint) error
}

type Image struct {
	extension string
	Size
	image.Image
}

type Size struct {
	Width  uint
	Height uint
}

func (i *Image) Extension() string {
	return i.extension
}

const (
	ErrImageRequired  = "ERR_IMAGE_REQUIRED"
	ErrFileReadFailed = "ERR_FILE_READ_FAILED"
)

func New(file multipart.File, validator IValidator) (*Image, error) {
	if file == nil {
		return nil, &field.FieldError{
			FieldName: "image",
			CodeError: ErrImageRequired,
		}
	}
	defer file.Close()

	buff, err := io.ReadAll(file)
	if err != nil {
		return nil, &field.FieldError{
			FieldName: "image",
			CodeError: ErrFileReadFailed,
		}
	}

	img, format, err := image.Decode(bytes.NewReader(buff))
	if err != nil {
		return nil, &field.FieldError{
			FieldName: "image",
			CodeError: ErrFileReadFailed,
		}
	}

	err = validateImageType(format)
	if err != nil {
		return nil, err
	}

	width := uint(img.Bounds().Dx())
	height := uint(img.Bounds().Dy())

	var errs = &field.FieldError{
		FieldName: "image",
	}

	dimensionErr := validator.Dimension(width, height)
	resolutionerr := validator.Resolution(width, height)

	if errs.HasError() {
		return nil, errs
	}

	return &Image{
		extension: format,
		Size: Size{
			Width:  width,
			Height: height,
		},
		Image: img,
	}, nil
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
		return &field.FieldError{
			FieldName: "image",
			CodeError: ErrInvalidFileExtension,
			InvalidValue: map[string]interface{}{
				"file_type": fileType,
			},
		}
	}
	return nil
}

type IResizer interface {
	Resize(img image.Image, size Size) (image.Image, error)
}

type ResizeMode int

const (
	KeepAspectRatioByWidth ResizeMode = iota
	KeepAspectRatioByHeight
	KeepExactDimensions
)

type ResizeConfig struct {
	Name string
	Mode ResizeMode
	Size
	ConverterConfig
}

func Resize(img Image, resizer IResizer, converter IConverter, configs []ResizeConfig) (map[string]image.Image, error) {
	var resizedImages = make(map[string]image.Image, len(configs))

	for _, config := range configs {

		newSize := calculateDimensions(img, config.Mode, config.Size)

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
		return Size{Width: width, Height: width / aspectRatio}
	},

	KeepAspectRatioByHeight: func(aspectRatio, width, height uint) Size {
		return Size{Width: width, Height: height * aspectRatio}
	},

	KeepExactDimensions: func(aspectRatio, width, height uint) Size {
		return Size{Width: width, Height: height}
	},
}

// mudar nome de size para não coincidir com o nome do pacote
func calculateDimensions(img Image, mode ResizeMode, newSize Size) Size {
	aspectRatio := uint(img.Width / img.Height)
	return dimensionCalculator[mode](aspectRatio, newSize.Width, newSize.Height)
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
