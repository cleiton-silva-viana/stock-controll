package promotion

import "github.com/stretchr/testify/mock"

type PromotionMock struct {
	mock.Mock
}

func Mock() *PromotionMock {
	return &PromotionMock{}
}
