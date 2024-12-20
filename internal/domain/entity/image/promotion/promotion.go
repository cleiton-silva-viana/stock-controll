package promotionimg

// abaixo o formato almejado
/* {
    "promotion_image": {
        "base_image": {
            "name": "promotion_hygiene_day",
            "title": "Promoção - dia da higiene",
            "description": "PROMOÇÃO - todos os produtos higiénicos com ao menos 5% de desconto",
            "key_words": [
                "promoção",
                "higiene",
                "desconto"
            ]
        },
        "devices": {
            "announcement": {
                "500x500": {
                    "image_id": 123,
                    "url": "https:localhost:1323.com/promotion_hygiene_day_500x500.png",
                    "size": {
                        "width": 500,
                        "heigth": 500
                    },
                    "ImageSize": 128,
                    "extension": "png",
                    "status": "active",
                    "UploadedAt": "2024-22-10"
                },
                "1200x1200": {
                    "image_id": 124,
                    "url": "https:localhost:1323.com/promotion_hygiene_day_1200x1200.png",
                    "size": {
                        "width": 1200,
                        "heigth": 1200
                    },
                    "ImageSize": 320,
                    "extension": "png",
                    "status": "active",
                    "UploadedAt": "2024-22-10"
                }
            },
            "banner": {
                "728x90": {},
                "300x250": {},
                "160x600": {},
                "468x60": {}
            },
            "carousel": {},
            "card": {}
        }
    }
} */

// As promoções não necessáriamente devem possuir uma imagem associada
// Mas quando o tiverem, devem respeitar uma série de critérios, tais quais:
// Devem possuir imagens para diversas resoluções diferentes
// 		Desktop
//		Modile
// Devem possuir formatos variados de acordo com o tipo de promoção
// 		Formato para:
//					banner,
//					anúncio
//					card


