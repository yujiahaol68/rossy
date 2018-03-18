package database_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/entity"
)

const (
	tableSource = "./data/Rossy_Table_source.json"
)

func setup() error {
	database.Open()
	return database.Sync()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func Test_MockDB(t *testing.T) {
	// Source Table
	sourceRows := make([]entity.Source, 500)
	byt, err := ioutil.ReadFile(tableSource)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(byt, &sourceRows); err != nil {
		panic(err)
	}

	for i, row := range sourceRows {
		row.Alias = fmt.Sprintf("%s-%d", row.Alias, i+1)
		_, err = database.Conn().Insert(&row)
		if err != nil {
			t.Fatal(err)
		}
	}
}
