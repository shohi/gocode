package array_test

import (
	"log"
	"testing"
)

type session struct {
	name  string
	value int
}

func TestZero(t *testing.T) {
	type Arr [10]session
	var a Arr
	log.Printf("len: %v, cap: %v, first: [%#v]", len(a), cap(a), a[0])

}
