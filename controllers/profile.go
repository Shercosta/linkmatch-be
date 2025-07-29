package controllers

import (
	"linkmatch-be/responses"

	"github.com/gin-gonic/gin"
)

// @Summary Get Profile
// @Description Returns a JSON with the username that requested it
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Security Bearer
// @Router /profile [get]
func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("username")
		responses.JSONSuccess(c.Writer, gin.H{
			"username": username,
			"message":  "You have access to a protected route",
		}, nil, nil)
	}
}
