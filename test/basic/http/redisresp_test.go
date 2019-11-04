package http_test

import (
	"bytes"
	"flag"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rsAddr = flag.String("rsaddr", "", `redis compatible server address, e.g. "http://localhost:9301/api/broadcast`)

func TestHTTPRequest_SendRESP(t *testing.T) {
	if len(*rsAddr) == 0 {
		t.Skip("Should start server first")
	}

	assert := assert.New(t)
	reader := bytes.NewReader([]byte("*3\r\n$7\r\nPUBLISH\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"))
	req, err := http.NewRequest("POST", *rsAddr, reader)

	assert.Nil(err)

	var client http.Client

	resp, err := client.Do(req)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.Equal(204, resp.StatusCode)
}
