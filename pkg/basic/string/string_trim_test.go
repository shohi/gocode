package string_test

import (
	"log"
	"strings"
	"testing"
)

func TestStringTrim(t *testing.T) {
	str := "     hello   "
	log.Println(strings.TrimSpace(str))

	str = "a/b/c/d/e///"
	log.Println(strings.TrimRight(str, "/"))
	log.Println(strings.TrimRight("/a/b/c/d/", "//")) // ==> /a/b/c/d
}

func TestStringTrimLeft(t *testing.T) {
	str := "/abH/a/b"

	// ==> "H/a/b"
	log.Println(strings.TrimLeft(str, "/ba"))
}

func TestStringTrimSuffix(t *testing.T) {
	aa := "aaa/bbb"
	bb := strings.TrimPrefix(aa, "aaa/")

	log.Println(aa, bb)
}
