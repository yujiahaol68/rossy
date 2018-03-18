package service

import (
	"github.com/go-xorm/xorm"
	"github.com/yujiahaol68/rossy/app/entity"
)

type Source struct{ *xorm.Engine }

func (db Source) FindOne(sourceID int64) *entity.Source {
	s := new(entity.Source)
	db.Id(sourceID).Get(s)
	return s
}
