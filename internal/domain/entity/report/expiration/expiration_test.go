package expiration

import (
	"fmt"
	"stock-controll/test/unitary"
	"testing"
	"time"
	
	"stock-controll/internal/domain/valueobject/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type expirationConfig struct {
	quantity     int
	reporterUUID string
	productUUID  string
	expireDate   time.Time
}

func Setup() *expirationConfig {
	return &expirationConfig{
		quantity:     10,
		reporterUUID: uuid.New().String(),
		productUUID:  uuid.New().String(),
		expireDate:   time.Now().AddDate(0, 0, 10),
	}
}

func TestNewExpirationNoError(t *testing.T) {
	tests := []unitary.TestField[expirationConfig]{
		{
			Description: "with min quantity allowed",
			Handler:     func(ec *expirationConfig) { ec.quantity = minQuantity },
		},
		{
			Description: "with max quantity allowed",
			Handler:     func(ec *expirationConfig) { ec.quantity = maxQuantity },
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Test function New of package expiration %s", tt.Description), func(t *testing.T) {
			// Arrange
			e := Setup()
			tt.Handler(e)

			// Act
			result, err := New(e.reporterUUID, e.productUUID, e.quantity, e.expireDate)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, result)
		})
	}
}

func TestNewExpirationWithError(t *testing.T) {
	tests := []unitary.TestField[expirationConfig]{
		{
			Description: "quantity is less than min allowed",
			Handler:     func(ec *expirationConfig) { ec.quantity = minQuantity  - 1},
		},
		{
			Description: "quantity is greater than allowed",
			Handler:     func(ec *expirationConfig) { ec.quantity = maxQuantity + 1 },
		},
		{
			Description: "expire date is in the past",
			Handler: func(ec *expirationConfig) { ec.expireDate = time.Now().AddDate(0, 0, -1) },
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Test function New of package expiration - error is expected, as the %s", tt.Description), func(t *testing.T) {
			// Arrange
			e := Setup()
			tt.Handler(e)

			// Act
			result, err := New(e.reporterUUID, e.productUUID, e.quantity, e.expireDate)

			// Assert
			assert.NoError(t, err)
			require.NotNil(t, result)
		})
	}
}