package regexp_test

import (
	"log"
	"regexp"
	"testing"
)

func TestReMatch(t *testing.T) {
	ptn := "reading"
	str := "Reading book is good"

	re := regexp.MustCompile("(?i).*" + ptn + ".*")
	log.Println(re.Match([]byte(str)))
}

func TestPatternReCompile(t *testing.T) {
	ptn := "reading"
	str := "reading golang book"

	re := regexp.MustCompile(".*" + ptn + ".*")
	log.Printf("matched: %v", re.MatchString(str))
}
