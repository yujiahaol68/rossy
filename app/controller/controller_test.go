package controller_test

import (
	"log"
	"os"
	"testing"

	"github.com/yujiahaol68/rossy/app/controller"
	"github.com/yujiahaol68/rossy/app/database"
)

const (
	exampleFeedURL = "https://blog.golang.org/feed.atom"
	noAvailbleURL  = "https://nurybe4652.cn/nothing"
)

func setup() error {
	return database.OpenForTest()
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

func Test_AddNewSource(t *testing.T) {
	_, err := controller.Source.Add(exampleFeedURL, 100)
	if err == nil {
		t.Fatalf("Should have error because ID out of range")
	}

	//_, er := controller.Source.Add(noAvailbleURL, 2)
	//log.Fatal(er)
	//assert.NotNil(t, er)

	_, e := controller.Source.Add(exampleFeedURL, 4)
	if e != nil {
		t.Fatal(e)
	}
}
