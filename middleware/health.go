package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	if c.Request.URL.Path == "/health" {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Golang zt server is running ~",
		})
		return
	}
	c.Next()
}
