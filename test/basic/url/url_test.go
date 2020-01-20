package url_test

import (
	"log"
	"net/url"
	"testing"
)

func TestUrlResolveReference(t *testing.T) {
	urlStr := "http://localhost:9090/hello/newyorker?season=summer"
	base, _ := url.Parse(urlStr)

	log.Println(base)

	urlStr1 := "a/b/c/d.ts"
	url1, err := url.Parse(urlStr1)
	log.Println(err)

	log.Println(base.ResolveReference(url1))
}

func TestURLPathPrefix(t *testing.T) {
	urlStr := "http://ip/stream1/segment1"
	base, _ := url.Parse(urlStr)

	log.Println(base.Path)
}
