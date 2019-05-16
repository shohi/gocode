package http_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestMuxRoute(t *testing.T) {
	r := mux.NewRouter()
	defaultHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("default"))
	}
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	}

	// NOTE: default handler must be the last one
	r.HandleFunc("/hello", helloHandler)
	// r.NotFoundHandler = http.HandlerFunc(defaultHandler)
	r.PathPrefix("/").Handler(http.HandlerFunc(defaultHandler))

	// Test Default Handler
	req, _ := http.NewRequest("GET", "http://localhost/world", nil)
	var match mux.RouteMatch
	ok := r.Match(req, &match)

	log.Printf("match ===> %v, matches: %T", ok, match.Handler)

	w := httptest.NewRecorder()
	match.Handler.ServeHTTP(w, req)

	log.Printf("content ==> %v", w.Body.String())

	// Test Overlay
	match = mux.RouteMatch{}
	req, _ = http.NewRequest("GET", "http://localhost/hello", nil)
	r.Match(req, &match)
	w = httptest.NewRecorder()
	match.Handler.ServeHTTP(w, req)

	log.Printf("content ==> %v", w.Body.String())

}
