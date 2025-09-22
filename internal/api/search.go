package api

import (
	"SmartSearch/internal/search"
	"SmartSearch/internal/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Query string `json:"query"`
}

func SetupRouter(cfg utils.Config) *gin.Engine {
	r := gin.Default()
	r.POST("/search", func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
		}
		var query Query
		err = json.Unmarshal(jsonData, &query)
		results := search.FullTextSearch(query.Query)
		c.JSON(http.StatusOK, gin.H{"results": results})
	})
	return r
}
