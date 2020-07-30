package channel_test

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestChan_CloseOrNot(t *testing.T) {
	ch := make(chan error, 10)
	close(ch)
	select {
	case err, ok := <-ch:
		log.Printf("error: [%v], ok: [%v]", err, ok)
	}
}

func TestChan_ReadFromClosedBuffered(t *testing.T) {
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

func TestChan_WriteToClosed(t *testing.T) {
	ch := make(chan string, 10)
	close(ch)

	// NOTE: sending on closed channel causes panic
	ch <- "hello"

	fmt.Println(ch)
}

type MyStruct struct {
	Key string
	Val int
}

func TestChan_RxFromClosed(t *testing.T) {
	ch := make(chan MyStruct, 10)
	ch <- MyStruct{Key: "k1", Val: 1}

	v := <-ch
	fmt.Printf("=====> %v\n", v)
	close(ch)

	v = <-ch
	fmt.Printf("=====> %v\n", v)
}
