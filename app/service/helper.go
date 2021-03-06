package service

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/go-xorm/xorm"
)

func CheckTableCountLT(minCount int64, table string, db *xorm.Engine) (bool, error) {
	sql := fmt.Sprintf("SELECT COUNT(id) AS count FROM %s", table)

	var actualCount int64
	results, err := db.Query(sql)
	actualCount, _ = binary.Varint(results[0]["count"])

	if err != nil {
		return false, err
	}

	return minCount <= actualCount, nil
}

func ToInt64(b []byte) int64 {
	i, _ := strconv.ParseInt(string(b), 10, 64)
	return i
}
