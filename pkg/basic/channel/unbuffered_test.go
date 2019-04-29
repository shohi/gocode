package channel_test

import (
	"log"
	"testing"
)

func TestChannel_Unbuffer_Sender(t *testing.T) {
	ch := make(chan struct{}, 0)
	go func() {
		// NOTE: sender doesn't block for the first time
		ch <- struct{}{}
		log.Printf("send complete")
	}()

	<-ch
	log.Printf("receive complete")
}
