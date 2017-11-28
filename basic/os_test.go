package basic

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHostname(t *testing.T) {
	log.Println(os.Hostname())
}

func TestExit(t *testing.T) {
	log.Println("going to exit")
	os.Exit(-21)
}

func TestDefer(t *testing.T) {
	defer func() {
		log.Println("exit")
	}()
	os.Exit(-1)
}

func TestGetenv(t *testing.T) {
	log.Println(os.Getenv("GOPATH"))
}

func TestOSRemove(t *testing.T) {
	fp := "test/test/test.txt"
	_, err := os.Stat(fp)
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(fp), 0777) // for mock, use 0777 for now
	}

	err = ioutil.WriteFile(fp, []byte("Test test"), 0777)
	rootDir := strings.Split(fp, string(filepath.Separator))[0]
	err = os.RemoveAll(rootDir)
	if err != nil {
		log.Println(err)
	}

}
