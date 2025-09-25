package index

import (
	"SmartSearch/internal/service/index_service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndexHandler CreateIndexHandler parameters: name
func DeleteIndexHandler(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	indexName := queryParams.Get("indexName")
	if indexName == "" {
		log.Println("failed to read query parameters:", "no index name provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	indexName, err := index_service.DeleteIndex(indexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"Index deleted": indexName})
}
