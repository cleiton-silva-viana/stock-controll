package productimg

import (
	"strings"

	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/error/list"
	"stock-controll/internal/domain/services/validate"
)

type Metadatas struct {
	title       string
	description string
	keywords    map[string]struct{}
}

func NewMetadatas(title, description string, keywords []string) (*Metadatas, []error) {
	var keys = make(map[string]struct{}, len(keywords))
	for _, keyword := range keywords {
		keys[strings.ToLower(keyword)] = struct{}{}
	}

	var errs = make([]error, 0, 3)

	err := validateTitle(title)
	if err != nil {
		errs = append(errs, err)
	}

	err = validateDescription(description)
	if err != nil {
		errs = append(errs, err)
	}

	err = validateKeywords(keys)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return &Metadatas{
		title:       title,
		description: description,
		keywords:    keys,
	}, nil
}

func (m *Metadatas) Title() string {
	return m.title
}

func (m *Metadatas) Description() string {
	return m.description
}

func (m *Metadatas) Keywords() []string {
	keywordsCopy := make([]string, len(m.keywords))
	for key := range m.keywords {
		keywordsCopy = append(keywordsCopy, key)
	}
	return keywordsCopy
}

const (
	minMetadatasTitleLength = 3
	maxMetadatasTitleLength = 24
)

func validateTitle(title string) error {
	err := validate.New(
		"title", title,
		validate.IsBlank(),
		validate.IsLengthInRange(minMetadatasTitleLength, maxMetadatasTitleLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err != nil {
		return err
	}
	return nil
}

const (
	minMetadatasDescriptionLength = 8
	maxMetadatasDescriptionLength = 54
)

func validateDescription(description string) error {
	err := validate.New(
		"description", description,
		validate.IsBlank(),
		validate.IsLengthInRange(
			minMetadatasDescriptionLength,
			maxMetadatasDescriptionLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
	if err != nil {
		return err
	}
	return nil
}

const (
	minKeyWordLength = 2
	maxKeyWordLength = 16
)

func validateKeyword(keyword string) error {
	return validate.New(
		"keyword", keyword,
		validate.IsBlank(),
		validate.IsLengthInRange(minKeyWordLength, maxKeyWordLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
}

const (
	minKeywordsQuantity      = 2
	maxKeywordsQuantity      = 10
	ErrKeywordsNotValidRange = "ERR_KEYWORDS_NOT_VALID_RANGE"
)

/*
	Definição: Geralmente, "keywords" refere-se a palavras ou frases que são usadas para descrever o conteúdo de uma imagem de forma mais técnica ou orientada a SEO (Search Engine Optimization).
	Uso: Se o foco é otimizar a busca e a indexação da imagem em motores de busca ou sistemas de gerenciamento de conteúdo, "keywords" pode ser mais apropriado.
	Exemplo: Palavras-chave como "natureza", "paisagem", "verão" podem ser usadas para melhorar a visibilidade da imagem em pesquisas.
*/

func validateKeywords(keywords map[string]struct{}) error {

	if len(keywords) < minKeywordsQuantity || len(keywords) > maxKeywordsQuantity {
		return &field.FieldError{
			FieldName: "keywords",
			CodeError: ErrKeywordsNotValidRange,
		}
	}

	var errs list.ListError
	for keyword := range keywords {
		err := validateKeyword(keyword)
		if err != nil {
			errs.AddError(keyword, err)
		}
	}

	if errs.HasError() {
		return &errs
	}
	return nil
}
