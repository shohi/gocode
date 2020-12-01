package race_test

import (
	"sync/atomic"
	"testing"
	"time"
)

type MyStruct struct {
	val int32
}

func (m *MyStruct) Load() int {
	return int(atomic.LoadInt32(&m.val))
}

func (m *MyStruct) Store(v int) {
	atomic.StoreInt32(&m.val, int32(v))
}

func TestAtomicRace(t *testing.T) {
	m := MyStruct{
		val: 16,
	}

	// getter - multiple goroutine
	for i := 0; i < 10; i++ {
		go func(k int) {
			_ = m.Load()
		}(i)
	}

	// setter - one goroutine
	go func() {
		for {
			m.Store(10)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
}
