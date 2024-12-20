package productimg

import (
	"stock-controll/internal/domain/services/error/field"
)

type ProductValidator struct{}

const ErrImageNotSquare = "ERR_IMAGE_NOT_SQUARE"

func (piv *ProductValidator) Dimension(width, height uint) error {
	if width != height {
		return &field.FieldError{
			FieldName: "image",
			CodeError: ErrImageNotSquare,
			InvalidValue: map[string]interface{}{
				"width": width,
				"height": height,
			},
		}
	}
	return nil
}

const ErrImageDimensionInvalid = "ERR_IMAGE_DIMENSIONS_INVALID"

func (piv *ProductValidator) Resolution(width, height uint) error {
	const minLength = 800

	if width < minLength || height < minLength {
		return &field.FieldError{
			FieldName: "image",
			CodeError: ErrImageDimensionInvalid,
			InvalidValue: map[string]interface{}{
				"min_length": minLength,
			},
		}
	}
	return nil
}
