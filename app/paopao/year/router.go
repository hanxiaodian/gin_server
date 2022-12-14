package year

import (
	"gin_server/conf/cors"
	"log"

	"github.com/gin-gonic/gin"
)

func Year_router(g *gin.RouterGroup) (whitelists cors.Whitelists) {
	groupRouter := g.Group("year")

	groupRouter.GET("year-login", LoginHandler)

	var origins = cors.Whitelists{
		".seayoo.com",
		".seayoo.io",
		".kingsoft.com",
		".localhost(?::\\d*)?",
	}

	log.Print("Year_router origins:  ", origins)

	return origins
}
