package time_test

import (
	"log"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	log.Println("world")
	time.AfterFunc(200*time.Millisecond, func() {
		log.Println("hello")
		wg.Done()
	})

	wg.Wait()
}

func TestTimer_Goroutine(t *testing.T) {
	grStart := runtime.NumGoroutine()
	log.Printf("goroutine start number: %v", grStart)
	memStart := &runtime.MemStats{}
	runtime.ReadMemStats(memStart)
	log.Printf("memstat start: %v", memStart.NumGC)

	var wg sync.WaitGroup
	cnt := 1
	wg.Add(cnt)
	for k := 0; k < cnt; k++ {
		go func(v int) {
			for j := 0; j < 1000; j++ {
				tr := time.NewTimer(1 * time.Microsecond)
				<-tr.C
			}
			wg.Done()
		}(k)
	}

	wg.Wait()

	grEnd := runtime.NumGoroutine()
	memEnd := &runtime.MemStats{}
	runtime.ReadMemStats(memEnd)
	log.Printf("goroutine end number: %v", grEnd)
	log.Printf("memstat end: %v", memEnd.NumGC)
}

func TestTimer_StopTwice(t *testing.T) {
	tm := time.NewTimer(2 * time.Second)
	res := tm.Stop()
	res2 := tm.Stop()

	log.Printf("1st stop: %v, 2nd stop: %v", res, res2)
}
