package feed

import (
	"fmt"
	"sync"

	"github.com/yujiahaol68/rossy/logger"
)

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

// AddNewSource bind with CMD: rossy add [url] [url] ...
func (c CmdController) AddNewSource(tunnel chan *logger.Message, category string, urls ...string) (s []*Source, reason error) {
	defer close(tunnel)
	s = make([]*Source, len(urls))

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
			source.Category = category
			s[index] = source
		}(i, u)
	}

	wg.Wait()
	return
}
