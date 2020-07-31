package map_test

import (
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/shohi/gocode/test/util"
)

var globalMap map[int]int

func TestMap_Expand_Mem(t *testing.T) {

	// disable GC
	debug.SetGCPercent(-1)

	m := make(map[int]int, 100)
	var s1 runtime.MemStats
	var s2 runtime.MemStats

	// runtime.GC()

	runtime.ReadMemStats(&s1)
	util.PrintMemStats("before", &s1)

	/*
		for i := 0; i < 100; i++ {
			m[i] = i
		}
	*/

	runtime.ReadMemStats(&s2)
	util.PrintMemStats("after", &s2)

	globalMap = m
}
