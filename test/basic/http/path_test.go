package http_test

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"testing"
)

func TestPathJoin(t *testing.T) {
	u, err := url.Parse("../../..//search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	base, err := url.Parse("http://example.com/directory/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base.ResolveReference(u))
}

func TestPathClean(t *testing.T) {
	p := "/usr/data//1"
	log.Println(path.Clean(p))
}

func TestRequestPath(t *testing.T) {
	r, err := http.NewRequest("GET", "http://localhost:9090/app/path/to/file", nil)

	log.Printf("create request err: %v", err)
	log.Printf("request path: %v", r.URL.Path)
}
