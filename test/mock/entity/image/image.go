package imagemock

import (
	"github.com/stretchr/testify/mock"
)

type Validator struct {
	mock.Mock
}

func (v *Validator) Dimension(width, height uint) error {
	args := v.Called(width, height)
	return args.Error(0)
}

func (v *Validator) Resolution(width, height uint) error {
	args := v.Called(width, height)
	return args.Error(0)
}
