package search

import (
	"SmartSearch/internal/models/requests"
	"SmartSearch/internal/service/search_service"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchHandler(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("failed to read body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var searchQuery requests.SearchRequest
	if err := json.Unmarshal(jsonData, &searchQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	results := search_service.FullTextSearch(searchQuery.Query, searchQuery.IndexName)
	c.JSON(http.StatusOK, gin.H{"results": results})
}
