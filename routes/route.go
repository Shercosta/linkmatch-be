package routes

import (
	"linkmatch-be/controllers"
	"linkmatch-be/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteInit(r *gin.Engine, db *gorm.DB) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login(db))
		auth.POST("/register", controllers.Register(db))
	}

	secured := r.Group("/secure").Use(middlewares.AuthMiddleware())
	{
		secured.GET("/profile", controllers.Profile())
	}
}
