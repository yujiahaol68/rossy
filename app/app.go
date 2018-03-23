package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/routers/api"
)

func register(r *gin.Engine) {
	api.Router(r)
}

func Run() {
	database.Open()

	err := database.Sync()
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	register(router)
	router.Run(":3456")
}
