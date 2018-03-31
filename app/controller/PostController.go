package controller

import (
	postService "github.com/yujiahaol68/rossy/app/service/post"
	"github.com/yujiahaol68/rossy/feed"
)

type PostController struct{}

var post PostController = PostController{}

func (ctrl *PostController) AddNewFeeder(f feed.Feeder, source int64) {
	pl := f.Convert()

	for _, post := range pl {
		post.From = source
		postService.Insert(post)
		// error inform via web-socket
	}
}
