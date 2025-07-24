package controllers

import (
	"linkmatch-be/requests"
	"linkmatch-be/responses"
	"linkmatch-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body requests.RegisterRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, err.Error(), nil)
			return
		}

		result := services.Register(
			db,
			c,
			&body,
		)

		if result != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, result.Error(), nil)
			return
		}

		responses.JSONSuccess(c.Writer, "Success", nil, nil)
	}
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}
