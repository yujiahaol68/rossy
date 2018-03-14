package feed

import (
	"fmt"
	"sync"

	"github.com/yujiahaol68/rossy/logger"
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
	Type         string
}

// SourceController will response to the costumer like CMD etc
type SourceController interface {
	AddNewSource(url ...string) ([]Source, error)
	SaveSource() error
}

type CmdController struct{}

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

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

// AddNewSource bind with CMD: rossy add [url] [url] ...
func (c CmdController) AddNewSource(tunnel chan *logger.Message, urls ...string) (s []*Source, reason error) {
	defer close(tunnel)
	s = make([]*Source, len(urls))
	// TODO: validate url before access network

	wg.Add(len(urls))
	for i, u := range urls {
		go func(index int, url string) {
			defer wg.Done()

			source, err := getSourceByURL(url)
			if err != nil {
				if err == ErrUserNetwork {
					reason = err
				}

				tunnel <- logger.NewErrMsg(err)
				return
			}
			tunnel <- &logger.Message{Level: "info", Msg: fmt.Sprintf("Found %s feed: %s", source.Type, source.Alias)}

			source.URL = url
			s[index] = source
		}(i, u)
	}

	wg.Wait()
	return
}
