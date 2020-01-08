package string_test

import (
	"log"
	"testing"
)

func TestStringFormat(t *testing.T) {
	str := `"hello world"`
	log.Printf("%s\n", str)
}
