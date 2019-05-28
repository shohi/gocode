package runtime_test

import (
	"runtime"
	"testing"

	"github.com/shohi/gocode/test/util"
)

func TestMemStat_String(t *testing.T) {
	var m1 runtime.MemStats // before string allocation
	var m2 runtime.MemStats // after string allocation
	var m3 runtime.MemStats // call GC without reset string
	var m4 runtime.MemStats // call GC and reset string

	runtime.GOMAXPROCS(1)
	/*
		var b strings.Builder
		b.Grow(1024 * 1024)
		for k := 0; k < 1024*1024; k++ {
			b.WriteByte('s')
		}
	*/

	runtime.GC()
	var b [1024]byte
	c := b[:0]

	// 1. before
	runtime.ReadMemStats(&m1)

	// 2. assgin string
	// str := strings.Repeat("S", 1024*1024)
	// str := b.String()
	c = append(c, []byte("12345678")...)
	str := string(c)
	// b = append(b, 'c')
	runtime.ReadMemStats(&m2)

	// 3. gc
	runtime.GC()
	runtime.ReadMemStats(&m3)

	// 4. set string to zero-value
	str = ""
	runtime.GC()
	runtime.ReadMemStats(&m4)

	util.PrintMemStats("before", &m1)
	util.PrintMemStats("after", &m2)
	util.PrintMemStats("gc-without-reset", &m3)
	util.PrintMemStats("gc-with-reset", &m4)

	_ = str
}

type Field struct {
	Name  string
	Value interface{}
}

func appendField(buf []Field, v interface{}) []Field {
	if f, ok := v.(Field); ok {
		buf = append(buf, f)
	}
	return buf
}

func TestMemStat_Interface(t *testing.T) {
	var m1 runtime.MemStats // before assignment
	var m2 runtime.MemStats // after assignment

	var buf [1024]Field
	b := buf[:0]

	runtime.GC()
	runtime.ReadMemStats(&m1)

	v := Field{"Hello", "world"}
	b = appendField(b, v)
	runtime.ReadMemStats(&m2)

	util.PrintMemStats("before", &m1)
	util.PrintMemStats("after", &m2)

	_ = v
	_ = b
}
