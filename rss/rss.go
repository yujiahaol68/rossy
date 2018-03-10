package rss

import (
	"encoding/xml"
	"html/template"
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
	PubDate  string `xml:"channel>pubDate"`
	ItemList []Item `xml:"channel>item"`
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
}

// New return Rss type pointer
func New() *Rss {
	r := Rss{}
	return &r
}
