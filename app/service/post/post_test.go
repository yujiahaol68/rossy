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
	li, err := post.List(20, 0)
	if err != nil {
		t.Fatal(err)
	}
	for i, p := range li {
		t.Logf("\n%d: %s", i, p.Title)
	}
}

func Test_unreadList(t *testing.T) {
	li, err := post.UnreadList(20, 0)
	if err != nil {
		t.Fatal(err)
	}
	for i, p := range li {
		t.Logf("\n%d: %s", i, p.Title)
	}
}

func Test_sourceList(t *testing.T) {
	li, err := post.SourceList(2, 10, 0)
	if err != nil {
		t.Fatal(err)
	}
	for i, p := range li {
		t.Logf("\n%d: %s", i, p.Title)
	}
}

func Test_Del(t *testing.T) {
	err := post.DelBySource(int64(1))
	if err != nil {
		t.Fatal(err)
	}
}
