package channel_test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestBufferedChan_Same_Goroutine(t *testing.T) {
	ch := make(chan string, 1)

	log.Printf("++++ send to channel")
	ch <- "hello"

	msg := <-ch
	log.Printf("++++ receive from channel, msg: %v", msg)
}

func TestChan_Same_Block(t *testing.T) {
	t.Skip("unbuffered channel will not work when sending and receiving in the same goroutine.")
	ch := make(chan string)
	ch <- "hello"
	<-ch
}

func TestChan_Unbuffered_continue(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- "hello"
		fmt.Printf("====> send success")

	}()

	time.Sleep(10 * time.Second)
}
