package util

import (
	"log"
	"runtime"
)

func PrintMemStats(name string, m *runtime.MemStats) {
	log.Printf("%v: sys - %v, heap - %v, heap_objects: %v,malloc - %v, free - %v",
		name,
		m.Sys,
		m.HeapObjects,
		m.HeapAlloc,
		m.Mallocs,
		m.Frees)
}
