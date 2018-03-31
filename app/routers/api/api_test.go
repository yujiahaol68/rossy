package api_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yujiahaol68/rossy/app/model/checkpoint"

	"github.com/gin-gonic/gin"

	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/routers/api"
)

var router = gin.Default()

func setup() error {
	api.Router(router)
	//router.Use(cors.Default())
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

func toReader(t interface{}) *bytes.Reader {
	b, _ := json.Marshal(t)
	return bytes.NewReader(b)
}

func makeJSONReq(method, path string, data interface{}) *http.Request {
	req, _ := http.NewRequest(method, path, toReader(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req
}

func Test_Source(t *testing.T) {
	t.Log("POST: /api/source")
	check := checkpoint.PostSource{"http://www.infoq.com/cn/feed", 2}

	w := httptest.NewRecorder()
	req := makeJSONReq("POST", "/api/source/", &check)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	w = httptest.NewRecorder()
	check.URL = "http://www.ruanyifeng.com/blog/atom.xml"
	req = makeJSONReq("POST", "/api/source/", &check)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusFound, w.Code)

	w = httptest.NewRecorder()
	check.Category = 0
	req = makeJSONReq("POST", "/api/source/", &check)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_category(t *testing.T) {
	t.Log("GET: /api/categories")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/categories/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}