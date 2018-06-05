package http

import (
	"log"
	"net/http"
	"testing"
)

func TestParseRequest(t *testing.T) {
	url := "http://localhost?season=summer&season=spring&show=tony&nokey&srcurl=http://localhost:8082"

	req, _ := http.NewRequest("GET", url, nil)
	values := req.URL.Query()

	key := "season"
	log.Printf("key: %v, value: %v", key, values.Get(key))

	log.Printf("key: %v, value: %v", "srcurl", values.Get("srcurl"))

	// case sensitive!!!
	key = "SHOW"
	log.Printf("key: %v, value: %v", key, values.Get(key))
}
