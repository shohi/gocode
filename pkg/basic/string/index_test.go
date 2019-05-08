package string_test

import (
	"log"
	"testing"
)

func TestIndexAlloc(t *testing.T) {
	s := "hello"
	// Compile error: string is immutable
	// s[0] = 'w'
	b := []rune(s)
	b[0] = 'w'
	log.Printf("%v", string(b))
}
