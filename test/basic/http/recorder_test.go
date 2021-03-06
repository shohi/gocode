package http_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTP_ResponseRecorder(t *testing.T) {
	// data := []byte{}
	var data []byte
	log.Println(data)

	data = []byte{}
	log.Println(data)

	handler := func(w http.ResponseWriter, r *http.Request) {
		n, err := w.Write(data)
		log.Println(n, err)
		// io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	log.Println("response body", resp.Body)
	resp.Body = nil
	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.StatusCode)
	log.Println(resp.Header.Get("Content-Type"))
	log.Println(string(body))

	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}
