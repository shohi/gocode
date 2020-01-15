package sync_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// refer, https://my.oschina.net/u/553243/blog/1799695
func TestCondWithSimpleCase(t *testing.T) {
	var locker = new(sync.Mutex)
	var cond = sync.NewCond(locker)
	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			fmt.Printf("get lock => %v\n", x)

			cond.Wait() //等待通知,阻塞当前goroutine, 得到唤醒后会持有Lock
			fmt.Println(x)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	// NOTE: 100ms后下发一个通知给已经获取锁的goroutine
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Signal 1...")
	cond.Signal()

	// NOTE: 100ms后再下发一个通知给已经获取锁的goroutine
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Signal 2...")
	cond.Signal()

	// NOTE: 100ms后下发广播给所有等待的goroutine
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Broadcast...")
	cond.Broadcast()

	// NOTE: wait for other goroutines completed
	time.Sleep(2 * time.Second)
}
