package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/controller"
)

func GetCategoryList(c *gin.Context) {
	cl := controller.Category.List()
	ResultList(c, cl, int64(len(cl)))
}
