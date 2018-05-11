package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/handlers"
	"github.com/yujiahaol68/rossy/socket"
)

// Router will register all the routes
func Router(router *gin.Engine) {
	apiRouter := router.Group("api")

	{
		source := apiRouter.Group("source")
		source.POST("/", handlers.PostSource)
		source.GET("/unread", handlers.GetUnreadSourceList)
		source.DELETE("/:id", handlers.Unsubscribe)
	}

	{
		category := apiRouter.Group("categories")
		category.POST("/", handlers.PostNewCategory)
		category.GET("/", handlers.GetCategoryList)
		category.PUT("/:id", handlers.UpdateCategoryName)
	}

	{
		post := apiRouter.Group("post")
		post.GET("/", handlers.GetPostList)
		post.PUT("/:id", handlers.MarkPost)
		post.GET("/unread", handlers.GetUnreadPostList)
		post.GET("/source/:id", handlers.GetSourcePostList)
	}

	{
		socket.Enable = true
		ws := apiRouter.Group("ws")
		ws.GET("/", func(c *gin.Context) {
			socket.Wshandler(c.Writer, c.Request)
		})
	}
}
