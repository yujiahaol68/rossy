package post_test

import (
	"encoding/json"
	"fmt"
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

const sql1 = `SELECT source_id, category_id, alias, COUNT(unread)
FROM (post LEFT JOIN source ON source_id = source.id)
LEFT JOIN category ON category.id = category_id
WHERE unread = 0
GROUP BY source_id
ORDER BY COUNT(unread)`

const sql2 = ``

func Test_CountUnreadSQL(t *testing.T) {
	db := database.Conn()
	r, err := db.Query(sql2)
	if err != nil {
		t.Fatal(err)
	}
	printResult(&r)
	i := string(r[5]["COUNT(unread)"])
	fmt.Println(i)
	//fmt.Println(len(r))
}

func printResult(r *[]map[string][]byte) {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
