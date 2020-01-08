package strings_test

import (
	"log"
	"strings"
	"testing"
)

func TestStringAffix(t *testing.T) {
	log.Println(strings.HasPrefix("/aaa", "/"))
	log.Println(strings.HasSuffix("bbbb/", "/"))
}

func TestStringContains(t *testing.T) {
	str := "【求】"
	substr := "求"

	log.Println(strings.Contains(str, substr))
}
