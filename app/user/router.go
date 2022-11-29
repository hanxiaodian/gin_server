package user

import (
	"github.com/gin-gonic/gin"
)

func Users(e *gin.Engine) {
	e.GET("/user", userHandler)
}
