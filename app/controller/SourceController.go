package controller

import (
	"errors"

	"github.com/yujiahaol68/rossy/app/database"

	"github.com/yujiahaol68/rossy/app/entity"
	"github.com/yujiahaol68/rossy/app/service/category"
	postService "github.com/yujiahaol68/rossy/app/service/post"
	sourceService "github.com/yujiahaol68/rossy/app/service/source"
	"github.com/yujiahaol68/rossy/feed"
)

type SourceController struct{}

var Source SourceController = SourceController{}

func (ctrl *SourceController) Add(url string, categoryID int64) (*entity.Source, error) {
	source, err := feed.GetSourceByURL(url)

	if err != nil {
		return nil, err
	}

	c := category.FindOne(categoryID)
	if c == nil {
		return nil, errors.New("ID out of range")
	}

	s := &entity.Source{
		Category:     categoryID,
		URL:          url,
		ETag:         source.ETag,
		LastModified: source.LastModified,
		Alias:        source.Alias,
		Kind:         source.Type,
	}

	es := sourceService.FindByURL(url)
	if es.ID != 0 {
		return nil, errors.New("Source exist")
	}

	_, err = database.Conn().InsertOne(s)

	go post.AddNewFeeder(feed.RequestCache[url], s.ID)

	return s, err
}

func (ctrl *SourceController) Unsubscribe(id int64) error {
	es := sourceService.FindOne(id)
	if es.ID == 0 {
		return errors.New("Not Found")
	}

	go sourceService.Del(es.ID)
	return postService.DelBySource(es.ID)
}
