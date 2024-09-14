package loadbalancer

import (
	"errors"
	"fmt"
	"net/url"
)

type Builder interface {
	Build(url string) (Resolver, error)
	Scheme() string
}
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

func ParseUrl(rawUrl string) (string, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	if parsedURL.Scheme == httpScheme || parsedURL.Scheme == httpsScheme {
		return "", nil
	}
	re := resolvers[parsedURL.Scheme]
	if re == nil {
		return "", errors.New(fmt.Sprintf("没有找到对应url的解析器%v", parsedURL.Scheme))
	}
	if !re.Running() {
		re.Start()
	}
	parsedHost := re.Next(parsedURL.Host)
	if parsedHost == "" {
		return "", errors.New("没有找到服务")
	}
	parsedURL.Host = parsedHost

	return parsedURL.String(), nil
}
