package server_test

import (
	"log"
	"net/http/httptest"
	"testing"
)

func TestHttptestServer(t *testing.T) {
	server := httptest.NewServer(nil)
	log.Println(server.URL)
}
