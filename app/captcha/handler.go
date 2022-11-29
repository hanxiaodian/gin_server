package captcha

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func captchaHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "your captcha is 123456!",
	})
}
