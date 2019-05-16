package channel_test

import (
	"fmt"
	"log"
	"strconv"
	"sync"
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

func TestChannelGetFromClosed(t *testing.T) {
	// channel will send default value of the type immediately
	// when it get closed.
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		count := 0
		for {
			j := <-ch
			log.Printf("int ==> %v", j)

			count++
			if count > 20 {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- "str" + strconv.Itoa(i)
		}
		close(ch)
	}()

	wg.Wait()
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1

	// channel will be always block until available
	for range ch {
		fmt.Println(<-ch)
	}

}

func TestSendNilToChannel(t *testing.T) {
	ch := make(chan error, 2)
	ch <- nil
	ch <- nil

	log.Printf("channel - len: [%v], cap: [%v]", len(ch), cap(ch))
	res := <-ch
	log.Printf("get result: %v", res)
}

// Reading from nil channel will block forever
/*
func TestReceiveFromNilChannel(t *testing.T) {
	var ch chan error
	ch = nil

	val := <-ch
	log.Println(val)
}
*/

func TestChannelCopy(t *testing.T) {
	ch := make(chan struct{}, 10)

	var ch2 chan struct{}
	ch2 = ch
	log.Printf("%T, %v, %p", ch, ch, &ch)
	log.Printf("%T, %v, %p", ch2, ch2, &ch2)

	t.Logf("hello")
}
