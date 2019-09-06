package channel_test

import (
	"log"
	"testing"
	"time"
)

func TestChannel_Unbuffer_Sender(t *testing.T) {
	ch := make(chan struct{}, 0)
	go func() {
		// NOTE: sender will block until element is taken out
		ch <- struct{}{}
		log.Printf("send complete")
	}()

	time.Sleep(5 * time.Second)
	log.Printf("receive complete")
}