package basic

import (
	"fmt"
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
}

func TestBitOperation(t *testing.T) {
	var a, b uint8
	a = 1
	b = 1

	fmt.Println(a ^ b)
}
