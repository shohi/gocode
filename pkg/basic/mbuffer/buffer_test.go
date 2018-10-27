package mbuffer_test

import (
	"bytes"
	"log"
	"testing"
)

func TestBufferTruncate(t *testing.T) {

	var buf bytes.Buffer
	log.Printf("buf: cap - [%d], len - [%d]", buf.Cap(), buf.Len())
	buf.Write([]byte("h"))
	log.Printf("buf: cap - [%d], len - [%d]", buf.Cap(), buf.Len())

	// NOTE: will panic because of truncation out of range
	// n := 100 // buf.Len() = 1
	// buf.Truncate(100)
}
