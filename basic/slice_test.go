package basic

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var aa []int
	for i := 0; i < 5; i++ {
		aa = append(aa, 0)
	}

	fmt.Println(aa)
}

func TestSliceCapAndLen(t *testing.T) {
	aa := make([]byte, 5)
	fmt.Println(len(aa), cap(aa))

	cc := make([]byte, 10, 20)
	fmt.Println(len(cc), cap(cc))

	var bb []byte
	fmt.Println(len(bb), cap(bb), bb, bb == nil)
	bb = append(bb, 0x10)
	fmt.Println(len(bb), cap(bb), bb, bb == nil)
}

func TestSliceAppend(t *testing.T) {
	var aa []int
	aa = append(aa, 10, 20)
	fmt.Println(aa)
}
