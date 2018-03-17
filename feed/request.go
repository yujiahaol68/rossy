package feed

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"

	"github.com/yujiahaol68/rossy/atom"
	"github.com/yujiahaol68/rossy/rss"
	"golang.org/x/net/html/charset"
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
	sourceReg = regexp.MustCompile(sourceHeaderReg)
)

func getSourceByURL(url string) (*Source, error) {
	resp, err := http.Get(url)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return nil, ErrUserNetwork
		}
		return nil, err
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

	d := xml.NewDecoder(bytes.NewReader(bodyContent))
	d.CharsetReader = func(s string, reader io.Reader) (io.Reader, error) {
		return charset.NewReader(reader, s)
	}

	s := new(Source)
	s.LastModified = rsp.Header.Get("last-modified")
	s.ETag = rsp.Header.Get("etag")
	s.Type = feedType

	switch feedType {
	case "rss":
		r := rss.New()
		err = d.Decode(r)

		if err != nil {
			return nil, CmdErrXMLParse
		}

		s.Alias = r.Description
		return s, nil

	case "atom":
		a := atom.New()
		err = d.Decode(a)

		if err != nil {
			return nil, CmdErrXMLParse
		}

		s.Alias = a.Title
		return s, nil

	// case "xml":
	default:
		rp := rss.New()
		ap := atom.New()

		if bytes.Contains(bodyContent, []byte("<feed")) {
			err = d.Decode(&ap)
			if err != nil {
				log.Fatal(err)
				return nil, CmdErrXMLParse
			}

			s.Alias = ap.Title
			s.Type = "atom"
			return s, nil
		}

		err = d.Decode(&rp)
		if err != nil {
			log.Fatal(err)
			return nil, CmdErrXMLParse
		}

		s.Alias = rp.Description
		s.Type = "rss"
		return s, nil
	}
}

// func CheckUpdate() {

// }
