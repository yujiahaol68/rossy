package database

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func Conn() *xorm.Engine {
	return engine
}

func Open() {
	eng, err := xorm.NewEngine("sqlite3", "rossy_data.db")
	if err != nil {
		panic(err)
	}
	engine = eng

	err = Sync()
	if err != nil {
		panic(err)
	}
}

func OpenForTest() error {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "app") {
		wd = filepath.Dir(wd)
	}

	dbPath := filepath.Join(wd, "database", "rossy_data.db")

	eng, err := xorm.NewEngine("sqlite3", dbPath)
	if err != nil {
		return err
	}
	engine = eng

	err = Sync()
	if err != nil {
		return err
	}
	return nil
}
