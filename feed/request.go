package feed

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/encoding/korean"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/yujiahaol68/rossy/app/entity"
	sourceService "github.com/yujiahaol68/rossy/app/service/source"

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

func GetSourceByURL(url string) (*Source, error) {
	resp, err := http.Get(url)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return nil, ErrUserNetwork
		}
		return nil, err
	}

	if feedType := sourceType(resp.Header.Get("content-type")); feedType != "" {
		return getSourceDesc(resp, url, feedType)
	}
	return nil, ErrInvalidSource
}

func sourceType(header string) string {
	return sourceReg.FindString(header)
}

func detectCharset(headerVal string) string {
	pairs := strings.Split(headerVal, ";")
	defaultCharset := "utf-8"

	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			continue
		}
		if strings.Contains(kv[0], "charset") {
			return strings.ToLower(strings.TrimSpace(kv[1]))
		}
	}

	return defaultCharset
}

func decodeReader(rsp *http.Response) io.Reader {
	switch detectCharset(rsp.Header.Get("content-type")) {
	case "gb2312":
		fallthrough
	case "gbk":
		return transform.NewReader(rsp.Body, simplifiedchinese.GBK.NewDecoder())
	case "shift_jis":
		return transform.NewReader(rsp.Body, japanese.ShiftJIS.NewDecoder())
	case "euc-jp":
		return transform.NewReader(rsp.Body, japanese.EUCJP.NewDecoder())
	case "euc-kr":
		return transform.NewReader(rsp.Body, korean.EUCKR.NewDecoder())
	default:
		return rsp.Body
	}
}

func getSourceDesc(rsp *http.Response, url, feedType string) (*Source, error) {
	r := decodeReader(rsp)
	defer rsp.Body.Close()

	bodyContent, err := ioutil.ReadAll(r)
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
		RequestCache[url] = r

		s.Alias = r.Description
		return s, nil

	case "atom":
		a := atom.New()
		err = d.Decode(a)

		if err != nil {
			return nil, CmdErrXMLParse
		}
		RequestCache[url] = a

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
			RequestCache[url] = ap

			s.Alias = ap.Title
			s.Type = "atom"
			return s, nil
		}

		err = d.Decode(&rp)
		if err != nil {
			log.Fatal(err)
			return nil, CmdErrXMLParse
		}
		RequestCache[url] = rp

		s.Alias = rp.Description
		s.Type = "rss"
		return s, nil
	}
}

func Update(ss []*entity.Source, c chan *[]*entity.Post) {
	client := http.Client{}

	for _, s := range ss {
		wg.Add(1)
		go func(source *entity.Source) {
			req, err := http.NewRequest("GET", source.URL, nil)
			if err != nil {
				log.Fatalf("newRequest: %v", err)
				wg.Done()
				return
			}

			var hasCondition bool
			if source.ETag != "" || source.LastModified != "" {
				hasCondition = true
				req.Header.Add("If-Modified-Since", source.LastModified)
				req.Header.Add("If-None-Match", source.ETag)
			}

			rq, err := client.Do(req)
			if err != nil {
				log.Fatalf("client request fail: %v", err)
				wg.Done()
				return
			}

			r := decodeReader(rq)
			defer rq.Body.Close()

			newRawBody, err := ioutil.ReadAll(r)
			if err != nil {
				log.Fatalf("Read body fail: %v", err)
				wg.Done()
				return
			}

			d := xml.NewDecoder(bytes.NewReader(newRawBody))
			d.CharsetReader = func(s string, reader io.Reader) (io.Reader, error) {
				return charset.NewReader(reader, s)
			}

			var pl []*entity.Post

			switch source.Kind {
			case "rss":
				r := rss.New()
				err = d.Decode(r)
				if err != nil {
					fmt.Printf("Decode err: %v", err)
					wg.Done()
					return
				}
				pl = r.Diff(source.Updated, hasCondition)

				// Both PubDate and LastBuildDate are optional
				latest, err := time.Parse(time.RFC822, r.PubDate)
				if err == nil {
					latest, _ = time.Parse(time.RFC3339, latest.Format(time.RFC3339))
					sourceService.UpdateDate(source.ID, latest)
				} else {
					latest, err = time.Parse(time.RFC822, r.LastBuildDate)
					if err == nil {
						latest, _ = time.Parse(time.RFC3339, latest.Format(time.RFC3339))
						sourceService.UpdateDate(source.ID, latest)
					}
					sourceService.UpdateDate(source.ID, time.Now())
				}

			case "atom":
				a := atom.New()
				err = d.Decode(a)
				if err != nil {
					fmt.Printf("Decode err: %v", err)
					wg.Done()
					return
				}
				pl = a.Diff(source.Updated, hasCondition)
				// assume the <updated> is provided
				latest, _ := time.Parse(time.RFC3339, a.Updated)
				sourceService.UpdateDate(source.ID, latest)
			}

			if len(pl) == 0 {
				wg.Done()
				return
			}
			for _, p := range pl {
				p.From = source.ID
			}

			c <- &pl
		}(s)
	}

	wg.Wait()
	close(c)
}
