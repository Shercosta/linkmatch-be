package controllers

import (
	"linkmatch-be/responses"
	"linkmatch-be/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Users(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users := services.GetUsers(db)

		responses.JSONSuccess(c.Writer, users, nil, nil)
	}
}
