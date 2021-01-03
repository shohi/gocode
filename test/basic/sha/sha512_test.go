package sha_test

import (
	"archive/tar"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/atrox/homedir"
	"github.com/mholt/archiver/v3"
)

func TestSHA512(t *testing.T) {
	filename := "testdata/apache-maven-3.6.2-bin.tar.gz"
	fp, _ := homedir.Expand(filename)
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(err)
	}
	h := sha512.New()
	h.Write(data)

	fmt.Printf("sum: %x\n", h.Sum(nil))
	fmt.Printf("sum: %s\n", hex.EncodeToString(h.Sum(nil)))
}

// TODO: move
func TestArchiveList(t *testing.T) {
	filename := "testdata/apache-maven-3.6.2-bin.tar.gz"
	archiver.Walk(filename, func(f archiver.File) error {
		if h, ok := f.Header.(*tar.Header); ok {
			filepath.Dir(h.Name)
			fmt.Println(h.Name)
			// return errors.New("Found")
		}

		return nil
	})
}
