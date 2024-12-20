package tag

import (
	"time"

	"stock-controll/internal/domain/services/error/entity"
	"stock-controll/internal/domain/services/error/field"
	"stock-controll/internal/domain/services/validate"
	"stock-controll/internal/domain/valueobject/uuid"
)

/*
/*
Estrutura de uma Tag
ID da Tag:
Um identificador único para cada tag, que pode ser um número ou uma string. Isso ajuda a evitar duplicações e facilita a referência.

Nome da Tag:
O nome descritivo da tag, que deve ser claro e conciso. Por exemplo, "Eletrônicos", "Promoção", "Novo", "Verão", etc.

Descrição (opcional):
Uma breve descrição da tag, explicando seu propósito ou contexto. Isso pode ser útil para administradores ou para SEO.

Tipo de Tag (opcional):
A data em que a tag foi criada. Isso pode ser útil para rastrear a relevância e a atualização das tags ao longo do tempo.
Data de Atualização:

Status:
Um campo que indica se a tag está ativa ou inativa. Isso é útil para gerenciar tags que não são mais relevantes.

Número de Produtos Associados:
Um contador que indica quantos produtos estão associados a essa tag. Isso pode ajudar na gestão e na visualização da popularidade da tag.

6. Relevância
Validação de Conteúdo: As tags devem ser relevantes para o produto. Você pode implementar um sistema de validação que verifique se a tag está relacionada ao tipo de produto ao qual está sendo atribuída.

7. Status
Ativação/Inativação: Permita que apenas tags ativas sejam atribuídas a produtos. Tags inativas não devem ser exibidas na interface do usuário.

8. Aprovação
Revisão de Tags: Considere implementar um sistema de aprovação onde novas tags precisam ser revisadas e aprovadas por um administrador antes de serem publicadas.

11. Proibição de Palavras Proibidas
Lista de Palavras Proibidas: Crie uma lista de palavras ou frases que não podem ser usadas em tags, como termos ofensivos, marcas registradas ou palavras irrelevantes.

Tag example:

	{
	  "id": "1",
	  "nome": "Eletrônicos",
	  "descricao": "Produtos eletrônicos como smartphones, laptops e acessórios.",
	  "tipo": "categoria",
	  "data_criacao": "2023-01-01T12:00:00Z",
	  "data_atualizacao": "2023-01-15T12:00:00Z",
	  "status": "ativo",
	  "numero_produtos_associados": 150
	}

O que é Tipo de Tag?
O "Tipo de Tag" refere-se à categorização das tags que você pode usar em seu sistema de e-commerce. Diferentes tipos de tags podem servir a diferentes propósitos e ajudar a organizar os produtos de maneiras específicas. A definição de tipos de tags permite que você categorize e filtre produtos de forma mais eficaz.

Exemplos de Tipos de Tag
- Categorias:
Descrição: Tags que representam categorias amplas de produtos. Por exemplo,

"Eletrônicos", "Roupas", "Móveis".
Uso: Ajuda os clientes a navegar pelo site e encontrar produtos dentro de uma categoria específica.

Atributos:
Descrição: Tags que descrevem características específicas de um produto. Por exemplo, "Tamanho: M", "Cor: Azul", "Material: Algodão".
Uso: Permite que os clientes filtrem produtos com base em características específicas.

Promoções:
Descrição: Tags que indicam que um produto está em promoção ou em uma oferta especial. Por exemplo, "Desconto de 20%", "Oferta do Dia".
Uso: Ajuda a destacar produtos que estão em promoção, atraindo a atenção dos clientes.

Temporais:
Descrição: Tags que são relevantes apenas em determinados períodos, como "Coleção de Verão", "Natal 2023".
Uso: Permite que você destaque produtos que são sazonais ou que fazem parte de uma coleção específica.

Marcas:
Descrição: Tags que representam a marca do produto. Por exemplo, "Nike", "Apple", "Samsung".
Uso: Ajuda os clientes a encontrar produtos de marcas específicas.

Definindo Quais Tipos São Permitidos e Obrigatórios
Ao implementar tipos de tags, você deve definir quais tipos são permitidos e quais são obrigatórios para cada tag. Aqui estão algumas considerações:

Tipos Permitidos:
Flexibilidade: Permita que os administradores do sistema criem tags dentro dos tipos definidos. Por exemplo, se você tem um tipo "Atributo", os administradores podem criar tags como "Tamanho: P", "Cor: Vermelho", etc.
Limitação: Você pode restringir a criação de tags a apenas alguns tipos predefinidos, evitando a criação de tags irrelevantes ou confusas.
Tipos Obrigatórios:
Obrigatoriedade: Para garantir que todos os produtos sejam categorizados de maneira eficaz, você pode tornar certos tipos de tags obrigatórios. Por exemplo, ao adicionar um novo produto, pode ser obrigatório atribuir uma tag de "Categoria".

Exemplo: Se um produto é adicionado, ele deve ter pelo menos uma tag de "Categoria" e uma tag de "Atributo" (se aplicável), mas as tags de "Promoção" podem ser opcionais.

Implementação Prática
Interface de Usuário: Ao criar ou editar um produto, a interface deve permitir que o usuário selecione tipos de tags apropriados. Por exemplo, um menu suspenso pode listar os tipos de tags disponíveis.

Lista de Tags Pré-definidas:
Banco de Dados de Tags: Mantenha um banco de dados de tags pré-definidas que são relevantes para cada categoria de produto. Quando um administrador ou usuário tenta atribuir uma tag a um produto, o sistema pode verificar se a tag está na lista permitida para aquele tipo de produto.

Exemplo: Se um produto é uma "Camisa", as tags permitidas podem incluir "Tamanho: M", "Cor: Azul", "Material: Algodão". Se alguém tentar adicionar uma tag como "Eletrônicos", o sistema deve rejeitar essa tag.

Regras de Associação:
Regras Lógicas: Defina regras que determinam quais tags são apropriadas para quais produtos. Isso pode incluir a criação de um mapeamento entre categorias de produtos e suas tags relevantes.
Exemplo: Para a categoria "Calçados", as tags relevantes podem incluir "Tamanho: 42", "Estilo: Casual", enquanto tags como "Cor: Vermelho" podem ser permitidas apenas se a cor for realmente uma opção para aquele calçado.

Validação em Tempo Real:
Feedback Imediato: Ao adicionar ou editar tags, forneça feedback em tempo real ao usuário. Se uma tag não for relevante, o sistema pode exibir uma mensagem de erro ou aviso, ajudando o usuário a corrigir a entrada antes de salvar.
Exemplo: Se um usuário tentar adicionar uma tag irrelevante, como "Promoção" a um produto que não está em promoção, o sistema pode alertá-lo.

Aprovação de Tags:
Sistema de Revisão: Considere implementar um sistema onde novas tags precisam ser revisadas e aprovadas por um administrador antes de serem atribuídas a produtos. Isso garante que apenas tags relevantes sejam usadas.
*/

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)

type Type string

const (
	Promotion   Type = "promotion"
	Attribute   Type = "attribute"
	Category    Type = "category"
	Subcategory Type = "subcategory"
	Brand       Type = "brand"
	Seasonal    Type = "seasonal"
)

type Tag struct {
	uuid              uuid.UUID
	name              string // Atributo deve ser único no banco de dados
	description       string
	tagType           Type
	createdAt         time.Time
	updatedAt         time.Time
	status            Status
	associatedProduct map[string]struct{}
}

func New(name, description, tagType string) (*Tag, error) {

	tagErrors := entity.Error("tag").
		AddValidationError(validateName(name)).
		AddValidationError(validateDescription(description)).
		AddValidationError(validateTagType(tagType))

	if tagErrors.HasError() {
		return nil, tagErrors
	}

	return &Tag{
		uuid:        *uuid.New(),
		name:        name,
		description: description,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
		status:      Active,
		tagType:     Type(tagType),
	}, nil
}

func (t *Tag) UUID() string {
	return t.uuid.String()
}

func (t *Tag) Name() string {
	return t.name
}

func (t *Tag) Description() string {
	return t.description
}

func (t *Tag) TagType() string {
	return string(t.tagType)
}

func (t *Tag) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Tag) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Tag) Status() string {
	return string(t.status)
}

func (t *Tag) UpdateTag(name, description string, status Status) error {
	tagErrors := entity.Error("tag")
	tagErrors.
		AddValidationError(validateName(name)).
		AddValidationError(validateDescription(description))
	if tagErrors.HasError() {
		return tagErrors
	}
	t.name = name
	t.description = description
	t.status = status
	t.updatedAt = time.Now()
	return nil
}

func (t *Tag) AssociatedProductsUUIDs() []string {
	uuids := make([]string, 0, len(t.associatedProduct))
	for id := range t.associatedProduct {
		uuids = append(uuids, id)
	}
	return uuids
}

const ErrProductAlreadyAssociatedWithTag = "ERR_PRODUCT_ALREADY_ASSOCIATED_WITH_TAG"

func (t *Tag) AssociateProduct(productUUID string) error {
	err := uuid.IsValid("product_uuid", productUUID)
	if err != nil {
		return err
	}
	if t.IsAssociated(productUUID) {
		return &field.FieldError{
			CodeError: ErrProductAlreadyAssociatedWithTag,
		}
	}
	t.associatedProduct[productUUID] = struct{}{}
	t.updatedAt = time.Now()
	return nil
}

const ErrProductNotAssociatedWithTag = "ERR_PRODUCT_NOT_ASSOCIATED_WITH_TAG"

func (t *Tag) DisassociateProduct(productUUID string) error {
	if err := uuid.IsValid("product_uuid", productUUID); err != nil {
		return err
	}
	if t.IsAssociated(productUUID) {
		delete(t.associatedProduct, productUUID)
		t.updatedAt = time.Now()
		return nil
	}
	return &field.FieldError{
		FieldName: "product_uuid",
		CodeError: ErrProductNotAssociatedWithTag,
	}
}

func (t *Tag) IsAssociated(productUUID string) bool {
	_, exists := t.associatedProduct[productUUID]
	return exists
}

const (
	minNameLength = 3
	maxNameLength = 50
)

func validateName(name string) error {
	return validate.New("name", name,
		validate.IsBlank(),
		validate.IsLengthInRange(minNameLength, maxNameLength),
		validate.CheckSpecialChars(validate.Disallow),
	)
}

const (
	minDescriptionLength = 12
	maxDescriptionLength = 120
)

func validateDescription(description string) error {
	return validate.New("description", description,
		validate.IsBlank(),
		validate.IsLengthInRange(minDescriptionLength, maxDescriptionLength),
	)
}

const ErrInvalidTagType = "ERR_INVALID_TAG_TYPE"

func validateTagType(tagType string) error {
	switch Type(tagType) {
	case Promotion, Attribute, Category, Brand, Seasonal:
		return nil
	default:
		return &field.FieldError{
			FieldName: "tag_type",
			CodeError: ErrInvalidTagType,
		}
	}
}
