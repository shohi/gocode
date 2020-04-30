package switch_test

import (
	"fmt"
	"testing"
)

func TestSwitch_Order(t *testing.T) {
	// NOTE: random order
	for i := 0; i < 5; i++ {
		ch1 := make(chan string, 2)
		ch1 <- "v1"

		ch2 := make(chan string, 2)
		ch2 <- "v2"

		select {
		case v := <-ch1:
			fmt.Printf("====> %v:  %v\n", i, v)
		case v := <-ch2:
			fmt.Printf("====> %v:  %v\n", i, v)
		}
	}

}
