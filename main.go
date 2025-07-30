// @title LinkMatch API
// @version 1.0
// @description This is the API documentation for the LinkMatch application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@linkmatch.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5995
// @BasePath /
// @schemes http

// @securityDefinitions.apikey Bearer
// @type apiKey
// @name Authorization
// @in header
// @description Enter your bearer token in the format **Bearer &lt;token>**

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
