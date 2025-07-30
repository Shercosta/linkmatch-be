package routes

import (
	"linkmatch-be/controllers"
	"linkmatch-be/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteInit(r *gin.Engine, db *gorm.DB) {
	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login(db))
		auth.POST("/register", controllers.Register(db))
	}

	// Protected API routes
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		api.GET("/profile", controllers.Profile(db))

		// Nested group under /api/profile
		profile := api.Group("/profile")
		{
			profile.POST("/parse-cv", controllers.ParseResume())
		}
	}
}
