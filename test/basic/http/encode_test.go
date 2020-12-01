package http_test

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL_Encode(t *testing.T) {
	assert := assert.New(t)
	baseURL := "http://localhost:8080/"
	subPath := "hello:world:test"

	req, err := http.NewRequest("GET", "http://localhost/hello:world:test", nil)

	assert.Nil(err)

	fmt.Printf("====> req: %v, query escape: %v\n", req.URL, url.QueryEscape(subPath))

	u2 := path.Join(baseURL, url.PathEscape(subPath))
	req, err = http.NewRequest("GET", u2, nil)
	fmt.Printf("====> req2: %v, raw: %v, u2: %v\n", req.URL, req.URL.Path, u2)

	//
	u3 := "http://localhost/key=hello:world:test"
	fmt.Printf("====> req3: %v\n", url.QueryEscape(u3))

}

func TestURL_fast(t *testing.T) {
	assert := assert.New(t)

	baseURL := "http://localhost:8080/"

	reqUrl, _ := url.Parse(baseURL)
	query := url.Values{
		"key": []string{"hello:world"},
	}
	reqUrl.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	assert.Nil(err)

	fmt.Printf("====> req: %v, raw: %v\n", req.URL.Query(), req.URL.RawQuery)
}

func TestURL_path(t *testing.T) {
	assert := assert.New(t)

	baseURL := "http://localhost:8080/"
	reqUrl, _ := url.Parse(baseURL)
	reqUrl.Path = path.Join(reqUrl.Path, url.PathEscape("urn:linear"))

	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	assert.Nil(err)

	fmt.Printf("====> req: %v, raw: %v\n", req.URL.Path, req.URL.RawPath)
}

func TestURL_server(t *testing.T) {

}
