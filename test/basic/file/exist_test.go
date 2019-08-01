package file_test

import (
	"log"
	"os"
	"testing"
)

func TestFile_Exists(t *testing.T) {
	if _, err := os.Stat("testdata/nonexist.txt"); os.IsNotExist(err) {
		log.Printf("====> not exist, err: %v", err)
	}
}

func TestFile_Empty(t *testing.T) {

	if fi, err := os.Stat("testdata/empty.txt"); os.IsNotExist(err) {
		log.Printf("====> not exist, err: %v", err)
	} else {
		log.Printf("===> file content length: %v", fi.Size())
	}

}
