package escape

import (
	"log"
	"runtime"
	"testing"

	"github.com/shohi/gocode/test/util"
)

func TestEscape(t *testing.T) {
	var m1 runtime.MemStats
	var m2 runtime.MemStats

	var base [12]Field
	runtime.GC()
	runtime.ReadMemStats(&m1)
	buf := base[:0]
	f := Field{"key", "value"}
	buf = appendField(buf, f)
	runtime.ReadMemStats(&m2)

	log.Printf("%v", Displayer(f))
	util.PrintMemStats("before", &m1)
	util.PrintMemStats("after", &m2)
}
