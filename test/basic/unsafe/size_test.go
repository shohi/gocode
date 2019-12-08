package unsafe_test

import (
	"fmt"
	"runtime"
	"testing"
)

var tmpCollect = []chan int{}

func escape_fn(arr []chan int) {
	tmpCollect = arr
}

func TestUnsafe_Channel_Size(t *testing.T) {
	// baseSize := getChanAlloc(0)
	// Output: 96
	// log.Printf("unbuffered channel size: %v", baseSize)

	fmt.Printf("size, alloc\n")
	for k := 0; k < (2 << 5); k++ {
		fmt.Printf("%v, %v\n", k, getChanAlloc(k))
	}

	// log.Printf("arr size: %v", unsafe.Sizeof(arr))
}

func getChanAlloc(bufSize int) uint64 {
	runtime.GC()

	var m1 runtime.MemStats
	var m2 runtime.MemStats

	arr := make([]chan int, 1)
	runtime.ReadMemStats(&m1)
	for k := 0; k < len(arr); k++ {
		if bufSize <= 0 {
			arr[k] = make(chan int)
		} else {
			arr[k] = make(chan int, bufSize)
		}
	}
	runtime.ReadMemStats(&m2)

	escape_fn(arr)

	return m2.Alloc - m1.Alloc
}
