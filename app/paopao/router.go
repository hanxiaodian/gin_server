package paopao

import (
	"github.com/gin-gonic/gin"

	"gin_server/app/paopao/year"
	"gin_server/conf/cors"
)

type PaoPaoWhitelist struct {
	YEAR cors.Whitelists

	GetOrigins func(path string) cors.Whitelists
}

var projectWhitelist PaoPaoWhitelist

func Routers(r *gin.RouterGroup) PaoPaoWhitelist {
	projectWhitelist.GetOrigins = getPaoPaoOrigins

	// 活动路由加载
	projectWhitelist.YEAR = year.Year_router(r)

	return projectWhitelist
}

func getPaoPaoOrigins(path string) cors.Whitelists {
	switch path {
	case "YEAR":
		return projectWhitelist.YEAR
	default:
		return cors.Whitelists{}
	}
}
