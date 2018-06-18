package httpclient

import (
	"crypto/tls"
	"net/http"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *http.Client
)

func New() *http.Client {
	once.Do(func() {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		instance = &http.Client{
			Timeout:   8 * time.Second,
			Transport: tr,
		}
	})
	return instance
}
