package order

import (
	"testing"
	"time"

	"stock-controll/internal/domain/valueobject/uuid"
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type orderConfig struct {
	buyerUUID          string
	supplierUUID       string
	expectedDeliveryOn time.Time
	products           []Product
	status             OrderStatus
}

func Setup() *orderConfig {
	return &orderConfig{
		buyerUUID:          uuid.New().String(),
		supplierUUID:       uuid.New().String(),
		expectedDeliveryOn: time.Now().AddDate(0, 1, 0),
	}
}

func TestNewOrderNoError(t *testing.T) {
	tests := []unitary.TestField[orderConfig]{
		{
			Description: "",
			Handler:     func(oc *orderConfig) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			o := Setup()
			tt.Handler(o)

			// Act
			result, err := New(o.buyerUUID, o.supplierUUID, o.expectedDeliveryOn, nil)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, result)

		})
	}
}
