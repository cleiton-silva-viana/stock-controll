package tag

import (
	"stock-controll/internal/domain/entity/common"
	validationError "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
	"time"
)

type TagStatus string

const (
	Active   TagStatus = "active"
	Inactive TagStatus = "inactive"
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
type TagType string

const (
	Promotion TagType = "promotion"
	Attribute TagType = "attribute"
	Category  TagType = "category"
	Brand     TagType = "brand"
	Seasonal  TagType = "seasonal" // achar nome análogo em inglês
)

type Tag struct {
	id                int
	name              string
	description       string
	tagType           TagType
	createdAt         time.Time
	updatedAt         time.Time
	status            TagStatus
	associatedProduct map[string]struct{}
}

func NewTag(name, description string, tagType TagType) (*Tag, *validationError.IValidationError) {

	tagErrors := validationError.NewValidationError("tag")
	tagErrors.
		AddValidationError(validateName(name)).
		AddValidationError(validateDescription(description)).
		AddValidationError(validateTagType(tagType))

	if tagErrors.HasError() {
		return nil, &tagErrors
	}

	return &Tag{
		name:        name,
		description: description,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
		status:      Active,
		tagType:     tagType,
	}, nil
}

func (t *Tag) GetName() string {
	return t.name
}

func (t *Tag) GetDescirption() string {
	return t.description
}

func (t *Tag) GetTagType() string {
	return string(t.tagType)
}

func (t *Tag) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t *Tag) GetUpdateAt() time.Time {
	return t.updatedAt
}

func (t *Tag) GetStatus() string {
	return string(t.status)
}

func (t *Tag) UpdateTag(name, description string, status TagStatus) *validationError.IValidationError {

	tagErrors := validationError.NewValidationError("tag")
	tagErrors.
		AddValidationError(validateName(name)).
		AddValidationError(validateDescription(description))

	if tagErrors.HasError() {
		return &tagErrors
	}

	t.name = name
	t.description = description
	t.status = status
	t.updatedAt = time.Now()
	return nil
}

func (t *Tag) GetAssociatedProductsUUIDs() []string {
	uuids := make([]string, 0, len(t.associatedProduct))
	for uuid := range t.associatedProduct {
		uuids = append(uuids, uuid)
	}
	return uuids
}

func (t *Tag) AssociateProduct(productID string) *validation.FieldError {
	uuidValid := common.IsValidUUUID(productID)
	if !uuidValid {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	if t.IsAssociated(productID) {
		return &validation.FieldError{
			CodeErrors: []string{string(validation.ErrUnknown)},
		}
	}

	t.associatedProduct[productID] = struct{}{}
	t.updatedAt = time.Now()
	return nil
}

func (t *Tag) DisassociateProduct(productID string) {
	if t.IsAssociated(productID) {
		delete(t.associatedProduct, productID)
		t.updatedAt = time.Now()
	}
}

func (t *Tag) IsAssociated(productID string) bool {
	_, exists := t.associatedProduct[productID]
	return exists
}

func validateName(name string) *validation.FieldError {
	return validation.Validate("name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(3, 50, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
}

func validateDescription(description string) *validation.FieldError {
	return validation.Validate("description", description,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(12, 120, validation.ErrUnknown),
	)
}

// TODO: implementar
func validateTagType(tagType TagType) *validation.FieldError {
	return nil
}
