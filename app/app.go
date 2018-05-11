package app

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/controller"
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

	// update every 20 mins
	go func() {
		fmt.Println("Feed update...")
		controller.Source.UpdateAll()
		fmt.Println("Wait for 20 min to update again")
		// TODO: User can choose update frequency
		updateTimer := time.NewTimer(20 * time.Minute)
		<-updateTimer.C
	}()

	router := gin.Default()
	router.Use(cors.Default())

	register(router)
	router.Run(":3456")
}
