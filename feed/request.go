package feed

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/yujiahaol68/rossy/atom"
	"github.com/yujiahaol68/rossy/rss"
)

const (
	sourceHeaderReg = `rss|atom|xml`
)

var (
	ErrUserNetwork    = errors.New("network unavailable")
	ErrInvalidSource  = errors.New("invalid feed source URL")
	CmdErrXMLParse    = errors.New("Oops.. Something wrong when parsing XML feed")
	CmdNotSupportXML  = errors.New("Encounter unsupport XML type")
	CmdErrUserNetwork = errors.New("Contains unavailble URL or network unavailble")
)

var (
	// CheckList contains all the new feed that can be showed
	CheckList []Feed
	sourceReg = regexp.MustCompile(sourceHeaderReg)
)

func getSourceByURL(url string) (*Source, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrUserNetwork
	}

	if feedType := sourceType(resp.Header.Get("content-type")); feedType != "" {
		return getSourceDesc(resp, feedType)
	}
	return nil, ErrInvalidSource
}

func sourceType(header string) string {
	return sourceReg.FindString(header)
}

func getSourceDesc(rsp *http.Response, feedType string) (*Source, error) {
	// TODO: b maybe very big, consider using buffer
	bodyContent, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()

	if err != nil {
		return nil, CmdErrXMLParse
	}

	switch feedType {
	case "rss":
		r := rss.New()

		err = xml.Unmarshal(bodyContent, &r)
		if err != nil {
			return nil, CmdErrXMLParse
		}

		return &Source{
			"",
			rsp.Header.Get("last-modified"),
			rsp.Header.Get("etag"),
			r.Description,
			feedType,
		}, nil

	case "atom":
		a := atom.New()

		err = xml.Unmarshal(bodyContent, &a)
		if err != nil {
			return nil, CmdErrXMLParse
		}

		return &Source{
			"",
			rsp.Header.Get("last-modified"),
			rsp.Header.Get("etag"),
			a.Title,
			feedType,
		}, nil

	//case "xml":
	// TODO: header like Content-Type: text/xml; charset=utf-8 is unknown type
	default:
		return nil, CmdNotSupportXML
	}
}

// func CheckUpdate() {

// }
