package basic

import (
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
