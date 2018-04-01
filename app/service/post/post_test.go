package post_test

import (
	"log"
	"os"
	"testing"

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
