package httpx

import (
	"go-zero-learn/loadbalancer"
	"net/http"
)

type HttpClient struct {
	client http.Client
}

var DefaultClient = &HttpClient{}

func Get(url string) (resp *http.Response, err error) {
	parsedUrl, err := loadbalancer.ParseUrl(url)
	if err != nil {
		panic(err)
	}
	return DefaultClient.client.Get(parsedUrl.String())
}

func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {
	parsedUrl, err := loadbalancer.ParseUrl(req.URL.String())
	if err != nil {
		panic(err)
	}
	req.URL = parsedUrl
	return DefaultClient.client.Do(req)
}
