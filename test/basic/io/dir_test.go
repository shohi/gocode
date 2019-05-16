package io_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDir_TempDir(t *testing.T) {
	tmpDir, err := ioutil.TempDir(".", "data_server_")

	log.Printf("tmp dir: %v, err: %v", tmpDir, err)

	if err == nil {
		defer os.RemoveAll(tmpDir)
	}
}

func TestDir_ReadDir(t *testing.T) {
	dirpath := "../.."
	fileList, _ := ioutil.ReadDir(dirpath)
	abspath, _ := filepath.Abs(dirpath)
	log.Println(abspath)
	for _, f := range fileList {
		log.Println(filepath.Join(dirpath, f.Name()))
	}
}

func TestDir_ReadFileInDir(t *testing.T) {
	dirpath := "."
	fileList, _ := ioutil.ReadDir(dirpath)
	for _, f := range fileList {
		log.Println(filepath.Join(dirpath, f.Name()))
	}
}

func TestDir_RemoveDir(t *testing.T) {
	dirpath := "non-exist/non-exist"
	err := os.RemoveAll(dirpath)
	log.Println(err)
}
