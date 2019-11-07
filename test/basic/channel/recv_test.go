package channel_test

import (
	"log"
	"testing"
)

// NOTE: A receive from a closed channel returns the *zero value* immediately
// https://dave.cheney.net/2014/03/19/channel-axioms
func TestRecvFromClosedChannel(t *testing.T) {
	ch := make(chan int)
	close(ch)

	err := <-ch
	log.Printf("err: [%+v]", err)
}
