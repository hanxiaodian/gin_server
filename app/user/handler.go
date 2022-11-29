package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello friends!",
	})
}
