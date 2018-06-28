package http

import (
	"fmt"
	"log"
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
