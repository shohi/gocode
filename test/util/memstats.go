package util

import (
	"log"
	"runtime"
)

func PrintMemStats(name string, m *runtime.MemStats) {
	log.Printf("%v: sys - %v, heap_alloc - %v, heap_objects: %v, heap_inuse: %v, malloc - %v, free - %v",
		name,
		m.Sys,
		m.HeapAlloc,
		m.HeapObjects,
		m.HeapInuse,
		m.Mallocs,
		m.Frees)
}

func PrintMemNow(name string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	PrintMemStats(name, &m)
}
