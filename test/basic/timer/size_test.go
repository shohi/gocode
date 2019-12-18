package timer_test

import (
	"log"
	"runtime"
	"testing"
	"time"
	"unsafe"

	"github.com/shohi/gocode/test/util"
)

var testTimerArr interface{}

//go:noinline
func escape(v []*time.Timer) {
	testTimerArr = v
}

func TestTimer_size(t *testing.T) {
	runtime.GC()

	var m1 runtime.MemStats
	var m2 runtime.MemStats

	runtime.ReadMemStats(&m1)

	size := 65536
	timerArr := make([]*time.Timer, size)
	for i := 0; i < size; i++ {
		timerArr[i] = time.NewTimer(1 * time.Minute)
	}

	runtime.ReadMemStats(&m2)

	util.PrintMemStats("before", &m1)
	util.PrintMemStats("after", &m2)

	log.Printf("size: %v", unsafe.Sizeof(*timerArr[0]))
	escape(timerArr)
}
