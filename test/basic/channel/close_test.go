package channel_test

import (
	"log"
	"testing"
	"time"
)

func TestCloseChannel(t *testing.T) {
	ch := make(chan error, 10)
	close(ch)
	select {
	case err, ok := <-ch:
		log.Printf("error: [%v], ok: [%v]", err, ok)
	}
}

func TestReadFromClosedBufferedChannel(t *testing.T) {
	ch := make(chan int, 10)
	for k := 0; k < 5; k++ {
		ch <- k
	}

	close(ch)

	for k := range ch {
		log.Println(k)
	}
}

func TestChan_Close(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- "hello"
		log.Printf("++++ running in goroutine")
	}()

	<-ch
	time.Sleep(1 * time.Second)
	close(ch)

	log.Printf("+++++ running in main")
}
