package fmt_test

import (
	"log"
	"testing"
)

func TestFormatVerb_Q(t *testing.T) {
	// int: 10, '\n'
	log.Printf("int: %[1]v, %[1]q", 10)

	// string: hello, "hello"
	log.Printf("string: %[1]v, %[1]q", "hello")
}
