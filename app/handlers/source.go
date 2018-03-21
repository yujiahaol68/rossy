package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yujiahaol68/rossy/app/controller"
	"github.com/yujiahaol68/rossy/app/model/checkpoint"
)

func PostSource(c *gin.Context) {
	var newFeedSource checkpoint.PostSource
	err := c.ShouldBindJSON(&newFeedSource)
	if err != nil {
		log.Fatal(err)
		ResultFail(c, err)
		return
	}

	_, err = controller.Source.Add(newFeedSource.URL, newFeedSource.Category)

	if err != nil {
		log.Fatal(err)
		ResultFail(c, err)
		return
	}

	ResultOk(c, nil)
}
