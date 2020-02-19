package strings_test

import (
	"log"
	"strings"
	"testing"
)

func TestStrings_Split(t *testing.T) {
	str := "a.b.c.d.e.f"

	tokens := strings.Split(str, ".")
	log.Println(tokens)

	log.Println(strings.Join(tokens[len(tokens):], ","))
}
