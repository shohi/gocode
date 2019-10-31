package http_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPGet(t *testing.T) {

	//
	c := &http.Client{}

	// success one
	resp, err := c.Get("https://www.douban.com")
	if err == nil {
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
	}
	// log.Println(resp, err)

	// failure one
	resp, err = c.Get("https://localhost:12345")
	log.Println(resp, err)

	log.Println(resp == nil)
}

func TestCreateHTTPRequest(t *testing.T) {
	req, err := http.NewRequest("PUT", "localhost:8080", nil)
	if err != nil {
		log.Println(err)
	}
	log.Println(req.Body == nil)

	// ErrorCase
	// leading space error, ref: https://github.com/golang/go/issues/24246
	req, err = http.NewRequest("GEt", " http:/localhost:8080", nil)
	log.Printf("err: %v", err)
}

func TestHTTPRequest_SendRESP(t *testing.T) {
	t.Skip("Should start server first")

	assert := assert.New(t)
	reader := bytes.NewReader([]byte("*3\r\n$7\r\nPUBLISH\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"))
	req, err := http.NewRequest("POST", "http://localhost:9301/api/broadcast", reader)

	assert.Nil(err)

	var client http.Client

	resp, err := client.Do(req)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.Equal(204, resp.StatusCode)
}
