package main

import (
	"jwt-auth/database"
	"jwt-auth/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load .env variables
	database.Load()

	// Connect to Database
	database.ConnectToDB()

	// Create Gin Router
	router := gin.Default()

	// define HTML pattern
	router.LoadHTMLGlob("views/*.html")
	router.Static("/public", "./public")

	// Allow for CORS so we can send information across origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	// Setup Routes
	routes.Setup(router)

	// Run application on port 3000
	router.Run(":3000")
}
