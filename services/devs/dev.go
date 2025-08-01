package devs

import (
	"linkmatch-be/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedImage(db *gorm.DB) {
	images := []map[string]interface{}{
		{
			"image_url": "https://c.inilah.com/reborn/2025/05/Profil_Christy_No_Na_Foto_mnowid_44a6f015c3.webp",
			"username":  "a",
			"position":  1,
		},
		{
			"image_url": "https://akcdn.detik.net.id/community/media/visual/2025/05/20/nona-christy-1747714126527_34.jpeg?w=375",
			"username":  "a",
			"position":  2,
		},
		{
			"image_url": "https://image.idntimes.com/post/20250430/tangkapan-layar-2025-04-30-pukul-083847-48ba257d2481528c8db14fb610168355.png",
			"username":  "a",
			"position":  3,
		},
	}

	for _, image := range images {
		db.Create(&models.Image{
			ID:       uuid.New().String(),
			ImageURL: image["image_url"].(string),
			Username: image["username"].(string),
			Position: int(image["position"].(int)),
		})
	}
}
