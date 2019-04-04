package path_test

import (
	"log"
	"net/http"
	"path"
	"testing"
)

func TestBase_FilePath(t *testing.T) {
	log.Println(path.Base("/id/123"))
}

func TestBase_RequestPath(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:80/path/to/base", nil)
	log.Printf("request base: %v", path.Base(req.URL.Path))
}
