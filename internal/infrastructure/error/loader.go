package loader

/*
	Este pacote tem por finalidade retornar uma esturutra contendo um map de erros usados na aplicação
	Tal função que retorna um map, deve ser utilizada na inicialização do sistema
	Deve ser um singleton

	o map:
	validationErrs: {
		"ERR_FILED_IS_EMPTY": {

		},
		"ERR_FIELD_IS_SHORT": {

		},
		...
	}
*/

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var (
	once           sync.Once
	ValidationErrs ValidationErrors
)

type ValidationErrors struct {
	FieldName string
	Erros     map[string]ValidationError `json:"validationErrors"`
}

func LoadValidationErrors() {
	once.Do(func() {
		filePath := getValidationErrorsFilePath()
		loadValidationErrors(filePath)
	})
}

func loadValidationErrors(filePath string) {
	jsonDatas, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonDatas, &ValidationErrs)
	if err != nil {
		log.Fatal(err)
	}
}

func getValidationErrorsFilePath() string {
	/* 	exePath, err := os.Executable()
	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   	exeDir := filepath.Dir(exePath)
	   	return filepath.Join(exeDir, "../../../i18n/errors_pt.json") */
	// return "c:\\Users\\Cleit\\OneDrive\\Documentos\\stock-controll\\i18n\\errors_pt.json"
	return "C:\\Users\\Inara\\Documents\\stock-controll\\internal\\infrastructure\\i18n\\errors_pt.json"
}

// TODO: Implementar
func (ve *ValidationErrors) Error() string {
	return " "
}

func (ve *ValidationErrors) ContainsErrors() bool {
	return len(ve.Erros) > 0
}

type ValidationError struct {
	Message  string `json:"Message"`
	Solution string `json:"Solution"`
	Details  string `json:"Details"`
}

// Função abaixo deve se rusada em pacote diferente ...
func GetValidationErrorDescription(code errorCode) *ValidationError {
	LoadValidationErrors()
	err, exists := ValidationErrs.Erros[string(code)]
	if !exists {
		return &ValidationError{
			Message: string(ErrUnknow),
			Details: "report this occurency to suport area",
		}
	}
	return &err
}

// Refatorar
func GetValidationError(fieldName string, code errorCode) *ValidationCodeErrors {
	LoadValidationErrors()

	var ve = ValidationCodeErrors{
		FieldName:  fieldName,
		CodeErrors: make([]string, 5),
	}

	return &ve
}
