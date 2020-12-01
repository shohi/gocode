package string_test

import (
	"fmt"
	"testing"
)

func TestString_Rune(t *testing.T) {
	fmt.Printf("val: %v, rune: %v, string: %v\n",
		10, rune(10), string(rune(10)))

}
