package api

import (
	"SmartSearch/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg utils.Config) *gin.Engine {
	r := gin.Default()
	r.GET("/search", func(c *gin.Context) {
		results := "Hello smart search!"
		c.JSON(http.StatusOK, gin.H{"results": results})
	})
	return r
}
