package basic

import (
	"fmt"
	"strings"
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

func TestStringAndBytes(t *testing.T) {
	str := "12"
	b := []byte(str)

	fmt.Println(str, b)

	b = []byte{0x01, 0x02}
	str = string(b)

	fmt.Println(str, b)
}

func TestStringsFold(t *testing.T) {
	want := true
	got := strings.EqualFold("Get", "GET")

	if want != got {
		t.Errorf("strings.EqualFold(%q, %q) = %v, want %v", "Get", "GET", got, want)
	}
}
