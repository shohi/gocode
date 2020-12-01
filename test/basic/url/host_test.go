package url_test

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestParseHost_Schema(t *testing.T) {
	// NOTE: will panic
	// urlStr := "127.0.0.1:8080"
	urlStr := "pulsar://127.0.0.1:6650"
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Host ==> %v\n", u.Host)
}

func TestParseHost_NoSchema(t *testing.T) {
	urlStr := "localhost:6379"

	// NOTE: can't parse correctly
	// url:
	// {
	// Scheme:localhost
	// Opaque:6379
	// User: Host:
	// Path:
	// RawPath:
	// ForceQuery:false
	// RawQuery:
	// Fragment:
	// RawFragment:
	// }
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Printf("url parse error: %v", err)
	} else {
		log.Printf("url host: %v, path: %v, url: %+v\n",
			u.Host, u.Path, *u)
	}
}
