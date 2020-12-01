package fasthttp_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
)

func newRouter(latency time.Duration) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("===> receive request, time: %v\n", time.Now())
		if latency > 0 {
			time.Sleep(latency)
		}
		fmt.Printf("===> write response start, time: %v\n", time.Now())

		w.Write([]byte("OK"))
	})

	return router
}

func newTestServer(port int) *http.Server {
	addr := fmt.Sprintf(":%d", port)
	router := newRouter(0)
	server := &http.Server{
		Addr:        addr,
		Handler:     router,
		IdleTimeout: 2 * time.Second,
	}

	return server
}

func newTestFasthttpClient(addr string) *fasthttp.HostClient {

	client := &fasthttp.HostClient{
		Addr:                      addr,
		Name:                      "fasthttp-client",
		MaxIdleConnDuration:       10 * time.Second,
		MaxConns:                  1,
		MaxConnDuration:           20 * time.Second,
		MaxIdemponentCallAttempts: 1,
		// RetryIf:                   func(req *fasthttp.Request) bool { return false },
	}

	return client
}

func newTestHttpClient(addr string) *http.Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return client
}

func getHttpResponse(client *http.Client, reqURL string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", reqURL, nil)
	resp, err := client.Do(req)
	return resp, err
}

func getFasthttpResponse(client *fasthttp.HostClient, reqURL string) (*fasthttp.Response, error) {

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.URI().Update(reqURL)
	resp := fasthttp.AcquireResponse()

	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)

	return resp, err
}

func fasthttpRespToString(resp *fasthttp.Response) string {
	return fmt.Sprintf("header: [%v], content: [%v]",
		resp.Header.String(), string(resp.Body()))
}

func httpRespToString(resp *http.Response) string {
	var content []byte
	if resp.Body != nil {
		defer func() {
			resp.Body.Close()
		}()
		content, _ = ioutil.ReadAll(resp.Body)
	}
	return fmt.Sprintf("header: [%v], content: [%v]",
		resp.Header, string(content))
}

func TestFasthttpClient(t *testing.T) {
	port := 8086
	server := newTestServer(port)
	errCh := make(chan error, 2)

	go func() {
		errCh <- server.ListenAndServe()
	}()

	time.Sleep(50 * time.Millisecond)

	// first request
	client := newTestFasthttpClient(fmt.Sprintf("localhost:%d", port))
	reqURL := fmt.Sprintf("http://localhost:%d/", port)
	resp, err := getFasthttpResponse(client, reqURL)
	fmt.Printf("first request - err: %v, resp: %v\n", err, fasthttpRespToString(resp))

	// wait 8s
	time.Sleep(8 * time.Second)

	// request again
	resp, err = getFasthttpResponse(client, reqURL)
	fmt.Printf("second request (8s later) - err: %v, resp: %v\n",
		err, fasthttpRespToString(resp))

	time.Sleep(30 * time.Second)
	server.Close()
}

func TestHttpClient(t *testing.T) {
	port := 8086
	server := newTestServer(port)
	errCh := make(chan error, 2)

	go func() {
		errCh <- server.ListenAndServe()
	}()

	time.Sleep(50 * time.Millisecond)

	// first request
	client := newTestHttpClient(fmt.Sprintf("localhost:%d", port))
	reqURL := fmt.Sprintf("http://localhost:%d/", port)
	resp, err := getHttpResponse(client, reqURL)
	fmt.Printf("first request - err: %v, resp: %v\n", err, httpRespToString(resp))

	// wait 8s
	time.Sleep(8 * time.Second)

	// request again
	resp, err = getHttpResponse(client, reqURL)
	fmt.Printf("second request (8s later) - err: %v, resp: %v\n",
		err, httpRespToString(resp))

	time.Sleep(30 * time.Second)
	server.Close()
}
