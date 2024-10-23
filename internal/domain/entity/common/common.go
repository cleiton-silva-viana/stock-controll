package common

// Remover da camada de domínio dependencia externa

import (
	"regexp"
	"stock-controll/internal/domain/validation"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuid, _ := uuid.NewV7()
	return uuid.String()
}

func IsValidUUUID(uuid string) bool {
	const re = `^(\w{8})(\-)(\w{4})(\-)(\w{4})(\-)(\w{4})(\-)(\w{12})`
	var err = validation.Validate("uuid", uuid,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsFormatValid(regexp.MustCompile(re), validation.ErrUnknown),
	)
	return err == nil
}
