package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChan_Send_ToNil(t *testing.T) {
	var ch chan string

	fmt.Printf("===> channel is nil == %v\n", ch == nil)

	ch = nil

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("send to nil channel blocked")
	case <-ch:
		fmt.Println("send to nil channel blocked")
	}
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
