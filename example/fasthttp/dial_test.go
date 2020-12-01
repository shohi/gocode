package fasthttp_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFasthttp_Dial(t *testing.T) {
	client := newTestFasthttpClient("localhost:9002")
	code, body, err := client.Get(nil, "http://localhost:9002/")
	// terr := err.(*net.OpError)
	// terr.Source = nil
	fmt.Printf("=====> code: %v, body: %v, err: %v\n", code, string(body), err)
}

func TestHttp_Dial(t *testing.T) {
	client := http.Client{}

	_, err := client.Get("http://localhost:9002/")

	fmt.Printf("====> err: %v\n", err)
}
