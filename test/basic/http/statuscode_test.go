package http_test

import (
	"log"
	"net/http"
	"testing"
)

func TestStatusCode_StatusText(t *testing.T) {
	log.Printf("status code: %v", http.StatusText(http.StatusTooManyRequests))
}
