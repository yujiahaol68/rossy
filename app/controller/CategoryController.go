package controller

import (
	"github.com/yujiahaol68/rossy/app/entity"
	categoryService "github.com/yujiahaol68/rossy/app/service/category"
)

type CategoryController struct{}

var Category CategoryController = CategoryController{}

func (ctrl *CategoryController) Create(name string) error {
	c := new(entity.Category)
	c.Name = name
	return categoryService.InsertOne(c)
}

func (ctrl *CategoryController) List() []*entity.Category {
	return categoryService.List()
}
