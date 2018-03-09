package feed

import (
	"fmt"
)

// post implement Feed interface and hide the detail to customer
type post struct {
	title    string
	url      string
	desc     string
	content  string
	category string
}

type category struct {
	Name      string
	Subscribe []string
}

// Feed is the entry point of every post consume by cmd
type Feed interface {
	GetName() string
	GetSource() string
	GetContent() string
	GetDesc() string
	From() string
	Display()
}

type Source struct {
	URL          string
	ETag         string
	LastModified string
	Alias        string
}

func (p post) GetName() string {
	return p.title
}

func (p post) GetSource() string {
	return p.url
}

func (p post) GetContent() string {
	return p.content
}

func (p post) GetDesc() string {
	return p.desc
}

func (p post) From() string {
	return p.category
}

func (p post) String() string {
	return fmt.Sprintf("%s\nsource:%s", p.title, p.url)
}

func (p post) Display() {
	fmt.Println(p)
}
