package file_test

import (
	"log"
	"path/filepath"
	"testing"
)

func TestFilepathJoin(t *testing.T) {
	aa := "key/"

	log.Println(filepath.Join("/", aa))
}
