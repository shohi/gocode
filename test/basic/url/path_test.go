package url_test

import (
	"log"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlPathEscape(t *testing.T) {
	log.Println(url.PathEscape("lang:>50"))
	log.Println(url.PathEscape("https://www.amazon.com"))
}

func TestPath_LeadingSlash(t *testing.T) {

	cases := []struct {
		name    string
		input   string
		expPath string
	}{
		{"normal-path", "https://github.com/pkg/errors.git", "/pkg/errors.git"},
		{"no-path", "https://github.com", ""},
		{"only-root-path", "https://github.com/", "/"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert := assert.New(t)
			u, err := url.Parse(c.input)
			assert.Nil(err)

			assert.Equal(c.expPath, u.Path)

		})
	}
}
