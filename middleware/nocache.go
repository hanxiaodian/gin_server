package middleware

import (
	"github.com/gin-gonic/gin"
)

func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Next()
}
