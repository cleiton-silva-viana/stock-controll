package image

// Os produtos devem possuir imagens para: 
// 		Desktop (imagens grandes)
//		Mobile (imagens pequenas)
// 		Carrinho de compras (imagens extra pequenas)
// A imagem dos produtos devem possuir o formato 1:1
// O fundo da imagem deve ser pensado posteriormente
type ProductImage struct {
	BaseImage
	productUUID string
	resolutions map[string]Image
}


