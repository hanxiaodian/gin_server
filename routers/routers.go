package routers

import (
	"github.com/gin-gonic/gin"

	"gin_server/app/my2"
	"gin_server/app/paopao"
)

// 初始化
func InitRouter(group *gin.RouterGroup, project string) {
	switch project {
	case "paopao":
		paopao.Routers(group)
	case "my2":
		my2.Routers(group)
	default:
	}
}
