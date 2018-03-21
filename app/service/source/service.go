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
