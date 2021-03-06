package string_test

import (
	"log"
	"strings"
	"testing"
)

func TestStringCompareWithEqual(t *testing.T) {
	a := "bb"
	b := "bb"

	log.Println(a == b)
}

func TestStringCompareWithFold(t *testing.T) {
	want := true
	got := strings.EqualFold("Get", "GET")

	if want != got {
		t.Errorf("strings.EqualFold(%q, %q) = %v, want %v", "Get", "GET", got, want)
	}
}
func TestStringCompare(t *testing.T) {
	a := "bb"
	b := "bb"

	log.Println(a == b)
}

func TestCmp_Numeric(t *testing.T) {
	log.Printf("value: %v", '9' >= '1')
}
