package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/routers/api"
)

func register(r *gin.Engine) {
	api.Router(r)
}

func Run() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	register(router)
	router.Run(":3456")
}
