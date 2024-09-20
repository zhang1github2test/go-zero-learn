package loadbalancer

import (
	"errors"
	"fmt"
	"net/url"
)

type Resolver interface {
	Next(serviceName string) string
	Scheme() string
	Running() bool
	Start() bool
}

const (
	httpScheme  = "http"
	httpsScheme = "https"
)

var (
	resolvers = make(map[string]Resolver)
)

func Register(r Resolver) {
	resolvers[r.Scheme()] = r
}

func ParseUrl(rawUrl string) (*url.URL, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	if parsedURL.Scheme == httpScheme || parsedURL.Scheme == httpsScheme {
		return nil, nil
	}
	re := resolvers[parsedURL.Scheme]
	if re == nil {
		return nil, errors.New(fmt.Sprintf("没有找到对应url的解析器%v", parsedURL.Scheme))
	}
	if !re.Running() {
		re.Start()
	}

	if !re.Running() {
		return nil, errors.New(fmt.Sprintf("初始化解析器失败%v", parsedURL.Scheme))
	}

	parsedHost := re.Next(parsedURL.Host)
	if parsedHost == "" {
		return nil, errors.New("没有找到服务")
	}
	parsedURL.Host = parsedHost
	parsedURL.Scheme = httpScheme

	return parsedURL, nil
}
