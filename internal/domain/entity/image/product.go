package img

import (
	"errors"
	"image"
	"mime/multipart"
)

/*
Os produtos devem possuir imagens nas seguintes tamanhos:
	- extra pequedas para serem utilizadas no carrinho de compras
	- pequenas para serem utilizadas nos cards
	- médias para serem utilizadas na página do produto

As imagens devem seguir o formato 1:1

Deve ser usado algum mecanimos que ofusque, borre ou torne transparente o fundo da imagem

exemplo
coca-cola: {
	"654896964-547987-49867-9879874.698-87"
	metadatas: {
		title: "coca-cola 2L s/a"
		description: "coca-cola sem açucar de 2L"
		keyWords: ["coca-cola", "sem açúcar", "2 litros", "refigerante"]
	}
	imageResolutions: {
		"400x400": "localhost:8000/image/product/123496514gre79"
		"800x800": "localhost:8000/image/product/1234654r9eger9"
		"1200x1200": "localhost:8000/image/product/65v4r5gre6545"
	}
}


/*
Keywords
Definição: Geralmente, "keywords" refere-se a palavras ou frases que são usadas para descrever o conteúdo de uma imagem de forma mais técnica ou orientada a SEO (Search Engine Optimization).
Uso: Se o foco é otimizar a busca e a indexação da imagem em motores de busca ou sistemas de gerenciamento de conteúdo, "keywords" pode ser mais apropriado.
Exemplo: Palavras-chave como "natureza", "paisagem", "verão" podem ser usadas para melhorar a visibilidade da imagem em pesquisas.


*/

type ProductImageMetadatas struct {
	title       string
	description string
	keyWords    []string
}

func (pim *ProductImageMetadatas) Title() string {
	return pim.title
}

func (pim *ProductImageMetadatas) Description() string {
	return pim.description
}

type ProductImage struct {
	productUUID      string
	metadatas        ProductImageMetadatas
	imageResolutions map[string]image.Image
}

func createSizeConfiguration(name string, width, heigth uint) ResizeConfig {
	return ResizeConfig{
		Name: name,
		NewSize: Size{
			width:  width,
			height: heigth,
		},
		Mode: KeepAspectRatioByWidth,
		ConverterConfig: ConverterConfig{
			TargetFormat:    "image/webp",
			Compress:        true,
			Lossless:        true,
			CompressQuality: 80,
		},
	}
}

var productImageSmallSizeConfig = createSizeConfiguration("small", 150, 150)
var productImageMediumSizeConfig = createSizeConfiguration("medium", 400, 400)
var productImageBigSizeConfig = createSizeConfiguration("big", 800, 800)

func NewProduct(productUUID, title, description string, keywords []string, file multipart.File, resizer IResizer, converter IConverter) (*ProductImage, error) {

	// Validar productUUID
	// validar título
	// validar descrição
	// validar as keywords

	img, err := new(file, &ProductImageValidator{})
	if err != nil {
		return nil, err
	}

	sizes := []ResizeConfig{
		productImageSmallSizeConfig,
		productImageMediumSizeConfig,
		productImageBigSizeConfig,
	}

	redizedImages, err := resize(*img, resizer, converter, sizes)
	if err != nil {
		return nil, err
	}

	return &ProductImage{
		productUUID: productUUID,
		metadatas: ProductImageMetadatas{
			title:       title,
			description: description,
			keyWords:    keywords,
		},
		imageResolutions: redizedImages,
	}, nil
}

type ProductImageValidator struct{}

// Retornar erro personalizado
func (piv *ProductImageValidator) Dimension(width, heigth uint) error {
	if width != heigth {
		return errors.New("Is not square")
	}
	return nil
}

func (piv *ProductImageValidator) Resolution(width, heigth uint) error {
	const minLength = 800

	if width < minLength || heigth < minLength {
		return errors.New("Image size is less than minimum allowed (%d)" /* minLength */)
	}
	return nil
}
