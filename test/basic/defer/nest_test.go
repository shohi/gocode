package defer_test

import (
	"log"
	"testing"
)

func TestDefer_InSelect(t *testing.T) {
	msgCh := make(chan string, 2)
	doneCh := make(chan struct{})

	msgCh <- "hello"

	select {
	case <-msgCh:
		defer log.Printf("hello msg")
	case <-doneCh:
	}

	log.Printf("running in main")
}
