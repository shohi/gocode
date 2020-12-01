package fasthttp_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

func newTestServerWithWriteTimeout(port int) *http.Server {
	addr := fmt.Sprintf(":%d", port)
	router := newRouter(10 * time.Second)
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		IdleTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	return server
}

func newTestFasthttpClientWithTimeout(addr string) *fasthttp.HostClient {

	client := &fasthttp.HostClient{
		Addr:                      addr,
		Name:                      "fasthttp-client",
		MaxIdleConnDuration:       20 * time.Second,
		MaxConns:                  10,
		MaxConnDuration:           20 * time.Second,
		MaxIdemponentCallAttempts: 3,
		RetryIf:                   func(r *fasthttp.Request) bool { return false },
	}

	return client
}

func TestFasthttpClient_WriteTimeout(t *testing.T) {
	port := 8086
	server := newTestServerWithWriteTimeout(port)
	errCh := make(chan error, 2)

	go func() {
		errCh <- server.ListenAndServe()
	}()

	time.Sleep(50 * time.Millisecond)

	// first request
	client := newTestFasthttpClientWithTimeout(fmt.Sprintf("localhost:%d", port))
	reqURL := fmt.Sprintf("http://localhost:%d/", port)
	resp, err := getFasthttpResponse(client, reqURL)
	fmt.Printf("first request - err: %v, resp: %v\n", err, fasthttpRespToString(resp))
}
