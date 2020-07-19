package channel_test

import (
	"testing"
	"time"
)

func TestInitChan_WithNegative(t *testing.T) {
	sz := -1
	ch := make(chan string, sz)

	go func() {
		ch <- "hello"
	}()

	time.Sleep(10 * time.Second)
}
