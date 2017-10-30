package basic

import (
	"fmt"
	"testing"
)

func TestGoroutine(t *testing.T) {
	ch := make(chan interface{})
	go func() {
		ch <- nil
	}()

	info := <-ch
	fmt.Println(info)
}
