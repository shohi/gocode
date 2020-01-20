package url_test

import (
	"log"
	"net/url"
	"testing"

	giturls "github.com/whilp/git-urls"
)

func TestUrlParse(t *testing.T) {
	urlStr := "http://localhost:9090/hello/newyorker?season=summer"
	log.Println(url.Parse(urlStr))

	urlStr = "http://ip/?action=save"

	myURL, _ := url.Parse(urlStr)
	log.Println(url.ParseQuery(myURL.RawQuery))
}

func TestURLParseWithoutScheme(t *testing.T) {
	urlStr := "172.17.7.1/path/to/file"

	u, err := url.Parse(urlStr)
	if err != nil {
		log.Printf("url parse error: %v", err)
	} else {
		log.Printf("url host: %v", u.Host)
	}
}

func TestParse_GitURL(t *testing.T) {
	urlStr := "git@github.com:pkg/errors/hello.git"
	u, err := giturls.Parse(urlStr)

	// u, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("url: %v, scheme: %v, path: %v, user: %v, err: %v",
		u,
		u.Scheme,
		u.Path,
		u.User,
		err)
}
