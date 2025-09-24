package api

import (
	"SmartSearch/internal/api/index"
	"SmartSearch/internal/api/search"
	"SmartSearch/internal/utils"

	"github.com/gin-gonic/gin"
)

func RegisterSearchRoutes(router *gin.Engine) {
	router.POST("/search", search.SearchHandler)
}

func RegisterIndexRoutes(router *gin.Engine) {
	router.POST("/index", index.IndexCreateHandler)
}

func SetupRouter(cfg utils.Config) *gin.Engine {
	r := gin.Default()

	RegisterSearchRoutes(r)
	RegisterIndexRoutes(r)

	return r
}
