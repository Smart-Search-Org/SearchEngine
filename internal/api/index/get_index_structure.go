package index

import (
	"SmartSearch/internal/models/requests"
	"SmartSearch/internal/service/user_database_service"
	"SmartSearch/internal/service/user_index_service"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndexStructure(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("failed to read body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var indexStructureRequest requests.GetIndexStructureRequest
	if err := json.Unmarshal(jsonData, &indexStructureRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	hasIndex, err := user_index_service.IsUserHasIndex(indexStructureRequest.UserId, indexStructureRequest.IndexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	if !hasIndex {
		c.JSON(http.StatusNotFound, gin.H{"error": "index not found"})
		return
	}

	indexStructure, err := user_database_service.GetIndexStructure(indexStructureRequest.UserId, indexStructureRequest.IndexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusCreated, indexStructure)
}
