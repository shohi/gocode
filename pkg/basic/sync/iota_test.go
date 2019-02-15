package sync_test

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
	// mutexMax         = iota
)

func TestIOTA(t *testing.T) {
	log.Printf("mutexLocked: [%d]", mutexLocked)
	log.Printf("mutexWoken: [%d]", mutexWoken)
	log.Printf("mutexStarving: [%d]", mutexStarving)
	log.Printf("mutexWaiterShift: [%d]", mutexWaiterShift)
	// log.Printf("mutexMax: [%d]", mutexMax)

	// shift
	log.Printf("mutexLocked shifted by mutexWaiterShift: [%d]", mutexLocked>>mutexWaiterShift)
	log.Printf("mutexWoken shifted by mutexWaiterShift: [%d]", mutexWoken>>mutexWaiterShift)
	log.Printf("mutexWaiterShift shifted by mutexWaiterShift: [%d]", mutexWaiterShift>>mutexWaiterShift)
	log.Printf("mutexStarving shifted by mutexWaiterShift: [%d]", mutexStarving>>mutexWaiterShift)
}

func TestCAS(t *testing.T) {
	var a int32 = 10
	var wg sync.WaitGroup
	count := 10
	wg.Add(count)
	for k := 0; k < count; k++ {
		go func(i int) {
			swapped := atomic.CompareAndSwapInt32(&a, 10, 20)
			log.Printf("goroutine-[%d]: swapped-[%v], value-[%v]", i, swapped, a)
			wg.Done()
		}(k)
	}

	wg.Wait()

	log.Printf("=====final result: %v", a)
}
