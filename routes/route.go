package routes

import (
	"linkmatch-be/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteInit(r *gin.Engine, db *gorm.DB) {
	r.Group("/auth").
		POST("/login", controllers.Login).
		POST("/register", controllers.Register(db))
}
