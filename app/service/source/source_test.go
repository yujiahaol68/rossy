package source_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/service/source"
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
	source := source.FindOne(int64(5))
	assert.Equal(t, int64(5), source.ID)
}

func Test_GetByURL(t *testing.T) {
	s := source.FindByURL("https://tonybai.com/feed/")
	assert.Equal(t, int64(4), s.ID)

	a := source.FindByURL("njvbbfhee")
	assert.Equal(t, int64(0), a.ID)
	assert.NotEqual(t, "njvbbfhee", a.URL)
}
