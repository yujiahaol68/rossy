package handlers

import (
	"net/http"

	"github.com/yujiahaol68/rossy/app/service/source"

	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/controller"
	"github.com/yujiahaol68/rossy/app/model/checkpoint"
)

func PostSource(c *gin.Context) {
	var newFeedSource checkpoint.PostSource
	err := c.ShouldBindJSON(&newFeedSource)
	if err != nil {
		ResultFail(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}

	_, err = controller.Source.Add(newFeedSource.URL, newFeedSource.Category)

	if err != nil {
		ResultFail(c, http.StatusFound, err)
	} else {
		ResultCreated(c)
	}
}

func GetUnreadSourceList(c *gin.Context) {
	data, err := source.UnreadList()
	if err != nil {
		ResultFail(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}

	ResultOk(c, data)
}
