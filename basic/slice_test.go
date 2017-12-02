package basic

import (
	"log"
	"testing"
)

func TestSlice(t *testing.T) {
	var aa []int
	for i := 0; i < 5; i++ {
		aa = append(aa, 0)
	}

	log.Println(aa)
}

func TestSliceCapAndLen(t *testing.T) {
	aa := make([]byte, 5)
	log.Println(len(aa), cap(aa))

	cc := make([]byte, 10, 20)
	log.Println(len(cc), cap(cc))

	var bb []byte
	log.Println(len(bb), cap(bb), bb, bb == nil)
	bb = append(bb, 0x10)
	log.Println(len(bb), cap(bb), bb, bb == nil)
}

func TestSliceAppend(t *testing.T) {
	var aa []int
	aa = append(aa, 10, 20)
	log.Println(aa)

	aa = make([]int, 5)
	aa = append(aa, 30, 40)
	log.Println(aa)
}

func TestSliceInitialization(t *testing.T) {
	// initialize slice without length parameter will cause error
	// aa := make([]int)
	aa := make([]int, 10)
	log.Println(aa)
}

func TestNilSliceTraverse(t *testing.T) {
	var a []*int

	for k, v := range a {
		log.Println(k, v)
	}
}
