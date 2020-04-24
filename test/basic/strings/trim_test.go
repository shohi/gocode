package strings_test

import (
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
}
