package basic

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "ABCDE"

	fmt.Printf("%v ==> %v\n", ([]byte)(s), s)

}

func TestStringCompare(t *testing.T) {
	a := "bb"
	b := "bb"

	fmt.Println(a == b)
}
