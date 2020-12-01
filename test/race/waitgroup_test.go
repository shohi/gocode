package race_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRace_Sleep(t *testing.T) {
	var total int64 = 0

	for i := 0; i < 10; i++ {
		go func(k int) {
			atomic.AddInt64(&total, int64(k))
		}(i)
	}

	// NOTE: race occurs, sleep time can't asure
	// computing goroutines exit before main goroutine
	// wait computing done
	time.Sleep(10 * time.Second)

	fmt.Printf("total: %v\n", total)
}

func TestRace_WaitGroup(t *testing.T) {
	var total int64 = 0
	var count = 10
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(k int) {
			defer wg.Done()
			atomic.AddInt64(&total, int64(k))
		}(i)
	}
	// NOTE: no race here
	// wait computing done
	wg.Wait()

	fmt.Printf("total: %v\n", total)
}
