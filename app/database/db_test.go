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
	tableSource   = "./data/source.json"
	tablePost     = "./data/post.json"
	tableCategory = "./data/category.json"
)

var (
	realFeedUrl = [...]string{
		"http://feeds.bbci.co.uk/news/rss.xml",
		"http://www.ruanyifeng.com/blog/atom.xml",
		"https://xiequan.info/feed/",
		"https://tonybai.com/feed/",
		"http://tech.qq.com/web/rss_web.xml",
		"https://www.reddit.com/.rss",
		"https://www.reddit.com/r/news/.rss",
		"https://www.reddit.com/r/food/.rss",
		"https://www.reddit.com/r/football/.rss",
		"https://www.reddit.com/r/golang/.rss",
	}
)

func setup() error {
	database.Open()
	return database.Sync()
}

func TestMain(m *testing.M) {
	if _, e := os.Stat("rossy_data.db"); e == nil {
		os.Remove("rossy_data.db")
	}

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
	sourceRows := make([]entity.Source, 10)
	byt, err := ioutil.ReadFile(tableSource)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(byt, &sourceRows); err != nil {
		panic(err)
	}

	for i, row := range sourceRows {
		row.URL = realFeedUrl[i]
		row.Alias = fmt.Sprintf("%s-%d", row.Alias, i+1)
		_, err = database.Conn().Insert(&row)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Post Table
	postRows := make([]entity.Post, 200)
	byt, err = ioutil.ReadFile(tablePost)

	if err != nil {
		t.Fatal(err)
	}

	if err = json.Unmarshal(byt, &postRows); err != nil {
		panic(err)
	}

	for _, row := range postRows {
		_, err = database.Conn().Insert(&row)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Categopry Table
	categoryRows := make([]entity.Category, 6)
	byt, err = ioutil.ReadFile(tableCategory)

	if err != nil {
		t.Fatal(err)
	}

	if err = json.Unmarshal(byt, &categoryRows); err != nil {
		panic(err)
	}

	for _, row := range categoryRows {
		_, err = database.Conn().Insert(&row)
		if err != nil {
			t.Fatal(err)
		}
	}
}
