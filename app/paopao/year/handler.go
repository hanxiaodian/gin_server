package year

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "your captcha is 123456!",
	})
}
