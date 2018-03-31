package post

import (
	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/entity"
)

func Insert(p *entity.Post) error {
	db := database.Conn()
	_, err := db.InsertOne(p)
	return err
}
