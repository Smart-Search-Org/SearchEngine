package user_database

import (
	"SmartSearch/internal/models/requests"
	"SmartSearch/internal/service/index_doc_service"
	"SmartSearch/internal/service/user_database_service"
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

	extract, err := user_database_service.Extract(pir.Driver, pir.DSN, pir.Table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	err = index_doc_service.AddDocs(pir.IndexName, extract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": "index populated"})
}
