package devc

import (
	"linkmatch-be/responses"
	"linkmatch-be/services/devs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SeedImage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		devs.SeedImage(db)

		responses.JSONSuccess(c.Writer, "Image seeded", nil, nil)
	}
}
