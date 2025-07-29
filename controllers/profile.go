package controllers

import (
	"linkmatch-be/responses"
	"linkmatch-be/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Get Profile
// @Description Returns a JSON with the user that requested it
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Security Bearer
// @Router /api/profile [get]
func Profile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("username")

		user := services.GetUserPublic(db, username)

		responses.JSONSuccess(c.Writer, user, nil, nil)
	}
}
