package discard

import (
	"testing"
	"time"

	"stock-controll/internal/domain/valueobject/uuid"
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type discardConfig struct {
	uuid        string
	checkerUUID string
	productUUID string
	batchUUID   string
	quantity    int
}

func Setup() *discardConfig {
	return &discardConfig{
		uuid:        uuid.New().String(),
		checkerUUID: uuid.New().String(),
		productUUID: uuid.New().String(),
		batchUUID:   uuid.New().String(),
		quantity:    10,
	}
}

func TestNewDiscardNoError(t *testing.T) {
	tests := []unitary.TestField[discardConfig]{
		{
			Description: "min quantity allowed",
			Handler:     func(dc *discardConfig) { dc.quantity = minQuantity },
		},
		{
			Description: "max quantity allowed",
			Handler:     func(dc *discardConfig) { dc.quantity = maxQuantity },
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			d := Setup()
			tt.Handler(d)

			// Act
			result, err := New(d.checkerUUID, d.productUUID, d.batchUUID, d.quantity)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, result)
			assert.NotEmpty(t, result.UUID.String())
			assert.Equal(t, d.productUUID, result.productUUID.String())
			assert.Equal(t, d.checkerUUID, result.checkerUUID.UUID())
			assert.Equal(t, d.batchUUID, result.batchUUID.UUID())
			assert.Equal(t, d.quantity, result.quantity)
		})
	}
}

func TestNewDiscarWithError(t *testing.T) {
	tests := []unitary.TestField[discardConfig]{
		{
			Description: "invalid batch uuid",
			Handler:     func(dc *discardConfig) { dc.batchUUID = "                " },
		},
		{
			Description: "invalid product uuid",
			Handler:     func(dc *discardConfig) { dc.productUUID = "fdq2f3f3928fr42873r342479832rf23r" },
		},
		{
			Description: "invalid checker uuid",
			Handler:     func(dc *discardConfig) { dc.checkerUUID = "_________________________________" },
		},
		{
			Description: "quantity is less than min allowed",
			Handler:     func(dc *discardConfig) { dc.quantity = minQuantity - 1 },
		},
		{
			Description: "quantity is greater than max allowed",
			Handler:     func(dc *discardConfig) { dc.quantity = maxQuantity + 1 },
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			d := Setup()
			tt.Handler(d)

			// Act
			result, err := New(d.checkerUUID, d.productUUID, d.batchUUID, d.quantity)

			// Assert
			assert.Nil(t, result)
			assert.Error(t, err)
		})
	}
}

func TestSetStatusDiscardNoError(t *testing.T) {
	// Arrange
	d := Setup()
	discardInstance, _ := New(d.checkerUUID, d.productUUID, d.batchUUID, d.quantity)

	// Act
	err := discardInstance.UpdateStatus(done)

	// Assert
	assert.NoError(t, err)
	assert.LessOrEqual(t, discardInstance.HeldIn(), time.Now())
	assert.GreaterOrEqual(t, discardInstance.HeldIn(), time.Now().Add(-100000))
	assert.Equal(t, done, discardInstance.Status())
}

func TestSetStatusDiscardWithError(t *testing.T) {
	// Arrange
	d := Setup()
	discardInstance, _ := New(d.checkerUUID, d.productUUID, d.batchUUID, d.quantity)
	discardInstance.UpdateStatus(done)

	// Act
	err := discardInstance.UpdateStatus(toDo)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, done, discardInstance.Status())
}
