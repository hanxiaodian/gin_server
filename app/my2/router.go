package my2

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/captcha", captchaHandler)
}
