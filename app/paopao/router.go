package paopao

import (
	"github.com/gin-gonic/gin"

	"gin_server/app/paopao/year"
)

func Routers(r *gin.RouterGroup) {
	// 活动路由加载
	year.Year_router(r)
}
