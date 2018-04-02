package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/handlers"
)

// Router will register all the routes
func Router(router *gin.Engine) {
	apiRouter := router.Group("api")

	{
		source := apiRouter.Group("source")
		source.POST("/", handlers.PostSource)
		source.GET("/unread", handlers.GetUnreadSourceList)
	}

	{
		category := apiRouter.Group("categories")
		category.GET("/", handlers.GetCategoryList)
	}

	{
		post := apiRouter.Group("post")
		post.GET("/", handlers.GetPostList)
		post.GET("/unread", handlers.GetUnreadPostList)
		post.GET("/source/:id", handlers.GetSourcePostList)
	}
}
