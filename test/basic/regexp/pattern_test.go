package regexp_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatternMatch(t *testing.T) {
	cases := []struct {
		name string
		// input
		ptn     string
		content string

		expErr   bool
		expMatch bool
	}{
		{"all-numeric", `^\d+$`, "1234", false, true},
		{"all-numeric-w-leading-char", `^\d+$`, "a1234", false, false},
		{"path-asterisk", `/page/\d+$`, "news/page/1", false, true},
		{"empty-for-nonempty", "", "", false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert := assert.New(t)

			re, err := regexp.Compile(c.ptn)
			assert.Equal(c.expErr, err != nil)

			if re != nil {
				assert.Equal(c.expMatch, re.MatchString(c.content))
			}
		})
	}
}
