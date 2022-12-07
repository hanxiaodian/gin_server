package year

import "github.com/gin-gonic/gin"

func Year_router(g *gin.RouterGroup) {
	groupRouter := g.Group("year")

	groupRouter.GET("year-login", LoginHandler)
}
