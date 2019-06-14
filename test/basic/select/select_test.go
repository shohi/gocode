package select_test

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {

	fn := func() string {
		time.Sleep(10 * time.Second)
		return "Hi"
	}

	_ = fn

	r, _ := http.NewRequest("GET", "www.google.com", nil)
	r.WithContext(context.Background())

	// sigCh := make(chan struct{})

	select {
	case <-time.After(2 * time.Second):
		log.Printf("receive signal")

		// NOTE: not work, `select` assignment must have receive on right hand
		// case val := fn():
		// log.Printf("value: %v", val)
		// }
	}
}
