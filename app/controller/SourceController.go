package controller

import (
	"errors"

	"github.com/yujiahaol68/rossy/app/database"

	"github.com/yujiahaol68/rossy/app/entity"
	"github.com/yujiahaol68/rossy/app/service/category"
	"github.com/yujiahaol68/rossy/feed"
)

type SourceController struct{}

var Source SourceController = SourceController{}

func (sc SourceController) Add(url string, categoryID int64) (*entity.Source, error) {
	source, err := feed.GetSourceByURL(url)

	if err != nil {
		return nil, err
	}

	s := &entity.Source{
		Category:     categoryID,
		URL:          url,
		ETag:         source.ETag,
		LastModified: source.LastModified,
		Alias:        source.Alias,
	}

	c := category.FindOne(categoryID)
	if c == nil {
		return nil, errors.New("ID out of range")
	}

	db := database.Conn()
	_, err = db.InsertOne(s)
	return s, err
}