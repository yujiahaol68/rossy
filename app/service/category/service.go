package category

import (
	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/entity"
)

func FindOne(categoryID int64) *entity.Category {
	db := database.Conn()

	s := new(entity.Category)
	db.Id(categoryID).Get(s)
	if s.ID == 0 {
		return nil
	}

	return s
}