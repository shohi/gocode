package mstring

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

func TestStringSplit(t *testing.T) {
	str := ""
	strSlice := strings.Split(str, ",")
	log.Println(strSlice)
}

func TestStringSplitWithRegularExpression(t *testing.T) {
	// ptn := "[,，\\s+]"
	ptn := "\\s+|[,，]"
	strSlice := regexp.MustCompile(ptn).Split("a   b   c d  e   f,g,h，HHH", -1)

	for k, v := range strSlice {
		log.Printf("%d. %s\n", k, v)
	}
}

func TestStringSplitWithSlash(t *testing.T) {
	str := "a/b/c/d/"
	strs := strings.Split(str, "/")

	for k, v := range strs {
		log.Printf("%d => %s", k, v)
	}
}

func TestStringSplitWithDot(t *testing.T) {
	str := "a.b...d."
	strs := strings.Split(str, ".")

	for k, v := range strs {
		log.Printf("%d => %s", k, v)
	}
}
