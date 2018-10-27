package string_test

import (
	"log"
	"testing"
)

func TestStringByteSliceConversion(t *testing.T) {
	str := "12"
	b := []byte(str)
	log.Println(str, b)

	b = []byte{0x01, 0x02}
	str = string(b)
	log.Println(str, b)
}

func TestCopyBetweenStringAndByteSlice(t *testing.T) {
	// case 1 - copy string to slice
	a := make([]byte, 10)
	var str string = "hello"
	copy(a, str)
	log.Printf("slice - [%v], string - [%v]", string(a), str)

	// case 2 - copy slice to string
	// NOTE: not work, the first argument of copy MUST be slice
	// var str2 string
	// copy(str2, a)
}
