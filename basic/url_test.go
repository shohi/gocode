package basic

import (
	"log"
	"net/url"
	"testing"
)

func TestUrlPathEscape(t *testing.T) {
	log.Println(url.PathEscape("lang:>50"))
}

func TestUrlParse(t *testing.T) {
	urlStr := "http://localhost:9090/hello/newyorker?season=summer"
	log.Println(url.Parse(urlStr))
}

func TestUrlResolveReference(t *testing.T) {
	urlStr := "http://localhost:9090/hello/newyorker?season=summer"
	base, _ := url.Parse(urlStr)

	log.Println(base)

	urlStr1 := "a/b/c/d.ts"
	url1, err := url.Parse(urlStr1)
	log.Println(err)

	log.Println(base.ResolveReference(url1))
}

func TestURLQuery(t *testing.T) {
	urlStr := "http://localhost:9090/hello/newyorker?season=summer&season=spring&show=tony"
	base, _ := url.Parse(urlStr)

	for key, value := range base.Query() {
		log.Printf("key ==> %s, value ==> %v", key, value)
	}
	log.Println(base.Query())
}
