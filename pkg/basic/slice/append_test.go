package slice_test

import (
	"log"
	"testing"
)

func TestAppendWithInitCap(t *testing.T) {
	// case 1 - omit len parameter
	// len === cap
	b := make([]int, 10)
	b = append(b, 99)
	log.Printf("len: [%d], cap: [%d], content - %v", len(b), cap(b), b)

	// case 2 - use len parameter
	b = make([]int, 0, 10)
	b = append(b, 99)
	log.Printf("len: [%d], cap: [%d], content - %v", len(b), cap(b), b)
}

func TestAppendWithNil(t *testing.T) {
	var s []*int
	var a *int
	s = append(s, a)
	log.Printf("slice: %v, len: %v, cap: %v", s, len(s), cap(s))

	//
	var chSlice []chan struct{}
	chSlice = append(chSlice, nil)
	log.Printf("slice: %v, len: %v, cap: %v", s, len(s), cap(s))

	// slice append slice with expandation
	aa := []int{1, 2, 3}
	var bb []int
	aa = append(aa, bb...)
	log.Printf("slice aa: %v", aa)

	bb = []int{4, 5}
	aa = append(aa, bb...)
	log.Printf("slice aa: %v", aa)

	bb = nil
	aa = append(aa, bb...)
	log.Printf("slice aa: %v", aa)
}

func TestSliceModification(t *testing.T) {
	var dst []byte

	fn := func(d []byte) []byte {
		d = append(d, []byte("hello")...)
		return d
	}

	dst = fn(dst)

	log.Printf("value: %v", string(dst))
}

func TestSliceCopy(t *testing.T) {
	// case 1 - empty dst
	dst := make([]byte, 20)
	log.Printf("cap: %d, len: %d, content: [%s]\n", cap(dst), len(dst), string(dst))
	// src = append(src, []byte("hello")...)
	bs := []byte("hello")
	copy(dst, bs)

	log.Printf("cap: %d, len: %d, content: [%s]\n", cap(dst), len(dst), string(dst))

	// case 2 - non-empty dst
	dst = []byte("this-is-a-long-byte-slice")
	bs = []byte("world")
	n := copy(dst, bs)
	log.Printf("cap: %d, len: %d, content: [%s], c: %d\n", cap(dst), len(dst), string(dst), n)
}

func TestSliceNilReset(t *testing.T) {
	var dst []byte
	dd := dst[:0]

	log.Printf("slice: %v", dd)
}

func TestSliceAppendNil(t *testing.T) {
	dst := []byte("hello")
	var app []byte
	dst = append(dst, app...)

	log.Printf("value ==> %v", dst)
}
