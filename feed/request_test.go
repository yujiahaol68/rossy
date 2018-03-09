package feed

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	stdRSSurl = "https://news.microsoft.com/feed/"
)

func Test_Headers(t *testing.T) {
	resp, err := http.Get(stdRSSurl)
	if err != nil {
		t.Fatal(err)
	}

	etag := resp.Header.Get("etag")
	lastDate := resp.Header.Get("last-modified")
	fmt.Printf("etag: %s\nlast-modified: %s\n", etag, lastDate)
	if etag == "" || lastDate == "" {
		t.Errorf("expect etag and last-modified headers but got nothing")
	}
}

func Test_conditionalGET(t *testing.T) {
	resp, err := http.Get(stdRSSurl)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	beforeDataSize := len(rawBody)

	etag := resp.Header.Get("etag")
	lastDate := resp.Header.Get("last-modified")

	client := &http.Client{}
	req, err := http.NewRequest("GET", stdRSSurl, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("If-Modified-Since", lastDate)
	req.Header.Add("If-None-Match", etag)

	rq, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer rq.Body.Close()
	newRawBody, err := ioutil.ReadAll(rq.Body)

	if rq.StatusCode != 304 {
		t.Fatalf("expect conditional GET StatusCode to be 304, got %d", rq.StatusCode)
	}

	if len(newRawBody) >= beforeDataSize {
		t.Fatalf("expect condition GET can reduce data size when request it again")
	}

	fmt.Printf("Original GET size: %d\nConditional GET size: %d\n", beforeDataSize, len(newRawBody))
}

func Test_lookup(t *testing.T) {

}
