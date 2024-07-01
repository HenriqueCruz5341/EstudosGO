package main

import (
	"hexagonal-news-api/adapter/input/routes"
	"hexagonal-news-api/configuration/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger.Info("Starting Hexagonal News API")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error trying to init application on port 8080", err)
		return
	}
}
