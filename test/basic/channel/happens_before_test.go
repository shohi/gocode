package channel_test

/*
 *
 * go happens-before常用的三原则是：
 *
 * 1. 对于不带缓冲区的channel，对其写happens-before对其读
 * 2. 对于带缓冲区的channel,对其读happens-before对其写
 * 3. 对于不带缓冲的channel的接收操作 happens-before 相应channel的发送操作完成
 */

import (
	"fmt"
	"testing"
	"time"
)

func TestHappensBefore_buffered(t *testing.T) {
	var c = make(chan int, 10)
	f := func() {
		c <- 0
		fmt.Println("send done")
	}

	go f()
	<-c

	fmt.Println("recv done")
	time.Sleep(1 * time.Second)
}

func TestHappensBefore_unbuffered(t *testing.T) {
	var c = make(chan int)
	f := func() {
		c <- 0
		fmt.Println("send done")
	}

	go f()
	<-c

	fmt.Println("recv done")
	time.Sleep(1 * time.Second)
}

func TestHappensBefore_3(t *testing.T) {
	var c = make(chan int)
	f := func() {
		<-c
		fmt.Println("recv done")
	}

	go f()
	c <- 0
	fmt.Println("send done")

	time.Sleep(1 * time.Second)
}

func TestHappensBefore_4(t *testing.T) {
	var c = make(chan int, 1)
	f := func() {
		<-c
		fmt.Println("recv done")
	}

	go f()
	c <- 0
	fmt.Println("send done")

	time.Sleep(1 * time.Second)
}
