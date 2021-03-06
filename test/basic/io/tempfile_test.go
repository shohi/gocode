package io_test

import (
	"io/ioutil"
	"log"
	"os"
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

func TestFile_TempFile_Empty(t *testing.T) {
	fs, err := ioutil.TempFile("", "subs")
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
