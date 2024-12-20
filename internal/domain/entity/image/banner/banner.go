package bannerimg

import "stock-controll/internal/domain/entity/image/img"

/*
O tamanho padrão de um banner para desktop pode variar dependendo do tipo de banner e da plataforma onde será exibido. No entanto, alguns tamanhos comuns incluem:

	- Banner Leaderboard: 728 x 90 pixels
	- Banner Medium Rectangle: 300 x 250 pixels
	- Banner Wide Skyscraper: 160 x 600 pixels
	- Banner Full Banner: 468 x 60 pixels

Esses tamanhos são frequentemente utilizados em publicidade online, mas é sempre bom verificar as diretrizes específicas da plataforma onde o banner será veiculado, pois elas podem ter requisitos diferentes.

Para tanto, podemos adotar os seguintes tamanhos e formatos para banners:
	- Full width para desktops (e.g., 1920x1080 pixels)
	- Medium width para tablets (e.g., 1280x720 pixels)
	- Small width para smartphones (e.g., 800x450 pixels)
*/

type BannerImage struct {
	// BaseImage
	promotionUUID string // um banner deverá ser atribuído a uma promoção? pensar a respeito
	resolutions   map[string]img.Image
}
