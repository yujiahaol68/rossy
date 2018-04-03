package post_test

import (
	"log"
	"os"
	"testing"

	"github.com/yujiahaol68/rossy/app/service/post"

	"github.com/yujiahaol68/rossy/app/database"
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

func Test_readList(t *testing.T) {
	_, err := post.List(1, 4)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Del(t *testing.T) {
	err := post.DelBySource(int64(1))
	if err != nil {
		t.Fatal(err)
	}
}
