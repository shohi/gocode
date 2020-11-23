package channel_test

import (
	"fmt"
	"testing"
	"time"
)

// NOTE: read from nil channel will block
func TestChan_ReadFromNil(t *testing.T) {
	var ch chan string

	fmt.Printf("===> channel is nil == %v\n", ch == nil)
	ch = nil

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("read from nil channel: blocked")
	case <-ch:
		fmt.Println("read from nil channel: not blocked")
	}
}

// NOTE: send to nil channel will block
func TestChan_SendToNil(t *testing.T) {
	var ch chan struct{}

	select {
	case ch <- struct{}{}:
		fmt.Println("send to nil channel: not blocked")
	case <-time.After(5 * time.Second):
		fmt.Println("send to nil channel: blocked")
	}
}

// NOTE: send to closed channel will PANIC
func TestChan_SendToClosed(t *testing.T) {
	var ch chan struct{}
	close(ch)

	ch <- struct{}{}
}

func TestChan_Reset(t *testing.T) {
	type myStruct struct {
		ch chan string
	}

	ms := myStruct{
		ch: make(chan string, 10),
	}
	aa := ms.ch

	aa = nil

	fmt.Printf("aa is nil: %v, ms.ch is nil: %v\n",
		aa == nil,
		ms.ch == nil)

}
