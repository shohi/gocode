package basic

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrlPathEscape(t *testing.T) {
	fmt.Println(url.PathEscape("lang:>50"))
}

func TestUrlParse(t *testing.T) {
	urlStr := "http://localhost:9090/hello/newyorker?season=summer"
	fmt.Println(url.Parse(urlStr))
}

func TestUrlResolveReference(t *testing.T) {
	urlStr := "http://localhost:9090/hello/newyorker?season=summer"
	base, _ := url.Parse(urlStr)

	fmt.Println(base)

	urlStr1 := "a/b/c/d.ts"
	url1, err := url.Parse(urlStr1)
	fmt.Println(err)

	fmt.Println(base.ResolveReference(url1))

}
