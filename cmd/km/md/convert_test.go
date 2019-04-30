package md

import (
	"log"
	"testing"
)

// TODO: refactor
func TestConvert(t *testing.T) {
	fp := "testdata/sample.xlsx"
	md := XlsxToMd(fp)

	log.Printf("md: %v", md)
}
