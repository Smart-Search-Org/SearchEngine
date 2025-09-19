package main

import (
	"SmartSearch/internal/api"
	"SmartSearch/internal/utils"
	"log"
)

func main() {
	cfg := utils.LoadConfig()
	router := api.SetupRouter(cfg)
	log.Println("AI Search running on", cfg.Server.Port)

	err := router.Run(":" + cfg.Server.Port)
	if err != nil {
		log.Println("Error starting server:", err)
	}
}
