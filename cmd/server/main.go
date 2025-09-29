package main

import (
	"SmartSearch/internal/api"
	"SmartSearch/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	utils.LoggingMiddleware()
	api.RegisterRoutes(r)

	return r
}

func main() {
	cfg := utils.LoadConfig()
	router := SetupRouter()
	log.Println("Smart Search running on", cfg.Server.Port)

	utils.LoadIndexesFromDisk()
	log.Println("Loaded indexed in memory")

	err := router.Run(":" + cfg.Server.Port)
	if err != nil {
		log.Println("Error starting server:", err)
	}
}
