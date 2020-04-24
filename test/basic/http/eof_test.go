package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHTTPClient_EOF(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "", nil)

	assert.Nil(err)
	req.Close = true
}

func TestXXX(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// rwc, _, _ := w.(http.Hijacker).Hijack()
				// defer rwc.Close()
				time.Sleep(200 * time.Millisecond)
				w.Write([]byte("OK"))
			},
		),
	)
	defer server.Close()

	client := &http.Client{
		Timeout: 20 * time.Millisecond,
	}

	var wg sync.WaitGroup
	for k := 0; k < 100; k++ {
		go func(index int) {
			wg.Add(1)
			req, _ := http.NewRequest("GET", server.URL, nil)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("index: %v, GET returned error: %v\n",
					index, err)
			}
			if resp != nil && resp.Body != nil {
				resp.Body.Close()
			}
			wg.Done()
		}(k)
	}

	wg.Wait()
}
