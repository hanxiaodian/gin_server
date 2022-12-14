package routers

import (
	"github.com/gin-gonic/gin"

	"gin_server/app/my2"
	"gin_server/app/paopao"
	"gin_server/conf/cors"
)

type ProjectWhitelist struct {
	PAOPAO       paopao.PaoPaoWhitelist
	ACQUIESCENCE cors.Whitelists
}

var ProjectWhitelists = new(ProjectWhitelist)

// 初始化
func InitRouter(group *gin.RouterGroup, project string) {
	switch project {
	case "paopao":
		ProjectWhitelists.PAOPAO = paopao.Routers(group)
	case "my2":
		my2.Routers(group)
	default:
		ProjectWhitelists.ACQUIESCENCE = cors.Whitelists{
			".seayoo.com",
			".seayoo.io",
			".kingsoft.com",
			".localhost(?::\\d*)?",
		}
	}
}

// 获取跨域白名单配置
func GetAllowOrigins(project string, path string) (origins cors.Whitelists) {
	switch project {
	case "paopao":
		var origins = ProjectWhitelists.PAOPAO.GetOrigins(path)
		var len = len(origins)
		if len == 0 {
			return ProjectWhitelists.ACQUIESCENCE
		}
		return origins
	// case "my2":
	// 	my2.Routers(group)
	default:
		return ProjectWhitelists.ACQUIESCENCE
	}
}
