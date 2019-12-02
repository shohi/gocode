package lambda_test

import (
	"log"
	"runtime"
	"testing"

	"github.com/shohi/gocode/test/util"
)

func TestLambda_Init(t *testing.T) {
	var valueSlice []string
	log.Printf("len: %v", len(valueSlice))

	var m1 runtime.MemStats
	var m2 runtime.MemStats

	runtime.ReadMemStats(&m1)

	fn := func() {
		valueSlice = []string{"1", "2", "3"}
	}

	var aa = 2
	if aa > 3 {
		fn()
	}

	runtime.ReadMemStats(&m2)

	util.PrintMemStats("start", &m1)
	util.PrintMemStats("end", &m2)

	log.Printf("len: %v", len(valueSlice))
}
