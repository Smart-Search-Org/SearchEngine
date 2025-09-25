package index

import (
	"SmartSearch/internal/models/requests"
	"SmartSearch/internal/service/index_service"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateIndexHandler parameters: name
func CreateIndexHandler(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("failed to read body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var createIndexRequest requests.CreateIndexRequest
	if err := json.Unmarshal(jsonData, &createIndexRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	indexName, err := index_service.CreateIndex(createIndexRequest.IndexName)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Index created": indexName})
}
