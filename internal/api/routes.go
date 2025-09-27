package api

import (
	"SmartSearch/internal/api/index"
	"SmartSearch/internal/api/search"
	"SmartSearch/internal/api/user_database"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/search", search.SearchHandler)

	router.POST("/index", index.CreateIndexHandler)
	router.DELETE("/index", index.DeleteIndexHandler)

	router.POST("/index/docs", user_database.PopulateIndexHandler)
}
