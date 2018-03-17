package feed_test

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/yujiahaol68/rossy/logger"

	"github.com/yujiahaol68/rossy/feed"
)

var (
	testUrls = []string{
		"http://www.ftchinese.com/rss/feed",
		"http://feeds.bbci.co.uk/news/rss.xml",
		"https://xiequan.info/feed/",
	}
)

func Test_addNewFeed(t *testing.T) {
	cc := new(feed.CmdController)
	tunnel := make(chan *logger.Message)

	fmt.Println("Should list the feeds that got:")

	var actualMsgCount int64
	expectMsgCount := int64(len(testUrls))

	// channel go first for collecting any log
	go func() {
		for m := range tunnel {
			m.ShowInCmd()
			// Avoid DATA RACE
			atomic.AddInt64(&actualMsgCount, 1)
		}

		if expectMsgCount != atomic.LoadInt64(&actualMsgCount) {
			t.Fatalf("every url must have feedback msg, expect %d pieces of msg, but got %d", expectMsgCount, actualMsgCount)
		}
	}()

	_, err := cc.AddNewSource(tunnel, "default", testUrls...)
	if err != nil {
		t.Fatal(err)
	}
}
