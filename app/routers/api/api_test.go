package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/yujiahaol68/rossy/app/model/thirdparty"

	"github.com/yujiahaol68/rossy/app/entity"

	"github.com/yujiahaol68/rossy/app/controller"
	"github.com/yujiahaol68/rossy/socket"

	"github.com/stretchr/testify/assert"
	"github.com/yujiahaol68/rossy/app/model/checkpoint"

	"github.com/gin-gonic/gin"

	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/routers/api"
)

var router = gin.Default()

func setup() error {
	api.Router(router)
	socket.Enable = false
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

	w = httptest.NewRecorder()
	req = makeJSONReq("GET", "/api/source/unread", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/api/source/3", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_category(t *testing.T) {
	t.Log("GET: /api/categories")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/categories/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	t.Log(w.Body.String())

	t.Log("PUT: /api/categories/:id?name=xxx")
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/categories/1?name=abcd", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	t.Log(w.Body.String())

	t.Log("POST: /api/categories")
	checkNew := checkpoint.PostCategory{"DS520"}
	w = httptest.NewRecorder()
	req = makeJSONReq("POST", "/api/categories/", &checkNew)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	t.Log(w.Body.String())
}

func Test_post(t *testing.T) {
	w := httptest.NewRecorder()
	req := makeJSONReq("GET", "/api/post/?offset=1&limit=10", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	t.Log("GET: /api/post/")
	t.Log(w.Body.String())

	w = httptest.NewRecorder()
	req = makeJSONReq("GET", "/api/post/unread?offset=1&limit=4", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	t.Log("GET: /api/post/unread")
	t.Log(w.Body.String())

	w = httptest.NewRecorder()
	req = makeJSONReq("GET", "/api/post/source/2?offset=1&limit=4", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	t.Log("GET: /api/post/source/:id")
	t.Log(w.Body.String())

	w = httptest.NewRecorder()
	req = makeJSONReq("PUT", "/api/post/1", nil)
	router.ServeHTTP(w, req)
	t.Log(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	p := new(entity.Post)
	database.Conn().Id(1).Get(p)
	if p.Unread {
		t.Fatalf("Expect post.Unread is false, But got %v", p.Unread)
	}

	thirdparty.Key = os.Getenv("MERCURY_API_TOKEN")
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/post/full?url=https://www.bbc.com/news/world-asia-44772783", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	if w.Body.Len() == 0 {
		t.Fatalf("Expect has body but got nothing")
	}
}

func Test_updateFeed(t *testing.T) {
	controller.Source.UpdateAll()
}
