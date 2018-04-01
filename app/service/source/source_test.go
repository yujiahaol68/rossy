package source_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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

const sql1 = `SELECT source_id, category_id, alias, name, COUNT(unread) AS count
FROM (post LEFT JOIN source ON source_id = source.id)
LEFT JOIN category ON category.id = category_id
WHERE unread = 0
GROUP BY source_id
ORDER BY count;`

const sql2 = `SELECT COUNT(id) AS count FROM post WHERE unread = 0;`

func Test_CountUnreadSQL(t *testing.T) {
	db := database.Conn()
	r, err := db.Query(sql1)
	if err != nil {
		t.Fatal(err)
	}
	r2, _ := db.Query(sql2)

	var unreadCount int64
	for _, v := range r {
		c, _ := strconv.ParseInt(string(v["count"]), 10, 64)
		unreadCount += c
	}

	actualCount, _ := strconv.ParseInt(string(r2[0]["count"]), 10, 64)
	assert.Equal(t, actualCount, unreadCount)
}

func printResult(r *[]map[string][]byte) {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
