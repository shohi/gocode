package channel_test

import (
	"log"
	"testing"
	"time"
)

func TestChan_WaitOnClosed(t *testing.T) {

	ch := make(chan string)

	go func() {
		<-ch
		log.Printf("++++++ in goroutine")
	}()

	time.Sleep(1 * time.Second)
	ch = nil

	log.Printf("++++++ in main")

}
