package mstring

import (
	"log"
	"testing"
)

func TestStringSliceLiteral(t *testing.T) {
	m := map[string][]string{
		"hello": {"apple", "orange"},
	}
	log.Printf("%T ==> %v", m["hello"], m["hello"])
}
