package routes

import (
	"linkmatch-be/controllers"
	"linkmatch-be/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteInit(r *gin.Engine, db *gorm.DB) {

	r.Group("/auth").
		POST("/login", controllers.Login(db)).
		POST("/register", controllers.Register(db))

	r.Group("/secure").Use(middlewares.AuthMiddleware()).
		GET("/profile", controllers.Profile())

}
