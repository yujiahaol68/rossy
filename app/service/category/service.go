package category

import (
	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/entity"
)

func FindOne(categoryID int64) *entity.Category {
	db := database.Conn()

	s := new(entity.Category)
	db.Id(categoryID).Get(s)
	if s.ID == categoryID {
		return s
	}

	return nil
}

func List() []*entity.Category {
	db := database.Conn()

	cl := make([]*entity.Category, 0)

	db.Find(&cl)
	return cl
}
