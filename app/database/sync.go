package database

import (
	"errors"

	"github.com/go-xorm/core"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yujiahaol68/rossy/app/entity"
)

// Sync every entity table
func Sync() error {
	if engine == nil {
		return errors.New("Lost DB connection")
	}

	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.ShowSQL(true)

	// Sync every table
	return engine.Sync2(
		new(entity.Source),
		new(entity.Post),
		new(entity.Category),
	)
}
