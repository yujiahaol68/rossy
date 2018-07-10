package thirdparty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/yujiahaol68/rossy/feed/httpclient"

	"github.com/yujiahaol68/rossy/app/entity"
)

var (
	ParserURL  = "https://mercury.postlight.com/parser?url=%s"
	Key        string
	once       sync.Once
	crawler    entity.Crawler
	charsetReg = "charset=.{1,}?\""
)

var _ entity.Crawler = new(mercury)

type mercury struct {
	Title         string      `json:"title"`
	Content       string      `json:"content"`
	DatePublished time.Time   `json:"date_published"`
	LeadImageURL  string      `json:"lead_image_url"`
	Dek           string      `json:"dek"`
	URL           string      `json:"url"`
	Domain        string      `json:"domain"`
	Excerpt       string      `json:"excerpt"`
	WordCount     int         `json:"word_count"`
	Direction     string      `json:"direction"`
	TotalPages    int         `json:"total_pages"`
	RenderedPages int         `json:"rendered_pages"`
	NextPageURL   interface{} `json:"next_page_url"`
}

func NewParser() entity.Crawler {
	once.Do(func() {
		crawler = new(mercury)
	})
	return crawler
}

func (m *mercury) ParseURL(u string) error {
	client := httpclient.New()

	req, err := http.NewRequest("GET", fmt.Sprintf(ParserURL, u), nil)
	if err != nil {
		return err
	}

	req.Header.Add("x-api-key", Key)
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, m)
}

func (m *mercury) Bytes() ([]byte, error) {
	re, _ := regexp.Compile(charsetReg)
	s := re.FindString(m.Content)
	kv := strings.Split(s, "=")
	if len(kv) == 2 {
		s = strings.Trim(kv[1], "\"")
		srd := strings.NewReader(m.Content)
		var reader io.Reader
		switch s {
		case "gb2312":
			fallthrough
		case "gbk":
			reader = transform.NewReader(srd, simplifiedchinese.GBK.NewDecoder())
		case "shift_jis":
			reader = transform.NewReader(srd, japanese.ShiftJIS.NewDecoder())
		case "euc-jp":
			reader = transform.NewReader(srd, japanese.EUCJP.NewDecoder())
		case "euc-kr":
			reader = transform.NewReader(srd, korean.EUCKR.NewDecoder())
		default:
			goto unknownCharset
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		m.Content = buf.String()
	}
unknownCharset:
	return json.Marshal(m)
}
