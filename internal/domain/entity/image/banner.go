package image

/*
O tamanho padrão de um banner para desktop pode variar dependendo do tipo de banner e da plataforma onde será exibido. No entanto, alguns tamanhos comuns incluem:

Banner Leaderboard: 728 x 90 pixels
Banner Medium Rectangle: 300 x 250 pixels
Banner Wide Skyscraper: 160 x 600 pixels
Banner Full Banner: 468 x 60 pixels
Esses tamanhos são frequentemente utilizados em publicidade online, mas é sempre bom verificar as diretrizes específicas da plataforma onde o banner será veiculado, pois elas podem ter requisitos diferentes.
*/

type BannerImage struct {
	BaseImage
	promotionUUID string // um banner deverá ser atribuído a uma promoção? pensar a respeito
	resolutions   map[string]Image
}
