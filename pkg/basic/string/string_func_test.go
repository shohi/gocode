package string_test

import (
	"log"
	"strings"
	"testing"
)

func TestStringAffix(t *testing.T) {
	log.Println(strings.HasPrefix("/aaa", "/"))
	log.Println(strings.HasSuffix("bbbb/", "/"))
}

func TestStringRepeat(t *testing.T) {
	str := "na"
	log.Println("ba" + strings.Repeat(str, 2))
}

func TestStringJoin(t *testing.T) {
	strs := []string{"Hello", "World"}

	joinedStr := strings.Join(strs, "|")

	log.Println(joinedStr)
	log.Println(strings.Split(joinedStr, "|"))
}

func TestStringContains(t *testing.T) {
	str := "【求】"
	substr := "求"

	log.Println(strings.Contains(str, substr))
}

func TestStringFormat(t *testing.T) {
	str := `"hello world"`
	log.Printf("%s\n", str)
}
