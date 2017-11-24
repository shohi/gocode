package basic

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestSelectForChannel(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			log.Println("tick.")
		case <-boom:
			log.Println("BOOM!")
			return
		default:
			log.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func TestSelectForOrder(t *testing.T) {

	// Select Not keep the order of declaration
	ch := make(chan string, 1)
	ch <- "1, not default"

	ch2 := make(chan string, 1)
	ch2 <- "2, not default"

	select {
	case v2 := <-ch2:
		log.Println(v2)
	case v1 := <-ch:
		log.Println(v1)
	default:
		log.Println("default")
	}
}

func TestSelectForDefault(t *testing.T) {
	ch := make(chan string)
	flag := false
	start := time.Now()
	log.Printf("process start at: %v\n", start)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cnt := 0
		for {
			select {
			case <-ch:
				log.Printf("after times: %d, finally return\n", cnt)
				log.Printf("process end at: %v\n", time.Now())
				return
			default:
				if !flag {
					flag = true
				} else {
					cnt++
				}
			}
		}
	}()

	time.AfterFunc(10*time.Second, func() {
		ch <- "hello"
	})
	wg.Wait()
}
