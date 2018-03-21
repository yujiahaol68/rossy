package feed

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	stdRSSurl          = "https://news.microsoft.com/feed/"
	atomURL            = "http://www.ruanyifeng.com/blog/atom.xml"
	notFeedurl         = "https://bing.com"
	notUTF8encodingURL = "http://news.qq.com/newsgn/rss_newsgn.xml"
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

func Test_sourceByURL(t *testing.T) {
	s, err := GetSourceByURL(stdRSSurl)
	if err != nil {
		t.Fatal(err)
	}

	if s.Type != "rss" {
		t.Fatalf("expect RSS type feed, but got %v", s.Type)
	}

	fmt.Printf("Alias: %s\nType: %s\nEtag: %s\nLast-modified:%s\n", s.Alias, s.Type, s.ETag, s.LastModified)

	s, err = GetSourceByURL(atomURL)

	if err != nil {
		t.Fatal(err)
	}

	if s.Type != "atom" {
		t.Fatalf("expect atom feed type, but got %s", s.Type)
	}

	_, err = GetSourceByURL(notFeedurl)
	if err != ErrInvalidSource {
		t.Fatal(err)
	}
}

func Test_notUTF8(t *testing.T) {
	s, err := GetSourceByURL(atomURL)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(s.Alias)
}

func Test_lookup(t *testing.T) {

}
