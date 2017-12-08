package basic

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBitShiftRight(t *testing.T) {
	var b uint8
	b = 255

	fmt.Println(b >> 1)

	b = 128
	fmt.Println(b & uint8(1))
}

func TestBitShiftLeft(t *testing.T) {
	var b uint8
	b = 1

	fmt.Println(b)
	fmt.Println(b << 1)
	fmt.Println(b << 1)
}

func TestBitClear(t *testing.T) {
	n := 127
	fmt.Println("b", strconv.FormatInt(int64(n), 2))
	pos := uint(2)
	mask := ^(1 << pos)
	n &= mask

	fmt.Println("a", strconv.FormatInt(int64(n), 2))
}

func TestBitOperation(t *testing.T) {
	var a, b uint8
	a = 1
	b = 1

	fmt.Println(a ^ b)
}
