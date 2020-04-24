package channel_test

import (
	"log"
	"testing"
	"time"
)

func goChanLoop(strCh chan string, pause time.Duration) {
	go func() {
		for {
			time.Sleep(pause)
			err, ok := <-strCh
			log.Printf("====> str: %v, ok: %v\n", err, ok)
			if !ok {
				return
			}
		}
	}()
}

func TestBufferedOne(t *testing.T) {
	strCh := make(chan string, 1)
	goChanLoop(strCh, 0)

	log.Println("### send event")
	strCh <- "s"
	close(strCh)
	log.Println("### send event done")

	time.Sleep(5 * time.Second)
}

func TestBufferedOne_Select_Default(t *testing.T) {
	ch := make(chan string, 1)
	goChanLoop(ch, 100*time.Millisecond)

	ch <- "hello"
	select {
	case ch <- "world":
		log.Printf("### another string: %v\n", "world")
	default:
		log.Printf("### buffer full")
	}

	time.Sleep(1 * time.Second)
	close(ch)
	time.Sleep(1 * time.Second)
}
