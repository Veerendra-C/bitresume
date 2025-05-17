package main

import (
	"bitresume/config"
	"bitresume/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	// CORS config
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	// Apply CORS before routing
	r.Use(cors.New(corsConfig))

	// Register routes to existing engine
	routes.RegisterRoutes(r)

	r.Run(":6001")
}
