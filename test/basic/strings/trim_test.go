package strings_test

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestTrimLeft(t *testing.T) {
	prefix := "_test."
	raw := "_raft"

	log.Printf("raw: %v, left: %v\n",
		raw, strings.TrimLeft(raw, prefix))
}

func TestTrimSuffix(t *testing.T) {
	prefix := "_test."
	raw := "_raft"

	log.Printf("raw: %v, left: %v\n",
		raw, strings.TrimSuffix(raw, prefix))

	// remove all trailing new lines
	str := "newline\n\n\n"
	tStr := strings.TrimRight(str, "\n")
	log.Printf("raw: '%s', after: '%s'\n", str, tStr)

	bytes.TrimRight([]byte(str), "\n")
}
