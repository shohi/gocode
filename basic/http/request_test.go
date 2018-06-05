package http

import (
	"log"
	"net/http"
	"testing"
)

func TestParseRequest(t *testing.T) {
	url := "http://localhost:9090/hello/newyorker?season=summer&season=spring&show=tony&nokey"

	req, _ := http.NewRequest("GET", url, nil)
	values := req.URL.Query()

	key := "season"
	log.Printf("key: %v, value: %v", key, values.Get(key))

}
