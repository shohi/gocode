package channel_test

import (
	"log"
	"testing"
)

func TestBufferedChannel_Same_Goroutine(t *testing.T) {
	ch := make(chan string, 1)

	log.Printf("++++ send to channel")
	ch <- "hello"

	msg := <-ch
	log.Printf("++++ receive from channel, msg: %v", msg)
}

func TestChannel_Same_Block(t *testing.T) {
	t.Skip("unbuffered channel will not work when sending and receiving in the same goroutine.")
	ch := make(chan string)
	ch <- "hello"
	<-ch
}
