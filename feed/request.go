package feed

import (
	"errors"
	"io"
	"net/http"
	"regexp"
)

const (
	sourceHeaderReg  = `rss|atom`
	ErrUserNetwork   = "network issue"
	ErrInvalidSource = "invalid feed source URL"
)

var (
	// CheckList contains all the new feed that can be showed
	CheckList []Feed
	sourceReg = regexp.MustCompile(sourceHeaderReg)
)

func IsDirectSourceURL(url string) (error, io.Reader) {
	resp, err := http.Get(url)
	if err != nil {
		return errors.New(ErrUserNetwork), nil
	}

	if header := resp.Header.Get("content-type"); sourceReg.MatchString(header) {
		return nil, resp.Body
	}
	return errors.New(ErrInvalidSource), nil
}

func CheckUpdate() {

}
