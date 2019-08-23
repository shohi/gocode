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

func TestReFindString(t *testing.T) {
	ptn := "reading(.*)golang"
	str := "reading golang book, reading again"

	re := regexp.MustCompile(ptn)

	log.Printf("substring indexes: %v", re.FindStringSubmatchIndex(str))
	log.Printf("all substring match: %q\n", re.FindAllStringSubmatch(str, -1))
}
