package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup route group for the API
func api() {

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}
