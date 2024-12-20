package credential

import (
	"testing"
	"time"

	"stock-controll/internal/domain/valueobject/uuid"
	"stock-controll/test/unitary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type credentialData struct {
	password string
	userUUID string
}

func Setup() *credentialData {
	return &credentialData{
		password: "AbCd24@$Aefg&56dv",
		userUUID: uuid.New().String(),
	}
}

func TestNewCredentialNoError(t *testing.T) {
	// Arrange
	var c = Setup()

	// Act
	credential, err := New(c.userUUID, c.password)

	// Arrange
	assert.NoError(t, err)
	require.NotNil(t, credential)
	assert.Equal(t, c.userUUID, credential.UUID())
	assert.LessOrEqual(t, credential.CreatedAt(), time.Now())
	assert.LessOrEqual(t, credential.UpdatedAt(), time.Now())
	assert.GreaterOrEqual(t, credential.CreatedAt().Unix(), time.Now().Add(-time.Minute*5).Unix())
	assert.GreaterOrEqual(t, credential.UpdatedAt().Unix(), time.Now().Add(-time.Minute*5).Unix())
}

func TestNewCredentialWithError(t *testing.T) {
	tests := []unitary.TestField[credentialData]{
		{
			Description: "invalid uuid",
			Handler:     func(cd *credentialData) { cd.userUUID = "#%$#%¨$#¨#$ERGERGERge6r5ger4" },
		},
		{
			Description: "invalid password",
			Handler:     func(cd *credentialData) { cd.password = "this is invalid pass" },
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			// Arrange
			c := Setup()
			tt.Handler(c)

			// Act
			result, err := New(c.userUUID, c.password)

			// Assert
			assert.Nil(t, result)
			assert.Error(t, err)
		})
	}
}
