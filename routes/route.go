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

	r.Group("/api").Use(middlewares.AuthMiddleware()).
		GET("/profile", controllers.Profile(db)).
		POST("/profile/parse-cv", controllers.ParseResume())
}
