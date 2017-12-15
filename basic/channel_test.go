package basic

import (
	"log"
	"testing"
)

func TestChannelCapacity(t *testing.T) {
	ch := make(chan int, 100)

	log.Println(cap(ch))
	log.Println(len(ch))
}

func TestChannelSendNil(t *testing.T) {
	ch := make(chan int)
	if 2 > 1 {
		close(ch)
	}

	// write to closed channel will raise panic
	// ch <- 10

}
