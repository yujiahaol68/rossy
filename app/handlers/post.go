package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/yujiahaol68/rossy/app/service/post"

	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/model/checkpoint"
	"github.com/yujiahaol68/rossy/app/model/thirdparty"
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

	ResultList(c, spl, int64(len(spl)))
}

func MarkPost(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResultFail(c, http.StatusBadRequest, errors.New("invalid ID"))
		c.Abort()
		return
	}

	err = post.MarkRead(id)
	if err != nil {
		log.Fatal(err)
		ResultFail(c, http.StatusBadRequest, err)
		c.Abort()
		return
	}

	ResultOk(c, nil)
}

func ParseFullPost(c *gin.Context) {
	u := c.Query("url")
	if u == "" {
		ResultOk(c, nil)
		c.Abort()
		return
	}

	p := thirdparty.NewParser()
	err := p.ParseURL(u)
	if err != nil {
		ResultFail(c, http.StatusInternalServerError, errors.New("parse URL Failed"))
		c.Abort()
		return
	}

	bs, err := p.Bytes()
	if err != nil {
		ResultFail(c, http.StatusInternalServerError, errors.New("parse text Failed"))
		c.Abort()
		return
	}

	c.Data(200, "application/json; charset=utf-8", bs)
}
