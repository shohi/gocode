package runtime_test

import (
	"log"
	"runtime"
	"runtime/debug"
	"testing"
)

func TestRuntime_Debug(t *testing.T) {
	gs := &debug.GCStats{}
	debug.ReadGCStats(gs)
	log.Printf("gc stats: %v", gs)

	info, ok := debug.ReadBuildInfo()

	log.Printf("build info: %v, ok: %v", info, ok)
}

func TestRuntime_GoMaxProcs(t *testing.T) {
	currentGMP := runtime.GOMAXPROCS(0)

	log.Printf("current gmp: %v", currentGMP)

}
