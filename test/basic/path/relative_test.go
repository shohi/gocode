package path_test

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

func TestRelativePath(t *testing.T) {
	t.Skip("filepath can't handle homedir")

	p := "~/tmp.yaml"
	absp, err := filepath.Abs(p)
	if err != nil {
		t.Fatalf("failed to get absolute path for [%v]", p)
	}
	data, err := ioutil.ReadFile(absp)
	if err != nil {
		t.Fatalf("failed to read file - %v", absp)
	}

	log.Printf("file: [%v], content: [%v]", p, string(data))
}
