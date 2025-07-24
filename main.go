package main

import (
	"fmt"
	"linkmatch-be/database"
	"linkmatch-be/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "linkmatch-be/docs" // Import generated docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	findEnv()
	port := os.Getenv("PORT")

	db := database.Connect()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.Default())

	routes.RouteInit(router, db)

	fmt.Println("Starting server on port", port)
	router.Run(":" + port)
}

func findEnv() { // find the .env file on the current root project ('../') and load it
	possibleEnv := []string{
		"./.env.example",
		"./.env.production",
		"./.env.development",
		"./.env.local",
		"./.env",
	}

	for _, env := range possibleEnv {
		if _, err := os.Stat(env); err == nil {
			fmt.Println("Loading", env)
			_ = godotenv.Load(env)
			return
		}
	}
}
