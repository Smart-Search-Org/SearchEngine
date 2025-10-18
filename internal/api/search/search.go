package search

import (
	"SmartSearch/internal/models/requests"
	"SmartSearch/internal/service/search_service"
	"SmartSearch/internal/service/user_index_service"
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

	hasIndex, err := user_index_service.IsUserHasIndex(searchQuery.UserId, searchQuery.IndexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	if !hasIndex {
		c.JSON(http.StatusNotFound, gin.H{"error": "index not found"})
		return
	}

	results := search_service.FullTextSearch(searchQuery.Query, searchQuery.Filters, searchQuery.IndexName)
	c.JSON(http.StatusOK, gin.H{"results": results})
}
