package regexp_test

import (
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexp_Raw(t *testing.T) {
	assert := assert.New(t)

	raw := `hello\+world`

	re, err := regexp.Compile(".*" + raw + ".*")
	log.Printf("compile err: %v, expression: %v", err, re.String())

	assert.Nil(err)

	// NOTE: if raw string is `hello+world`, the following `Match` will return false
	// as `hello+world` is a normal regular expression, `+` has special meanings.
	log.Printf("matched: %v", re.Match([]byte(`hello+world`)))
}
