// Package atom defines XML data structures for an Atom feed.
package atom

import (
	"encoding/xml"
	"time"

	"github.com/yujiahaol68/rossy/app/entity"
)

type Atom struct {
	XMLName   xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Title     string   `xml:"title"`
	Subtitle  string   `xml:"subtitle"`
	Id        string   `xml:"id"`
	Updated   string   `xml:"updated"`
	Rights    string   `xml:"rights"`
	Link      Link     `xml:"link"`
	Author    Author   `xml:"author"`
	EntryList []Entry  `xml:"entry"`
}

type Link struct {
	Href string `xml:"href,attr"`
}

type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type Entry struct {
	Title   string `xml:"title"`
	Summary string `xml:"summary"`
	Content string `xml:"content"`
	Id      string `xml:"id"`
	Updated string `xml:"updated"`
	Link    Link   `xml:"link"`
	Author  Author `xml:"author"`
}

func New() *Atom {
	a := Atom{}
	return &a
}

func (a *Atom) Convert() []*entity.Post {
	if len(a.EntryList) == 0 {
		return nil
	}

	pl := make([]*entity.Post, len(a.EntryList))

	for i, entry := range a.EntryList {
		t, err := time.Parse(time.RFC3339, entry.Updated)
		if err != nil {
			t = time.Now()
		}

		p := new(entity.Post)
		p.Unread = true
		p.Title = entry.Title
		p.CreateAt = t
		p.Desc = string(entry.Summary)
		if c := string(entry.Content); c != "" {
			p.Content = c
		}
		p.Link = entry.Link.Href
		if entry.Author.Name == "" {
			p.Author = a.Author.Name
		} else {
			p.Author = entry.Author.Name
		}

		pl[i] = p
	}

	return pl
}

func (a *Atom) Diff(latest time.Time, underCondition bool) []*entity.Post {
	if underCondition {
		return a.Convert()
	}

	var diffIndex int
	var curItemPubDate time.Time

	for i, entry := range a.EntryList {
		curItemPubDate, _ = time.Parse(time.RFC3339, entry.Updated)
		if latest.After(curItemPubDate) {
			diffIndex = i
			break
		}
	}

	if diffIndex == 0 {
		return a.Convert()
	}
	a.EntryList = a.EntryList[:diffIndex]
	return a.Convert()
}
