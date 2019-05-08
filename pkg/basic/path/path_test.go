package path_test

import (
	"log"
	"net/http"
	"path"
	"path/filepath"
	"testing"
)

func TestBase_FilePath(t *testing.T) {
	log.Println(path.Base("/id/123"))
}

func TestBase_RequestPath(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:80/path/to/base", nil)
	log.Printf("request base: %v", path.Base(req.URL.Path))
}

func TestFilePath_Dir(t *testing.T) {
	p := "/"
	log.Printf("Dir: %v, Base: %v", filepath.Dir(p), filepath.Base(p))
}
