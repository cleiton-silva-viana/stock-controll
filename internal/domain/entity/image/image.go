package image

/*

Acessibilidade: Considere adicionar atributos de acessibilidade, como alt text, para as imagens. Isso é importante para usuários com deficiência visual e também ajuda no SEO.

Validação de URLs: Verifique se os URLs das imagens estão corretos e acessíveis. URLs inválidos podem resultar em imagens quebradas na promoção.

Versionamento: Considere implementar um sistema de versionamento para as imagens, especialmente se as promoções forem atualizadas com frequência. Isso ajuda a manter um histórico das alterações.

Data de Upload: O formato da data deve ser consistente e preferencialmente em um padrão ISO (YYYY-MM-DD) para facilitar a comparação e a ordenação.

*/

import (
	"errors"
	"mime/multipart"
	validationerrors "stock-controll/internal/domain/entity/error"
	"stock-controll/internal/domain/validation"
	"time"
)

// Sempre manter proporção 1:1
/*
Imagens de Produto no Carrinho:
	Tamanho: 150x150 pixels a 300x300 pixels
	Esse tamanho é geralmente suficiente para exibir uma visualização clara do produto
	sem ocupar muito espaço na interface do carrinho.

Imagens de Produto em Alta Resolução:
	Se o seu carrinho de compras permite que os usuários vejam detalhes
	do produto ao passar o mouse ou clicar,
	você pode considerar usar imagens de 300x300 pixels a 500x500 pixels.
	Isso permite uma visualização mais detalhada sem comprometer a performance.

Otimização: Certifique-se de que as imagens estejam otimizadas para a web.
Isso significa que elas devem ter um tamanho de arquivo reduzido para garantir
tempos de carregamento rápidos, sem sacrificar a qualidade visual.

Responsividade: Se o seu site for responsivo,
considere como as imagens do carrinho se comportarão em diferentes tamanhos de tela.
Você pode precisar ajustar o tamanho das imagens para dispositivos móveis.

Aspecto Visual: Além do tamanho, preste atenção à proporção da imagem.
Uma proporção de 1:1 (quadrada) é comum, mas você também pode usar proporções retangulares,
dependendo do design do seu site. */

/*
Keywords
Definição: Geralmente, "keywords" refere-se a palavras ou frases que são usadas para descrever o conteúdo de uma imagem de forma mais técnica ou orientada a SEO (Search Engine Optimization).
Uso: Se o foco é otimizar a busca e a indexação da imagem em motores de busca ou sistemas de gerenciamento de conteúdo, "keywords" pode ser mais apropriado.
Exemplo: Palavras-chave como "natureza", "paisagem", "verão" podem ser usadas para melhorar a visibilidade da imagem em pesquisas.
*/

type ImageExtension string

const (
	JPG  ImageExtension = "jpg"
	JPEG ImageExtension = "jpeg"
	PNG  ImageExtension = "png"
	GIF  ImageExtension = "gif"
)

type ImageSize struct {
	Width  int
	Heigth int
}

type ImageStatus string

const (
	Active   ImageStatus = "active"
	Pending  ImageStatus = "pending"
	Archived ImageStatus = "archived"
)

type Image struct {
	imageUUID  string
	url        string         // caminho de onde a imagem é armazenad ano servidor
	size       ImageSize      // largura e altura da imagem em pixels
	extension  ImageExtension // tipo de extensão do arquivo, ex: JPEG, PNG, GIF
	status     ImageStatus    // ativa ou pendente
	UploadedAt time.Time
}

type BaseImage struct {
	name        string   // nome original da imagem
	title       string   // título para a imagem
	description string   // descrição da imagem
	keyWords    []string // palavras chaves - definir estratégias
}

/*
Processo de cadastro de uma imagem no sistema:

	Recebemos uma imagem enviada pelo front-end
	Checamos se o tamanho da imagem não excede um limite razoável (e.g: 2mb)
	Checamos a extensão no qual a imagem está
	Checamos as dimensões da imagem (se está 1:1 caso seja imagem de produto)
	Checamos o título da imagem
	Checamos o nome da imagem
	Checamos a descrição da imagem
	Checamos as key workds

	Realizamos a operação de otimização das imagens
	Geramos cópias para os tamanhos e formatos necessários
	Criamos as urls
	referenciamos as url

	instânciamos a imagem
	salvamos no banco de dados
*/
func NewImage(image multipart.File,title, name, description string, heigth, width int) (*Image, validationerrors.IValidationError) {}


// Verificar se o arquivo foi enviado.
// Verificar se o arquivo é realmente uma imagem.
// Verificar o tamanho do arquivo (opcional).
// Decodificar a imagem para garantir que ela é válida.

func (i *Image) validate() *validationerrors.IValidationError {

}

const (
	minNameLength = 3
	maxNameLength = 30
)

func validateName(name string) *validation.FieldError {
	return validation.Validate(
		"name", name,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minNameLength, maxNameLength, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
}

const (
	minTitleLength = 3
	maxTitleLength = 30
)

func validateTitle(title string) *validation.FieldError {
	return validation.Validate(
		"title", title,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minTitleLength, maxTitleLength, validation.ErrUnknown),
		validation.CheckSpecialChars(validation.Disallow, validation.ErrUnknown),
	)
}

const (
	minDescirptionLength = 12
	maxDescriptionLength = 60
)

func validateDescription(description string) *validation.FieldError {
	return validation.Validate(
		"description", description,
		validation.IsBlank(validation.ErrUnknown),
		validation.IsLengthInRange(minDescirptionLength, maxDescriptionLength, validation.ErrUnknown),
	)
}

func validateSize(length, heigth int) {
}

func OptimizeImage(image string) string {
	return ""
}

func isValidExtension(image string, expectedExtensions []ImageExtension) bool {
	return false
}
