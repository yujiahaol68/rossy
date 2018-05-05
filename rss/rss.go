package rss

import (
	"encoding/xml"
	"html/template"
	"time"

	"github.com/yujiahaol68/rossy/app/entity"
)

// Rss Feed XML struct
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	// Required
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	// Optional
	PubDate       string `xml:"channel>pubDate"`
	LastBuildDate string `xml:"channel>lastBuildDate"`
	ItemList      []Item `xml:"channel>item"`
}

// Item is a sub struct in Rss
type Item struct {
	// Required
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// Optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Comments string        `xml:"comments"`
	Author   string        `xml:"dc:creator"`
}

// New return Rss type pointer
func New() *Rss {
	r := Rss{}
	return &r
}

func (r *Rss) Convert() []*entity.Post {
	if len(r.ItemList) == 0 {
		return nil
	}

	pl := make([]*entity.Post, len(r.ItemList))

	for i, item := range r.ItemList {
		t, err := time.Parse("2017-06-23T11:49:32Z", item.PubDate)
		if err != nil {
			t = time.Now()
		}

		p := new(entity.Post)
		p.Unread = true
		p.Title = item.Title
		p.CreateAt = t
		p.Desc = string(item.Description)
		if c := string(item.Content); c != "" {
			p.Content = c
		}
		p.Link = item.Link
		if item.Author == "" {
			p.Author = r.Link
		} else {
			p.Author = item.Author
		}
		// Still need source_id and category_id before insert into DB

		pl[i] = p
	}
	return pl
}

func (r *Rss) Diff(latest time.Time, underCondition bool) []*entity.Post {
	if underCondition {
		return r.Convert()
	}
	var diffIndex int
	var curItemPubDate time.Time

	for i, item := range r.ItemList {
		curItemPubDate, _ = time.Parse("2017-06-23T11:49:32Z", item.PubDate)

		if latest.After(curItemPubDate) {
			diffIndex = i
			break
		}
	}

	if diffIndex == 0 {
		return r.Convert()
	}
	r.ItemList = r.ItemList[:diffIndex]
	return r.Convert()
}
