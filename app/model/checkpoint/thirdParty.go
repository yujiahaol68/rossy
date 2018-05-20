package checkpoint

import (
	"time"

	"github.com/yujiahaol68/rossy/app/entity"
)

var (
	ParserURL = "https://mercury.postlight.com/parser?url=%s"
	Key       string
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
	return new(mercury)
}

func (m *mercury) ParseURL(u string) {

}

func (m *mercury) FullEssay() []byte {
	return []byte{}
}
