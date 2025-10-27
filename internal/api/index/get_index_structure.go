package index

import (
	"SmartSearch/internal/service/user_database_service"
	"SmartSearch/internal/service/user_index_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndexStructure(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	indexName := queryParams.Get("indexName")
	userId := queryParams.Get("userId")

	hasIndex, err := user_index_service.IsUserHasIndex(userId, indexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	if !hasIndex {
		c.JSON(http.StatusNotFound, gin.H{"error": "index not found"})
		return
	}

	indexStructure, err := user_database_service.GetIndexStructure(userId, indexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusCreated, indexStructure)
}
