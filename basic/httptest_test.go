package basic

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPTestServer(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	log.Println(server)

	time.Sleep(20 * time.Second)
}
