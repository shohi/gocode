package channel_test

import (
	"fmt"
	"log"
	"sync"
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

func TestChan_CloseAfterSend(t *testing.T) {
	ch := make(chan string, 8)
	ready := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		wg.Done()

		<-ch
		close(ready)

		for msg := range ch {
			fmt.Printf("======> %v\n", msg)
		}

		fmt.Printf("goroutine one over\n")
	}()

	ch <- "start"
	<-ready

	go func() {
		defer wg.Done()
		ch <- "hello"
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	wg.Wait()

}
