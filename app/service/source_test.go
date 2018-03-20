package service_test

import (
	"log"
	"os"
	"testing"

	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/service"
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

func Test_GetSourceById(t *testing.T) {
	s := service.Source{database.Conn()}
	source := s.FindOne(int64(5))

	if source.ID != int64(5) {
		t.Fatalf("expect source ID is 5, but got %v", source.ID)
	}
}
