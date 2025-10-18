package index

import (
	"SmartSearch/internal/service/index_service"
	"SmartSearch/internal/service/user_index_service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteIndexHandler(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	indexName := queryParams.Get("indexName")
	userId := queryParams.Get("userId")

	if indexName == "" || userId == "" {
		log.Println("failed to read query parameters: ", "no index name or user id provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	hasIndex, err := user_index_service.IsUserHasIndex(userId, indexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	if !hasIndex {
		c.JSON(http.StatusNotFound, gin.H{"error": "index not found"})
		return
	}

	indexName, err = index_service.DeleteIndex(indexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"Index deleted": indexName})
}
