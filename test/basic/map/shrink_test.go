package map_test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"
)

func clearMap(m map[int]struct{}) {
	for k := range m {
		delete(m, k)
	}
}

func TestMapShrink_Clear(t *testing.T) {
	v := struct{}{}

	a := make(map[int]struct{})

	for i := 0; i < 10000; i++ {
		a[i] = v
	}

	runtime.GC()
	printMemStats("After Map Add 100000")

	for i := 0; i < 10000-1; i++ {
		delete(a, i)
	}

	runtime.GC()
	printMemStats("After Map Delete 9999")

	for i := 0; i < 10000-1; i++ {
		a[i] = v
	}

	runtime.GC()
	printMemStats("After Map Add 9999 again")
	fmt.Printf("%d\n", len(a))

	runtime.GC()
	printMemStats("After Map Add 9999 again")
	fmt.Printf("%d\n", len(a))

	for k := range a {
		delete(a, k)
	}
	runtime.GC()

	printMemStats("After Map Clear")
	fmt.Printf("%d\n", len(a))

	a = nil
	runtime.GC()
	printMemStats("After Map Set nil")
}

func populateMap(m map[string]string, cnt int) {
	for i := 0; i < cnt; i++ {
		m[fmt.Sprintf("key-%d", i)] = fmt.Sprintf("value-%d", i)
	}
}

func TestMapClear_Mem(t *testing.T) {
	m := make(map[string]string)

	ticker := time.NewTicker(5 * time.Second)
	count := 0
	for range ticker.C {
		printMemStats(fmt.Sprintf("iter-%d-before", count))

		populateMap(m, 10000)

		printMemStats(fmt.Sprintf("iter-%d-populate", count))

		// delete
		for k := range m {
			delete(m, k)
		}

		printMemStats(fmt.Sprintf("iter-%d-delete", count))

		runtime.GC()
		printMemStats(fmt.Sprintf("iter-%d-GC", count))

		fmt.Printf(strings.Repeat("\n", 3))
	}
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v：memory = %vKB, GC Times = %v\n", mag, m.Alloc/1024, m.NumGC)
}
