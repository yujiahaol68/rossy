package handlers

import (
	"html"
	"net/http"
	"strconv"

	"github.com/yujiahaol68/rossy/app/model/checkpoint"

	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/controller"
	"github.com/yujiahaol68/rossy/app/service/category"
)

func PostNewCategory(c *gin.Context) {
	nc := new(checkpoint.PostCategory)
	err := c.BindJSON(nc)
	if err != nil {
		c.Abort()
		ResultFail(c, http.StatusBadRequest, err)
		return
	}

	err = controller.Category.Create(nc.Name)
	if err != nil {
		c.Abort()
		ResultFail(c, http.StatusBadRequest, "duplicated name")
		return
	}
	ResultOk(c, nil)
}

func GetCategoryList(c *gin.Context) {
	cl := controller.Category.List()
	ResultList(c, cl, int64(len(cl)))
}

func UpdateCategoryName(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Abort()
		ResultFail(c, http.StatusBadRequest, "params error")
		return
	}

	newName := c.Query("name")
	if newName == "" {
		c.Abort()
		ResultFail(c, http.StatusBadRequest, "empty query")
		return
	}
	newName = html.EscapeString(newName)

	err = category.ChangeName(id, newName)
	if err != nil {
		c.Abort()
		ResultFail(c, 500, "ID out of range")
		return
	}
	ResultOk(c, nil)
}
