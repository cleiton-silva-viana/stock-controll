package value_object

import (
	"testing"
	
	"stock-controll/internal/domain/failure"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var fake = faker.New()

func Test_NewEmail_WithValidEmail(t *testing.T) {
	// Arrange
	email := fake.Internet().Email()
	expected_type := Email{}

	// Act
	result, err := NewEmail(email)

	// Assert
	assert.Nil(t, err)
	assert.IsType(t, &expected_type, result)
}

func Test_NewEmail_WithInvalidEmailFormat(t *testing.T) {
	// Arrange
	tests := []struct {
		name string
		email string
	}{
		{"faltando símbolo @", "user.example.com"},
		{"faltando domínio", "user.example@"},
		{"faltando o nome de usuário", "@example.com"},
		{"domínio inválido", "@example..com"},
		{"espaços em branco", "user @example.com"},
		{"domínio sem TDL (Top-Level Domain)", "user@example"},
		{"uso de caracteres especiais no ínicio  do email", "@user@example.com"},
		{"uso de caracteres especiais no fim do email", "user@example.com."},
		{"incorrect format", "user@example,com."},
	}

	// Act
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := NewEmail(test.email)

			// Assert
			assert.Nil(t, result)
			assert.ErrorIs(t, err, failure.EmailWithInvalidFormat)
		})
	}
}

