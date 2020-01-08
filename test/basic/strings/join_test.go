package strings_test

import (
	"log"
	"strings"
	"testing"
)

func TestStringJoin(t *testing.T) {
	strs := []string{"Hello", "World"}

	joinedStr := strings.Join(strs, "|")

	log.Println(joinedStr)
	log.Println(strings.Split(joinedStr, "|"))
}

func TestStringJoin_EmptySlice(t *testing.T) {
	var strs []string

	log.Printf("empty string array join: [%v]", strings.Join(strs, ","))

}
