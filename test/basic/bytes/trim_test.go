package bytes_test

import (
	"bytes"
	"log"
	"testing"
)

func TestTrim_Right(t *testing.T) {
	// remove all trailing new lines
	var bs = []byte("newline\n\n\n")
	tbs := bytes.TrimRight(bs, "\n")
	log.Printf("raw: '%s', after: '%s'\n",
		string(bs), string(tbs))

	// trim only one trailing char
	tbs2 := bytes.TrimSuffix(bs, []byte{'\n'})
	log.Printf("raw: '%s', after: '%s'\n",
		string(bs), string(tbs2))
}
