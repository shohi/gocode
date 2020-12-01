package select_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestSelect_OneBranch(t *testing.T) {
	fn := func() string {
		time.Sleep(10 * time.Second)
		return "Hi"
	}

	_ = fn

	r, _ := http.NewRequest("GET", "www.google.com", nil)
	r.WithContext(context.Background())

	// sigCh := make(chan struct{})

	select {
	case <-time.After(2 * time.Second):
		log.Printf("receive signal")

		// NOTE: not work, `select` assignment must have receive on right hand
		// case val := fn():
		// log.Printf("value: %v", val)
		// }
	}
}

func TestSelect_Continue(t *testing.T) {
	strCh := make(chan string, 3)
	strCh <- "hello"
	strCh <- "world"

	intCh := make(chan int, 3)
	intCh <- 1
	intCh <- 2

	for {
		select {
		case v := <-strCh:
			fmt.Printf("===> ticker 1 - %v\n", v)
			// break
			continue
			fmt.Printf("==> unreachable")
			/*
				case v := <-intCh:
					fmt.Printf("===> ticker 2 - %v\n", v)
					continue
			*/

		}

		fmt.Println("====> Aha")
	}
}
