package basic

import (
	"log"
	"math"
	"testing"
)

func TestUintConvert(t *testing.T) {
	var a int
	a = -1
	log.Println(uint32(a))
}

func TestUint32(t *testing.T) {
	a := uint32(10)
	var b uint32
	b = math.MaxUint32

	log.Println(a, b+uint32(8))
	log.Println(a > (b + uint32(8)))
}
