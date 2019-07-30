package io_test

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestIOWriter(t *testing.T) {
	w := bytes.NewBuffer(make([]byte, 0, 1024))
	log.SetOutput(w)
	log.Println("hello world")

	log.SetOutput(os.Stderr)

	log.Printf("content: [%v], writer: [%+v]", w.String(), w)
}
