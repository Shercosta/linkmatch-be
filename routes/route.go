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
		// Nested group under /api/profile
		profile := api.Group("/profile")
		{
			profile.GET("/", controllers.Profile(db))
			profile.POST("/parse-cv", controllers.ParseResume())
			profile.PUT("/", controllers.UpdateProfile(db))
		}
	}
}
