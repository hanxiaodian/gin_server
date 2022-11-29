package captcha

import (
	"github.com/gin-gonic/gin"
)

func Captchas(e *gin.Engine) {
	e.GET("/captcha", captchaHandler)
}
