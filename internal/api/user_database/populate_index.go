package user_database

import (
	"SmartSearch/internal/models/requests"
	"SmartSearch/internal/service/index_doc_service"
	"SmartSearch/internal/service/user_database_service"
	"SmartSearch/internal/service/user_index_service"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PopulateIndexHandler(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("failed to read body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var pir requests.PopulateIndexRequest
	if err := json.Unmarshal(jsonData, &pir); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	hasIndex, err := user_index_service.IsUserHasIndex(pir.UserId, pir.IndexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	if !hasIndex {
		c.JSON(http.StatusNotFound, gin.H{"error": "index not found"})
		return
	}

	extract, err := user_database_service.Extract(pir.Driver, pir.DSN, pir.Table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	log.Println("Information from user database was extracted successfully")

	err = index_doc_service.AddDocs(pir.IndexName, extract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	log.Println("All the docs from the user database were added successfully")

	err = user_database_service.Update(pir.UserId, pir.IndexName, extract[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": "index populated"})
}
