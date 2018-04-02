package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/yujiahaol68/rossy/app/service/post"

	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/model/checkpoint"
)

func GetPostList(c *gin.Context) {
	page := new(checkpoint.PageArg)
	err := c.BindQuery(page)
	if err != nil {
		ResultFail(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}
	pl, err := post.List(page.Size, page.From)
	if err != nil {
		ResultFail(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}

	ResultList(c, pl, int64(len(pl)))
}

func GetUnreadPostList(c *gin.Context) {
	page := new(checkpoint.PageArg)
	err := c.BindQuery(page)
	if err != nil {
		ResultFail(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}

	upl, err := post.UnreadList(page.Size, page.From)
	if err != nil {
		ResultFail(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}

	ResultList(c, upl, int64(len(upl)))
}

func GetSourcePostList(c *gin.Context) {
	sourceID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResultFail(c, http.StatusBadRequest, errors.New("invalid ID"))
		c.Abort()
		return
	}

	page := new(checkpoint.PageArg)
	err = c.BindQuery(page)
	if err != nil {
		ResultFail(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}

	spl, err := post.SourceList(sourceID, page.Size, page.From)
	if err != nil {
		ResultFail(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}

	ResultOk(c, spl)
}
