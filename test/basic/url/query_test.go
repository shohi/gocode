package url_test

import (
	"log"
	"net/url"
	"testing"
)

func TestURLQuery(t *testing.T) {
	urlStr := "http://localhost:9090/hello/newyorker?season=summer&season=spring&show=tony&nokey"
	base, _ := url.Parse(urlStr)

	for key, value := range base.Query() {
		log.Printf("key ==> %s, value ==> %v", key, value)
	}

	log.Printf("raw query: %v, parsed query: %v", base.RawQuery, base.Query())
	log.Printf("Scheme: %v, Host: %v, Port: %v, Path: %v", base.Scheme, base.Host, base.Port(), base.Path)
}
