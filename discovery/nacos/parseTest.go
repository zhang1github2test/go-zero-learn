package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "nacos://nacos.rpc:8080/path?query=123#fragment"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
	fmt.Println("url:", parsedURL.String())
}
