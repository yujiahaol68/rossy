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

func InsertOne(c *entity.Category) error {
	db := database.Conn()

	_, err := db.Insert(c)
	return err
}

func List() []*entity.Category {
	db := database.Conn()

	cl := make([]*entity.Category, 0)

	db.Find(&cl)
	return cl
}

func ChangeName(id int64, newName string) error {
	db := database.Conn()

	s := new(entity.Category)
	s.Name = newName
	_, err := db.Id(id).Update(s)
	return err
}
