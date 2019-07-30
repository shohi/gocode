package io_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestFile_TempFile(t *testing.T) {
	rootDir := "testdata"
	fs, err := ioutil.TempFile(rootDir, "subs")
	if err != nil {
		log.Printf("create temp file err: %v", err)
	}
	if fs != nil {
		fs.Close()

		fp := fs.Name()
		log.Printf("filepath: %v", fp)
		os.Remove(fp)
	}
}

func TestOutputBytes(t *testing.T) {
	b := []byte("hello")
	err := createFile("testdata/test.dat")
	if err != nil {
		t.Errorf("fail to write bytes to file, err: %v", err)
	}

	err = ioutil.WriteFile("testdata/test.dat", b, 0644)

	if err != nil {
		t.Errorf("fail to write bytes to file, err: %v", err)
	}

}

func createFile(path string) error {
	// detect if file exists
	_, err := os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {

		err := os.MkdirAll(filepath.Dir(path), 0777)
		if err != nil {
			return err
		}

		file, err := os.Create(path)
		if err != nil {
			return err
		}
		file.Close()
	}

	return nil
}
