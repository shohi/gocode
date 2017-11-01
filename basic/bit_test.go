package basic

import (
	"fmt"
	"testing"
)

func TestBitShift(t *testing.T) {
	var b uint8
	b = 255

	fmt.Println(b >> 1)

	b = 128
	fmt.Println(b & uint8(1))
}
