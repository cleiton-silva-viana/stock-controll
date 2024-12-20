package promotion

import (
	"strings"
	"testing"
	"time"

	"stock-controll/internal/domain/services/error/field"

	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var config = ConfigPromotion{
	Name:              "COUPON10",
	Description:       "This promotion is destined for all employees and customers",
	ExtendedPromotion: 0,
	Status:            Scheduled, // TODO: implementar testes para o status!!!
	StartDate:         time.Now().AddDate(0, 0, 7),
	EndDate:           time.Now().AddDate(0, 0, 14),
	Discount:          nil,
}

func TestNewNoError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[ConfigPromotion]{
		{
			Description: "promotion name length is equals to min length allowed",
			Handler:     func(cp *ConfigPromotion) { cp.Name = strings.Repeat("b", minLengthForPromotionName) },
		},
		{
			Description: "promotion name length is equals to max length allowed",
			Handler:     func(cp *ConfigPromotion) { cp.Name = strings.Repeat("a", maxLengthForPromotionName) },
		},
		{
			Description: "promotion name with number",
			Handler:     func(cp *ConfigPromotion) { cp.Name = "CUPOM10" },
		},
		{
			Description: "description length is equal than minimum allowed",
			Handler:     func(cp *ConfigPromotion) { cp.Description = strings.Repeat("a", minDescriptionLength) },
		},
		{
			Description: "description length is equal than maximum allowed",
			Handler:     func(cp *ConfigPromotion) { cp.Description = strings.Repeat("b", maxDescriptionLength) },
		},
		{
			Description: "extend promotion is equal than maximum allowed",
			Handler:     func(cp *ConfigPromotion) { cp.ExtendedPromotion = 1 },
		},
		{
			Description: "start date is equal than minimum allowed",
			Handler: func(cp *ConfigPromotion) {
				startIn := time.Now()
				cp.StartDate = startIn
				cp.EndDate = startIn.AddDate(0, 0, 2)
			},
		},
		{
			Description: "start date is equal than maximum allowed",
			Handler: func(cp *ConfigPromotion) {
				startIn := time.Now().Add(maxStartDateOffset)
				cp.StartDate = startIn
				cp.EndDate = startIn.AddDate(0, 0, 6)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Description, func(t *testing.T) {
			dataCopy := config
			tt.Handler(&dataCopy)

			// Act
			promotion, err := New(dataCopy)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, promotion)
			assert.Equal(t, dataCopy.UUID, promotion.UUID())
			assert.Equal(t, dataCopy.Name, promotion.Name())
			assert.Equal(t, dataCopy.Description, promotion.Description())
			assert.Equal(t, dataCopy.ExtendedPromotion, promotion.ExtendedPromotion())
			assert.Equal(t, dataCopy.StartDate, promotion.StartDate())
			assert.Equal(t, dataCopy.EndDate, promotion.EndDate())
			assert.Equal(t, dataCopy.Status, promotion.Status())
			// TODO: testar o método get de promotion
		})
	}
}

func TestNewErrorCases(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[ConfigPromotion]{
		{
			Description: "invalid UUID format",
			Handler:     func(cp *ConfigPromotion) { cp.UUID = "invalid-uuid" },
		},
		{
			Description: "empty promotion name",
			Handler:     func(cp *ConfigPromotion) { cp.Name = "" },
		},
		{
			Description: "promotion name exceeds max length",
			Handler:     func(cp *ConfigPromotion) { cp.Name = strings.Repeat("a", maxLengthForPromotionName+1) },
		},
		{
			Description: "promotion name below min length",
			Handler:     func(cp *ConfigPromotion) { cp.Name = strings.Repeat("a", minLengthForPromotionName-1) },
		},
		{
			Description: "promotion name has special characters",
			Handler:     func(cp *ConfigPromotion) { cp.Name = "#bellow2" },
		},
		{
			Description: "empty description",
			Handler:     func(cp *ConfigPromotion) { cp.Description = "" },
		},
		{
			Description: "description is less than minimum allowed",
			Handler:     func(cp *ConfigPromotion) { cp.Description = strings.Repeat("a", minDescriptionLength-1) },
		},
		{
			Description: "description is greater than maximum allowed",
			Handler:     func(cp *ConfigPromotion) { cp.Description = strings.Repeat("b", maxDescriptionLength+1) },
		},
		{
			Description: "start date in the past",
			Handler:     func(cp *ConfigPromotion) { cp.StartDate = time.Now().AddDate(0, 0, -1) },
		},
		{
			Description: "end date before start date",
			Handler:     func(cp *ConfigPromotion) { cp.EndDate = cp.StartDate.AddDate(0, 0, -1) },
		},
		{
			Description: "start date and end date are the same",
			Handler:     func(cp *ConfigPromotion) { cp.EndDate = cp.StartDate },
		},
		{
			Description: "invalid promotion status",
			Handler:     func(cp *ConfigPromotion) { cp.Status = "invalid_status" },
		},
		{
			Description: "negative extended promotion value",
			Handler:     func(cp *ConfigPromotion) { cp.ExtendedPromotion = -1 },
		},
		{
			Description: "high extended promotion value",
			Handler:     func(cp *ConfigPromotion) { cp.ExtendedPromotion = 2 },
		},
		{
			Description: "null discount strategy",
			Handler:     func(cp *ConfigPromotion) { cp.Discount = nil },
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Description, func(t *testing.T) {
			dataCopy := config
			tt.Handler(&dataCopy)

			// Act
			promotion, err := New(dataCopy)

			// Assert
			assert.Nil(t, promotion)
			require.Error(t, err)
		})
	}
}

var data = Promotion{
	uuid:                             "0192ca23-cdc3-7dca-b23d-9593d8333f4f",
	name:                             "COUPON10",
	description:                      "This promotion is destined for all employees and customers",
	extendedPromotion:                0,
	status:                           Scheduled,
	startDate:                        time.Now().AddDate(0, 0, 7),
	endDate:                          time.Now().AddDate(0, 0, 14),
	IProductSpecificDiscountStrategy: nil, // TODO: concertar
}

func TestPromotion_CancelPromotion_NoError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Promotion]{
		{
			Description: "cancelling the promotion in progress",
			Handler: func(p *Promotion) {
				p.status = InProgress
			},
		},
		{
			Description: "cancelling the scheduled promotion",
			Handler: func(p *Promotion) {
				p.status = Scheduled
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			err := dataCopy.CancelPromotion()

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, Canceled, dataCopy.Status())
		})
	}
}

func TestPromotion_CancelPromotion_WithError(t *testing.T) {
	testCases := []unitary.TestField[Promotion]{
		{
			Description: "Promoção cancelada antes",
			Handler: func(p *Promotion) {
				p.status = Canceled
			},
		},
		{
			Description: "Promoção encerrada",
			Handler: func(p *Promotion) {
				p.status = Ended
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := data
			test.Handler(&dataCopy)

			err := dataCopy.CancelPromotion()

			assert.Error(t, err)
			assert.ErrorIs(t, err, &field.FieldError{})
			assert.Equal(t, data, dataCopy)
		})
	}

}

func TestPromotion_ExtendPromotionDate_NoError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Promotion]{
		{
			Description: "min date allowed",
			Handler: func(p *Promotion) {
				p.endDate = time.Now().AddDate(0, 0, 1)
			},
		},
		{
			Description: "max date allowed",
			Handler: func(p *Promotion) {
				p.endDate = time.Now().AddDate(0, 0, 21)
			},
		},
		{
			Description: "min date allowed",
			Handler: func(p *Promotion) {
				p.startDate = time.Now().AddDate(0, 0, 5)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			err := dataCopy.ExtendPromotionDate(dataCopy.endDate)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, data.EndDate(), dataCopy.EndDate())
			assert.Equal(t, dataCopy.ExtendedPromotion(), 1)
		})
	}
}

func TestPromotion_ExtendPromotionDate_WithError(t *testing.T) {
	// Arrange
	testCases := []unitary.TestField[Promotion]{
		{
			Description: "date in the past",
			Handler: func(p *Promotion) {
				p.endDate = time.Now().AddDate(0, 0, -1)
			},
		},
		{
			Description: "date is greather than max allowed",
			Handler: func(p *Promotion) {
				p.endDate = time.Now().AddDate(0, 0, 0).Add(maxDateExtension + 1)
			},
		},
		{
			Description: "date is zeroed",
			Handler: func(p *Promotion) {
				p.startDate = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			dataCopy := data
			test.Handler(&dataCopy)

			// Act
			err := dataCopy.ExtendPromotionDate(dataCopy.endDate)

			// Assert
			require.Error(t, err)
			assert.ErrorIs(t, err, &field.FieldError{})
			assert.Equal(t, data.EndDate(), dataCopy.EndDate())
			assert.Equal(t, dataCopy.ExtendedPromotion(), 0)
		})
	}
}
