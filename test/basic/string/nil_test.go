package string_test

import (
	"log"
	"testing"
)

func TestNilString(t *testing.T) {
	var str string
	log.Printf("%v", str)
}
