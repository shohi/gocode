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
	for i := 0; i < 40; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知,阻塞当前goroutine, 得到唤醒后会持有Lock
			fmt.Println(x)
			time.Sleep(time.Second * 1)
		}(i)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Broadcast() //3秒之后 下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 60)
}
