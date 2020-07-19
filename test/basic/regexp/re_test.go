package regexp_test

import (
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexp_Match(t *testing.T) {
	assert := assert.New(t)

	ptn := "reading"
	str := "Reading book is good"

	re := regexp.MustCompile("(?i).*" + ptn + ".*")
	assert.True(re.Match([]byte(str)))
}

func TestRegexp_PatternCompile(t *testing.T) {
	ptn := "reading"
	str := "reading golang book"

	re := regexp.MustCompile(".*" + ptn + ".*")
	log.Printf("matched: %v", re.MatchString(str))
}

func TestRegexp_FindString(t *testing.T) {
	ptn := "reading(.*)golang"
	str := "reading golang book, reading again"

	re := regexp.MustCompile(ptn)

	log.Printf("substring indexes: %v", re.FindStringSubmatchIndex(str))
	log.Printf("all substring match: %q\n", re.FindAllStringSubmatch(str, -1))
}

func TestRegexp_EmptyPattern(t *testing.T) {
	ptn := ""
	re := regexp.MustCompile(ptn)

	log.Printf("empty pattern matches [%v]: %v", "abc", re.MatchString("abc"))
}
