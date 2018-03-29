package source

import (
	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/entity"
)

func FindOne(sourceID int64) *entity.Source {
	db := database.Conn()
	s := new(entity.Source)
	db.Id(sourceID).Get(s)
	return s
}

func FindByURL(url string) *entity.Source {
	db := database.Conn()
	s := new(entity.Source)
	db.Cols("id", "url").Where("source.url = ?", url).Get(s)
	return s
}

func Del(sourceID int64) {
	db := database.Conn()
	db.Id(sourceID).Delete(new(entity.Source))
}
