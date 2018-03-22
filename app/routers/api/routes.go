package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/handlers"
)

func Router(router *gin.Engine) {
	apiRouter := router.Group("api")

	{
		source := apiRouter.Group("source")
		source.POST("/", handlers.PostSource)
	}
}
