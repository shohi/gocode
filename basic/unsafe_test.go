package basic

import (
	"log"
	"math"
	"testing"
	"unsafe"
)

func TestUnsafeSizeOf(t *testing.T) {
	var bb int
	aa := make([]int, math.MaxUint32)
	log.Println(int(unsafe.Sizeof(&bb)) * len(aa) / (1024 * 1024))

}
